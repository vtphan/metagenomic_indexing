
	Usage: go run reduce_dimension.go input_file

Format of each line in input_file: **kmer-id g1 f1 g2 f2** ...

Example: "**19 25 2 1 5**" means k-mer 19 occurs in genome 25, 2 times and in genome 1, 5 times.

Genome ids must between between 0 and M.  It is okay that some ids are missing.  The program will assume there are M+1 genomes, where M is the largest genome id.

Kmer ids must be non-negative.

Part of the output is a Lx(M+1) matrix, where L is the number of kmer groups.




