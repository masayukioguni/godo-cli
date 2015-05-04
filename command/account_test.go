package command

import (
	"reflect"
	"testing"
)

func TestCommond_AccountSynopsis(t *testing.T) {
	command := &AccountCommand{
		Client: nil,
	}

	wantSynopsis := "Show a account information."
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("AccountCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
