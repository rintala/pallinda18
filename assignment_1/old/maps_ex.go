package main

import ("fmt")

func main(){
	m := map[string]int{
		"Key1": 32,
		"Key2": 27,
	}

	v, ok := m["nonexistingkey"]
	fmt.Println(v)
	fmt.Println(ok)
}


