// Code generated by entc, DO NOT EDIT.

package profile

import (
	"blog/internal/blog-admin/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDesc), v))
	})
}

// Qq applies equality check predicate on the "qq" field. It's identical to QqEQ.
func Qq(v int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQq), v))
	})
}

// Wechat applies equality check predicate on the "wechat" field. It's identical to WechatEQ.
func Wechat(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWechat), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Img applies equality check predicate on the "img" field. It's identical to ImgEQ.
func Img(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImg), v))
	})
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDesc), v))
	})
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDesc), v))
	})
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDesc), v...))
	})
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDesc), v...))
	})
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDesc), v))
	})
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDesc), v))
	})
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDesc), v))
	})
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDesc), v))
	})
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDesc), v))
	})
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDesc), v))
	})
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDesc), v))
	})
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDesc), v))
	})
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDesc), v))
	})
}

// QqEQ applies the EQ predicate on the "qq" field.
func QqEQ(v int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQq), v))
	})
}

// QqNEQ applies the NEQ predicate on the "qq" field.
func QqNEQ(v int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQq), v))
	})
}

// QqIn applies the In predicate on the "qq" field.
func QqIn(vs ...int) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldQq), v...))
	})
}

// QqNotIn applies the NotIn predicate on the "qq" field.
func QqNotIn(vs ...int) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldQq), v...))
	})
}

// QqGT applies the GT predicate on the "qq" field.
func QqGT(v int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQq), v))
	})
}

// QqGTE applies the GTE predicate on the "qq" field.
func QqGTE(v int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQq), v))
	})
}

// QqLT applies the LT predicate on the "qq" field.
func QqLT(v int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQq), v))
	})
}

// QqLTE applies the LTE predicate on the "qq" field.
func QqLTE(v int) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQq), v))
	})
}

// WechatEQ applies the EQ predicate on the "wechat" field.
func WechatEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWechat), v))
	})
}

// WechatNEQ applies the NEQ predicate on the "wechat" field.
func WechatNEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWechat), v))
	})
}

// WechatIn applies the In predicate on the "wechat" field.
func WechatIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWechat), v...))
	})
}

// WechatNotIn applies the NotIn predicate on the "wechat" field.
func WechatNotIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWechat), v...))
	})
}

// WechatGT applies the GT predicate on the "wechat" field.
func WechatGT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWechat), v))
	})
}

// WechatGTE applies the GTE predicate on the "wechat" field.
func WechatGTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWechat), v))
	})
}

// WechatLT applies the LT predicate on the "wechat" field.
func WechatLT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWechat), v))
	})
}

// WechatLTE applies the LTE predicate on the "wechat" field.
func WechatLTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWechat), v))
	})
}

// WechatContains applies the Contains predicate on the "wechat" field.
func WechatContains(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWechat), v))
	})
}

// WechatHasPrefix applies the HasPrefix predicate on the "wechat" field.
func WechatHasPrefix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWechat), v))
	})
}

// WechatHasSuffix applies the HasSuffix predicate on the "wechat" field.
func WechatHasSuffix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWechat), v))
	})
}

// WechatEqualFold applies the EqualFold predicate on the "wechat" field.
func WechatEqualFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWechat), v))
	})
}

// WechatContainsFold applies the ContainsFold predicate on the "wechat" field.
func WechatContainsFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWechat), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// ImgEQ applies the EQ predicate on the "img" field.
func ImgEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImg), v))
	})
}

// ImgNEQ applies the NEQ predicate on the "img" field.
func ImgNEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImg), v))
	})
}

// ImgIn applies the In predicate on the "img" field.
func ImgIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImg), v...))
	})
}

// ImgNotIn applies the NotIn predicate on the "img" field.
func ImgNotIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImg), v...))
	})
}

// ImgGT applies the GT predicate on the "img" field.
func ImgGT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImg), v))
	})
}

// ImgGTE applies the GTE predicate on the "img" field.
func ImgGTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImg), v))
	})
}

// ImgLT applies the LT predicate on the "img" field.
func ImgLT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImg), v))
	})
}

// ImgLTE applies the LTE predicate on the "img" field.
func ImgLTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImg), v))
	})
}

// ImgContains applies the Contains predicate on the "img" field.
func ImgContains(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImg), v))
	})
}

// ImgHasPrefix applies the HasPrefix predicate on the "img" field.
func ImgHasPrefix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImg), v))
	})
}

// ImgHasSuffix applies the HasSuffix predicate on the "img" field.
func ImgHasSuffix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImg), v))
	})
}

// ImgEqualFold applies the EqualFold predicate on the "img" field.
func ImgEqualFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImg), v))
	})
}

// ImgContainsFold applies the ContainsFold predicate on the "img" field.
func ImgContainsFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImg), v))
	})
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvatar), v))
	})
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAvatar), v...))
	})
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.Profile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Profile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAvatar), v...))
	})
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvatar), v))
	})
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvatar), v))
	})
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvatar), v))
	})
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvatar), v))
	})
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAvatar), v))
	})
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAvatar), v))
	})
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAvatar), v))
	})
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAvatar), v))
	})
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAvatar), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Profile) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Profile) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Profile) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		p(s.Not())
	})
}
