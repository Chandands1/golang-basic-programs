
import "fmt"

func checkPrime(number int){
	if number<= 1 {
        fmt.Println("Not Prime")
        return
    }

    for i := 2; i < number; i++ {
        if number % i == 0 {
            fmt.Println("Not Prime")
            return
        }
    }

    fmt.Println("Prime")
}