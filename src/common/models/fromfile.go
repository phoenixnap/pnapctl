package models

import (
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

func CreateRequestFromFile[T any](filename string) (*T, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var item T

	err = files.Unmarshal(data, &item)

	if err != nil {
		return nil, err
	}

	return &item, nil
}
