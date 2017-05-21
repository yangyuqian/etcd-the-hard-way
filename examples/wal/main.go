package main

import (
	"github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/pbutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/wal"
	"github.com/coreos/etcd/wal/walpb"
	"log"
	"os"
)

var waldir = "./_test/test.wal"

func main() {
	createWAL()
	printMeta(waldir)
}

func createWAL() {
	os.RemoveAll(waldir)

	meta := pbutil.MustMarshal(&etcdserverpb.Metadata{NodeID: 1, ClusterID: 2})
	w, err := wal.Create(waldir, meta)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	w.Save(raftpb.HardState{
		Term:   3,
		Vote:   1,
		Commit: 1,
	}, nil)
}

func printMeta(path string) {
	w, err := wal.Open(path, walpb.Snapshot{})
	if err != nil {
		panic(err)
	}
	defer w.Close()

	meta, state, _, _ := w.ReadAll()
	pbmeta := &etcdserverpb.Metadata{}
	pbutil.MustUnmarshal(pbmeta, meta)

	log.Printf("cluster: %s node: %s term: %s commit: %s vote: %s",
		types.ID(pbmeta.ClusterID),
		types.ID(pbmeta.NodeID),
		types.ID(state.Term),
		types.ID(state.Commit),
		types.ID(state.Vote),
	)
}
