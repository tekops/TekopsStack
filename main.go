package main

import (
	"flag"
	"github.com/golang/glog"
	"io/ioutil"
	"launchpad.net/goyaml"
	"net"
	"time"
)

const (
	CFG_FILE = "conf/master.yaml"
)

var (
	config  *Config
	listner net.Listener
	err     error
)

type Config struct {
	ListenAddr string   "listen_addr"
	Master     bool     "is_master"
	Minion     bool     "is_minion"
	Masters    []string "masters"
	NodeId     string   "node_id"
}

func main() {
	flag.Parse()
	defer glog.Flush()

	readConfig()
	if glog.V(2) {
		glog.Infof("%#v\n", config)
	}

	if config.Master {
		startMaster()
	}
	time.Sleep(1 * time.Second)

	if config.Minion {
		startMinion()
	}

	time.Sleep(5 * time.Second)
}

func readConfig() {
	file, err := ioutil.ReadFile(CFG_FILE)
	if err != nil {
		glog.Fatal("error reading config file: ", err)
	}
	err = goyaml.Unmarshal(file, &config)
	if err != nil {
		glog.Fatal("error reading config file: ", err)
	}
	if config.Master && config.ListenAddr == "" {
		glog.Fatal("Master should have port defined")
	}
}
