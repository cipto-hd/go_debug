package go_debug

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

var Version string = "1.0"

func ErrorHandler() {
	f, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Println("Error open log file: ", err)
	}

	if r := recover(); r != nil {
		fmt.Println("Recovered ", r)

		errorLog := log.New(f, "[Error] ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
		errorLog.Printf("\"%s\"\n %s\n", r, string(debug.Stack()))
	}

	f.Close()
}

/*
Just to print message to standard output
ex: go_debug.Log("yes", "pilihanku nomor ", 1, "amin")
*/

func Log(v ...any) {
	fmt.Println(v...)
}
