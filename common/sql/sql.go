package sql

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Map map[string]string

func (r Map) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Map) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}

type Strings []string

func (r Strings) Value() (driver.Value, error) {
	return Value(r)
}

func (r *Strings) Scan(input interface{}) error {
	return Scan(r, input)
}

func Value(r interface{}) (driver.Value, error) {
	return json.Marshal(r)
}

func Scan(r interface{}, input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}
