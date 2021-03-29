package main

import (
	"fmt"

	"github.com/anusha/aspiration_mapper/mapper"
)

// And change your code to look like this:

func main() {
	s := mapper.NewSkipString(3, "Aspiration.com")
	mapper.MapString(&s)
	fmt.Println(s)
}
