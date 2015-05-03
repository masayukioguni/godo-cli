package command

import (
	"reflect"
	"testing"
)

func TestCommond_RegionsSynopsis(t *testing.T) {
	command := &RegionsCommand{
		Cli: nil,
	}

	wantSynopsis := "Show regions"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("RegionsCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
