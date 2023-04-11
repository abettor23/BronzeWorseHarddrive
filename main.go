package main

import (
	"fmt"
)

func main() {
	fmt.Println("Привет! Введите любое целое число")

  var i int

  fmt.Scan(&i)

  fmt.Println("Введенное вами число:", i)
}
