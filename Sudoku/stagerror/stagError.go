package stagerror

import (
	"os"
)

// Problem is we can't do anything when
// an error occurs in here
func SaveToLog(err error, isRelease bool) {
	if err == nil {
		return
	}

	// Make sure that when we hit an error
	// in development we know about it
	if !isRelease {
		panic(err)
	}

	// However, if a user hits an error,
	// don't crash, we don't want it
	// ruining their experience. Instead,
	// we just log it for them to tell us
	// about

	data, err := os.ReadFile("errors.log")
	if err != nil {
		data = []byte{}
		err := os.WriteFile("errors.log", data, 0644)

		// Make sure to hard error if we can't make the log
		if err != nil {
			panic(err)
		}
	}

	data = append(data, []byte(err.Error())...)
	os.WriteFile("errors.log", data, 0644)
}
