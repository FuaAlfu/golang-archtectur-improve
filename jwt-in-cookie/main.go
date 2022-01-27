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
	afterVeritification, err := jwt.ParseWithClaims(ss,&myClaims{},func(beforeVeritification *jwt.Token)(interface{},error){ //t *jwt.Token
		return []byte(myKey), nil
	})

	//isEqual := true

	message := "Not logged in"
	if isEqual{
		message = "Logged in"
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