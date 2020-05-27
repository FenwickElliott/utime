package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	switch len(os.Args) {
	case 1:
		t := time.Now()
		fmt.Println(t, t.Unix())
	case 2:
		utimeAsInt, err := strconv.ParseInt(os.Args[1], 10, 64)
		fatal(err)

		fmt.Println(utimeAsInt, "=>", time.Unix(utimeAsInt, 0))
	default:
		log.Fatal("unexpected number of arguments")
	}
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
