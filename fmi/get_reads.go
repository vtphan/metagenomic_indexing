package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
)

//-----------------------------------------------------------------------------
func main() {
	f, err := os.Open(os.Args[1])
	check_for_error(err)
	r := bufio.NewReader(f)
	i := 0
	var read1, read2 []byte
	re := regexp.MustCompile(`SOURCE_1="[^"]+`)
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		if len(line) > 1 {
			if i%4 == 0 {
				header := string(bytes.TrimSpace(line))
				fmt.Println("HEADER:", header)
				fmt.Println(re.FindString(header)[10:])
			}
			if i%4 == 1 {
				read1 = bytes.TrimSpace(line)
			} else if i%4 == 3 {
				read2 = reverse_complement(bytes.TrimSpace(line))
				fmt.Println(string(read1))
				fmt.Println(string(read2))
			}
		}
		i++
	}
}

func reverse_complement(s []byte) []byte {
	rs := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			rs[len(s)-i-1] = 'T'
		} else if s[i] == 'T' {
			rs[len(s)-i-1] = 'A'
		} else if s[i] == 'C' {
			rs[len(s)-i-1] = 'G'
		} else if s[i] == 'G' {
			rs[len(s)-i-1] = 'C'
		} else {
			rs[len(s)-i-1] = s[i]
		}
	}
	return rs
}
func check_for_error(e error) {
	if e != nil {
		panic(e)
	}
}
