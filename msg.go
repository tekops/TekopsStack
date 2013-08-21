package main

import (
	"net"
)

type Msg struct {
	JobId   string
	JobType int
	Cmd     string
	OutPut  string
	NodeId  string
	conn    net.Conn
}

const (
	HELLO_JOBTYPE = iota
	CMD_JOBTYPE
	MON_JOB_TYPE
)
