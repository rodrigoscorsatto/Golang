// first you need to create the module with: go mod init example/rod
// second you run: go mod tidy (for download the dependencies)

package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
