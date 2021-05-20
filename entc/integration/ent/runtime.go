// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"net"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/card"
	"entgo.io/ent/entc/integration/ent/fieldtype"
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/group"
	"entgo.io/ent/entc/integration/ent/groupinfo"
	"entgo.io/ent/entc/integration/ent/pet"
	"entgo.io/ent/entc/integration/ent/schema"
	"entgo.io/ent/entc/integration/ent/task"
	"entgo.io/ent/entc/integration/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cardMixin := schema.Card{}.Mixin()
	cardMixinFields0 := cardMixin[0].Fields()
	_ = cardMixinFields0
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescCreateTime is the schema descriptor for create_time field.
	cardDescCreateTime := cardMixinFields0[0].Descriptor()
	// card.DefaultCreateTime holds the default value on creation for the create_time field.
	card.DefaultCreateTime = cardDescCreateTime.Default.(func() time.Time)
	// cardDescUpdateTime is the schema descriptor for update_time field.
	cardDescUpdateTime := cardMixinFields0[1].Descriptor()
	// card.DefaultUpdateTime holds the default value on creation for the update_time field.
	card.DefaultUpdateTime = cardDescUpdateTime.Default.(func() time.Time)
	// card.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	card.UpdateDefaultUpdateTime = cardDescUpdateTime.UpdateDefault.(func() time.Time)
	// cardDescBalance is the schema descriptor for balance field.
	cardDescBalance := cardFields[0].Descriptor()
	// card.DefaultBalance holds the default value on creation for the balance field.
	card.DefaultBalance = cardDescBalance.Default.(float64)
	// cardDescNumber is the schema descriptor for number field.
	cardDescNumber := cardFields[1].Descriptor()
	// card.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	card.NumberValidator = cardDescNumber.Validators[0].(func(string) error)
	// cardDescName is the schema descriptor for name field.
	cardDescName := cardFields[2].Descriptor()
	// card.NameValidator is a validator for the "name" field. It is called by the builders before save.
	card.NameValidator = cardDescName.Validators[0].(func(string) error)
	fieldtypeFields := schema.FieldType{}.Fields()
	_ = fieldtypeFields
	// fieldtypeDescValidateOptionalInt32 is the schema descriptor for validate_optional_int32 field.
	fieldtypeDescValidateOptionalInt32 := fieldtypeFields[15].Descriptor()
	// fieldtype.ValidateOptionalInt32Validator is a validator for the "validate_optional_int32" field. It is called by the builders before save.
	fieldtype.ValidateOptionalInt32Validator = fieldtypeDescValidateOptionalInt32.Validators[0].(func(int32) error)
	// fieldtypeDescMAC is the schema descriptor for mac field.
	fieldtypeDescMAC := fieldtypeFields[27].Descriptor()
	// fieldtype.MACValidator is a validator for the "mac" field. It is called by the builders before save.
	fieldtype.MACValidator = fieldtypeDescMAC.Validators[0].(func(string) error)
	// fieldtypeDescDir is the schema descriptor for dir field.
	fieldtypeDescDir := fieldtypeFields[31].Descriptor()
	// fieldtype.DefaultDir holds the default value on creation for the dir field.
	fieldtype.DefaultDir = fieldtypeDescDir.Default.(func() http.Dir)
	// fieldtypeDescNdir is the schema descriptor for ndir field.
	fieldtypeDescNdir := fieldtypeFields[32].Descriptor()
	// fieldtype.NdirValidator is a validator for the "ndir" field. It is called by the builders before save.
	fieldtype.NdirValidator = fieldtypeDescNdir.Validators[0].(func(string) error)
	// fieldtypeDescStr is the schema descriptor for str field.
	fieldtypeDescStr := fieldtypeFields[33].Descriptor()
	// fieldtype.DefaultStr holds the default value on creation for the str field.
	fieldtype.DefaultStr = fieldtypeDescStr.Default.(func() sql.NullString)
	// fieldtypeDescNullStr is the schema descriptor for null_str field.
	fieldtypeDescNullStr := fieldtypeFields[34].Descriptor()
	// fieldtype.DefaultNullStr holds the default value on creation for the null_str field.
	fieldtype.DefaultNullStr = fieldtypeDescNullStr.Default.(func() *sql.NullString)
	// fieldtypeDescLink is the schema descriptor for link field.
	fieldtypeDescLink := fieldtypeFields[35].Descriptor()
	// fieldtype.LinkValidator is a validator for the "link" field. It is called by the builders before save.
	fieldtype.LinkValidator = fieldtypeDescLink.Validators[0].(func(string) error)
	// fieldtypeDescIP is the schema descriptor for ip field.
	fieldtypeDescIP := fieldtypeFields[41].Descriptor()
	// fieldtype.DefaultIP holds the default value on creation for the ip field.
	fieldtype.DefaultIP = fieldtypeDescIP.Default.(func() net.IP)
	// fieldtypeDescPair is the schema descriptor for pair field.
	fieldtypeDescPair := fieldtypeFields[54].Descriptor()
	// fieldtype.DefaultPair holds the default value on creation for the pair field.
	fieldtype.DefaultPair = fieldtypeDescPair.Default.(func() schema.Pair)
	// fieldtypeDescVstring is the schema descriptor for vstring field.
	fieldtypeDescVstring := fieldtypeFields[56].Descriptor()
	// fieldtype.DefaultVstring holds the default value on creation for the vstring field.
	fieldtype.DefaultVstring = fieldtypeDescVstring.Default.(func() schema.VString)
	// fieldtypeDescTriple is the schema descriptor for triple field.
	fieldtypeDescTriple := fieldtypeFields[57].Descriptor()
	// fieldtype.DefaultTriple holds the default value on creation for the triple field.
	fieldtype.DefaultTriple = fieldtypeDescTriple.Default.(func() schema.Triple)
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescSize is the schema descriptor for size field.
	fileDescSize := fileFields[0].Descriptor()
	// file.DefaultSize holds the default value on creation for the size field.
	file.DefaultSize = fileDescSize.Default.(int)
	// file.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	file.SizeValidator = fileDescSize.Validators[0].(func(int) error)
	filetypeFields := schema.FileType{}.Fields()
	_ = filetypeFields
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescActive is the schema descriptor for active field.
	groupDescActive := groupFields[0].Descriptor()
	// group.DefaultActive holds the default value on creation for the active field.
	group.DefaultActive = groupDescActive.Default.(bool)
	// groupDescType is the schema descriptor for type field.
	groupDescType := groupFields[2].Descriptor()
	// group.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	group.TypeValidator = func() func(string) error {
		validators := groupDescType.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_type string) error {
			for _, fn := range fns {
				if err := fn(_type); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// groupDescMaxUsers is the schema descriptor for max_users field.
	groupDescMaxUsers := groupFields[3].Descriptor()
	// group.DefaultMaxUsers holds the default value on creation for the max_users field.
	group.DefaultMaxUsers = groupDescMaxUsers.Default.(int)
	// group.MaxUsersValidator is a validator for the "max_users" field. It is called by the builders before save.
	group.MaxUsersValidator = groupDescMaxUsers.Validators[0].(func(int) error)
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[4].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = func() func(string) error {
		validators := groupDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	groupinfoFields := schema.GroupInfo{}.Fields()
	_ = groupinfoFields
	// groupinfoDescMaxUsers is the schema descriptor for max_users field.
	groupinfoDescMaxUsers := groupinfoFields[1].Descriptor()
	// groupinfo.DefaultMaxUsers holds the default value on creation for the max_users field.
	groupinfo.DefaultMaxUsers = groupinfoDescMaxUsers.Default.(int)
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescAge is the schema descriptor for age field.
	petDescAge := petFields[0].Descriptor()
	// pet.DefaultAge holds the default value on creation for the age field.
	pet.DefaultAge = petDescAge.Default.(float64)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescPriority is the schema descriptor for priority field.
	taskDescPriority := taskFields[0].Descriptor()
	// task.DefaultPriority holds the default value on creation for the priority field.
	task.DefaultPriority = schema.Priority(taskDescPriority.Default.(int))
	// task.PriorityValidator is a validator for the "priority" field. It is called by the builders before save.
	task.PriorityValidator = taskDescPriority.Validators[0].(func(int) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescOptionalInt is the schema descriptor for optional_int field.
	userDescOptionalInt := userMixinFields0[0].Descriptor()
	// user.OptionalIntValidator is a validator for the "optional_int" field. It is called by the builders before save.
	user.OptionalIntValidator = userDescOptionalInt.Validators[0].(func(int) error)
	// userDescLast is the schema descriptor for last field.
	userDescLast := userFields[2].Descriptor()
	// user.DefaultLast holds the default value on creation for the last field.
	user.DefaultLast = userDescLast.Default.(string)
	// userDescAddress is the schema descriptor for address field.
	userDescAddress := userFields[4].Descriptor()
	// user.DefaultAddress holds the default value on creation for the address field.
	user.DefaultAddress = userDescAddress.Default.(func() string)
}
