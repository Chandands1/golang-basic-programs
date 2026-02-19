import "fmt"

func reverseNumber(number int){
	reversed := 0
	
	for number != 0{
		digit := number % 10
		reversed = reversed *10 + digit
		number = number /10
	}
	fmt.Println(reversed)
}