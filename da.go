package main
import (
	"fmt"
	"math/rand"
)
func main() {
	da(rand.Intn(3))
	net(rand.Intn(11))
}
func da(x int) {
	for i := 1; i <= 5; i++ {
		if x == rand.Intn(3) {
			fmt.Println("CHANCE 33.3%")
		}
		fmt.Println(i)
	}

}
func net(k int)  {
	if k % 2 == 0{
		fmt.Println("delet", k)
	} else {
		fmt.Println("nedelet", k)
	}
}