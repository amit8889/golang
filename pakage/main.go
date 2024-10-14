package main

import (
	"fmt"
	"package/auth"
)

func main() {
	// your code here
	fmt.Println("===========inside package=====")
	auth.LoginWithCredential(121, "Amit verma")

}
