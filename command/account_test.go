package command

import (
	"reflect"
	"testing"
)

func TestCommond_AccountSynopsis(t *testing.T) {
	command := &AccountCommand{
		Client: nil,
	}

	wantSynopsis := "Show available SSH keys"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("AccountCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
