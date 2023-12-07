package main

import (
	"fmt"
	"strconv"
)

var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	var i int = 100
	fmt.Println(i)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	fmt.Println(i5)

	outer()

	var s5 string = "Not Use"
	fmt.Println(s5)

	fmt.Println(i)

	var i8 int64 = 200

	fmt.Println(i + 50)

	fmt.Printf("%T\n", i8)

	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + i2)

	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	var t, f bool = true, false
	fmt.Println(t, f)

	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	test
		test`)
	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))

	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte("Hi")
	fmt.Println(c)

	fmt.Println(string(c))

	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B", "C"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])

	arr2[2] = "D"
	fmt.Println(arr2)

	//var arr5 [4]int
	//arr5 = arr1
	//fmt.Println(arr5)

	fmt.Println(len(arr1))

	var x interface{}
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)
	x = 1
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	var i5 int = 1
	fl642 := float64(i5)
	fmt.Println(fl642)
	fmt.Printf("i = %T\n", i5)
	fmt.Printf("fl64 = %T\n", fl642)

	i6 := int(fl64)
	fmt.Printf("i2 = %T\n", i6)

	fl2 := 10.5
	i7 := int(fl2)
	fmt.Printf("i7 = %T\n", i7)
	fmt.Println(i7)

	var s6 string = "100"
	fmt.Printf("s6 = %T\n", s6)

	i9, _ := strconv.Atoi(s6)
	fmt.Println(i9)
	fmt.Printf("i9 = %T\n", i9)

	var i10 int = 200
	s7 := strconv.Itoa(i10)
	fmt.Println(s7)
	fmt.Printf("s7 = %T\n", s7)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)

}
