package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	A int
	b string
}

type Bar struct {
	C string
	F Foo
}

type Baz struct {
	D string
	Foo
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	f := Foo{}
	fmt.Println(f)

	g := Foo{10, "Hello"}
	fmt.Println(g)

	h := Foo{
		b: "Bye",
	}
	fmt.Println(h)

	h.A = 1000
	fmt.Println(h.A)

	ff := Foo{
		A: 20,
	}
	var f2 Foo
	f2 = ff
	f2.A = 100
	fmt.Println(f2.A)
	fmt.Println(ff.A)

	var f3 *Foo = &f
	f3.A = 200
	fmt.Println(f3.A)
	fmt.Println(f.A)

	fs := Foo{A: 10, b: "Hello"}
	b1 := Bar{C: "Fred", F: fs}
	fmt.Println(b1.F.A)
	b2 := Baz{D: "Nancy", Foo: fs}
	fmt.Println(b2.A)
	var f4 Foo = b2.Foo
	fmt.Println(f4)

	bob := `{"name": "Bob", "age": 30}`
	var b Person
	json.Unmarshal([]byte(bob), &b)
	fmt.Println(b)
	bob2, _ := json.Marshal(b)
	fmt.Println(string(bob2))
}
