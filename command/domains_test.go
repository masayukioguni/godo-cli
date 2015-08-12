package command

import (
    "reflect"
    "testing"
)

func TestCommond_DomainsSynopsis(t *testing.T) {
    command := &DomainsCommand{
        Client: nil,
    }

    wantSynopsis := "Retrieve the list of domains."
    if !reflect.DeepEqual(command.Synopsis(), wantSynopsis) {
        t.Errorf("DomainsCommand.Synopsis returned %+v, want %+v", command.Synopsis(), wantSynopsis)
    }

}
