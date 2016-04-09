package generator

import (
	"fmt"
	"testing"
)

func TestFuncs(t *testing.T) {
	c := &Config{}

	for {
		fmt.Println(FirstName(c), LastName(c))
	}
}
