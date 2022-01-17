package main

import(
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func enDecode(key []byte, input string)([]byte, error){
	b, err := aes.NewCipher(key)
	if err != nil{
		return nil, fmt.Errorf("couldn't newCipher %w",err)
	}

	//initialization vector
	iv := make([]byte, aes.BlockSize)

	//create a cipher
	s := cipher.NewCTR(b, iv)

	buff := &bytes.Buffer{}
	sw := cipher.StreamWriter{
		S: s,
		W: buff,
	}
	_,err = sw.Write([]byte(input))
	if err != nil{
		return nil, fmt.Errorf("couldn't sw.Write to streamwriter %w",err)
	}
	return buff.Bytes(), nil
}

func main() {
	msg := "my massage, sharing my ..."

	password := "ilovecats"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil{
		log.Fatalln("couldn't bcrypt password", err)
	}
	//bs, err := enDecode([]byte{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}) //old but gold
	bs = bs[:16]
	result, err := enDecode(bs,msg)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println("before base64",string(result))

	result2, err := enDecode(bs,string(result))
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(string(result2))
}