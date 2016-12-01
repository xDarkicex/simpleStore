package helper

import "io"

// Close Wrapper not yet finished
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		Logger.Println(err)
	}
}
