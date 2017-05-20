package main

import (
	"flag"
	"github.com/coreos/etcd/wal"
)

var (
	waldir string
)

func init() {
	flag.StringVar(&waldir, "waldir", "", "set wal directory")
}

func main() {
}
