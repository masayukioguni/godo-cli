package command

import (
	"reflect"
	"testing"
)

func TestCommond_SSHSynopsis(t *testing.T) {
	command := &SSHCommand{
		Client: nil,
	}

	wantSynopsis := "SSH into a droplet."
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("SSHCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
