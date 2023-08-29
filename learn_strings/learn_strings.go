package main

import "fmt"

func main() {
	s := "Hello, \n\t\"world!\" with a backslash \\"
	t := `Hello,
	"world!" with a backslash \`
	e := "😀 "
	fmt.Println(s+e)
	fmt.Println(t)
	j := "Hello, world!"
	b := j[0]
	b2 := j[4]
	fmt.Println(j, b, string(b), b2, string(b2))
	j2 := j[:5]
	//Just like python
	fmt.Println(j, len(j), j2, len(j2))
	fmt.Println(e[0])

	var r rune = 127757
	fmt.Println(j + string(r))

	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6
	fmt.Println(vals, vals[0], vals[1], vals[2])
	// Arrays have to be decralred with its size, due to which it is not used widely.
}
