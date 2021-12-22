package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Configuration struct {
	message string
}

const (
	UDP_LISTEN_CONNECTION = "192.168.3.157" // Matthew's Laptop
	broker                = "localhost"
	port                  = 1883
)

// TODO
// Config based choosing between MQTT and ???
// XpertMessage formatting in publishing
// ACKs?
// Maintenence protocol?
// Configuration protocol?

func yourApp(s server) {
	s.winlog.Info(1, "In Xpert PointSix Parse")

	//load configs
	var configuration Configuration = loadConfigs()
	s.winlog.Info(1, configuration.message)

	packet := make([]byte, 65536)
	addr := net.UDPAddr{
		Port: 8557,
		IP:   net.ParseIP(UDP_LISTEN_CONNECTION), //Ethernet on Go Dev Server Side?
	}
	s.winlog.Info(1, "Address set to: "+addr.IP.String())

	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		s.winlog.Info(1, "Error on server set up: "+err.Error())
		panic(err)
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		s.winlog.Info(1, "Error connecting to mqtt")
	}

	for {
		_, remoteaddr, err := server.ReadFromUDP(packet) //is this blocking waiting for a UDP message to come in?
		if err != nil {
			s.winlog.Info(1, "Error on UDP read: "+err.Error()+remoteaddr.Network())
		}
		go client.Publish("topic/test", 0, false, parse(packet))
	}
}

func loadConfigs() Configuration {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
