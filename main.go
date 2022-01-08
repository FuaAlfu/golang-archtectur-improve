package main

import(
	"fmt"
	"encoding/base64"
	"log"
	_"encoding/json"
	_"net/http"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string)([]byte,error){
	bs, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating bcrypt hash from password: %w",err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass[]byte) error{
	err := bcrypt.CompareHashAndPassword(hashedPass,[]byte(password))
	if err != nil{
		return fmt.Errorf("Invalid password: %w",err)
	}
	return nil
}

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))

	pass := "qwerty"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass,hashedPass)
	if err != nil{
		log.Fatalln("Not logged in")
	}
	log.Println("logged in")
}