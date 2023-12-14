package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/icrowley/fake"
	"github.com/mewzax/gocolors"
	"github.com/shopspring/decimal"
)

func HelloWorld() string {
	return "Hello world!"
}

func SecondString() string {
	return "This is second line!"
}

func ThirdString() string {
	return "This is third line!"
}
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
func Floor(x float64) float64 {
	return math.Floor(x)
}
func Sin(x float64) float64 {
	result := math.Sin(x)
	return math.Floor(result*1000000000000000) / 1000000000000000
}

func Cos(x float64) float64 {
	result := math.Cos(x)
	return math.Floor(result*1000000000000000) / 1000000000000000
}

func DecimalSum(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", fmt.Errorf("invalid input for a: %v", err)
	}

	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", fmt.Errorf("invalid input for b: %v", err)
	}

	sum := decA.Add(decB)
	return sum.String(), nil
}

func DecimalSubtract(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", fmt.Errorf("invalid input for a: %v", err)
	}

	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", fmt.Errorf("invalid input for b: %v", err)
	}

	sum := decA.Sub(decB)
	return sum.String(), nil
}
func DecimalMultiply(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", fmt.Errorf("invalid input for a: %v", err)
	}

	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", fmt.Errorf("invalid input for a: %v", err)
	}
	multiply := decA.Mul(decB)
	return multiply.String(), nil
}

func DecimalDivide(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", fmt.Errorf("invalid input for a: %v", err)
	}

	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", fmt.Errorf("invalid input for a: %v", err)
	}
	multiply := decA.Div(decB)
	return multiply.String(), nil
}
func ColorizeRed(a string) string {
	return gocolors.Colorize(gocolors.Red, "Hello World!")
}

func GenerateFakeData() string {
	data := fmt.Sprintf("Name: %s\nAddress: %s\nPhone: %s\nEmail: %s", fake.FullName(), fake.StreetAddress(), fake.Phone(), fake.EmailAddress())
	return data
}

func AddPointer(a, b int) *int {
	result := new(int)
	*result = a + b
	return result
}

func MaxPointer(numbers []int) *int {
	maxNumber := new(int)
	*maxNumber = numbers[0]
	for _, number := range numbers {
		if number > *maxNumber {
			*maxNumber = number
		}
	}
	return maxNumber
}

func IsPrimePointer(number int) *bool {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resultNumber := new(int)
	*resultNumber = array[0]
	flag := false
	for _, numberArray := range array {
		if numberArray == number {
			flag = true
		}
	}
	return &flag
}

func ConcatenateStringsPointer(strs []string) *string {
	result := ""
	for _, str := range strs {
		result += str
	}
	return &result
}

func FactorialPonter(n *int) int {
	result := 1
	for i := 1; i <= *n; i++ {
		result *= i
	}
	return result
}

func isPalindrome(str *string) bool {
	s := strings.ToLower(*str)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func CountOccurrences(numbers *[]int, target *int) int {
	count := 0
	for _, num := range *numbers {
		if num == *target {
			count++
		}
	}
	return count
}

func ReverseString(str *string) string {
	s := []rune(*str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(SecondString())
	fmt.Println(ThirdString())
	result := Floor(Sqrt(13))
	fmt.Println(result) // 2

	sinResult := Sin(math.Pi / 2)
	fmt.Println(sinResult) // 1

	//cosResult := Cos(0)
	fmt.Println(math.Floor(math.Pi*1000000000) / 1000000000)

	resultImport, errImport := DecimalSum("10.5", "20.3")
	if errImport != nil {
		fmt.Println("Error:", errImport)
	} else {
		fmt.Println("Sum:", resultImport)
	}

	resultImport1, errImport1 := DecimalSubtract("20.5", "10.3")
	if errImport1 != nil {
		fmt.Println("Error:", errImport1)
	} else {
		fmt.Println("Sum:", resultImport1)
	}

	resultImport2, errImport2 := DecimalMultiply("20.5", "10.3")
	if errImport2 != nil {
		fmt.Println("Error:", errImport2)
	} else {
		fmt.Println("Sum:", resultImport2)
	}

	resultImport3, errImport3 := DecimalDivide("20", "10")
	if errImport3 != nil {
		fmt.Println("Error:", errImport3)
	} else {
		fmt.Println("Sum:", resultImport3)
	}

	resultColor := ColorizeRed("Hello world")
	fmt.Println("color", resultColor)

	resultFakeData := GenerateFakeData()
	fmt.Println(resultFakeData)

	resultAddSumPointer := AddPointer(9, 5)
	fmt.Println(*resultAddSumPointer)

	numbers := []int{2, 7, 4, 9, 1}
	resultMaxPointer := MaxPointer(numbers)
	fmt.Println(*resultMaxPointer)

	resultIsPrimePointer := IsPrimePointer(9)
	fmt.Println(*resultIsPrimePointer)

	strs := []string{"Hello", " ", "123456", " ", "world!"}
	resultConcatenateStringsPointer := ConcatenateStringsPointer(strs)
	fmt.Println(*resultConcatenateStringsPointer)

	n := 5
	fmt.Println(FactorialPonter(&n))

	str := "Racecar"
	fmt.Println(isPalindrome(&str))

	numbersPointer := []int{1, 2, 3, 4, 5, 2, 3, 2}
	target := 2
	fmt.Println(CountOccurrences(&numbersPointer, &target))

	str2 := "Hello, world!"
	fmt.Println(ReverseString(&str2))

	const (
    Monday    = "Понедельник"
    Tuesday   = "Вторник"
    Wednesday = "Среда"
    Thursday  = "Четверг"
    Friday    = "Пятница"
    Saturday  = "Суббота"
    Sunday    = "Воскресенье"
)
}
