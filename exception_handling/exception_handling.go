package main

import (
	"fmt"
	"os"
	"strconv"
)

// builin interface
/*
type interface error{
	Error()string
}
*/

type FileNotFoundException struct{}
type PermissionDeniedException struct{}

func (ex *FileNotFoundException) Error() string {
	return "File not found exception"
}

func (ex *PermissionDeniedException) Error() string {
	return "Permission denied exception"
}

func main() {
	//try
	defer func() {
		//catch
		if ex := recover(); ex != nil {
			fmt.Println("exit abnormally")
			if e, ok := ex.(error); ok {
				fmt.Println("caused by ", e.Error())
			} else {
				fmt.Println("caused by Unknown exception")
			}
			os.Exit(1) // exit in C++
		} else {
			fmt.Println("exit normally")
			os.Exit(0)
		}
	}()

	//os.Args:command line arguments
	if len(os.Args) == 1 {
		fmt.Println("no args")
		os.Exit(0)
	}

	if code, err := strconv.Atoi(os.Args[1]); err == nil {
		switch code {
		case 1:
			panic(&FileNotFoundException{}) //throw Exception
		case 2:
			panic(&PermissionDeniedException{})
		}
	} else {
		fmt.Println("invalid args")
		panic(err)
	}
}
