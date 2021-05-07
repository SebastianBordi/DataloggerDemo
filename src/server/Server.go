package server

import (
	"net/http"
	"time"
)

func GetServer(port string) *http.Server {
	server := http.Server{
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return &server
}
