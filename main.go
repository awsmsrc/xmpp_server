package main

import (
	"net"
	"github.com/lukeroberts1990/llog"
	"time"
	"fmt"
)

func HeartBeat(conn net.Conn) {
	defer conn.Close()
	hb := time.Tick(time.Duration(1) * time.Second)
	for {
		<- hb
		fmt.Fprint(conn, "Heartbeat\n")
	}
}

func main () {
	listener, _ := net.Listen("tcp", ":6666")
	for {
		llog.Info("Awaiting connection")
		conn, _ := listener.Accept()
		llog.Infof("Got connection from: %+v", conn)
		go HeartBeat(conn)
	}
}
