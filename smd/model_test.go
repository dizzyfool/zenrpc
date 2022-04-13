package smd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPropertyList_UnmarshalJSON(t *testing.T) {
	list := PropertyList{
		Property{
			Name: "Z",
			Type: "string",
		},
		Property{
			Name: "X",
			Type: "string",
		},
		Property{
			Name: "Y",
			Type: "string",
		},
	}

	for i := 0; i < 1000; i++ {
		list = append(list, Property{Name: fmt.Sprintf("name_%d", i)})
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })

	data, err := json.Marshal(list)
	if err != nil {
		t.Fatalf("json marshall error: %s", err)
	}

	check := PropertyList{}
	err = json.Unmarshal(data, &check)
	if err != nil {
		t.Fatalf("json unmarshall error: %s", err)
	}

	for i := range list {
		if list[i].Name != check[i].Name {
			t.Fatalf("order error: list: %s; check: %s; i: %d", list[i].Name, check[i].Name, i)
		}
	}
}
