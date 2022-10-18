package models

import (
	"encoding/json"
	"fmt"
)

func typename[A any]() string {
	var a A
	return fmt.Sprintf("%T", a)
}

func GetFromAllOf[Inner any, AllOf any](allOf AllOf) *Inner {
	var inner Inner
	bin, err := json.Marshal(allOf)

	if err != nil {
		// TODO replace with log
		// fmt.Printf("Error when extracting JSON from (%s): %v\n", typename[AllOf](), err)
		return nil
	}

	err = json.Unmarshal(bin, &inner)

	if err != nil {
		// TODO replace with log
		// fmt.Printf("Error when parsing (%s) from (%s): %v\n", typename[Inner](), typename[AllOf](), err)
		return nil
	}

	return &inner
}
