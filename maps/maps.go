package maps

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("could not find word in dictionary")

func (d Dictionary) Search(key string) (string, error) {
	// check validity of the key
	definition, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, definition string) error {
	d[key] = definition
	return nil
}
