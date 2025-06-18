package main
import (
	"fmt"
	"os"
)
func main() {
    var name string
    var age int
    fmt.Print("Введите имя: ")
    fmt.Fscan(os.Stdin, &name) 
     
    fmt.Print("Введите возраст: ")
    fmt.Fscan(os.Stdin, &age)
     
    fmt.Println(name, age)
}
// func da(x int) {
// 	for i := 1; i <= 5; i++ {
// 		if x <= rand.Intn(3) {
// 			fmt.Println("CHANCE 33.3%")
// 		}
// 	}

// }

// da(rand.Intn(3))
//nino
