package command

import (
	"reflect"
	"testing"
)

func TestCommond_PowerSynopsis(t *testing.T) {
	command := &PowerCommand{
		Client: nil,
	}

	wantSynopsis := "Power off/on/cycle a droplet."
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("PowerCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
