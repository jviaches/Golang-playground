package main

import "fmt"
import "math"

func main() {
	fmt.Println(fibbRecc(6))
	fmt.Println(fibShortIterative(6))

	f:=fibIterFunc()
	for i := 0; i < 10; i++ {
		fmt.Print(f() )
	}
}

func fibbRecc(numIndex int) (numValue int){
	if numIndex == 0 || numIndex == 1 {
		return numIndex
	}

	return fibbRecc(numIndex - 2) + fibbRecc(numIndex -1)
}

// Short iterative mathematics form
func fibShortIterative(numIndex float64) (float64){
	Phi := (1 + math.Sqrt(5)) / 2
	phi := (1 - math.Sqrt(5)) / 2

	return (math.Pow(Phi, numIndex) - math.Pow(phi, numIndex)) / math.Sqrt(5)
}

func fibIterFunc() func() int{
	x:=0
	y:=1

	return func() int{
		x,y = y,x+y
		return  x
	}
}