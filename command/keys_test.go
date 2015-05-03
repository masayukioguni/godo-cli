package command

import (
	"reflect"
	"testing"
)

func TestCommond_KeysSynopsis(t *testing.T) {
	command := &KeysCommand{
		Client: nil,
	}

	wantSynopsis := "Show available SSH keys"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("KeysCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
