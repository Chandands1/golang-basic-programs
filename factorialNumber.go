import "fmt"

func factorial(number int){
    fact := 1
	for i := number ; i > 0 ; i--{
		fact = fact * i
	}
	fmt.Println(fact)
}