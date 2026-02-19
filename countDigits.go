import "fmt"

func countDigits(num int) int {
    if num == 0 {
        return 1
    }

    count := 0
    if num < 0 {
        num = -num   // handle negative numbers
    }

    for num != 0 {
        num /= 10
        count++
    }

    return count
}
