package auth

import "fmt"

func LoginWithCredential[T interface{}, V string](pass T, name V) {
	// login logic
	fmt.Println("=======>login credential pass and name :", pass, name)
}
