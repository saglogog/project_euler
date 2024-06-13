package main

import (
	"fmt"
)

func main() {

	
	// check project euler solution of prime factorization

	// how I solved: simple factorization, I calculated the least amount of factors by hand 
	// (primes, largest numbers that contain smaller factors) and checked for a very large number

		for i := 0; i < 10000000000000; i++ {
			if (i%11==0) && (i%12==0) && (i%13==0) && (i%14==0) && (i%15==0) && (i%16==0) && (i%17==0) && (i%18==0) && (i%19==0) && (i%20==0) {
				fmt.Println(i)
	
			}
	
		}


	
}