package balance

import (
	"encoding/json"
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	file, err := os.Open("testData.json")
	if err != nil {
		t.Error(err)
	}
	decoder := json.NewDecoder(file)
	var data [][2]string
	if err := decoder.Decode(&data); err != io.EOF && err != nil {
		t.Error(err)
	}
	for _, v := range data {
		if answer := Balance(v[0]); answer != v[1] {
			t.Error(v, answer)
		} else {
			t.Log("Solved: ", answer)
		}
	}
}
