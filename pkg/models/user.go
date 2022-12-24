package models

import "errors"

type User struct {
    Id          string
	Username    string
	Email       string
	Password    string
    Role        Role
    Permissions []Permission 
	Fuzziemons  []Fuzziemon
}

func NewUser(id string, username string, email string, password string, fuzziemons []Fuzziemon) *User {
	return &User{
		Username:   username,
		Email:      email,
		Password:   password,
		Fuzziemons: fuzziemons,
	}
}

func (user *User) ToString() (string, error) {
    //TODO: Implement

	return "", errors.New("Not yet implemented");
}

func (user *User) Equals(userToCompare *User) bool {
    
    if user.Username == userToCompare.Username {
        return true;
    }
    
    return false;
}
