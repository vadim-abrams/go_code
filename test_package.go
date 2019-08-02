package greetings

import "fmt"

func Hello(name string) {
  greeting := "Hello, %v!\n"

  fmt.Printf(greeting, name)
}
