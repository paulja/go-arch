package config

import (
	"fmt"
	"os"
)

func GetServePort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return fmt.Sprintf(":%s", port)
}

func GetSeachAddr() string {
	addr := os.Getenv("SEARCH_ADDR")
	if addr == "" {
		return "localhost:4000"
	}
	return addr
}
