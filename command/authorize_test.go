package command

import (
	"reflect"
	"testing"
)

func TestCommond_Authorize(t *testing.T) {
	command := &AuthorizeCommand{}

	wantSynopsis := "Authorize a DigitalOcean account with godo-cli"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("AuthorizeCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

	if !reflect.DeepEqual(command.Help(), "") {
		t.Errorf("AuthorizeCommand.Help returned %+v, want %+v", command.Help(), "")
	}
}
