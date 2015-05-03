package command

import (
	"reflect"
	"testing"
)

func TestCommond_DropletsSynopsis(t *testing.T) {
	command := &DropletsCommand{
		Client: nil,
	}

	wantSynopsis := "Show available droplet sizes"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("DropletsCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
