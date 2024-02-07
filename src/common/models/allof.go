package models

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
)

func typename[A any]() string {
	var a A
	return fmt.Sprintf("%T", a)
}

func GetFromAllOf[Inner any, AllOf any](allOf AllOf) *Inner {
	var inner Inner
	bin, err := json.Marshal(allOf)

	if err != nil {
		log.Err(err).Msg(fmt.Sprintf("Error when extracting JSON from (%s)\n", typename[AllOf]()))
		return nil
	}

	err = json.Unmarshal(bin, &inner)

	if err != nil {
		log.Err(err).Msg(fmt.Sprintf("Error when parsing (%s) from (%s)\n", typename[Inner](), typename[AllOf]()))
		return nil
	}

	return &inner
}
