//ref: https://gobyexample.com/json

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type response3 struct {
	Name     string    `json:"name"`
	Capacity int       `json:"capacity"`
	Time     time.Time `json:"time"`
}

func main() {
	// marshell boolean -> []byte
	bolB, _ := json.Marshal(true)
	fmt.Println("true marshalled: ", string(bolB))

	// marshell int -> []byte
	intB, _ := json.Marshal(1)
	fmt.Println("int 1 marshalled: ", string(intB))

	// marshell float -> []byte
	fltB, _ := json.Marshal(2.34)
	fmt.Printf("float %f marshalled: %s\n", 2.34, string(fltB))

	// marshell string -> []byte
	strB, _ := json.Marshal("gopher")
	fmt.Println("gopher marshelled", strB)

	// marshell []string -> []byte
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("slcD", string(slcB))

	// marshell map -> []byte
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("mapD", string(mapB))

	// marshell struct -> []byte
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("res1D", string(res1B))

	// marshell struct with struct json tag -> []byte
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("res2B", string(res2B))

	fmt.Println("=== json.Unmarshal === ")
	rawstr := `{"num":6.13,"strs":["a","b"]}`
	byt := []byte(rawstr)

	// 1. unmarshal to unstructured data interface{}
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Printf("raw json: %s, unmarsheled to: %v\n", rawstr, dat)

	// need to cast the type because we don't know the type in advance
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 2. Unmarshal to struct
	jsonStr := `{"page": 1, "fruits": ["apple", "peach"], "will": "ignore"}`
	res := response2{}
	json.Unmarshal([]byte(jsonStr), &res)

	// field will be ignored if not defined in struct
	fmt.Printf("unmarshel raw json: %s to struct: \n%v", jsonStr, res)
	fmt.Println(res.Fruits[0])

	// 3. Unmarshal data with timestamp
	jsonStr = `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	data3 := response3{}
	if err := json.Unmarshal([]byte(jsonStr), &data3); err != nil {
		panic(err)
	}
	fmt.Printf("unmarshel raw json: %s to struct: \n%v\n", jsonStr, data3)
	fmt.Println("time field in milli", data3.Time.UnixMilli())

	// 4. encode to writer
	enc := json.NewEncoder(os.Stdout) // write to stdout
	d := map[string]int{"apple": 5, "lettuce": 7}
	fmt.Println("4. encode to writer: ")
	enc.Encode(d)
}
