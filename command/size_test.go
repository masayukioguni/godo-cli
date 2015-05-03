package command

import (
	"reflect"
	"testing"
)

func TestCommond_SizeSynopsis(t *testing.T) {
	command := &SizeCommand{
		Cli: nil,
	}

	wantSynopsis := "Show available droplet sizes"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("SizesCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
