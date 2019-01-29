package errors

/*
Implements error handling for the BGP connector
*/

// For now, this kills the program but in the future this should signal the broker
// that this particular goroutine is screwed
func RaiseError(e interface{}) {
	//log.Print(e)
	//log.Fatal(e)
	panic(e)
}

// Check to see if the passed err value is non-nil and handle accordingly
func CheckError(err error) {
	if err != nil {
		RaiseError(err.Error())
	}
}
