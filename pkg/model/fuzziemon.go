package model


type Fuzziemon struct {
    Id               string
	Name             string
	EntryDescription string
	EntryNumber      uint
	Stats            [6]uint8
}

func NewFuzziemon(id string, name string, entryDescription string, entryNumber uint, stats [6]uint8) *Fuzziemon {
	return &Fuzziemon {
        Id:               id,
		Name:             name,
		EntryDescription: entryDescription,
		EntryNumber:      entryNumber,
		Stats:            stats,
	}
}
