package prometheus

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"git.fd.io/govpp.git/adapter"
	"git.fd.io/govpp.git/adapter/statsclient"
	"github.com/projectcalico/vpp-dataplane/calico-vpp-agent/cni/storage"
	"github.com/projectcalico/vpp-dataplane/vpplink"
	"github.com/sirupsen/logrus"
)

var exported_data string
var sc *statsclient.StatsClient
var g_server Server

func viewHandler(w http.ResponseWriter, r *http.Request) {
	ifNames, dumpStats, _ := vpplink.Statsclientfunc(sc)
	exported_data = g_server.statsToPrometheusFormat(ifNames, dumpStats)
	fmt.Fprint(w, exported_data)
}

type Server struct {
	log             *logrus.Entry
	vpp             *vpplink.VppLink
	podInterfaceMap *map[string]storage.LocalPodSpec
}

func NewServer(vpp *vpplink.VppLink, l *logrus.Entry, podIfMap *map[string]storage.LocalPodSpec) (Server, error) {

	server := &Server{
		log:             l,
		vpp:             vpp,
		podInterfaceMap: podIfMap,
	}
	return *server, nil
}

func (s *Server) statsToPrometheusFormat(ifNames adapter.NameStat, dumpStats []adapter.StatEntry) (text string) {
	for _, sta := range dumpStats {
		if string(sta.Name) != "/if/names" {
			metricNames := []string{strings.Replace(string(sta.Name[4:]), "-", "_", -1)}
			if sta.Type == adapter.CombinedCounterVector {
				metricName := metricNames[0]
				metricNames = []string{metricName + "_packet", metricName + "_bytes"}
			}
			if sta.Type == adapter.SimpleCounterVector {
				values := sta.Data.(adapter.SimpleCounterStat)
				for worker := range values {
					for j := range values[worker] {
						if string(ifNames[j]) != "" {
							s.update_text(ifNames, &text, metricNames[0], strconv.Itoa(int(values[worker][j])), worker, j)
						}
					}
				}
			} else if sta.Type == adapter.CombinedCounterVector {
				values := sta.Data.(adapter.CombinedCounterStat)
				for k := range metricNames {
					for worker := range values {
						for j := range values[worker] {
							if string(ifNames[j]) != "" {
								s.update_text(ifNames, &text, metricNames[k], strconv.Itoa(int(values[worker][j][k])), worker, j)
							}
						}
					}
				}
			}
		}
	}
	return text
}

func (s *Server) update_text(ifNames adapter.NameStat, text *string, metricName string, value string, worker int, j int) *string {
	_, swifindex := s.vpp.SearchInterfaceWithName(string(ifNames[j]))
	for _, pod := range *s.podInterfaceMap {
		if swifindex == pod.SwIfIndex {
			namespace := pod.WorkloadID[:strings.Index(pod.WorkloadID, "/")]
			podName := pod.WorkloadID[strings.Index(pod.WorkloadID, "/")+1:]
			nameInPod := pod.InterfaceName
			*text = *text + "# HELP " + metricName + " any help msg\n"
			*text = *text + "# TYPE " + metricName + " counter\n"
			*text = *text + metricName + "{" + "interfaceName=\"" + string(ifNames[j]) + "\",worker=\"" + strconv.Itoa(worker) +
				"\",namespace=\"" + namespace + "\",podName=\"" + podName + "\",nameInPod=\"" + nameInPod + "\"} " + value + "\n"
		}
	}
	return text
}

func (s *Server) Serve() {
	s.log.Infof("Serve() Prometheus exporter")
	sc = statsclient.NewStatsClient("")
	err := sc.Connect()
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	g_server = *s

	http.HandleFunc("/metrics", viewHandler)
	err = http.ListenAndServe(":2112", nil)
	if err != nil {
		fmt.Println(err)
	}
	sc.Disconnect()

}
