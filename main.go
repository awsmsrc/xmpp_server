package main

import (
	"net"
	"github.com/awsmsrc/llog"
	"encoding/xml"
	"io"
)

func Handle(conn net.Conn) {
	defer conn.Close()
	for {
		decoder := xml.NewDecoder(conn)
		t, err := decoder.Token() 
		if err != nil {
			if err == io.EOF {
				llog.Success("Client quit")
				return
			}
			llog.Error(err)
		} else {
			llog.Successf("%v", t)
		}
	}
}

func main () {
	listener, _ := net.Listen("tcp", ":6666")
	for {
		llog.Info("Awaiting connection")
		conn, _ := listener.Accept()
		llog.Infof("Got connection from: %+v", conn)
		go Handle(conn)
	}
}
