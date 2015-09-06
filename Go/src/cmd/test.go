package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lib/util"
)

// read json file to map
func readFile(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}

	return bytes, nil
}

// Global xxxx
type Global struct {
	Owner string
	Group string
	Mode  string
}

// Special define special mode
type Special map[string]string

//Detail define source --> destination
type Detail map[string]string

type apiR struct {
	Global  Global
	Special Special
	Detail  Detail
}

func main() {
	body, _ := readFile("./distribute.json")
	var r apiR

	if err := json.Unmarshal([]byte(body), &r); err != nil {
		fmt.Printf("err was %v", err)
	}

	// Global parameter
	for key, val := range r.Special {
		fmt.Println(key, val)
	}

	if err := util.CheckUG("administrator"); err != nil {
		fmt.Println("Error", err)
	}

}
