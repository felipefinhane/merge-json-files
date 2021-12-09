package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {
	jsonFile1, err := os.Open(os.Args[1])
	defer jsonFile1.Close()
	if err != nil {
		panic(err)
	}

	jsonFile2, err := os.Open(os.Args[2])
	defer jsonFile2.Close()
	if err != nil {
		panic(err)
	}

	jsonData1, err := ioutil.ReadAll(jsonFile1)
	if err != nil {
		panic(err)
	}

	jsonData2, err := ioutil.ReadAll(jsonFile2)
	if err != nil {
		panic(err)
	}

	var joined1 []map[string]interface{}
	if err := json.Unmarshal(jsonData1, &joined1); err != nil {
		panic(err)
	}

	var joined2 []map[string]interface{}
	if err := json.Unmarshal(jsonData2, &joined2); err != nil {
		panic(err)
	}

	joined1 = append(joined1, joined2[0])

	marshalled, err := json.Marshal(joined1)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(os.Args[1], marshalled, 0777)
	if err != nil {
		panic(err)
	}
}
