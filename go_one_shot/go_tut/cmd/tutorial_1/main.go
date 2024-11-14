package main
import "fmt"

// func main()  {
// 	fmt.Println(("Hello world!"))
// 	var intNum int16 = 32767
// 	intNum = intNum + 1;
// 	fmt.Println((intNum))

// 	var floatNum float64 = 12345678.9
// 	fmt.Println(floatNum)

// 	var floatNum32 float32 = 10.1
// 	var intNum32  int32 = 2
// 	var result float32 = floatNum32 + float32(intNum32)
// 	fmt.Println(result) // 12.1

// 	var intNum1 int = 2
// 	var intNum2 int = 1
// 	fmt.Println((intNum1/intNum2)) // 1
// 	fmt.Println(intNum1%intNum2) // 0


// 	myVar := "testText"
// 	fmt.Println((myVar))

// 	var1, var2  := 1,2
// 	fmt.Println(var1,var2)

// 	// functions and control structures in golang
	
// 	printMe()
// }

// func  printMe(){
// 		fmt.Println("Hello world")
		
// 	}

func main(){
	var printValue string = "hello world"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result int = intDivision(numerator, denominator)
	fmt.Println(result)


}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) int {
	var result int = numerator/denominator
	return result
}