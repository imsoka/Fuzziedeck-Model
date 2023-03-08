package model

type Fuzziemon struct {
	Id               string        // A unique identifier for the Fuzziemon
	EntryNumber      uint          // The number assigned to the Fuzziemon in the fuzziedeck
	Name             string        // The name of the Fuzziemon
	EntryDescription string        // A brief description of the Fuzziemon in the fuzziedeck
	Stats            [6]uint8      // The Fuzziemon's attributes.
}

// NewFuzziemon creates a new instance of Fuzziemon with the given values
func NewFuzziemon(id string, entryNumber uint, name string, entryDescription string, stats [6]uint8) *Fuzziemon {
	return &Fuzziemon{
		Id:               id,
		EntryNumber:      entryNumber,
		Name:             name,
		EntryDescription: entryDescription,
		Stats:            stats,
	}
}
