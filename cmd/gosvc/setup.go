package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"golang.org/x/sys/windows/svc/debug"
)

type Configuration struct {
	Message string `json: "message"`

	UseMqtt  bool `json: "useMqtt"`
	UseKafka bool `json: "useKafka"`

	MqttServer  string `json: "mqttServer"`
	KafkaServer string `json: "kafkaServer"`

	Broker string `json: "broker"`
	Port   int    `json: "port"`
}

// if setup returns an error, the service doesn't start
func setup(wl debug.Log, svcName, sha1ver string) (server, Configuration, error) {
	var s server

	// did we get a full SHA1?
	if len(sha1ver) == 40 {
		sha1ver = sha1ver[0:7]
	}

	if sha1ver == "" {
		sha1ver = "dev"
	}

	s.winlog = wl

	// Note: any logging here goes to Windows App Log - I suggest you setup local logging
	s.winlog.Info(1, fmt.Sprintf("%s: setup (%s)", svcName, sha1ver))

	// read configuration
	configs := loadConfigs(s)

	return s, configs, nil
}

func loadConfigs(s server) Configuration {
	var defaultConfigs Configuration
	defaultConfigs.Message = "error loading"

	pathString := "C:\\temp\\conf.json"
	file, err3 := ioutil.ReadFile(pathString)
	if err3 != nil {
		s.winlog.Info(1, err3.Error())
		return defaultConfigs
	}
	fileConf := Configuration{}
	_ = json.Unmarshal([]byte(file), &fileConf)
	return fileConf
}
