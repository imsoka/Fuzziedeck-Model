package tests

import (
	"testing"

	"github.com/imsoka/fuzziedeck-model/pkg/models"
);

func TestNewFuzziemon(t *testing.T) {
    name := "FuzzieTest";
    entryDescription := "Este fuzziemon es para testear";
    entryNumber := 69;
    stats := [6]uint8{ 50, 21, 43, 98, 12, 120};

    _, err := models.NewFuzziemon(name, entryDescription, uint(entryNumber), stats);

    if(err != nil) {
        t.Fatalf("Error while creating fuzziemon: %s", err.Error())
    }
}
