package main

func main() {

}

func maxOnesAfterRemoveItem(slice []byte) uint {
	var maxLengh uint = 0
	var currentLengh uint = 0
	var hasZero bool = false
	var headLentgh uint = 0
	var tailLengh uint = 0

	for i := 0; i < len(slice); i++ {
		currentLengh = headLentgh + tailLengh

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
			//error
		}
	}

	currentLengh = headLentgh + tailLengh
	if currentLengh > maxLengh {
		maxLengh = currentLengh
	}

	if maxLengh == uint(len(slice)) {
		maxLengh--
	}

	return maxLengh
}
