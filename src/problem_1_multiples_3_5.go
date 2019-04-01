// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func main() {
    fmt.Println("Searching :")

    const LIMIT = 1000
    var sum uint = 0
    var totalFound uint = 0

    for number := 0; number <= LIMIT; number++ {
        if number % 3 == 0 || number % 5 == 0 {
	     fmt.Println(number)
	     sum += uint(number)
	     totalFound += uint(1)
	}
    }

     fmt.Println("Found", totalFound, "with a sum of", sum)
}

