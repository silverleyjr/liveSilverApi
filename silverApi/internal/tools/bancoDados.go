package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	usersByName = map[string]LoginDetails{}
)

func InitDatabase() {
	users := readUsers("usersDetails.txt")
	generateUserMap(users)

	for i, user := range users {
		fmt.Printf("user number %d is %s\n", (i+1), user.Authorization)
	}
}

func ResetUsers() {
	_, err := os.Create("usersDetails.txt")
	if err != nil {
		fmt.Println("erro no resetUsers")
		panic(err)
	}
	silverley := LoginDetails{"Silverley", "bostolao"}
	camilla := LoginDetails{"Camilla", "bostolao"}
	murillo := LoginDetails{"Murillo", "bostolao"}
	NewUser(silverley)
	NewUser(camilla)
	NewUser(murillo)
}

func NewUser(newUser LoginDetails) error {
	users := readUsers("usersDetails.txt")
	fmt.Println(newUser.Authorization)
	if getUserByName(newUser.Authorization).Authorization == "" {
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
func getUserByName(name string) LoginDetails {
	return usersByName[name]
}

func generateUserMap(users []LoginDetails) {
	for _, user := range users {
		usersByName[user.Authorization] = user
	}
}
func readUsers(fileName string) []LoginDetails {
	dataByte, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("erro no read user 1")
		log.Error(err)
		fmt.Println(err)
	}

	var usersFromFile []LoginDetails
	err = json.Unmarshal(dataByte, &usersFromFile)
	if err != nil {
		fmt.Println("erro no read user 2")
		log.Error(err)
		fmt.Println(err)
	}

	return usersFromFile
}
