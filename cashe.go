package cash

import "errors"

var cash = make(map[string]interface{})

func Set(key string, value interface{}) {
	cash[key] = value
}

func Get(key string) (interface{}, bool) {
	_, ok := cash[key]
	if !ok {
		return nil, false
	}
	return cash[key], true
}

func Delete(key string) error {
	_, ok := cash[key]
	if !ok {
		err := errors.New("key not found")
		if err != nil {
			return err
		}
	}
	delete(cash, key)
	return nil

}
