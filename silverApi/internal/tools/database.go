package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Authorization string
	Password  string
}
//type PasswordDetails struct {
//	Username string
//	Password string
//}

//type User struct {
//	Name          string
//	Password      string
//	Authorization string
//}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	//GetUserPassword(username string) *PasswordDetails
	//GetAllUserDetails(username string) *User
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}

//type LoginDetails struct {
//	AuthToken string
//	Username  string
//}
//type PasswordDetails struct {
//	Username string
//	Password string
//}
//
//type AllUserDetails struct {
//	Username string
//	Password string
//	AuthToken string
//}
//
//type DatabaseInterface interface {
//	GetUserLoginDetails(username string) *LoginDetails
//	GetUserPassword(username string) *PasswordDetails
//	SetupDatabase() error
//}
//
//func NewDatabase() (*DatabaseInterface, error) {
//
//	var database DatabaseInterface = &mockDB{}
//
//	var err error = database.SetupDatabase()
//	if err != nil {
//		log.Error(err)
//		return nil, err
//	}
//
//	return &database, nil
//}
