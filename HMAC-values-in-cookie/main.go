package main

import(
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
)

func getCode(msg string)string{
	h := hmac.New(sha256.New, []byte("I love you so much..."))
	h.Write([]byte(msg))
	return fmt.Sprintf("%x", h.Sum(nil))
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

	code := getCode(email)

	// *"hash / message digest / digest / hash value | "what we storge"
	c := http.Cookie{
		Name: "session",
		Value: code + "|" + email,
	}

	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func foo(w http.ResponseWriter, r *http.Request){
	c, err := r.Cookie("session")
	if err != nil{
		c = &http.Cookie{}
	}

	isEqual := true
	xs := strings.SplitN(c.Value, "|", 2)
	if len(xs) == 2{
		cCode := xs[0] //c means client
		cEmail := xs[1]

		code := getCode(cEmail)
		isEqual = hmac.Equal(cCode, code)
	}

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