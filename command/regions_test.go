package command

import (
	"reflect"
	"testing"
)

func TestCommond_RegionsSynopsis(t *testing.T) {
	command := &RegionsCommand{
		Client: nil,
	}

	wantSynopsis := "Show available droplet regions"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("RegionsCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
