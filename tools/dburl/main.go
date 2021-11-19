package main

import (
	"fmt"
	"github.com/xo/dburl"
	"os"
)

func main() {
	url := os.Args[1]

	dbUrl, err := dburl.Parse(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(dbUrl.DSN)
}
