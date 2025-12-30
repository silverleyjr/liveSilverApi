package tools

import (
	"fmt"
)

type mockDB struct{}

func (d *mockDB) GetAllUserDetails(username string) *User {
	var clientData = User{}
	clientData = getUserByName(username)
	fmt.Println(clientData)
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	InitDatabase()
	return nil
}

//type mockDB struct{}
//
//var mockLoginDetails = map[string]LoginDetails{
//	"Silverley": {
//		AuthToken: "SilverleyFodao",
//		Username:  "Silverley",
//	},
//	"Camilla": {
//		AuthToken: "SilverleyFodao",
//		Username:  "Camilla",
//	},
//	"Murillo": {
//		AuthToken: "SilverleyFodao",
//		Username:  "Murillo",
//	},
//}
//
//var mockPasswordDetails = map[string]PasswordDetails{
//	"Silverley": {
//		Username: "Silverley",
//		Password: "junior",
//	},
//	"Camilla": {
//		Username: "Camilla",
//		Password: "bostola",
//	},
//	"Murillo": {
//		Username: "Murillo",
//		Password: "chupamanga",
//	},
//}
//
//func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
//	time.Sleep(time.Second * 1)
//	var clientData = LoginDetails{}
//	clientData, ok := mockLoginDetails[username]
//	if !ok {
//		return nil
//	}
//
//	return &clientData
//}
//
//func (d *mockDB) GetUserPassword(username string) *PasswordDetails {
//	time.Sleep(time.Second * 1)
//	var clientData = PasswordDetails{}
//	clientData, ok := mockPasswordDetails[username]
//	if !ok {
//		return nil
//	}
//
//	return &clientData
//}
//
//func (d *mockDB) SetupDatabase() error {
//	return nil
//}
