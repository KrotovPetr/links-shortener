package config

import "os"

func getCurrentPort() (port string) {
	port = os.Getenv("SERVER_ADDRESS")

	if port == "" {
		port = "8888"
	}

	return
}

func getCurrentHost() (host string) {
	host = os.Getenv("BASE_URL")

	if host == "" {
		host = "http://localhost"
	}

	return
}
