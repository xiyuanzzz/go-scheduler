package main

import (
	"fmt"
	"log"
	"os/exec"
)

func callTestCase(filename string) []byte {
	cmd := exec.Command("python", filename)
	//err := cmd.Run()
	//if err != nil {
	//	log.Printf("Call python script error: %v", err)
	//}

	out, err := cmd.Output()
	if err != nil {
		log.Printf("Get script out error: %v", err)
	}
	return out
}

func main() {
	out := callTestCase("hello.py")
	fmt.Println(string(out))
}
