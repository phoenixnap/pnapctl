package models

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

// Gets the name of the type A passed.
func typename[A any]() string {
	var a A
	return fmt.Sprintf("%T", a)
}

// Extracts an Inner type from the AllOf type by marshaling and unmarshaling.
// For example, if Inner type has a field "Country", it should be populated by
// the "Country" field from the AllOf type.
func GetFromAllOf[Inner any, AllOf any](allOf AllOf) *Inner {
	var inner Inner
	bin, err := json.Marshal(allOf)

	if err != nil {
		log.Err(err).Msgf("Error when extracting JSON from (%s)\n", typename[AllOf]())
		return nil
	}

	err = json.Unmarshal(bin, &inner)

	if err != nil {
		// The unmarshaling step usually fails due to invalid validation. In most cases, this is due to enums.
		// If tests are panicking - this is usually the culprit.
		log.Err(err).Msgf("Error when parsing (%s) from (%s)\n", typename[Inner](), typename[AllOf]())
		return nil
	}

	return &inner
}
