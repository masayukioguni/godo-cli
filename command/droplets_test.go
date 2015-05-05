package command

import (
	"reflect"
	"testing"
)

func TestCommond_DropletsSynopsis(t *testing.T) {
	command := &DropletsCommand{
		Client: nil,
	}

	wantSynopsis := "Retrieve a list of your droplets."
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("DropletsCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
