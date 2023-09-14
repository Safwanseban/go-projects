// Extend program (5) to accept 3 inputs: “num1”, “num2” and
// “operation” where operation could be “+”, “-”, “*” or “/” to represent sum,
//  difference, multiplication or division. The output will be output of  “num1”
//  <operation> “num2”. The output shall be “num1=<num1> num2=<num2> <operation>=<result>”
//   where “<operation>” will be replaced  by the operation name. Use “sum”, “difference”,
//   “multiply” and “divide” as an operation name when printing the result.
//   [If/then/else OR switch statement, 3 hours]

package main

import "fmt"

func main() {

	var num1, num2 float32
	var operand string

	fmt.Scanf("%f", &num1)
	fmt.Scanf("%f", &num2)
	fmt.Scanf("%s", &operand)
	returnResult(num1, num2, operand)

}

func returnResult(num1, num2 float32, operand string) (n int, err error) {

	switch operand {

	case "+":
		return fmt.Printf("num1=%f\t num2=%f\t sum=%f", num1, num2, num1+num2)
	case "-":
		return fmt.Printf("num1=%f\t num2=%f\t diffrence=%f", num1, num2, num1-num2)
	case "*":
		return fmt.Printf("num1=%f\t num2=%f\t multiply=%f", num1, num2, num1*num2)
	case "/":
		if num2==0{
			return fmt.Println("0 division error")
		}
		return fmt.Printf("num1=%f\t num2=%f\t division=%f", num1, num2, num1/num2)

	default:
		return fmt.Printf("invalid %s", operand)
	}

}
