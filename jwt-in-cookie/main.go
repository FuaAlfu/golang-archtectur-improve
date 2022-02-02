package main

import(
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijavala/jwt-go"
)

type myClaims struct{
	jwt.StandardClaims
	Email string
}

const myKey = "I love you so much..."

func getJWT(msg string)(string, error){
	claims := myClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(), //old: 15000
		},
		Email: msg,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	ss, err := token.SignedString([]byte(myKey))
	if err != nil{
		return "", err.Errorf("couldnt SignedString %w", err)
	}
	return ss, nil
}

func bar(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	ss,err := getJWT(email)
	if err != nil{
		http.Error(w, "couldn't getjwt", http.StatusInternalServerError)
		return
	}

	// *"hash / message digest / digest / hash value | "what we storge"
	c := http.Cookie{
		Name: "session",
		Value: ss,
	}

	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func foo(w http.ResponseWriter, r *http.Request){
	c, err := r.Cookie("session")
	if err != nil{
		c = &http.Cookie{}
	}

	ss := c.Value
	afterVeritificationToken, err := jwt.ParseWithClaims(ss,&myClaims{},func(beforeVeritificationToken *jwt.Token)(interface{},error){ //t *jwt.Token
		if beforeVeritificationToken.Method.Alg() != jwt.SigningMethodH256.Alg(){
			return nil, fmt.Errorf("SOMEONE TRIED TO HACK enchanged signing method")
		}
		return []byte(myKey), nil
	})

	/*
	StandaredClaims has the ..
	Valid() error
	...method witch means it implements the Claims interface...

	type Claims interface{
		Valid() error
	}

	...when you ParseWithClaims ...
	the Valid() method gets run
	...and if all is well, then returns no "error" and
	type TOKEN witch has a field VALID will be true
	*/
	isEqual := err == nil && afterVeritificationToken.Valid

	message := "Not logged in"
	if isEqual{
		message = "Logged in"
		claims := afterVeritificationToken.Claims.(*myClaims)
	    fmt.Println(claims.Email)
	    fmt.Println(claims.ExpiredAt)
	}

	html := `<!DOCTYPE html>
	<html lang="en">
	
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="stylesheet" href="style.css">
		<title>HMAC Example</title>
	</head>
	
	<body>
	     <p>Cookie value: `+ c.Value +`</p>
		 <p>Cookie value: `+ message +`</p>
	     <form action="submit" method="post">
		      <input type="email" name="email" />
			  <input type="submit" />
		 </form>
	</body>
	
	</html>`
	io.WriteString(w,html)
}

func main() {
	http.HandleFunc("/",foo)
	http.HandleFunc("/submit",bar)
	http.ListenAndServe(":8080",nil)
}