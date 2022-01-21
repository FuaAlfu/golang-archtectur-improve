package main

import(
	_"fmt"
	"io"
	"net/http"
)

func getCode(msg string)string{}

func foo(w http.ResponseWriter, r *http.Request){
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
	     <form action="submit" method="post">
		      <input type="email" name="email" />
			  <input type="submit" />
		 </form>
	</body>
	
	</html>`
	io.WriteString(w,html)
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

	c := http.Cookie{
		Name: "session",
		Value: "",
	}
}

func main() {
	http.HandleFunc("/",foo)
	http.HandleFunc("/",bar)
	http.ListenAndServe(":8080",nil)
}