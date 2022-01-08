package schema

import (
	"database/sql/driver"
	"encoding/json"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"fmt"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

type Property struct {
}

func (p Property) Value() (driver.Value, error) {
	jsonStr, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return jsonStr, nil
}

func (p *Property) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("Property should not be nil")
	}
	switch value := src.(type) {
	default:
		return fmt.Errorf("database type error %T, expected []byte", src)
	case []byte:
		err := json.Unmarshal(value, &p)
		if err != nil {
			return err
		}
	}
	return nil
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("property", Property{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
