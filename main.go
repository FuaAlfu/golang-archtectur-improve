package main

import(
	"fmt"
	"encoding/base64"
	"log"
	"time"
	"io"
	"crypto/rand"
	"crypto/hmac"
	"crypto/sha512"
	_"encoding/json"
	_"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type(
	UserClaims struct{
	   jwt.StandardClaims
	   SessionID int64
	}
    key struct{
		key []byte
		created time.Time
	}
)

//var key []byte{0,1,2,3}

//example model
var keys = map[string]key{} //or var keys = map[string][]byte{}
var currentKid =  ""

func generateNewKey() error{
	//setup currentjob, that run at some point, it is chronologically every hour every day
	//and how to setup that, it's depends on where you got your code deployed..
	newKey := make([]byte, 64)
	_,err := io.ReadFull(rand.Reader,newKey)
	if err != nil{
		return fmt.Errorf("Error in generatorNewKey while generating key %w",err)
	}
	uid,err := uuid.NewV4()
	if err != nil{
		return fmt.Errorf("Error in generatorNewKey while generating kid %w",err)
	}
	//uuid has a string function to get in the canonical string represntation like that canonical
	keys[uid.String()] = key{
		key: newKey,
		created: time.Now(),
	}
	currentKid = uid.String()
	return nil
}

func (u *UserClaims) Valid() error{
	//check if token were expired or not..
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired..")
	}
	if u.SessionID == 0{
		return fmt.Errorf("invalid session ID")
	}

	return nil
}

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

func signMessage(msg []byte)([]byte, error){
	h := hmac.New(sha512.New, keys[currentKid].key)
	_,err := h.Write(msg)
	if err != nil{
		return nil,fmt.Errorf("Error in signMessage while hashing message: %w",err)
	}
	signature := h.Sum(nil)
	return signature,nil
}

func checkSig(msg, sig []byte)(bool,error){
	newSig, err := signMessage(msg)
	if err != nil{
		return false, fmt.Errorf("Error in checkSig while getting signature of message %w",err)
	}
	same := hmac.Equal(newSig,sig)
	return same,nil
}

func createToken(c *UserClaims)(string, error){
	t := NewUserClaims(jwt.SigningMethodHS512,c)
	signedToken, err := t.SignedString(keys[currentKid].key)
	if err != nil{
		return "", fmt.Errorf("Error in createToken while signing token %w",err)
	}
	return signedToken, nil
}

func parseToken(signedToken string)(*UserClaims, error){
	//claims := &UserClaims
	t, err := jwt.ParseWithClaims(signedToken,claims,func(t *jwt.Token)(interface{},error){
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg(){
			return nil, fmt.Errorf("Invalid signing algorithm")
		}
		//k = key, and id
		kid,ok := t.Header["kid"].("string")
		if !ok{
			return nil, fmt.Errorf("Invalid key ID")
		}
		//look for database
		k,ok := key[kid]
		if !ok{
			return nil, fmt.Errorf("Invalid key ID")
		}
		return k.key, nil
	})
	if err != nil{
		return nil, fmt.Errorf("Error in parseToken while parsing token: %w",err)
	}
	if !t.Valid{
		return nil, fmt.Errorf("Error in parseToken, token is not valid")
	}
	//assert, this is actually a type pointer to use claims
	return t.Claims.(*UserClaims), nil
}

func main() {
	/*
	//for hmac example
	for i  := 1; 0 <= 64; i++{
		key = append(key,byte(i))
	}
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
    */
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