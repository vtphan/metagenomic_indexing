package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Goal: generate a random partition of a given population.
	OUTPUT is []int of size population
	OUTPUT[i] is a random int between 0 and N-1
	sum of all OUTPUT[i] = population
*/
func Partition(population int, N int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(population)
	pop := make([]int, population)
	for i := 0; i < len(perm); i++ {
		pop[perm[i]] = rand.Intn(N)
		if pop[perm[i]] >= population {
			pop[perm[i]] = population
			break
		}
		population -= pop[perm[i]]
	}
	return pop
}

func main() {
	pop := Partition(50, 11)
	sum := 0
	for i := 0; i < len(pop); i++ {
		if pop[i] > 0 {
			fmt.Println(i, pop[i])
			sum += pop[i]
		}
	}
	fmt.Println("Population:", sum)
}
