package main

import "fmt"

/*
Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.

var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}

create map using make() function

var a = make(map[KeyType]ValueType)
b := make(map[KeyType]ValueType)
*/

func main() {
	fmt.Println("")

	grades := map[string]int{
		"Physics":  89,
		"Comp Sci": 99,
	}

	fmt.Println(grades)

	// initialized empty map object
	info := map[string]string{}

	// add values to map
	info["name"] = "benaiah"
	info["age"] = "20"

	fmt.Println(info)

	// acessing map data
	fmt.Println(info["name"])

	fmt.Println("")
}
