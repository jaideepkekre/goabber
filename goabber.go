package abber

import (
	"fmt"
	"unicode/utf8"
)

func abber(input_abber string, target string) bool {

	return _isAbber(input_abber, target, false)

}

func _isAbber(input_abber string, target string, returnvalue bool) bool {

	fmt.Println("abber: ", utf8.RuneCountInString(input_abber))
	fmt.Println("target:" + target)
	if utf8.RuneCountInString(input_abber) == 0 || returnvalue == true {
		returnvalue = true

	} else if utf8.RuneCountInString(input_abber) > 0 && utf8.RuneCountInString(target) == 0 {
		returnvalue = false
	} else if input_abber[0] == target[0] {
		fmt.Println("Matched : " + string(input_abber[0]))
		input_abber = input_abber[1:]
		target = target[1:]
		returnvalue = isAbber(input_abber, target, returnvalue)

	} else {
		fmt.Println("skipping : " + string(target[0]))
		target = target[1:]
		returnvalue = isAbber(input_abber, target, returnvalue)
	}
	fmt.Println(returnvalue)
	return returnvalue
}
