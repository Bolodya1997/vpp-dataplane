// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package crypto_sw_scheduler

import (
	"context"

	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service crypto_sw_scheduler.
type RPCService interface {
	CryptoSwSchedulerSetWorker(ctx context.Context, in *CryptoSwSchedulerSetWorker) (*CryptoSwSchedulerSetWorkerReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) CryptoSwSchedulerSetWorker(ctx context.Context, in *CryptoSwSchedulerSetWorker) (*CryptoSwSchedulerSetWorkerReply, error) {
	out := new(CryptoSwSchedulerSetWorkerReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
