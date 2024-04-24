package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

// func TestFactorial(t *testing.T) {
// 	result := Factorial(10)

// 	if result != 3628801 {
// 		t.Errorf("Factorial(10) = %d; want 3628800", result)
// 	}
// }

func TestHelloWorldFunc(t *testing.T) {
	var buf bytes.Buffer

	fmt.Fprint(&buf, HelloWorld())

	output := buf.String()
	expected := "Hello world!"

	if output != expected {
		t.Errorf("Unexpected output: %s", output)
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range tests {
		result := Fibonacci(test.input)
		if result != test.expected {
			t.Errorf("Fibonacci(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func average(xs []float64) float64 {
	sum := 0.0
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}

func generateSlice(size int) []float64 {
	rand.Seed(time.Now().UnixNano())
	slice := make([]float64, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Float64() * 100
	}
	return slice
}

func TestAverage(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3},
		{[]float64{10, 20, 30, 40, 50}, 30},
		{[]float64{-1, -2, -3, -4, -5}, -3},
	}

	for _, test := range tests {
		result := average(test.input)
		if result != test.expected {
			t.Errorf("average(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}

	for i := 0; i < 10; i++ {
		size := rand.Intn(100)
		slice := generateSlice(size)
		expected := average(slice)
		result := average(slice)
		if result != expected {
			t.Errorf("average(%v) = %v; expected %v", slice, result, expected)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibonacciFormula(7)
	}
}

func TestSumArray(t *testing.T) {
	nums := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := 36
	result := sumArray(nums)
	if result != expected {
		t.Errorf("sum(%v) = %d; expected %d", nums, result, expected)
	}
}

func TestAverageArray(t *testing.T) {
	nums := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := 4.5
	result := averageArray(nums)
	if result != expected {
		t.Errorf("average(%v) = %f; expected %f", nums, result, expected)
	}
}

func TestAverageFloatArray(t *testing.T) {
	floatNums := [8]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	expected := 4.5
	result := averageFloatArray(floatNums)
	if result != expected {
		t.Errorf("averageFloat(%v) = %f; expected %f", floatNums, result, expected)
	}
}

func TestReverseArray(t *testing.T) {
	nums := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := [8]int{8, 7, 6, 5, 4, 3, 2, 1}
	result := reverseArray(nums)
	if result != expected {
		t.Errorf("reverse(%v) = %v; expected %v", nums, result, expected)
	}
}

func TestSortDescIntAndFloat(t *testing.T) {
	numsInt := []int{3, 1, 4, 1, 5, 9, 2, 6}
	expectedInt := []int{9, 6, 5, 4, 3, 2, 1, 1}
	sortDesc(numsInt)
	if !equal(numsInt, expectedInt) {
		t.Errorf("sortDesc(%v) = %v; expected %v", numsInt, numsInt, expectedInt)
	}

	numsFloat := []float64{3.3, 1.1, 4.4, 1.1, 5.5, 9.9, 2.2, 6.6}
	expectedFloat := []float64{9.9, 6.6, 5.5, 4.4, 3.3, 2.2, 1.1, 1.1}
	sortDesc(numsFloat)
	if !equal(numsFloat, expectedFloat) {
		t.Errorf("sortDesc(%v) = %v; expected %v", numsFloat, numsFloat, expectedFloat)
	}
}

func TestSortAscIntAndFloat(t *testing.T) {
	numsInt := []int{3, 1, 4, 1, 5, 9, 2, 6}
	expectedInt := []int{1, 1, 2, 3, 4, 5, 6, 9}
	sortAsc(numsInt)
	if !equal(numsInt, expectedInt) {
		t.Errorf("sortAsc(%v) = %v; expected %v", numsInt, numsInt, expectedInt)
	}

	numsFloat := []float64{3.3, 1.1, 4.4, 1.1, 5.5, 9.9, 2.2, 6.6}
	expectedFloat := []float64{1.1, 1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 9.9}
	sortAsc(numsFloat)
	if !equal(numsFloat, expectedFloat) {
		t.Errorf("sortDesc(%v) = %v; expected %v", numsFloat, numsFloat, expectedFloat)
	}
}

func TestGetSubSlice(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	start := 1
	end := 4
	expected := []int{2, 3, 4}

	actual := getSubSlice(xs, start, end)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestMaxDifference(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{5, 4, 3, 2, 1}, 0},
		{[]int{1, 2, 3, 2, 5}, 4},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 0},
		{[]int{1}, 0},
		{[]int{}, 0},
	}

	for _, test := range tests {
		result := MaxDifference(test.numbers)
		if result != test.expected {
			t.Errorf("MaxDifference(%v) = %d; want %d", test.numbers, result, test.expected)
		}
	}
}

func TestFindSingleNumber(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1, 6, 7}
	expected := 6 ^ 7 // ожидаемый результат - XOR двух уникальных чисел
	result := findSingleNumber(numbers)
	if result != expected {
		t.Errorf("findSingleNumber(%v) = %d; expected %d", numbers, result, expected)
	}
}

func TestAppendInt(t *testing.T) {
	// Тестовые данные
	xs := []int{1, 2, 3}
	x := []int{4, 5}

	// Вызываем функцию
	result := appendInt(xs, x...)

	// Проверяем результат
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("appendInt(%v, %v) = %v, ожидалось %v", xs, x, result, expected)
	}

	// Проверяем, что исходный срез не изменился
	if !reflect.DeepEqual(xs, []int{1, 2, 3}) {
		t.Errorf("appendInt(%v, %v) изменил исходный срез", xs, x)
	}
}

func TestAppendInt2(t *testing.T) {
	xs := []int{1, 2, 3}
	appendInt2(&xs, 4, 5)

	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(xs, expected) {
		t.Errorf("Expected %v but got %v", expected, xs)
	}
}

func TestCut(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}

	// Тестирование получения подмножества среза
	expected := []int{2, 3, 4}
	result := Cut(xs, 1, 4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Cut(%v, %d, %d) = %v; expected %v", xs, 1, 4, result, expected)
	}

	// Тестирование получения пустого среза при выходе за границы
	expected = []int{}
	result = Cut(xs, 2, 2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Cut(%v, %d, %d) = %v; expected %v", xs, 2, 2, result, expected)
	}

	// Тестирование получения всего среза
	expected = []int{1, 2, 3, 4, 5}
	result = Cut(xs, 0, 5)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Cut(%v, %d, %d) = %v; expected %v", xs, 0, 5, result, expected)
	}
}

func TestGetUsers(t *testing.T) {
	users := getUsers()

	if len(users) != 10 {
		t.Errorf("Expected 10 users, but got %d", len(users))
	}

	for _, user := range users {
		if user.Name == "" {
			t.Error("User name is empty")
		}
		if user.Age < 18 || user.Age > 60 {
			t.Errorf("User age %d is out of range", user.Age)
		}
	}
}

func TestPreparePrint(t *testing.T) {
	users := []User{
		{Name: "John Doe", Age: 25},
		{Name: "Jane Smith", Age: 35},
	}

	result := preparePrint(users)

	expected := "Имя: John Doe, Возраст: 25\nИмя: Jane Smith, Возраст: 35\n"

	if strings.Compare(result, expected) != 0 {
		t.Errorf("Expected:\n%s\nBut got:\n%s", expected, result)
	}
}

func TestGetAnimals(t *testing.T) {
	animals := getAnimals()

	if len(animals) != 3 {
		t.Errorf("Expected 10 users, but got %d", len(animals))
	}

	for _, animal := range animals {
		if animal.Type == "" {
			t.Error("Animal type is empty")
		}
		if animal.Name == "" {
			t.Error("Animal name is empty")
		}
		if animal.Age < 18 || animal.Age > 60 {
			t.Errorf("User age %d is out of range", animal.Age)
		}
	}
}

func TestPreparePrintAnimals(t *testing.T) {
	animals := []Animal{
		{Type: "cat", Name: "Tom", Age: 2},
		{Type: "dog", Name: "Rex", Age: 5},
	}

	result := preparePrintAnimals(animals)

	expected := "Тип:cat, Имя: Tom, Возраст: 2\nТип:dog, Имя: Rex, Возраст: 5\n"

	if strings.Compare(result, expected) != 0 {
		t.Errorf("Expected:\n%s\nBut got:\n%s", expected, result)
	}
}

func TestAddDish(t *testing.T) {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	if len(order.Dishes) != 1 {
		t.Errorf("Expected 1 dish, got %d", len(order.Dishes))
	}

	order.AddDish(dish2)
	if len(order.Dishes) != 2 {
		t.Errorf("Expected 2 dishes, got %d", len(order.Dishes))
	}
}

func TestRemoveDish(t *testing.T) {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	order.RemoveDish(dish1)
	if len(order.Dishes) != 1 {
		t.Errorf("Expected 1 dish, got %d", len(order.Dishes))
	}

	order.RemoveDish(dish2)
	if len(order.Dishes) != 0 {
		t.Errorf("Expected 0 dishes, got %d", len(order.Dishes))
	}
}

func TestCalculateTotal(t *testing.T) {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	order.CalculateTotal()
	if order.Total != 16.98 {
		t.Errorf("Expected total of 16.98, got %.2f", order.Total)
	}

	order.RemoveDish(dish1)

	order.CalculateTotal()
	if order.Total != 5.99 {
		t.Errorf("Expected total of 5.99, got %.2f", order.Total)
	}
}

func TestMergeMaps(t *testing.T) {
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"c": 3, "d": 4}
	expected := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	mergedMap := mergeMaps(map1, map2)

	if !reflect.DeepEqual(mergedMap, expected) {
		t.Errorf("Expected %v but got %v", expected, mergedMap)
	}
}

func TestCreateUniqueText(t *testing.T) {
	input := "the quick brown fox jumps over the lazy dog"
	expected := "the quick brown fox jumps over lazy dog"

	result := createUniqueText(input)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestFilterSentence(t *testing.T) {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}
	expected := "Lorem dolor sit amet consectetur adipiscing"

	result := filterSentence(sentence, filter)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestCountBytes(t *testing.T) {
	s := "hello"
	expected := 5
	result := countBytes(s)
	if result != expected {
		t.Errorf("countBytes(%q) = %d; expected %d", s, result, expected)
	}

	s = "привет"
	expected = 12
	result = countBytes(s)
	if result != expected {
		t.Errorf("countBytes(%q) = %d; expected %d", s, result, expected)
	}
}

func TestCountSymbols(t *testing.T) {
	s := "hello"
	expected := 5
	result := countSymbols(s)
	if result != expected {
		t.Errorf("countSymbols(%q) = %d; expected %d", s, result, expected)
	}

	s = "привет"
	expected = 6
	result = countSymbols(s)
	if result != expected {
		t.Errorf("countSymbols(%q) = %d; expected %d", s, result, expected)
	}
}

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character string",
			input:    "a",
			expected: "a",
		},
		{
			name:     "string with even length",
			input:    "abcd",
			expected: "dcba",
		},
		{
			name:     "string with odd length",
			input:    "abcde",
			expected: "edcba",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ReverseStrings(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %q but got %q", tc.expected, actual)
			}
		})
	}
}

func TestGetType(t *testing.T) {
	// Тест для пустого интерфейса
	var i interface{}
	expected := "Пустой интерфейс"
	if getType(i) != expected {
		t.Errorf("getType(%v) = %v, ожидается %v", i, getType(i), expected)
	}

	// Тест для значения типа int
	i = 42
	expected = "Тип: int"
	if getType(i) != expected {
		t.Errorf("getType(%v) = %v, ожидается %v", i, getType(i), expected)
	}

	// Тест для значения типа string
	i = "hello"
	expected = "Тип: string"
	if getType(i) != expected {
		t.Errorf("getType(%v) = %v, ожидается %v", i, getType(i), expected)
	}

	// Тест для значения типа bool
	i = true
	expected = "Тип: bool"
	if getType(i) != expected {
		t.Errorf("getType(%v) = %v, ожидается %v", i, getType(i), expected)
	}
}

func equal(a, b interface{}) bool {
	switch a.(type) {
	case []int:
		aa := a.([]int)
		bb := b.([]int)
		if len(aa) != len(bb) {
			return false
		}
		for i := range aa {
			if aa[i] != bb[i] {
				return false
			}
		}
		return true
	case []float64:
		aa := a.([]float64)
		bb := b.([]float64)
		if len(aa) != len(bb) {
			return false
		}
		for i := range aa {
			if aa[i] != bb[i] {
				return false
			}
		}
		return true
	}
	return false
}
func TestGetJSON(t *testing.T) {
 users := []UserJSON{
  {Name: "Alice", Age: 25, Comments: []Comment{{Text: "Hello"}, {Text: "World"}}},
  {Name: "Bob", Age: 30, Comments: []Comment{{Text: "Goodbye"}}},
 }

 expectedJSON := `[{"name":"Alice","age":25,"comments":[{"text":"Hello"},{"text":"World"}]},{"name":"Bob","age":30,"comments":[{"text":"Goodbye"}]}]`

 jsonData, err := getJSON(users)
 if err != nil {
  t.Errorf("Unexpected error: %v", err)
 }

 if jsonData != expectedJSON {
  t.Errorf("JSON does not match. Expected: %s, Got: %s", expectedJSON, jsonData)
 }

 var decodedUsers []UserJSON
 err = json.Unmarshal([]byte(jsonData), &decodedUsers)
 if err != nil {
  t.Errorf("Error decoding JSON: %v", err)
 }

 if !reflect.DeepEqual(users, decodedUsers) {
  t.Errorf("Decoded users do not match original users")
 }
}