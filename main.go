package main

import (
	"fmt"
	"github.com/josheyr/finishjson/pkg/finishjson"
)

func main() {
	var unfinishedJSON = `{
  "person": {
    "name": "Alice",
    "age": 30,
    "address": {
      "city": "Wonderland"
    },
    "hobbies": ["reading", "gaming`

	var finishedJSON = finishjson.FinishJSON(unfinishedJSON)
	fmt.Println(finishedJSON)
}
