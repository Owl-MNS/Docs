package main

import (
	"documentation/models"
	"encoding/json"
	"fmt"
)

func main() {
	m := models.Package{}
	mJSON, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(mJSON))
}
