package repositories;

import (
    "github.com/imsoka/fuzziedeck-model/pkg/models";
)

type FuzziemonRepositoryInterface interface {
    GetAll() *[]models.Fuzziemon;
    GetById(id string) *models.Fuzziemon;
    GetByEntryNumber(entryNumber uint) *models.Fuzziemon;
    GetByName(name string) *models.Fuzziemon;
    CreateFuzziemon(name string, entryDescription string, entryNumber uint, stats [6]uint8) *models.Fuzziemon;
    UpdateFuzziemon(name string, entryDescription string, entryNumber uint, stats [6]uint8) *models.Fuzziemon;
    DeleteFuzziemon(fuzziemon *models.Fuzziemon) uint;
}
