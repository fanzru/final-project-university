package debugger

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/TylerBrock/colorjson"
)

func Print(v interface{}) {
	log.Println("----------------------------------------------------", v)
	str := fmt.Sprintf("%v", v)
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
	log.Println("----------------------------------------------------")
}

func PrintFatal(v interface{}) {
	log.Println("----------------------------------------------------")
	str := fmt.Sprintf("%v", v)
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
	log.Println("----------------------------------------------------")
	panic(v)
}

func PrintJson(v interface{}, name string) {
	log.Println("----------------------------------------------------")
	fmt.Println("Object : ", name)
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
	log.Println("----------------------------------------------------")
}
