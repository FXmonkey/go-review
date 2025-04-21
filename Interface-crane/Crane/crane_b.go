package Crane

import (
	"fmt"
)

type CraneB struct {
	Name string
}

func (c *CraneB) SayHello() string {
	fmt.Println("Hello from CraneB")
	return "Hello from CraneB"
}

func (c *CraneB) SayGoodbye() string {
	fmt.Println("Goodbye from CraneB")
	return "Goodbye from CraneB"
}
