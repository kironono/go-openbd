package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/kironono/go-openbd"
)

func main() {
	ids := flag.String("ids", "", "ISBNs")
	flag.Parse()

	client := openbd.DefaultClient()

	books, err := client.Books(strings.Split(*ids, ","))
	if err != nil {
		log.Panicf("error: %v", err)
	}

	for i, book := range books {
		if i > 0 {
			fmt.Println("====")
		}
		fmt.Printf("ISBN: %s\n", book.Summary.Isbn)
		fmt.Printf("TITLE: %s\n", book.Summary.Title)
	}
}
