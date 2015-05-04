package command

import (
	"reflect"
	"testing"
)

func TestCommond_CreateSynopsis(t *testing.T) {
	command := &CreateCommand{
		Client: nil,
	}

	wantSynopsis := "Create a droplet."
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("CreateCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
