package main

import(
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	_"crypto/rand"
	"fmt"
	"log"
	"io"

	"golang.org/x/crypto/bcrypt"
)

func encryptWriter(wtr io.Writer, key []byte)(io.Writer, error){
	/*
	this code make it wrapper, a little more abstracted..
	*/
	b, err := aes.NewCipher(key)
	if err != nil{
		return nil, fmt.Errorf("couldn't newCipher %w",err)
	}

	//initialization vector
	iv := make([]byte, aes.BlockSize)

	//create a cipher
	s := cipher.NewCTR(b, iv)

	return cipher.StreamWriter{
		S: s,
		W: wtr,
	},nil

}

func enDecode(key []byte, input string)([]byte, error){
	/*
	this code does it all in memory
	*/
	b, err := aes.NewCipher(key)
	if err != nil{
		return nil, fmt.Errorf("couldn't newCipher %w",err)
	}

	//initialization vector
	iv := make([]byte, aes.BlockSize)
	/*
	//under test :: add more randomness
	func enDecode(key []byte, input string)([]byte, []byte, error //add new arg to this func
	_ err = io.ReadFull(rand.Reader, iv)
	if err != nil{
		return nil, fmt.Errorf("couldn't sw.Write to streamwriter %w",err)
	}*/

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

	//you can put your own writer, it could be buffer, it could be the respone writer, could be some file or whatever
	wtr := &bytes.Buffer{}
	encWriter, err := encryptWriter(wtr,bs)
	if err != nil{
		log.Fatalln(err)
	}

	_,err = io.WriteString(encWriter,msg)
	if err != nil{
		log.Fatalln(err)
	}
	encrypted := wtr.String()
	fmt.Println(encrypted)

	/*
	result, err := enDecode(bs,msg)
	if err != nil{
		log.Fatalln(err)
	}*/
	//fmt.Println("before base64",string(result))
	fmt.Println("before base64",encrypted)

	/*
	result2, err := enDecode(bs,string(result))
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(string(result2))
	*/
	result2, err := enDecode(bs,encrypted)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(string(result2))
}