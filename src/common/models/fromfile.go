package models

import files "phoenixnap.com/pnapctl/common/fileprocessor"

func CreateRequestFromFile[T any](filename string, commandname string) (*T, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var item T

	err = files.Unmarshal(data, &item, commandname)

	if err != nil {
		return nil, err
	}

	return &item, nil
}
