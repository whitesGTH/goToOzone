package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func main() {
	//fmt.Println("Hello, playground")
	//mass := []byte{1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1}
	//mass := []byte{0, 1}
	//res, e := maxOnesAfterRemoveItem(mass)
	//res := maxOnesAfterRemoveItem(mass)

	//if e == nil {
	//fmt.Println("Result is: ", res)
	//} else {
	//	fmt.Println(e)
	//}
}

func TestSomething(t *testing.T) {

	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")

}

//func maxOnesAfterRemoveItem(slice []byte) (uint, error) {
func maxOnesAfterRemoveItem(slice []byte) uint {
	var maxLengh uint = 0
	var currentLengh uint = 0
	var hasZero bool = false
	var headLentgh uint = 0
	var tailLengh uint = 0

	fmt.Println("slice.len: ", len(slice))

	for i := 0; i < len(slice); i++ {
		currentLengh = headLentgh + tailLengh
		fmt.Printf("%02d: [i]:%d h:%02d t:%02d c:%02d m:%02d z:%t\n", i, slice[i], headLentgh, tailLengh, currentLengh, maxLengh, hasZero)

		if slice[i] == 0 {
			if hasZero {
				if currentLengh > maxLengh {
					maxLengh = currentLengh
				}
				headLentgh = 0
				tailLengh = 0
				continue
			}

			hasZero = true
		} else if slice[i] == 1 {
			if hasZero {
				//currentLengh = headLentgh + tailLengh //дублирование
				if currentLengh > maxLengh {
					maxLengh = currentLengh
				}

				if headLentgh == 0 {
					headLentgh++
				} else if tailLengh > 0 {
					headLentgh = tailLengh
					tailLengh = 1
				} else {
					tailLengh++
				}

			} else {
				if tailLengh == 0 {
					headLentgh++
				} else {
					tailLengh++
				}
			}

			hasZero = false
		} else {
			//return 0, errors.New("Wrong symbol in slice")
		}
	}

	currentLengh = headLentgh + tailLengh
	if currentLengh > maxLengh {
		maxLengh = currentLengh
	}

	if maxLengh == uint(len(slice)) {
		maxLengh--
	}

	fmt.Printf("%02d: [i]:%d h:%02d t:%02d c:%02d m:%02d z:%t\n", len(slice), slice[len(slice)-1], headLentgh, tailLengh, currentLengh, maxLengh, hasZero)
	fmt.Println(maxLengh)

	//return maxLengh, nil
	return maxLengh
}
