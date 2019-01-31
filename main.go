package main

import (
	"fmt"
	"os"

	"github.com/marcushaddon/anagolang/anagrams"
	"github.com/marcushaddon/anagolang/db"
)

func main() {
	connString := os.Getenv("CLEARDB_DATABASE_URL")
	fmt.Println(connString)
	wr := db.SQLWordRepo{
		ConnString: connString,
	}
	af := anagrams.AnagramFinder{
		WordRepo: wr,
	}

	thing := af.GetAnagrams("dog")
	fmt.Printf("%v", thing)
}
