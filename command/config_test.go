package command

import (
	"reflect"
	"testing"
)

func TestCommond_ConfigSynopsis(t *testing.T) {
	command := &ConfigCommand{
		Client: nil,
	}

	wantSynopsis := "configuration of the default setting."
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("ConfigCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
