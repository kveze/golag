package main
import (
	"fmt"
	"math/rand"
)
func main() {
	da(rand.Intn(3))
}
func da(x int) {
	for i := 1; i <= 5; i++ {
		if x <= rand.Intn(3) {
			fmt.Println("CHANCE 33.3%")
		}
	}

}
