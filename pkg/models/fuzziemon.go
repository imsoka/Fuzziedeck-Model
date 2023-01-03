package models

import "errors"

type Fuzziemon struct {
    Id               string
	Name             string
	EntryDescription string
	EntryNumber      uint
	Stats            [6]uint8
}

func NewFuzziemon(name string, entryDescription string, entryNumber uint, stats [6]uint8) *Fuzziemon {
	return &Fuzziemon{
		Name:             name,
		EntryDescription: entryDescription,
		EntryNumber:      entryNumber,
		Stats:            stats,
	}
}

func (fuzziemon *Fuzziemon) ToString() (string, error) {
	//TODO: Implement

	return "", errors.New("Not yet implemented");
}

func (fuzziemon *Fuzziemon) Equals(fuzziemonToCompare *Fuzziemon) bool {
    
    if fuzziemon.Name == fuzziemonToCompare.Name {
        return true;
    }
    
    return false;
}
