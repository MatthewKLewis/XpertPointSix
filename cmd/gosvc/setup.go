package main

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/sys/windows/svc/debug"
)

type Configuration struct {
	message string
	working bool
}

// if setup returns an error, the service doesn't start
func setup(wl debug.Log, svcName, sha1ver string) (server, error) {
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
	s.winlog.Info(1, "configs: "+configs.message)

	// configure more logging

	return s, nil
}

func loadConfigs(s server) Configuration {
	s.winlog.Info(1, "in configs")

	var defaultConfigs Configuration
	defaultConfigs.message = "error loading"
	defaultConfigs.working = false

	pathString := "C:\\temp\\conf.json"
	s.winlog.Info(1, pathString)

	file, err1 := os.Open(pathString)
	if err1 != nil {
		s.winlog.Info(1, "failed to load configs")
		return defaultConfigs
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err2 := decoder.Decode(&configuration)
	if err2 != nil {
		s.winlog.Info(1, "failed to decode configs")
		return defaultConfigs
	}
	return configuration
}
