package server

import (
	"testing"
	"time"
)

func TestServer(t *testing.T) {

	server := NewServer("23456")

	go func() {
		time.Sleep(1 * time.Second)
		server.Shutdown()
	}()

	err := server.Listen()
	if err != nil && err.Error() != "http: Server closed" {
		t.Errorf("unexpected error: %s", err.Error())
	}
}
