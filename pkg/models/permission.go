package models

import "errors"

type Permission struct {
    Id string
    Description string
    Action string
    Roles []Role
    Users []User
}

func NewPermission(id string, description string, action string) (*Permission, error) {
    if(len(id) == 0) {
        return nil, errors.New("Id cannot be empty");
    }
    if(len(description) == 0) {
        return nil, errors.New("Description cannot be empty");
    }
    if(len(action) == 0) {
        return nil, errors.New("Action cannot be empty");
    }

    return &Permission{
        Id: id,
        Action: action,
    }, nil;
}
