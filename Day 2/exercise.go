package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = "989133-1014784,6948-9419,13116184-13153273,4444385428-4444484883,26218834-26376188,224020-287235,2893-3363,86253-115248,52-70,95740856-95777521,119-147,67634135-67733879,2481098640-2481196758,615473-638856,39577-47612,9444-12729,93-105,929862406-930001931,278-360,452131-487628,350918-426256,554-694,68482544-68516256,419357748-419520625,871-1072,27700-38891,26-45,908922-976419,647064-746366,9875192107-9875208883,3320910-3352143,1-19,373-500,4232151-4423599,1852-2355,850857-889946,4943-6078,74-92,4050-4876"
var inputArray []string
var acc int64 = 0

func main() {

	f := func(c rune) bool {
		return c == ','
	}

	inputArray = strings.FieldsFunc(input, f)

	for i := range inputArray {
		f := func(c rune) bool {
			return c == '-'
		}
		currentRange := strings.FieldsFunc(inputArray[i], f)

		firstInput, err := strconv.ParseInt(currentRange[0], 10, 64)
		if err != nil {
			fmt.Print("Error")
		}
		secondInput, err := strconv.ParseInt(currentRange[1], 10, 64)
		if err != nil {
			fmt.Print("Error")
		}

		for j := firstInput; j <= secondInput; j++ {

			convertedString := strconv.FormatInt(j, 10)
			length := len(convertedString)

			if length > 0 && length%2 == 0 {

				firstHalf := convertedString[:length/2]
				secondHalf := convertedString[length/2:]

				if firstHalf == secondHalf {
					fmt.Println("Converted string is : ", convertedString)
					acc += j
				}
			}
		}

	}
	fmt.Println("Total should be", acc)
}
