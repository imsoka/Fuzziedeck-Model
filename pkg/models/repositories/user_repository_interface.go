package repositories;

import (
    "github.com/imsoka/fuzziedeck-model/pkg/models"
);

type UserRepositoryInterface interface {
    GetAll() *[]models.User;
    GetById(id string) (*models.User, error);
    GetByUsername(username string) (*models.User, error);
    GetByEmail(email string) (*models.User, error);
    CreateUser(id string, username string, email string, password string, fuzziemons []models.Fuzziemon) (string, error);
    UpdateUser(id string, username string, email string, password string, fuzziemons []models.Fuzziemon) (*models.User, error);
    DeleteUser(user *models.User) error;
}
