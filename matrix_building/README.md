
	Usage: go run reduce_dimension.go input_file

Format of each line in input_file: **kmer-id g1 f1 g2 f2** ...

Example: "**19 25 2 1 5**" means k-mer 19 occurs in genome 25, 2 times and in genome 1, 5 times.

Genome ids must be in [0,M].  The input might miss some genome ids.  The program will assume there are M+1 genomes, where M is the largest genome id.

Kmer ids must be non-negative.  Each kmer cannot appear in more than one line.  The input is invalid if the first numbers (kmer ids) of any two lines are the same.

Part of the output is a Lx(M+1) matrix, where L is the number of kmer groups.




