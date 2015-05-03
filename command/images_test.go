package command

import (
	"reflect"
	"testing"
)

func TestCommond_ImagesSynopsis(t *testing.T) {
	command := &ImagesCommand{
		Cli: nil,
	}

	wantSynopsis := "Retrieve a list of your images"
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("ImagesCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
