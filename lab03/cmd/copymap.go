package main

import "fmt"

func Copymap(m map[any]any) (result map[any]any) {
	result = make(map[any]any, len(m))
	for key, val := range m {
		result[key] = val
	}
	return
}

func main() {
	sourceMap := make(map[any]any)

	sourceMap["one"] = 1
	sourceMap["two"] = 2
	sourceMap["three"] = 3
	sourceMap["four"] = 4
	sourceMap["five"] = 5
	dest := Copymap(sourceMap)
	fmt.Printf("%v\n", dest)

}
