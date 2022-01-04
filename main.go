package main

import(
	"fmt"
	"log"
	"encoding/json"
)

type person struct{
	Name string
}

func main() {
	
	//constructing
	p1 := person{
		Name: "john",
	}
	p2 := person{
		Name: "doe",
	}

	bt := []person{p1, p2}

	bs, err := json.Marshal(bt)
	if err != nil{
		log.Panic(err)
	}
	fmt.Println(string(bs))
}