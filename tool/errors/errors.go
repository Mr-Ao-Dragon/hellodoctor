package errors

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
