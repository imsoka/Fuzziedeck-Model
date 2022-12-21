package repositories;

import (
    "github.com/imsoka/fuzziedeck-model/pkg/models";
)

type FuzziemonRepository interface {
    CreateFuzziemon(name string, entryDescription string, entryNumber uint, stats [6]uint8) *models.Fuzziemon;
    FindByEntryNumber(entryNumber uint) *models.Fuzziemon;
    FindByEntryName(entryNumber uint) *models.Fuzziemon;
    UpdateFuzziemon(name string, entryDescription string, entryNumber uint, stats [6]uint8) *models.Fuzziemon;
    DeleteFuzziemon(fuzziemon *models.Fuzziemon) uint;
}
