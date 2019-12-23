package services

import "gamestore/domain/users"

import "gamestore/utils/errors"

//CreateUser service สร้าง้ใช้
func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	return &user, nil

}
