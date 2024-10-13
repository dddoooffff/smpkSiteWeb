package service

import "errors"

func AddUser(name, email string) func() (string, string, int) {
	id := 0

	return func() (string, string, int) {
		id++
		return name, email, id
	}
}

func EmptyValues(name, email string) (bool, error) {
	if name == "" || email == "" {
		return false, errors.New("Name and email cannot be empty")
	}
	return true, nil
}
