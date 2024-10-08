package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Name string   `json:"name"`
	Type int      `json:"type"`
	Info []string `jsong:"info"`
}

func main() {

	d := &Data{
		"Dude",
		42,
		[]string{"one", "two", "three"},
	}
	fmt.Println(d)

	j, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		_ = fmt.Errorf("error: %s", err)
	}

	fmt.Println(string(j))

	enc := json.NewEncoder(os.Stdout)
	//enc.Encode(map[string]string{"one": "thing", "two": "shoes"})
	enc.Encode(d)

}
