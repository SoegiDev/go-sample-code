package main

import "fmt"

func main() {
	var i int

	fmt.Print("Type two numbers: ")
	fmt.Scanln(&i)
	if i == 1 {
		basic_1()
	}
}

func basic_1() {
	/// TIPE DATA GOLANG LANGUAGE TYPE
	fmt.Println("Tipe Data GO Language")
	// String
	t1 := "text"
	//Array String
	t2 := []string{"apple", "strawberry", "blueberry"}
	//Array String Float
	t3 := map[string]float64{"strawberry": 3.2, "blueberry": 1.2}
	// Integer
	t4 := 2
	// Flsoating
	t5 := 4.5
	//Bool
	t6 := true

	fmt.Printf("[]string{'apple', 'strawberry', 'blueberry'} %T\n", t1)
	fmt.Printf("t2: %T\n", t2)
	fmt.Printf("t3: %T\n", t3)
	fmt.Printf("t4: %T\n", t4)
	fmt.Printf("t5: %T\n", t5)
	fmt.Printf("t6: %T\n", t6)
}

func test_1() {
	// 2 DIMENSI ARRAY
	// Creating and initializing
	// 2-dimensional array
	// Using shorthand declaration
	// Here the (,) Comma is necessary
	arr := [3][3]string{{"C #", "C", "Python"}, {"Java", "Scala", "Perl"},
		{"C++", "Go", "HTML"}}
	// Accessing the values of the
	// array Using for loop
	fmt.Println("Elements of Array 1")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {

			fmt.Println(arr[x][y])
		}
	}

	// Creating a 2-dimensional
	// array using var keyword
	// and initializing a multi
	// -dimensional array using index
	var arr1 [2][2]int
	arr1[0][0] = 100
	arr1[0][1] = 200
	arr1[1][0] = 300
	arr1[1][1] = 400

	// Accessing the values of the array
	fmt.Println("Elements of array 2")
	for p := 0; p < 2; p++ {
		for q := 0; q < 2; q++ {
			fmt.Println(arr1[p][q])
		}
	}
}
