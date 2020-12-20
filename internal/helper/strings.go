package helper

func StringContainsFrom(theString, subString string, from int) bool {
	to := from + len(subString)
	if len(theString) < to {
		return false
	}
	return theString[from:to] == subString
}
