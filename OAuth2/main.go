package main

import(
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"
	"github.org/x/oauth2"
	"github.org/x/oauth2/github"
)

var githubLoginAttempts = make(map[string])

var githubOauthConfig = &oauth2.Config{
	ClientID: godotenv.Load(".env"),
	ClientSecret: godotenv.Load(".env"),
	Endpoint: github.Endpoint,
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
  }

func testPage(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, r.URL.Path[1:])
}

func index(w http.ResponseWriter, r *http.Request){
     fmt.Fprintln(w, `<!DOCTYPE html>
	 <html lang="en">
	 
	 <head>
		 <meta charset="UTF-8">
		 <meta http-equiv="X-UA-Compatible" content="IE=edge">
		 <meta name="viewport" content="width=device-width, initial-scale=1.0">
		 <link rel="stylesheet" href="style.css">
		 <title>my page</title>
	 </head>
	 
	 <body>
		 <form class="form" action="/ouath/github" method="post">
			 <input type="submit" value="login with github">
		 </form>
	 </body>

	 <style>
	      *{
			  box-border: box-sizing;
			  background-color: #ccc;
		  }
		  .form{
			background-color: #eee;
		  }
	 </style>
	 
	 </html>`)
}

func startGithubOauth(w http.ResponseWriter, r *http.Request){
	redirectURL := githubOauthConfig.AuthCodeURL("0000") //need to create database with uniq id
	http.Redirect(w,r, redirectURL, http.StatusSeeOther)
}

func complateGithubOauth(w http.ResponseWriter, r *http.Request){
	code := r.FormValue("code")
	state := r.FormValue("state")

	if state != "0000" {
		http.Error(w, "State is incorrect", http.StatusBadRequest)
		return
	}
	token, err :=  githubOauthConfig.Exchange(r.Context(), code)
	if err != nil{
		http.Error(w, "Couldn't login", http.StatusInternalServerError)
		return	
	}
	ts := githubOauthConfig.TokenSource(r.Context(), token)
	client :=  outh2.NewClient(r.Context(), ts)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "HomePage")
}
func server(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/login",index)
	http.HandleFunc("/oauth/github",startGithubOauth)
	http.HandleFunc("/oauth2/receive",complateGithubOauth)
	http.ListenAndServe(":8080", nil)
}

func main() {
	// godotenv package
	//dotenv := goDotEnvVariable(" ")

	server()
}