// Package struct1 for playing arround with a structure
//
// Nothing important
// just playing arround
package struct1

import "fmt"

// Str1 is the structure we play with in this package
// second line
// third line (of comment for Str1)
type Str1 struct {
	id   int
	name string
}

// NewStr1 constucts and returns new structure
func NewStr1(id int, name string) *Str1 {
	str := new(Str1)
	str.id = id
	str.name = name
	return str
}

// SayHello simply greets the supplied name
func (str *Str1) SayHello(friend string) {
	fmt.Println("Hello", friend, "from", str.name)
}
