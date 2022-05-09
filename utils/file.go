package utils

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
)

var (
	ErrFailReadingFile       = errors.New("fail reading file")
	ErrFailUnmarshallingFile = errors.New("fail unmarshalling file")
	ErrFailSavingFile        = errors.New("fail saving file")
	ErrFailMarshallingValue  = errors.New("fail marshalling value")
)

func GetObjectFromJSONFile[Type any](path string) (*Type, error) {
	var result Type

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, ErrFailReadingFile
	}

	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, ErrFailUnmarshallingFile
	}

	return &result, nil
}

func SaveFile[Type any](path string, object Type) error {
	marshalValue, err := json.Marshal(object)
	if err != nil {
		return ErrFailMarshallingValue
	}

	err = ioutil.WriteFile(path, marshalValue, fs.ModeAppend)
	if err != nil {
		return ErrFailSavingFile
	}

	return nil
}
