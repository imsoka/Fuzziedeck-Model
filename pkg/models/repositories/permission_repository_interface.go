package repositories

import "github.com/imsoka/fuzziedeck-model/pkg/models"


type PermissionRepositoryInterface interface {
    GetAll() ([]models.Permission, error);
    GetById(id string) (models.Permission, error);
    GetByUser(role models.Role) ([]models.Permission, error);
    GetByRole(role models.Role) ([]models.Permission, error);
    Create(id string, description string, action string, roles []models.Role, users []models.User) (models.Permission, error);
    Update(id string, description string, action string, roles []models.Role, users []models.User) (models.Permission, error);
    Delete(permission models.Permission) (models.Permission, error);
}
