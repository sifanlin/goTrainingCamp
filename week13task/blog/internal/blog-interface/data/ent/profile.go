// Code generated by entc, DO NOT EDIT.

package ent

import (
	"blog/internal/blog-interface/data/ent/profile"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Profile is the model entity for the Profile schema.
type Profile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// Qq holds the value of the "qq" field.
	Qq int `json:"qq,omitempty"`
	// Wechat holds the value of the "wechat" field.
	Wechat string `json:"wechat,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Img holds the value of the "img" field.
	Img string `json:"img,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Profile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case profile.FieldID, profile.FieldQq:
			values[i] = new(sql.NullInt64)
		case profile.FieldName, profile.FieldDesc, profile.FieldWechat, profile.FieldEmail, profile.FieldImg, profile.FieldAvatar:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Profile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Profile fields.
func (pr *Profile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case profile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case profile.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case profile.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				pr.Desc = value.String
			}
		case profile.FieldQq:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field qq", values[i])
			} else if value.Valid {
				pr.Qq = int(value.Int64)
			}
		case profile.FieldWechat:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wechat", values[i])
			} else if value.Valid {
				pr.Wechat = value.String
			}
		case profile.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pr.Email = value.String
			}
		case profile.FieldImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field img", values[i])
			} else if value.Valid {
				pr.Img = value.String
			}
		case profile.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				pr.Avatar = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Profile.
// Note that you need to call Profile.Unwrap() before calling this method if this Profile
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Profile) Update() *ProfileUpdateOne {
	return (&ProfileClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Profile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Profile) Unwrap() *Profile {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Profile is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Profile) String() string {
	var builder strings.Builder
	builder.WriteString("Profile(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", desc=")
	builder.WriteString(pr.Desc)
	builder.WriteString(", qq=")
	builder.WriteString(fmt.Sprintf("%v", pr.Qq))
	builder.WriteString(", wechat=")
	builder.WriteString(pr.Wechat)
	builder.WriteString(", email=")
	builder.WriteString(pr.Email)
	builder.WriteString(", img=")
	builder.WriteString(pr.Img)
	builder.WriteString(", avatar=")
	builder.WriteString(pr.Avatar)
	builder.WriteByte(')')
	return builder.String()
}

// Profiles is a parsable slice of Profile.
type Profiles []*Profile

func (pr Profiles) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
