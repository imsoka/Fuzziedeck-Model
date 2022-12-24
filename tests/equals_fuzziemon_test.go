package tests

import (
	"testing"

	"github.com/imsoka/fuzziedeck-model/pkg/models"
);

func TestEquals(t *testing.T) {
    name := "FuzzieTest";
    //name2 := "FuzzieTest2";
    entryDescription := "Este fuzziemon es para testear";
    entryNumber := 69;
    stats := [6]uint8{ 50, 21, 43, 98, 12, 120};


    fuzziemon, _ := models.NewFuzziemon(name, entryDescription, uint(entryNumber), stats);
    fuzziemon2, _ := models.NewFuzziemon(name, entryDescription, uint(entryNumber), stats);

    got := fuzziemon.Equals(fuzziemon2);
    want := true;

    if(got != want) {
        t.Fatalf("Error comparing fuzziemons, got %t, wanted %t", got, want);
    }
}
