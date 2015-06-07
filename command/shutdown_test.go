package command

import (
	"reflect"
	"testing"
)

func TestCommond_ShutdownSynopsis(t *testing.T) {
	command := &ShutdownCommand{
		Client: nil,
	}

	wantSynopsis := "Shutdown a droplet."
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("ShutdownCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
