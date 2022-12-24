package models

import "errors"

type Role struct {
    Id          string
    Name        string
    Permissions []Permission
    Users       []User
}

func NewRole(id string, name string) (*Role, error) {
    if(len(id) == 0) {
        return nil, errors.New("Id cannot be empty");
    }
    if(len(name) == 0) {
        return nil, errors.New("Name cannot be empty");
    }

    return &Role{
        Id: id,
        Name: name,
    }, nil;
}
