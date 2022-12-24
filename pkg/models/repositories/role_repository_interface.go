package repositories

import "github.com/imsoka/fuzziedeck-model/pkg/models"


type RoleRepositoryInterface interface {
    GetAll() ([]models.Role, error)
    GetBiId(id string) (models.Role, error);
    GetByPermission(permission models.Permission) (models.Role, error);
    GetByUser(user models.User) ([]models.Role, error);
    Create(id string, name string, permissions []models.Permission, users []models.User) (models.Role, error);
    Update(id string, name string, permissions []models.Permission, users []models.User) (models.Role, error);
    Delete(role models.Role) (models.Role, error);
}
