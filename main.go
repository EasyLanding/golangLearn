package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"

	"mymath/mymath"

	"encoding/json"

	testmymath "github.com/EasyLanding/testMyMath"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/icrowley/fake"
	"github.com/mattevans/dinero"
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

func WriteInfo() {
	var name string
	var age int
	var city string

	fmt.Print("Введите ваше имя: ")
	fmt.Scanln(&name)

	fmt.Print("Введите ваш возраст: ")
	fmt.Scanln(&age)

	fmt.Print("Введите ваш город: ")
	fmt.Scanln(&city)

	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Город:", city)
}

func calculate(a, b float64, operation string) (float64, error) {
	var result float64
	var err error

	switch operation {
	case "difference":
		result = a - b
	case "sum":
		result = a + b
	case "product":
		result = a * b
	case "quotient":
		if b == 0 {
			err = errors.New("division by zero")
		}
	default:
		err = errors.New("unknown operation")
	}

	return result, err
}

func hypotenuse(a, b *float64) float64 {
	return math.Sqrt(*a**a + *b**b)
}

func currencyPairRate(from string, to string, amount float64) (float64, error) {
	client := dinero.NewClient(
		"827dec1f583440849866926eb3f16106",
		from,
		20*time.Minute,
	)
	rate, err := client.Rates.Get(to)
	if err != nil {
		return 0, err
	}
	return amount * (*rate), nil
}

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	roundFactor := math.Pow(10, float64(decimalPlaces))
	aRounded := math.Round(a*roundFactor) / roundFactor
	bRounded := math.Round(b*roundFactor) / roundFactor
	isEqual = aRounded == bRounded
	difference = math.Abs(aRounded - bRounded)
	return isEqual, difference
}

func binaryStringToFloat(binary string) float32 {
	var number uint32
	// Преобразование строки в двоичной системе в целочисленное представление
	for _, digit := range binary {
		number <<= 1
		if digit == '1' {
			number |= 1
		}
	}
	// Преобразование целочисленного представления в число с плавающей точкой
	floatNumber := *(*float32)(unsafe.Pointer(&number))
	return floatNumber
}

func sizeOf(v interface{}) (int, error) {
	switch v.(type) {
	case bool:
		return int(unsafe.Sizeof(v.(bool))), nil
	case int:
		return int(unsafe.Sizeof(v.(int))), nil
	case int8:
		return int(unsafe.Sizeof(v.(int8))), nil
	case int16:
		return int(unsafe.Sizeof(v.(int16))), nil
	case int32:
		return int(unsafe.Sizeof(v.(int32))), nil
	case int64:
		return int(unsafe.Sizeof(v.(int64))), nil
	case uint:
		return int(unsafe.Sizeof(v.(uint))), nil
	case uint8:
		return int(unsafe.Sizeof(v.(uint8))), nil
	default:
		return 0, errors.New("type is not valid")
	}
}

func UserInfo(name, city, phone string, age, weight int) string {
	return fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s, Возраст: %d, Вес: %d", name, city, phone, age, weight)
}

func UserInfoMainCityes(name string, age int, cities ...string) string {
	result := "Имя: " + name + ", возраст: " + strconv.Itoa(age) + ", города: "
	for i, city := range cities {
		if i == len(cities)-1 {
			result += city
		} else {
			result += city + ", "
		}
	}

	return result
}

func ConcatenateStrings(sep string, str ...string) string {
	even := ""
	odd := ""
	for i, s := range str {
		if i != 0 {
			even += s + sep
		} else {
			odd += s + sep
		}
	}
	even = strings.TrimSuffix(even, sep)
	odd = strings.TrimSuffix(odd, sep)
	return fmt.Sprintf("even: %s, odd: %s,", even, odd)
}

func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {
	initial, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		return 0, errors.New("initial value is not a number")
	}

	final, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0, errors.New("final value is not a number")
	}

	if initial == 0 {
		return 0, errors.New("initial value cannot be zero")
	}

	percentageChange := (final - initial) / initial * 100

	if math.IsInf(percentageChange, 1) || math.IsInf(percentageChange, -1) {
		return 0, errors.New("division by zero")
	}

	return percentageChange, nil
}

func SumAn(a ...int) int {
	sum := 0
	for _, num := range a {
		sum += num
	}
	return sum
}

func MulAn(a ...int) int {
	mul := 1
	for _, num := range a {
		mul *= num
	}
	return mul
}

func SubAn(a ...int) int {
	sub := a[0]
	for i := 1; i < len(a); i++ {
		sub -= a[i]
	}
	return sub
}

func MathOperate(op func(a ...int) int, a ...int) int {
	return op(a...)
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func multiplier(factor float64) func(float64) float64 {
	return func(num float64) float64 {
		return num * factor
	}
}

func adder(initial int) func(int) int {
	return func(num int) int {
		return num + initial
	}
}

//Array learning

func sumArray(xs [8]int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}

func averageArray(xs [8]int) float64 {
	total := sumArray(xs)
	return float64(total) / float64(len(xs))
}

func averageFloatArray(xs [8]float64) float64 {
	total := 0.0
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func reverseArray(xs [8]int) [8]int {
	reversed := [8]int{}
	for i, j := 0, len(xs)-1; i < len(xs); i, j = i+1, j-1 {
		reversed[i] = xs[j]
	}
	return reversed
}

func sortDesc(nums interface{}) {
	switch nums.(type) {
	case []int:
		sort.Sort(sort.Reverse(sort.IntSlice(nums.([]int))))
	case []float64:
		sort.Sort(sort.Reverse(sort.Float64Slice(nums.([]float64))))
	}
}

func sortAsc(nums interface{}) {
	switch nums.(type) {
	case []int:
		sort.Sort(sort.IntSlice(nums.([]int)))
	case []float64:
		sort.Sort(sort.Float64Slice(nums.([]float64)))
	}
}

func fibonacciFormula(n int) int {
	phi := (1 + math.Sqrt(5)) / 2
	return int(math.Round(math.Pow(phi, float64(n)) / math.Sqrt(5)))
}

func getSubSlice(xs []int, start, end int) []int {
	subslice := xs[start:end]
	return subslice
}

func MaxDifference(numbers []int) int {
	if len(numbers) < 2 {
		return 0
	}

	min := numbers[0]
	maxDiff := 0

	for _, num := range numbers {
		if num < min {
			min = num
		}
		diff := num - min
		if diff > maxDiff {
			maxDiff = diff
		}
	}

	return maxDiff
}

func bitwiseXOR(n, res int) int {
	return n ^ res
}

func findSingleNumber(numbers []int) int {
	res := 0
	for _, num := range numbers {
		res = bitwiseXOR(num, res)
	}
	return res
}

func appendInt(xs []int, x ...int) []int {
	result := append(xs, x...)
	return result
}

func appendInt2(xs *[]int, x ...int) {
	*xs = append(*xs, x...)
}

func Cut(xs []int, start, end int) []int {
	result := xs[start:end]
	return result
}

type User struct {
	Name string
	Age  int
}

func getUsers() []User {
	var users []User
	for i := 0; i < 10; i++ {
		user := User{Name: gofakeit.Name(), Age: gofakeit.Number(18, 60)}
		users = append(users, user)
	}
	return users
}

func preparePrint(users []User) string {
	var result string
	for _, user := range users {
		result += fmt.Sprintf("Имя: %s, Возраст: %d\n", user.Name, user.Age)
	}
	return result
}

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {
	var animals []Animal
	for i := 0; i < 3; i++ {
		animal := Animal{Type: gofakeit.Animal(), Name: gofakeit.Name(), Age: gofakeit.Number(18, 60)}
		animals = append(animals, animal)
	}
	return animals
}

func preparePrintAnimals(animals []Animal) string {
	var result string
	for _, animal := range animals {
		result += fmt.Sprintf("Тип:%s, Имя: %s, Возраст: %d\n", animal.Type, animal.Name, animal.Age)
	}
	return result
}

type Dish struct {
	Name  string
	Price float64
}

type Order struct {
	Dishes []Dish
	Total  float64
}

func (order *Order) AddDish(dish Dish) {
	order.Dishes = append(order.Dishes, dish)
}

func (order *Order) RemoveDish(dish Dish) {
	for i, d := range order.Dishes {
		if d.Name == dish.Name && d.Price == dish.Price {
			order.Dishes = append(order.Dishes[:i], order.Dishes[i+1:]...)
			break
		}
	}
}

func (order *Order) CalculateTotal() {
	var total float64
	for _, dish := range order.Dishes {
		total += dish.Price
	}
	order.Total = total
}

func countWordOccurrences(text string) map[string]int {
	occurrences := make(map[string]int)
	words := strings.Fields(text) // разбиваем текст на слова

	for _, word := range words {
		occurrences[word]++ // увеличиваем счетчик для каждого слова
	}

	return occurrences
}

func mergeMaps(map1, map2 map[string]int) map[string]int {
	mergedMap := make(map[string]int)

	for key, value := range map1 {
		mergedMap[key] = value
	}

	for key, value := range map2 {
		mergedMap[key] = value
	}

	return mergedMap
}

func createUniqueText(text string) string {
	uniqueWords := make(map[string]bool)
	words := strings.Fields(text)
	var result []string

	for _, word := range words {
		if !uniqueWords[word] {
			uniqueWords[word] = true
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}

func filterSentence(sentence string, filter map[string]bool) string {
	words := strings.Fields(sentence) // разбиваем предложение на слова
	result := []string{}              // создаем пустой слайс для хранения отфильтрованных слов

	for _, word := range words {
		if !filter[word] { // если слово не находится в карте фильтра
			result = append(result, word) // добавляем его в результат
		}
	}

	return strings.Join(result, " ") // объединяем отфильтрованные слова в строку с пробелами
}

// type Word struct {
// 	Value string
// 	Upper bool
// }

// // splitSentences разделяет текст на предложения
// func splitSentences(text string) []string {
// 	// Регулярное выражение для разделения текста на предложения
// 	re := regexp.MustCompile(`"(?m)(?<=\.|\?|\!)\s"`)

// 	return re.Split(text, -1)
// }

// // WordsToSentence объединяет слова в предложение
// func WordsToSentence(words []Word) string {
// 	sentence := ""

// 	for i, word := range words {
// 		if i == 0 || word.Upper { // первое слово в предложении или слово с заглавной буквы
// 			sentence += CheckUpper(word.Value)
// 		} else {
// 			sentence += word.Value
// 		}

// 		if i < len(words)-1 { // добавляем пробелы между словами, кроме последнего
// 			sentence += " "
// 		}
// 	}

// 	return sentence
// }

// // CheckUpper проверяет, нужно ли добавить заглавную букву к слову
// func CheckUpper(word string) string {
// 	if len(word) == 0 {
// 		return word
// 	}

// 	firstChar := word[0]
// 	if firstChar >= 'A' && firstChar <= 'Z' {
// 		return word
// 	}

// 	return fmt.Sprintf("%c%s", firstChar-32, word[1:])
// }

// // filterWords Фильтрует текст, заменяя цензурные и повторяющиеся слова
// func filterWords(text string, censorMap map[string]string) string {
// 	sentences := splitSentences(text)

// 	if len(sentences) > 1 { // если предложений больше одного
// 		for i, sentence := range sentences {
// 			sentences[i] = filterWords(sentence, censorMap) // обработка каждого предложения рекурсивно
// 		}

// 		return strings.Join(sentences, " ")
// 	}

// 	words := strings.Fields(text)
// 	uniqueWords := make(map[string]Word)
// 	result := []Word{}

// 	for _, wordStr := range words {
// 		word := Word{Value: wordStr}

// 		if censorMap[wordStr] != "" { // если слово является цензурным
// 			word.Value = censorMap[wordStr]
// 			word.Upper = true
// 		} else if w, ok := uniqueWords[strings.ToLower(wordStr)]; ok { // если слово уже встречалось
// 			w.Value = ""
// 			word = w
// 		} else { // если слово уникальное
// 			uniqueWords[strings.ToLower(wordStr)] = word
// 		}

// 		result = append(result, word)
// 	}

// 	for i, word := range result { // замена в слайсе слов при помощи карты уникальных слов и их индекса
// 		if word.Value == "" {
// 			continue
// 		}

// 		if w, ok := uniqueWords[strings.ToLower(word.Value)]; ok {
// 			result[i] = w
// 		}
// 	}

// 	return WordsToSentence(result)
// }

func countBytes(s string) int {
	return len(s)
}

func countSymbols(s string) int {
	return utf8.RuneCountInString(s)
}

func ReverseStrings(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getType(i interface{}) string {
	switch i.(type) {
	case nil:
		return "Пустой интерфейс"
	default:
		return fmt.Sprintf("Тип: %T", i)
	}
}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func JsonEncode() []byte {
	person := Person{
		Name:  "Иван",
		Age:   30,
		Email: "ivan@example.com",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Ошибка сериализации в JSON:", err)
		return nil
	}

	return jsonData
}

type Comment struct {
	Text string `json:"text"`
}

type UserJSON struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

func getJSON(data []UserJSON) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
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

	const (
		Monday1    int = 1
		Tuesday2   int = 2
		Wednesday3 int = 3
		Thursday4  int = 4
		Friday5    int = 5
		Saturday6  int = 6
		Sunday7    int = 7
	)

	const (
		id1 = 2 * iota
		id2
		id3
		id4
		id5
		id6
		id7
		id8
		id9
		id10
		id11
		id12
		id13
	)
	fmt.Println(id1, id2, id3, id4, id5, id6, id7, id8, id9, id10, id11, id12, id13)

	//WriteInfo()

	a := 10.0
	b := 5.0

	difference, err := calculate(a, b, "difference")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Difference:", difference)
	}

	sum, err := calculate(a, b, "sum")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sum:", sum)
	}

	product, err := calculate(a, b, "product")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Product:", product)
	}

	quotient, err := calculate(a, b, "quotient")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Quotient:", quotient)
	}

	aHypotenuse := 3.0
	bHypotenuse := 4.0
	resultHypotenuse := hypotenuse(&aHypotenuse, &bHypotenuse)
	fmt.Println(resultHypotenuse)

	// rate, err := currencyPairRate("USD", "EUR", 100.0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(rate) // 82.73

	isEqual, difference := CompareRoundedValues(1.23456789, 1.13456788, 8)
	fmt.Println(isEqual, difference)

	binaryStr := "00111110001000000000000000000000"
	resultBinary := binaryStringToFloat(binaryStr)
	fmt.Println(resultBinary) // Выведет: 0.15625

	resultUserInfo := UserInfo("Jane", "Los Angeles", "987-654-3210", 25, 150)
	fmt.Println(resultUserInfo)

	resultUserInfoMainCity := UserInfoMainCityes("Alex", 34)
	fmt.Println(resultUserInfoMainCity)

	resultConcatStrings := ConcatenateStrings(":", "somes strings some strings")
	fmt.Println(resultConcatStrings)

	change, err := CalculatePercentageChange("8", "10")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(change)
	}

	fmt.Println(MathOperate(SumAn, 1, 1, 3))
	fmt.Println(MathOperate(MulAn, 1, 7, 3))
	fmt.Println(MathOperate(SubAn, 13, 2, 3))

	fmt.Println(Factorial(10))

	fmt.Println(Fibonacci(6))

	counter := createCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	m := multiplier(2.5)
	resultMulty := m(10)
	fmt.Println(resultMulty)

	addTwo := adder(2)
	resultAdder := addTwo(3)
	fmt.Println(resultAdder)

	nums := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sumArray(nums))     // Output: 36
	fmt.Println(averageArray(nums)) // Output: 4.5

	floatNums := [8]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	fmt.Println(averageFloatArray(floatNums)) // Output: 4.5

	reversedNums := reverseArray(nums)
	fmt.Println(reversedNums) // Output: [8 7 6 5 4 3 2 1]

	numsInt := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Println("Original numsInt:", numsInt)
	sortDesc(numsInt)
	fmt.Println("Sorted numsInt:", numsInt)

	numsFloat := []float64{1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9}
	fmt.Println("Original numsInt:", numsFloat)
	sortDesc(numsFloat)
	fmt.Println("Sorted numsInt:", numsFloat)

	numsIntAsc := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Println("Original numsInt:", numsIntAsc)
	sortDesc(numsIntAsc)
	fmt.Println("Sorted numsInt:", numsIntAsc)

	numsFloatAsc := []float64{1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9}
	fmt.Println("Original numsInt:", numsFloatAsc)
	sortDesc(numsFloatAsc)
	fmt.Println("Sorted numsInt:", numsFloatAsc)

	numbersSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	subSlice := getSubSlice(numbersSlice, 2, 6)
	fmt.Println(subSlice)

	numbersUnique := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1, 6, 7}
	singleNumber := findSingleNumber(numbersUnique)
	fmt.Println(singleNumber)

	numbersAppendSlice := []int{1, 2, 3}
	resultAppenSlice := appendInt(numbersAppendSlice, 4, 5, 6)
	fmt.Println(resultAppenSlice)

	xs := []int{1, 2, 3}
	appendInt2(&xs, 4, 5, 6)
	fmt.Println(xs)

	usersStruc := getUsers()
	resultUsersStruc := preparePrint(usersStruc)
	fmt.Println(resultUsersStruc)

	animalsStruc := getAnimals()
	resultAnimalsStruc := preparePrintAnimals(animalsStruc)
	fmt.Println(resultAnimalsStruc)

	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	order.CalculateTotal()
	fmt.Println("Total:", order.Total)

	order.RemoveDish(dish1)

	order.CalculateTotal()
	fmt.Println("Total:", order.Total)

	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	occurrences := countWordOccurrences(text)
	for word, count := range occurrences {
		fmt.Printf("%s: %d\n", word, count)
	}

	resultCreateUniqueText := createUniqueText("bar bar bar foo foo baz")
	fmt.Println(resultCreateUniqueText)

	// textUnique := "Lorem ipsum dolor sit amet consectetur adipiscing elit. Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium. Lorem ipsum dolor sit amet consectetur adipiscing elit."
	// censorMap := map[string]string{"ipsum": "****", "dolor": "****"}
	// filteredText := filterWords(textUnique, censorMap)
	// fmt.Println(filteredText)

	messageTest := "Привет, мир!"
	resultCountByteString := countBytes(messageTest)
	fmt.Println(resultCountByteString)
	resultCountSymbolsString := countSymbols(messageTest)
	fmt.Println(resultCountSymbolsString)

	var i interface{} = 42
	fmt.Println(getType(i)) // Вывод: "int"
	var j interface{} = "Hello, World!"
	fmt.Println(getType(j)) // Вывод: "string"
	var k interface{} = []int{1, 2, 3}
	fmt.Println(getType(k)) // Вывод: "[]int"
	var l interface{} = interface{}(nil)
	fmt.Println(getType(l))

	x := 4.0
	y := 2.0
	fmt.Println("Square root of", x, "is", mymath.MySqrt(x))
	fmt.Println("Ceil of", x, "is", mymath.MyCeil(x))
	fmt.Println("Floor of", x, "is", mymath.MyFloor(x))
	fmt.Println("Power of", x, "to", y, "is", mymath.MyPow(x, y))
	fmt.Println("Max of", x, "and", y, "is", mymath.MyMax(x, y))
	fmt.Println("Min of", x, "and", y, "is", mymath.MyMin(x, y))

	fmt.Println("Square root of", x, "is", testmymath.TestMySqrt(x))
	fmt.Println("Ceil of", x, "is", testmymath.TestMyCeil(x))
	fmt.Println("Floor of", x, "is", testmymath.TestMyFloor(x))
	fmt.Println("Power of", x, "to", y, "is", testmymath.TestMyPow(x, y))
	fmt.Println("Max of", x, "and", y, "is", testmymath.TestMyMax(x, y))
	fmt.Println("Min of", x, "and", y, "is", testmymath.TestMyMin(x, y))

	resultJsonEncode := JsonEncode()
	fmt.Println(resultJsonEncode)

	users := []UserJSON{
		{Name: "Alice", Age: 25, Comments: []Comment{{Text: "Hello"}, {Text: "World"}}},
		{Name: "Bob", Age: 30, Comments: []Comment{{Text: "Goodbye"}}},
	}

	jsonData, err := getJSON(users)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(jsonData)
	}
}
