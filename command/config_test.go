package command

import (
	"reflect"
	"testing"
)

func TestCommond_ConfigSynopsis(t *testing.T) {
	command := &ConfigCommand{
		Client: nil,
	}

	wantSynopsis := "Show a droplet's information."
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("ConfigCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
