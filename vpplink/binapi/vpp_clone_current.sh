#!/bin/bash
VPP_COMMIT=e988726cb

if [ ! -d $1/.git ]; then
	rm -rf $1
	git clone "https://gerrit.fd.io/r/vpp" $1
	cd $1
	git reset --hard ${VPP_COMMIT}
else
	cd $1
	git fetch "https://gerrit.fd.io/r/vpp" && git reset --hard ${VPP_COMMIT}
fi

# Run first so that it doesn't hide any patch failures
echo "diff --git a/src/vnet/ip/ip.api b/src/vnet/ip/ip.api
index f201ffbd8..548eb1a2f 100644
--- a/src/vnet/ip/ip.api
+++ b/src/vnet/ip/ip.api
@@ -256,7 +256,7 @@ autoreply define set_ip_flow_hash
     @param symmetric - include symmetry in flow hash
     @param flowlabel - include flowlabel in flow hash
 */
-enumflag ip_flow_hash_config
+enum ip_flow_hash_config
 {
   IP_API_FLOW_HASH_SRC_IP = 0x01,
   IP_API_FLOW_HASH_DST_IP = 0x02,
" | git apply -- && git add -A &&  git commit --author "Calico-vpp builder <>" -m "Remove enumflag not yet supported in govpp"


git fetch "https://gerrit.fd.io/r/vpp" refs/changes/86/29386/9 && git cherry-pick FETCH_HEAD # 29386: virtio: DRAFT: multi tx support | https://gerrit.fd.io/r/c/vpp/+/29386
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/04/30204/8 && git cherry-pick FETCH_HEAD # 30204: vxlan: add udp-port configuration support
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/18/31118/1 && git cherry-pick FETCH_HEAD # 31118: interface: Add promisc on/off in api | https://gerrit.fd.io/r/c/vpp/+/31118
# --------------- Cnat patches ---------------
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/75/30275/11 && git cherry-pick FETCH_HEAD # 30275: cnat: add input feature node | https://gerrit.fd.io/r/c/vpp/+/30275
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/87/28587/31 && git cherry-pick FETCH_HEAD # 28587: cnat: k8s extensions
# --------------- Cnat patches ---------------

# ------------- Policies patches -------------
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/83/28083/14 && git cherry-pick FETCH_HEAD # 28083: acl: acl-plugin custom policies |  https://gerrit.fd.io/r/c/vpp/+/28083
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/13/28513/17 && git cherry-pick FETCH_HEAD # 25813: capo: Calico Policies plugin | https://gerrit.fd.io/r/c/vpp/+/28513
# ------------- Policies patches -------------

git fetch "https://gerrit.fd.io/r/vpp" refs/changes/95/30695/2 && git cherry-pick FETCH_HEAD # 30695: wireguard: testing alternative timer dispatch | https://gerrit.fd.io/r/c/vpp/+/30695
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/17/31017/1 && git cherry-pick FETCH_HEAD # 31017: interface: automask interrupts to polling rxqs | https://gerrit.fd.io/r/c/vpp/+/31017
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/78/31078/3 && git cherry-pick FETCH_HEAD # 31078: devices: af-packet gso mtu | https://gerrit.fd.io/r/c/vpp/+/31078