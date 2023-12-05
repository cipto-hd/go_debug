# Simple Go Debug Logger to a log file

## How to use (sample usage)

1. Create a go project `mkdir my_go; cd $_`
2. Run `go mod init my_go`
3. Run `go get github.com/cipto-hd/go_debug`
4. Create file, ex: `main.go` with below sample content
5. Run `go run main.go`

```go
package main

import (
	"fmt"
	"log"

	"github.com/cipto-hd/go_debug"
)

func main() {

	log.Println("Starting program ...")

	fmt.Println("Return => ", Devide(4, 0)) // will cause error and log it
	fmt.Println("Return => ", Devide(10, 1))

	log.Println("Stopping program ...")

}

func Devide(nominator int, devider int) float32 {
	defer go_debug.ErrorHandler()

	if devider == 0 {
		panic("can't devide by 0")
	}

	return float32(nominator) / float32(devider)
}

```
