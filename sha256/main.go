package main

import(
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"log"
	"crypto/sha256"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	data, err := ioutil.ReadFile("file.txt")
    if err != nil {
        fmt.Errorf("File reading error %w", err)
        return
    }
    fmt.Println("Contents of file:", string(data))

	//--------------
	h := sha256.New()

	_, err = io.Copy(h, f)
	if err != nil {
		log.Fatalln("couldn't io.copy",err)
	}
	fmt.Printf("here is the type before sum: %T\n",h)
	fmt.Printf("%v\n",h)
	xb := h.Sum(nil)
	fmt.Printf("here is the type after sum: %T\n",xb)
	fmt.Printf("%x\n",xb)
	println("---")

	xb = h.Sum(nil)
	fmt.Printf("here is the type after second sum: %T\n",xb)
	fmt.Printf("%x\n",xb)
	println("---")

	xb = h.Sum(xb)
	fmt.Printf("here is the type after third sum and passing in xb: %T\n",xb)
	fmt.Printf("%x\n",xb)
	println("---")
}