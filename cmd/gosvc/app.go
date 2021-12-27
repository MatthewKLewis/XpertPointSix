package main

import (
	"fmt"
	"net"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// TODO
// Config based choosing between MQTT and ???
// XpertMessage formatting in publishing
// ACKs?
// Maintenence protocol?
// Configuration protocol?

func yourApp(s server, c Configuration) {
	s.winlog.Info(1, "In Xpert PointSix Parser")
	s.winlog.Info(1, c.Message)

	packet := make([]byte, 65536)
	addr := net.UDPAddr{
		Port: 8557,
		IP:   net.ParseIP(c.MqttServer), //Ethernet on Go Dev Server Side?
	}
	s.winlog.Info(1, "Address set to: "+addr.IP.String())

	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		s.winlog.Info(1, "Error on server set up: "+err.Error())
		panic(err)
	}

	if c.UseMqtt && c.UseKafka {
		s.winlog.Info(1, "Error: Both MQTT and Kafka Selected")
	} else if c.UseMqtt {
		//#region [ rgba(255,100,100,0.1) ] MQTT
		s.winlog.Info(1, "Publishing to MQTT")

		opts := mqtt.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("tcp://%s:%d", c.Broker, c.Port))
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
		//#endregion
	} else if c.UseKafka {
		//#region [ rgba(100,100,255,0.1) ] KAFKA
		s.winlog.Info(1, "Publishing to Kafka")

		//#endregion
	}

}
