package main

import(
	"fmt"
	"math/rand"
	"time"
)
func main() {

	rand.Seed(time.Now().Unix())
	//n := rand.Intn(100) + 1
	// fmt.Println(n)

	var count int
	for {
		count++
		n := rand.Intn(100) + 1
		if (n == 100){
			fmt.Println(count)
			break
		}
	}

}