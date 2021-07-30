#!/bin/bash
set -e

function git_cherry_pick ()
{
	refs=$1
	git fetch "https://gerrit.fd.io/r/vpp" ${refs}
	git cherry-pick FETCH_HEAD
	git commit --amend -m "gerrit:${refs#refs/changes/*/} $(git log -1 --pretty=%B)"
}

if [ -z "$1" ]; then
	echo "Missing VPP path"
	exit 1
fi

VPP_COMMIT=328b5dadb35916e8147237f2339dd5f4c38912fc
VPP_DIR="$1"

if [ ! -d $1/.git ]; then
	rm -rf $1
	git clone "https://gerrit.fd.io/r/vpp" $1
	cd $1
	git reset --hard ${VPP_COMMIT}
else
	cd $1
	git fetch "https://gerrit.fd.io/r/vpp" && git reset --hard ${VPP_COMMIT}
fi

git_cherry_pick refs/changes/68/32468/3 # 32468: buffers: fix buffer linearization | https://gerrit.fd.io/r/c/vpp/+/32468
git_cherry_pick refs/changes/33/32833/1 # 32833: ipsec: disable linearization | https://gerrit.fd.io/r/c/vpp/+/32833
git_cherry_pick refs/changes/86/29386/9 # 29386: virtio: DRAFT: multi tx support | https://gerrit.fd.io/r/c/vpp/+/29386
git_cherry_pick refs/changes/21/31321/11 # 31321: devices: add support for pseudo header checksum | https://gerrit.fd.io/r/c/vpp/+/31321

# Revert for now
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/68/32568/6 && git revert --no-edit FETCH_HEAD
git_cherry_pick refs/changes/69/31869/13 # 31869: gso: do not try gro on small packets | https://gerrit.fd.io/r/c/vpp/+/31869
git_cherry_pick refs/changes/76/32476/3 # 32476: gso: handle push flag in gro | https://gerrit.fd.io/r/c/vpp/+/32476
git_cherry_pick refs/changes/82/32482/1 # 32482: virtio: compute cksums in output no offload | https://gerrit.fd.io/r/c/vpp/+/32482
git_cherry_pick refs/changes/83/32483/1 # 32483: virtio: Still init unused txq | https://gerrit.fd.io/r/c/vpp/+/32483
git_cherry_pick refs/changes/71/32871/1 # 32871: devices: Add queues params in create_if | https://gerrit.fd.io/r/c/vpp/+/32871
git_cherry_pick refs/changes/71/32271/6 # 32271: memif: add support for ns abstract sockets | https://gerrit.fd.io/r/c/vpp/+/32271
git_cherry_pick refs/changes/68/31868/1 # 31868: ip6-nd: silent the source and target checks on given interface | https://gerrit.fd.io/r/c/vpp/+/31868
git_cherry_pick refs/changes/35/32235/1  # 32235: dpdk: enable ena interrupt support | https://gerrit.fd.io/r/c/vpp/+/32235
git_cherry_pick refs/changes/36/32636/11 # 32636: hsi: host stack intercept plugin | https://gerrit.fd.io/r/c/vpp/+/32636

# --------------- Dedicated plugins ---------------
git_cherry_pick refs/changes/64/33264/2 # 33264: pbl: Port based balancer | https://gerrit.fd.io/r/c/vpp/+/33264

git_cherry_pick refs/changes/88/31588/1 # 31588: cnat: [WIP] no k8s maglev from pods | https://gerrit.fd.io/r/c/vpp/+/31588

git_cherry_pick refs/changes/83/28083/16 # 28083: acl: acl-plugin custom policies |  https://gerrit.fd.io/r/c/vpp/+/28083
git_cherry_pick refs/changes/13/28513/20 # 25813: capo: Calico Policies plugin | https://gerrit.fd.io/r/c/vpp/+/28513
# --------------- Dedicated plugins ---------------

# Revert due to breaking check for unnumbered interfaces
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/57/33157/6 && git revert --no-edit FETCH_HEAD

git_cherry_pick refs/changes/01/33301/1 # 33301: session: make netns abtract name static | https://gerrit.fd.io/r/c/vpp/+/33301

# TODO : PUNT patch

