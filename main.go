package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"scrapper/internal/favorites"
	"scrapper/internal/login"
	"syscall"
)

func getUserName() (string, error) {
	fmt.Print("Enter Username: ")
	var username string
	_, err := fmt.Scanln(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

func getPassword() (string, error) {
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	password := string(bytePassword)
	return password, nil
}

func main() {
	client := login.NewClient()
	username, err := getUserName()
	if err != nil {
		log.Fatal("Error reading username:", err)
		return
	}
	password, err := getPassword()
	if err != nil {
		log.Fatal("Error reading password:", err)
		return
	}

	err = client.Login(username, password)
	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}

	// Create new instance of Favorites and pass in the client you've logged in
	favoritesInstance := favorites.Favorites{Client: client}
	listings, err := favoritesInstance.FetchAll()
	if err != nil {
		fmt.Println("Failed to fetch audio equipment:", err)
		return
	}

	fmt.Println("Listings:", listings)
}
