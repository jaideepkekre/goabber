//Jaideep Kekre
//Checks if word is valid abbreviation of target word

package goabber

import (
	"fmt"
	"unicode/utf8"
)

//Abber  is called with the abbereviation to check, the word to check that is a abbeviation of.
func Abber(inputAbber string, target string) bool {

	return isAbber(inputAbber, target, false)

}

func isAbber(inputAbber string, target string, returnvalue bool) bool {

	fmt.Println("abber: ", utf8.RuneCountInString(inputAbber))
	fmt.Println("target:" + target)
	if utf8.RuneCountInString(inputAbber) == 0 || returnvalue == true {
		returnvalue = true

	} else if utf8.RuneCountInString(inputAbber) > 0 && utf8.RuneCountInString(target) == 0 {
		returnvalue = false
	} else if inputAbber[0] == target[0] {
		fmt.Println("Matched : " + string(inputAbber[0]))
		inputAbber = inputAbber[1:]
		target = target[1:]
		returnvalue = isAbber(inputAbber, target, returnvalue)

	} else {
		fmt.Println("skipping : " + string(target[0]))
		target = target[1:]
		returnvalue = isAbber(inputAbber, target, returnvalue)
	}
	fmt.Println(returnvalue)
	return returnvalue
}
