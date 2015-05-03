package command

import (
	"reflect"
	"testing"
)

func TestCommond_SizesSynopsis(t *testing.T) {
	command := &SizesCommand{
		Cli: nil,
	}

	wantSynopsis := "Show available droplet sizes"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("SizesCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
