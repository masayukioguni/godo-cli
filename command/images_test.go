package command

import (
	"reflect"
	"testing"
)

func TestCommond_ImagesSynopsis(t *testing.T) {
	command := &ImagesCommand{
		Client: nil,
	}

	wantSynopsis := "Images command that are provided in the digitalocean."
	if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
		t.Errorf("ImagesCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
	}

}
