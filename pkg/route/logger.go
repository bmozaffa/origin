package route

import (
	"os"
	"runtime/debug"
	"fmt"
)

func LogToFile(text string) error {
	path := "/tmp/myorigin.log"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Println("LogToFile")
	f.Write(debug.Stack())
	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}
