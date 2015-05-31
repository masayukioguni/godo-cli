package command

import (
	"reflect"
	"testing"
)

func TestCommond_SnapshotSynopsis(t *testing.T) {
	command := &SnapshotCommand{
		Client: nil,
	}

	wantSynopsis := "Create a snapshot of a droplet."
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("SnapshotCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
