package repositories;

import (
    "github.com/imsoka/fuzziedeck-model/pkg/model";
)

type FuzziemonRepositoryInterface interface {
    GetAll() (*[]model.Fuzziemon, error);
    GetByEntryNumber(entryNumber uint) (*model.Fuzziemon, error);
    GetByName(name string) (*model.Fuzziemon, error);
    Create(name string, entryDescription string, entryNumber uint, stats [6]uint8) (*model.Fuzziemon, error);
    Update(name string, entryDescription string, entryNumber uint, stats [6]uint8) (*model.Fuzziemon, error);
    Delete(fuzziemon *model.Fuzziemon) (uint, error);
}
