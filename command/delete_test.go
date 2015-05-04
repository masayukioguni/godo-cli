package command

import (
	"reflect"
	"testing"
)

func TestCommond_DeleteSynopsis(t *testing.T) {
	command := &DeleteCommand{
		Client: nil,
	}

	wantSynopsis := "Destroy a droplet."
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("DeleteCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
