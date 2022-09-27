package logger

import "fmt"

func ErrPrint(output string) {
	fmt.Println("j:", output)
	fmt.Println(`For more information, use "j --help"`)
}
