package models;

type Fuzziemon struct {
    Name string;
    EntryDescription string;
    EntryNumber uint;
    Stats [6]uint8;
}

func NewFuzziemon(name string, entryDescription string, entryNumber uint, stats [6]uint8) *Fuzziemon {
    if (len(name) == 0) {
        return nil;
    }
    if(len(entryDescription) == 0) {
        return nil;
    }
    if(entryNumber <= 0) {
        return nil;
    }

    return &Fuzziemon{
        Name: name,
        EntryDescription: entryDescription,
        EntryNumber: entryNumber,
        Stats: stats,
    }
}
