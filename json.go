/**
exercising built-in support for JSON encoding and decoding, including to and from built-in and custom data types
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// using response1 and 2 structs to demonstrate encoding and decoding of custom types below
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// first looking at encoding basic data types to JSON strings.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(1.0)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// usage for slices and maps, which encode to json arrays and objects
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// json package can automatically encode your custom data types. It will only include exported fields in the encoded
	// output, and will by default use those names and the json keys
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// tags can be used on struct field declarations to customize the encoded json key names. check the definition of
	// response2 above to see an example of such tags
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// decoding json data into Go values

	// some generic data structure
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// need to provide a variable where the json package can put the decoded data. this map[string]interface{} will
	// hold a map of strings to arbitrary data types
	var dat map[string]interface{}

	// the actual decoding, and a check for associated errors
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// in order to use the values in the decoded map, we'll need to cast them to their appropriate type. for example,
	// here we cast the value in num to the expected float64 type
	num := dat["num"].(float64)
	fmt.Println(num)

	// accessing nested data requires a series of casts
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// we can also decode json into custom data types. this has the advantages of adding additional type-safety to our
	// programs and eliminating the need for type assertions when accessing the decoded data
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// in the examples above we always used bytes and strings as intermediates between the data and json representation
	// on standard out. we can also stream json encodings directly to os.Writers like os.Stdout or even HTTP response
	// bodies
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
