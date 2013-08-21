package main

import (
	"bufio"
	"github.com/golang/glog"
	"github.com/ugorji/go/codec"
	"net"
)

var (
	nodes         = &NodeList{nodes: make(map[string]Node)}
	masterMsgChan = make(chan *Msg, 4064)
)

func startMaster() {
	l, err := net.Listen("tcp", config.ListenAddr)
	if err != nil {
		glog.Fatal(err)
	}
	glog.Infoln("started master! ")
	go listen(l)
	go readMasterMsgChan()
}

func listen(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			glog.Infoln(err)
		}
		go readFromCon(c, masterMsgChan)
		sendToMinion(c, "du -h")
		sendToMinion(c, "df -h")
		sendToMinion(c, "ls /root")

	}
}

func readMasterMsgChan() {
	for msg := range masterMsgChan {
		if glog.V(2) {
			glog.Infof("Got from minion  %#v\n", msg)
		}
		print(msg.OutPut)
	}
}

func sendToMinion(c net.Conn, cmd string) {
	buf := bufio.NewWriter(c)
	j := &mh
	enc := codec.NewEncoder(buf, j)
	msg := Msg{JobId: "Jobid1", Cmd: cmd, JobType: CMD_JOBTYPE, NodeId: config.NodeId, conn: c}
	//err = enc.Encode([]byte("Hello\n"))
	err = enc.Encode(msg)
	if err != nil {
		glog.Errorln("error encoding", err)
	}
	buf.Flush()
	if glog.V(2) {
		glog.Infoln("sent command to minion")
	}
}
