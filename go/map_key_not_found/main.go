package main

import "fmt"

func main() {
	a := map[string]struct{}{}
	a["hoge"] = struct{}{}

	v := a["hoge"]
	fmt.Println(v)
	v = a["huge"]
	fmt.Println(v)
}
