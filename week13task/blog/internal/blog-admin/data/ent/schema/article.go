package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("title"),
		field.Int32("cid"),
		field.String("desc"),
		field.String("content"),
		field.String("img"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
