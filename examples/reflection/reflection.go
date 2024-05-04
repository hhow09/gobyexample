package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Part 1 - Law of reflections and Go reflect package basic operations
// Part 2 from 28:50 - Implementing a small JSON encoder using reflection

type User struct {
	Name string `en:"name" it:"nome"`
	Age  int64  `en:"age" it:"eta"`
}

type City struct {
	Name       string `en:"name" it:"nome"`
	Population int64  `en:"pop" it:"pop"`
	GDP        int64  `en:"gdp" it:"pil"`
	Mayor      string `en:"mayor" it:"sindaco"`
}

func main() {
	var x float64 = 3.14
	var u User = User{"bob", 10}

	// 1st law: you can go from interface value to reflection obj
	refObjValPtr := reflect.ValueOf(&x)
	refObjtyp := reflect.TypeOf(&x)
	fmt.Printf("ref obj val for × %s\n", refObjValPtr.String())
	fmt.Printf("ref obj typ for × %s\n", refObjtyp.String())
	// 2nd law: you can from reflection obj to interface value
	// deference the ptr to its actual value
	refObjVal := refObjValPtr.Elem()
	fmt.Printf("ref obj val for float64 x: %f\n", refObjVal.Float())
	fmt.Printf("can set new val for refl obj? %v\n", refObjVal.CanSet())
	// 3rd law when modifying reflection objects,
	// not all of them can be settable
	refObjVal.Set(reflect.ValueOf(4.25))
	fmt.Printf("updated × using ptr and reflection obj, now val %f \n", x)

	res, err := JSONEncode(u, "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	c := City{"sf", 5000000, 567896, "mr jones"}
	res, err = JSONEncode(c, "it")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func JSONEncode(v interface{}, tagKey string) ([]byte, error) {
	refObjVal := reflect.ValueOf(v)
	refObjTyp := reflect.TypeOf(v)
	buf := bytes.Buffer{}
	if refObjVal.Kind() != reflect.Struct {
		return buf.Bytes(), fmt.Errorf(
			"val of kind %s is not supported",
			refObjVal.Kind(),
		)
	}
	buf.WriteString("{")
	pairs := []string{}
	for i := 0; i < refObjVal.NumField(); i++ {
		structFieldRefObj := refObjVal.Field(i)
		structFieldRefObjTyp := refObjTyp.Field(i)

		tag := structFieldRefObjTyp.Tag.Get(tagKey)
		switch structFieldRefObj.Kind() {
		case reflect.String:
			strVal := structFieldRefObj.Interface().(string)
			pairs = append(pairs, `"`+tag+`":"`+strVal+`"`)
		case reflect.Int64:
			intVal := structFieldRefObj.Interface().(int64)
			pairs = append(pairs, `"`+tag+`":`+strconv.FormatInt(intVal, 10))
		default:
			return buf.Bytes(), fmt.Errorf(
				"struct field with name %s and kind %s is not supprted",
				structFieldRefObjTyp.Name,
				structFieldRefObj.Kind(),
			)
		}
	}

	buf.WriteString(strings.Join(pairs, ","))
	buf.WriteString("}")

	return buf.Bytes(), nil
}

// ref: [The Go Blog: The Laws of Reflection](https://go.dev/blog/laws-of-reflection)

// ref: [Go (Golang) Reflection Tutorial](https://youtu.be/f4aUrm7N5DU)
