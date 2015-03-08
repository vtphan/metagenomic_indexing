package main
import (
   "fmt"
   "bufio"
   "os"
   "strings"
   "strconv"
   "sort"
   "time"
)

const Dim = 5
type vector [Dim]int

type hashEntry struct {
   num_gids int
   kmers []int
   freqs []vector
}

type hash map[vector]*hashEntry

type kmerEntry struct {
   count int
   gids vector
   freqs vector
}

func (e *kmerEntry) Len() int { return e.count }
func (e *kmerEntry) Swap(i, j int) {
   e.gids[i], e.gids[j] = e.gids[j], e.gids[i]
   e.freqs[i], e.freqs[j] = e.freqs[j], e.freqs[i]
}
func (e *kmerEntry) Less(i, j int) bool { return e.gids[i] < e.gids[j] }

func ReadMatrix(input_file string, T hash) int {
   f, err := os.Open(input_file)
   if err != nil {
      panic("Error opening " + input_file)
   }
   defer f.Close()
   var line string
   var tokens []string
   var kmer, gid, freq int
   max_gid := 0
   scanner := bufio.NewScanner(f)
   for scanner.Scan() {
      line = scanner.Text()
      if line[0] != '#' {   // ignore comments, which start with #
         tokens = strings.SplitN(line," ",2)
         kmer, err = strconv.Atoi(tokens[0])
         if err != nil {
            panic("Trouble, trouble!!!")
         }
         tokens = strings.Split(tokens[1]," ")

         // Get the entry for this kmer
         entry := new(kmerEntry)
         for i:=0; i<len(tokens); i+=2 {
            gid, err = strconv.Atoi(tokens[i])
            if err != nil {
               panic("Trouble, trouble!!!")
            }
            freq, err = strconv.Atoi(tokens[i+1])
            if err != nil {
               panic("Trouble, trouble!!!")
            }
            entry.gids[i/2] = gid
            entry.freqs[i/2] = freq
            entry.count++

            if gid > max_gid {
               max_gid = gid
            }
         }

         // Make sure that genome ids in the k-mer's entry are sorted.
         sort.Sort(entry)

         _, exist := T[entry.gids]
         if ! exist {
            T[entry.gids] = new(hashEntry)
            T[entry.gids].num_gids = entry.count
         }
         T[entry.gids].kmers = append(T[entry.gids].kmers, kmer)
         T[entry.gids].freqs = append(T[entry.gids].freqs, entry.freqs)
      }
   }
   return max_gid
}

func ReduceMatrix(T hash, max_gid int) {
   h,m,s := time.Now().Hour(), time.Now().Minute(), time.Now().Second()
   row_filename := fmt.Sprintf("%s_row_%d_%d_%d", os.Args[1], h, m, s)
   matrix_filename := fmt.Sprintf("%s_mat_%d_%d_%d", os.Args[1], h, m, s)
   row_file, _ := os.Create(row_filename)
   defer row_file.Close()
   matrix_file, _ := os.Create(matrix_filename)
   defer matrix_file.Close()

   row_writer := bufio.NewWriter(row_file)
   matrix_writer := bufio.NewWriter(matrix_file)

   var sum_freq int
   for gids, e := range(T) {
      // save to file ColumnID
      for i:=0; i<len(e.kmers); i++ {
         fmt.Fprint(row_writer, e.kmers[i])
         if i<len(e.kmers)-1 {
            fmt.Fprint(row_writer, ",")
         } else {
            fmt.Fprint(row_writer, "\n")
         }
      }

      // save to file Matrix
      for i,j:=0,0; i<=max_gid; i++ {
         if i!=gids[j] || j>=e.num_gids{
            fmt.Fprint(matrix_writer, 0)
         } else if i==gids[j] {
            sum_freq = 0
            for _, f := range(e.freqs) {
               sum_freq += f[j]
            }
            j++
            fmt.Fprint(matrix_writer, sum_freq)
         }
         if i<max_gid {
            fmt.Fprint(matrix_writer, " ")
         } else {
            fmt.Fprint(matrix_writer, "\n")
         }
      }
   }
   row_writer.Flush()
   matrix_writer.Flush()
   fmt.Println(len(T), "kmer groups (rows) saved to file: ", row_filename)
   fmt.Printf("%d-by-%d matrix saved to file: %s\n", len(T), max_gid+1, matrix_filename)
   fmt.Printf("RowID and Matrix files must have the same timestamp: %d_%d_%d\n", h,m,s)
}

func main() {
   if len(os.Args) != 2 {
      fmt.Println("\n\tUsage: go run reduce_dimension.go input_file\n")
      fmt.Println("Format of each line in input_file: kmer-id g1-id f1 g2-id f2 ...")
      fmt.Println("Example: 19 25 2 1 5 means k-mer 19 occurs in genome 25, 2 times and in genome 1, 5 times.\n")
      os.Exit(0)
   }
   T := make(hash)
   max_gid := ReadMatrix(os.Args[1], T)

   // for k, e := range(T) {
   //    fmt.Println(k,e)
   // }

   ReduceMatrix(T, max_gid)
}
