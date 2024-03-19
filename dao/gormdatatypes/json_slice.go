package gormdatatypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type JSONSlice []string

// Value return json value, implement driver.Valuer interface
func (j JSONSlice) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	ba, err := j.MarshalJSON()
	return string(ba), err
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSONSlice) Scan(val interface{}) error {
	if val == nil {
		*j = make(JSONSlice, 0)
		return nil
	}
	var ba []byte
	switch v := val.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}
	t := []string{}
	err := json.Unmarshal(ba, &t)
	*j = JSONSlice(t)
	return err
}

// MarshalJSON to output non base64 encoded []byte
func (j JSONSlice) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	t := ([]string)(j)
	return json.Marshal(t)
}

// UnmarshalJSON to deserialize []byte
func (j *JSONSlice) UnmarshalJSON(b []byte) error {
	t := []string{}
	err := json.Unmarshal(b, &t)
	*j = JSONSlice(t)
	return err
}

// GormDataType gorm common data type
func (j JSONSlice) GormDataType() string {
	return "jsonslice"
}

// GormDBDataType gorm db data type
func (JSONSlice) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

func (j *JSONSlice) Has(item string) bool {
	for _, i := range *j {
		if i == item {
			return true
		}
	}
	return false
}

func (j *JSONSlice) Append(items ...string) {
	for _, item := range items {
		*j = append(*j, item)
	}
}
