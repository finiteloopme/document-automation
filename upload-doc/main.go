package main

import (
	log "github.com/finiteloopme/goutils/pkg/log"
	os "github.com/finiteloopme/goutils/pkg/os"
)

const (
	DEFAULT_SERVICE_NAME = "document-receiver"
	PORT                 = "8080"
	HOST                 = "0.0.0.0"
)

var hostname, port, serviceName string

func init() { // Get the service name
	serviceName = os.ReadEnvVarOptional("SERVICE_NAME")
	if serviceName == "" {
		// Set the service name to default value if ENV not set
		serviceName = DEFAULT_SERVICE_NAME
	}
	// HTTP port to be used for the service
	port = os.ReadEnvVarOptional("PORT")
	if port == "" {
		port = PORT
		log.Info("\tDefaulting to port: " + port)

	}
	// HOST name to be used to biind the service
	hostname = os.ReadEnvVarOptional("HOST")
	if hostname == "" {
		hostname = HOST
		log.Info("\tDefaulting to hostname: " + hostname)

	}
}

func main() {

	StartServer(hostname, port, serviceName)

}
