package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/vtphan/fmic"
	"math/rand"
	"os"
	"regexp"
	"runtime"
	"time"
)

//-----------------------------------------------------------------------------
func main() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
	re := regexp.MustCompile(`SOURCE_1="[^"]+`)
	genome_index := os.Args[1]
	read := os.Args[2]
	fmt.Println("loading index...")
	idx := fmic.LoadCompressedIndex(genome_index)
	genome_id := make(map[string]int)
	for i := 0; i < len(idx.GENOME_ID); i++ {
		genome_id[idx.GENOME_ID[i]] = i
	}

	f, err := os.Open(read)
	check_for_error(err)
	r := bufio.NewReader(f)
	i := 0
	tp, fp, fn := 0, 0, 0
	var read1, read2 []byte
	var cur_genome int
	var header string

	gid := make(map[string]int)
	for i := 0; i < len(idx.GENOME_ID); i++ {
		gid[idx.GENOME_ID[i]] = i
	}

	fmt.Println("querying reads...")
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		if len(line) > 1 {
			if i%4 == 0 {
				header = string(bytes.TrimSpace(line))
				cur_genome = gid[re.FindString(header)[10:]]
			}

			if i%4 == 1 {
				read1 = bytes.TrimSpace(line)
			} else if i%4 == 3 {
				read2 = bytes.TrimSpace(line)
				// fmt.Println("\n\n", cur_genome, header)
				seqs := idx.FindGenome(read1, reverse_complement(read2), 100, 1500)

				if _, ok := seqs[cur_genome]; ok {
					tp++
					fp += len(seqs) - 1
					// fmt.Println("Positive", seqs)
				} else {
					fn++
					// fmt.Println("\n\n", cur_genome, header)
					// fmt.Println("False Negative", seqs)
				}
			}
		}
		i++
	}
	fmt.Println(tp, fp, fn, float64(tp)/float64(tp+fp), float64(tp)/float64(tp+fn))
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
