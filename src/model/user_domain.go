package model

import (
	"LearnGo/src/setting/rest_err"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

type IUserDomain interface {
	FindUserUser(string) *rest_err.RestErr
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}

func NewUserDomain(email, password, name string, age int8) IUserDomain {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (u *UserDomain) FindUserUser(email string) *rest_err.RestErr {

	return nil
}

func (u *UserDomain) CreateUser() *rest_err.RestErr {
	u.encryptPassword(u.Password)
	fmt.Println(u.Password)
	return nil
}

func (u *UserDomain) UpdateUser(email string) *rest_err.RestErr {
	return nil
}

func (u *UserDomain) DeleteUser(email string) *rest_err.RestErr {
	return nil
}

func (u *UserDomain) encryptPassword(password string) {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}
