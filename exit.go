/*
 go run exit.go
 go build exit.go
$ ./exit
$ echo $?
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("!")

	os.Exit(3)
}
