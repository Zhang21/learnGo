package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Go offers built-in support for JSON encoding and decoding.
// Go的数据类型转换为对应JSON的数据类型

type response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encode/decode in JSON.
// Fields must start with capital(大写)letters to be exported.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// First we'll look at encoding basic data types to JSON strings.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Here are some for slices and maps, which encode to JSON arrays objects.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// The JSON package can automatically encode your custom data types.
	// It will only include exported fields in the encoded output and
	// will by default use those names as the JSON keys.
	// JSON的key这里是大写
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res11D := response1{
		Page:   11,
		Fruits: []string{"apple", "peach", "pear"}}
	res11B, _ := json.Marshal(res11D)
	fmt.Println(string(res11B))

	// You can use tags on struct field declarations to custom the encode JSON key names.
	// JSON的key这里是小写
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Decoding JSON data into Go values.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// We need to provide a variable where the JSON package can put the decoded data.
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	// In order to use the values in the decoded map, we'll need to convert them to appropriate type.
	num := dat["num"].(float64)
	fmt.Println(num)
	strs := dat["strs"].([]interface{})
	// Accessing nested data requires a series of conversions.
	str1 := strs[0].(string)
	fmt.Println(str1)

	// We can alson decode JSON into custom data types.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// We can alson stream JSON encodings directly to `os.Writer` like `os.Stdout` or HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
