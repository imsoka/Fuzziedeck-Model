package models

import "errors"

type Fuzziemon struct {
    Id               string
	Name             string
	EntryDescription string
	EntryNumber      uint
	Stats            [6]uint8
}

func NewFuzziemon(name string, entryDescription string, entryNumber uint, stats [6]uint8) (*Fuzziemon, error) {
	if len(name) == 0 {
		return nil, errors.New("Name cannot be empty");
	}
	if len(entryDescription) == 0 {
		return nil, errors.New("Entry description cannot be empty");
	}
	if entryNumber <= 0 {
		return nil, errors.New("Entry number cannot be empty");
	}

	return &Fuzziemon{
		Name:             name,
		EntryDescription: entryDescription,
		EntryNumber:      entryNumber,
		Stats:            stats,
	}, nil
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
