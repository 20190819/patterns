package flyWeight

import (
	"fmt"
	"testing"
)

var characters []string
var fonts []string

func init() {
	characters = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}

	fonts = []string{"Arial", "Times New Roman", "Verdana", "Helvetica"}
}

func TestFlyWeight(t *testing.T) {

	factory := NewFactory()

	for _, character := range characters {
		for _, font := range fonts {
			factory.get(character).Render(font)
		}
	}

	fmt.Println("pool len: ", factory.count())
}
