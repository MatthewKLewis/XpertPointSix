package main

import (
	"net"
)

const (
	UDP_LISTEN_CONNECTION = "192.168.3.157" // Matthew's Laptop
)

// The wrapper of your app
func yourApp(s server) {
	s.winlog.Info(1, "In Xpert PointSix Parse")

	packet := make([]byte, 65536)
	addr := net.UDPAddr{
		Port: 8552,
		IP:   net.ParseIP(UDP_LISTEN_CONNECTION), //Ethernet on Go Dev Server Side?
	}
	s.winlog.Info(1, "Address set to: "+addr.IP.String())

	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		s.winlog.Info(1, "Error on server set up: "+err.Error())
		panic(err)
	}

	s.winlog.Info(1, "Server string: "+server.LocalAddr().Network())

	for {
		//Read packets on the UDP connection
		_, remoteaddr, err := server.ReadFromUDP(packet) //is this 'blocking' waiting for a UDP message to come in?
		if err != nil {
			s.winlog.Info(1, "Error on UDP read: "+err.Error()+remoteaddr.Network())
		}
		s.winlog.Info(1, string(packet))
	}

	// Notice that if this exits, the service continues to run - you can launch web servers, etc.
}
