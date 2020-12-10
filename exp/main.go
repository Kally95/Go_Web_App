package main

import (
	"fmt"

	"github.com/Kally95/Go_Web_App/rand"
)

func main() {
	fmt.Println(rand.String(10))
	fmt.Println(rand.RememberToken())
}
