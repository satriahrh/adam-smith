package helper

import "fmt"

type ErrorDictionaryItem struct {
	From  int
	Value string
	Error error
}

func ErrorParseFromDictionary(err error, dictionary []ErrorDictionaryItem) error {
	for _, item := range dictionary {
		if StringContainsFrom(err.Error(), item.Value, item.From) {
			return item.Error
		}
	}
	return fmt.Errorf("unknown error")
}
