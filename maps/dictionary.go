package maps

import "errors"

type Dictionary map[string]string

var ErrNotfound = errors.New("could not find he word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	defintion, ok := d[word]
	if !ok {
		return "", ErrNotfound
	}
	return defintion, nil
}
