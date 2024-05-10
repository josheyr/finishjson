package main

import (
	"fmt"
	"github.com/josheyr/finishjson/pkg/finishjson"
)

func main() {
	var unfinishedJSON = `["hello", {"world": [4`

	var finishedJSON = finishjson.FinishJSON(unfinishedJSON)
	fmt.Println(finishedJSON)
}
