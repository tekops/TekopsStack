package main

import (
	"github.com/golang/glog"
	//"github.com/ugorji/go/codec"
	"net"
	"os/exec"
	"strings"
)

var (
	minionMsgChan = make(chan *Msg, 1000)
)

func startMinion() {
	go readMinionMsgChan()

	for _, server := range config.Masters {
		c, err := net.Dial("tcp", server)
		if err != nil {
			glog.Errorln("error opening connection to server ", err)
		} else {
			go readFromCon(c, minionMsgChan)
			if glog.V(2) {
				glog.Infoln("called readFromCon")
			}
		}

	}

	glog.Infoln("started minion")
}

func readMinionMsgChan() {
	for msg := range minionMsgChan {

		switch msg.JobType {
		case CMD_JOBTYPE:
			if glog.V(2) {
				glog.Infof("Got command %#v\n", msg)
			}

			go executeCommand(msg)
		}
	}
}

func executeCommand(msg *Msg) {

	str := strings.Fields(msg.Cmd)
	if glog.V(2) {
		glog.Infof("str is: %#v\n", str)
	}

	cmdName := str[0]
	aname, err := exec.LookPath(cmdName)
	if err != nil {
		aname = cmdName
	}

	args := str[1:]
	cmd := exec.Cmd{
		Path: aname,
		Args: append([]string{cmdName}, args...),
	}

	if glog.V(2) {
		glog.Infof("Command is: %#v\n", cmd)
	}

	out, err := cmd.Output()
	if err != nil {
		glog.Errorln("error executing ", err)

	} else {
		if glog.V(2) {
			glog.Infoln(string(out))
		}

		returnMsg := Msg{
			JobId:  msg.JobId,
			Cmd:    msg.Cmd,
			OutPut: string(out),
			NodeId: config.NodeId,
			conn:   msg.conn,
		}
		sendMsg(returnMsg)

	}
}
