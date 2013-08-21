package main

import (
	"bufio"
	"github.com/golang/glog"
	"github.com/ugorji/go/codec"
	"net"
)

// create and configure Handle
var (
	mh codec.MsgpackHandle
)

func readFromCon(c net.Conn, ch chan<- *Msg) {
	buf := bufio.NewReader(c)
	for {
		var v = &Msg{}
		h := &mh
		dec := codec.NewDecoder(buf, h)
		err = dec.Decode(v)
		if err != nil {
			glog.Errorln("error decoding", err)
		}
		v.conn = c
		ch <- v
	}

}

func sendMsg(msg Msg) {
	if glog.V(2) {
		glog.Infof("Msg in sendMsg is: %#v\n", msg)
	}
	conn := msg.conn
	h := &mh
	enc := codec.NewEncoder(conn, h)
	err := enc.Encode(msg)
	if err != nil {
		glog.Errorln("error decoding", err)
	}

}
