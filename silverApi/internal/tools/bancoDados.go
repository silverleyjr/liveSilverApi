package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	usersByName = map[string]User{}
)

func InitDatabase() {
	users := readUsers("usersDetails.txt")
	generateUserMap(users)

	i := 1
	for _, user := range users {
		fmt.Printf("user number %d is %s\n", i, user.Name)
		i++
	}
	//for _, user := range users {
	//	fmt.Printf("My user is %s\n", user.Name)
	//}
}

func ResetUsers() {
	_, err := os.Create("usersDetails.txt")
	if err != nil {
		fmt.Println("erro no resetUsers")
		panic(err)
	}
	silverley := User{"Murillo", "bostolao", "SilverleyFodao"}
	camilla := User{"Murillo", "bostolao", "SilverleyFodao"}
	murillo := User{"Murillo", "bostolao", "SilverleyFodao"}
	NewUser(silverley)
	NewUser(camilla)
	NewUser(murillo)
}

func NewUser(newUser User) error {
	users := readUsers("usersDetails.txt")
	fmt.Println(newUser.Name)
	if getUserByName(newUser.Name).Name == "" {
		users = append(users, newUser)

		usersBytes, err := json.Marshal(users)
		if err != nil {
			fmt.Println("erro no new user")
			log.Error(err)
			return err
		}

		err = os.WriteFile("usersDetails.txt", usersBytes, os.ModePerm)
		if err != nil {
			fmt.Println("erro no new user")
			log.Error(err)
			return err
		}

		return nil
	} else {
		return errors.New("User already exists")
	}
}
func getUserByName(name string) User {
	return usersByName[name]
}

func generateUserMap(users []User) {
	for _, user := range users {
		usersByName[user.Name] = user
	}
}
func readUsers(fileName string) []User {
	dataByte, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("erro no read user 1")
		log.Error(err)
		fmt.Println(err)
	}

	var usersFromFile []User
	err = json.Unmarshal(dataByte, &usersFromFile)
	if err != nil {
		fmt.Println("erro no read user 2")
		log.Error(err)
		fmt.Println(err)
	}

	return usersFromFile
}
