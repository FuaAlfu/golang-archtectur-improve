package main

import(
	"fmt"
	"log"
	"encoding/json"
	"net/http"
)

type person struct{
	Name string
}

func handlePage(){
	http.HandleFunc("/encode",foo)
	http.HandleFunc("/decode",bar)
	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter, r *http.Request){

}

func bar(w http.ResponseWriter, r *http.Request){

}

func main() {
	handlePage()

	//constructing
	p1 := person{
		Name: "john",
	}
	p2 := person{
		Name: "doe",
	}

	bt := []person{p1, p2}

  //marshal
	bs, err := json.Marshal(bt)
	if err != nil{
		log.Panic(err)
	}
	fmt.Println("printing json ",string(bs))

  //unmarshal
   xp := []person{}
   err = json.Unmarshal(bs,&xp)
   if err != nil{
	   log.Panic(err)
   }
   fmt.Println("back into a Go data structure ",xp)
}