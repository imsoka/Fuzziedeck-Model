package repositories;

import (
    "github.com/imsoka/fuzziedeck-model/pkg/models";
)

type FuzziemonRepositoryInterface interface {
    GetAll() (*[]models.Fuzziemon, error);
    GetByEntryNumber(entryNumber uint) (*models.Fuzziemon, error);
    GetByName(name string) (*models.Fuzziemon, error);
    Create(name string, entryDescription string, entryNumber uint, stats [6]uint8) (*models.Fuzziemon, error);
    Update(name string, entryDescription string, entryNumber uint, stats [6]uint8) (*models.Fuzziemon, error);
    Delete(fuzziemon *models.Fuzziemon) (uint, error);
}
