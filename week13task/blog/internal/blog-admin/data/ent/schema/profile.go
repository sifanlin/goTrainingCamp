package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.String("desc"),
		field.Int("qq"),
		field.String("wechat"),
		field.String("email"),
		field.String("img"),
		field.String("avatar"),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return nil
}
