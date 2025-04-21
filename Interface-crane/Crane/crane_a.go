package Crane

import (
	"fmt"
)

type CraneA struct {
	Name string
}

func (c *CraneA) Up() string {
	fmt.Println("Hello from CraneA")
	return "Hello from CraneA"
}

func (c *CraneA) Down() string {
	fmt.Println("Goodbye from CraneA")
	return "Goodbye from CraneA"
}
