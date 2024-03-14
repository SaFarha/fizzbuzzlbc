package application

import (
	"testing"
	"time"
)

func TestApplication(t *testing.T) {

	app := NewApplication()

	go func() {
		time.Sleep(1 * time.Second)
		app.ShutDown()
	}()

	err := app.Run()
	if err != nil && err.Error() != "http: Server closed" {
		t.Errorf("unexpected error: %s", err.Error())
	}

}
