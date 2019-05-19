package main

import "fmt"


func main(){
	// spefic the key type, and the value type
	var namePrefix map[string]string
	// we must initialize the map
	namePrefix = make(map[string]string)
	namePrefix["Daniel"] = "Mr "
	namePrefix["Naomy"] = "Ms "
	namePrefix["Florian"] = "Mr "

	var name string = "Daniel"

	fmt.Println(namePrefix[name])

	name = "Naomy"
	// shorthand notation
	prefix := map[string]string {
		"Daniel":"Mr ",
		"Florain":"Mr ",
		"Naomy":"Ms ",
	}

	fmt.Println(prefix[name])
	
	name = "Notin"
	fmt.Println(prefix[name])
}
