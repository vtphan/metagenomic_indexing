package main

import (
	"fmt"
	"github.com/vtphan/fmic"
	"os"
)

//-----------------------------------------------------------------------------
func main() {
	genome := os.Args[1]
	fmt.Println("building index...")
	idx := fmic.CompressedIndex(genome, true, 30)
	fmt.Println("saving index...")
	idx.SaveCompressedIndex(1)
	fmt.Println("done.")
}

func check_for_error(e error) {
	if e != nil {
		panic(e)
	}
}
