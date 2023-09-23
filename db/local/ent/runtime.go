// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/Creaft-JP/tit/db/local/ent/commit"
	"github.com/Creaft-JP/tit/db/local/ent/committedfile"
	"github.com/Creaft-JP/tit/db/local/ent/page"
	"github.com/Creaft-JP/tit/db/local/ent/remote"
	"github.com/Creaft-JP/tit/db/local/ent/schema"
	"github.com/Creaft-JP/tit/db/local/ent/section"
	"github.com/Creaft-JP/tit/db/local/ent/stagedfile"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commitFields := schema.Commit{}.Fields()
	_ = commitFields
	// commitDescNumber is the schema descriptor for number field.
	commitDescNumber := commitFields[0].Descriptor()
	// commit.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	commit.NumberValidator = commitDescNumber.Validators[0].(func(int) error)
	// commitDescMessage is the schema descriptor for message field.
	commitDescMessage := commitFields[1].Descriptor()
	// commit.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	commit.MessageValidator = commitDescMessage.Validators[0].(func(string) error)
	committedfileFields := schema.CommittedFile{}.Fields()
	_ = committedfileFields
	// committedfileDescPath is the schema descriptor for path field.
	committedfileDescPath := committedfileFields[0].Descriptor()
	// committedfile.PathValidator is a validator for the "path" field. It is called by the builders before save.
	committedfile.PathValidator = committedfileDescPath.Validators[0].(func(string) error)
	pageFields := schema.Page{}.Fields()
	_ = pageFields
	// pageDescPathname is the schema descriptor for pathname field.
	pageDescPathname := pageFields[0].Descriptor()
	// page.PathnameValidator is a validator for the "pathname" field. It is called by the builders before save.
	page.PathnameValidator = pageDescPathname.Validators[0].(func(string) error)
	// pageDescNumber is the schema descriptor for number field.
	pageDescNumber := pageFields[1].Descriptor()
	// page.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	page.NumberValidator = pageDescNumber.Validators[0].(func(int) error)
	remoteFields := schema.Remote{}.Fields()
	_ = remoteFields
	// remoteDescURL is the schema descriptor for url field.
	remoteDescURL := remoteFields[1].Descriptor()
	// remote.URLValidator is a validator for the "url" field. It is called by the builders before save.
	remote.URLValidator = remoteDescURL.Validators[0].(func(string) error)
	sectionFields := schema.Section{}.Fields()
	_ = sectionFields
	// sectionDescSlug is the schema descriptor for slug field.
	sectionDescSlug := sectionFields[0].Descriptor()
	// section.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	section.SlugValidator = sectionDescSlug.Validators[0].(func(string) error)
	// sectionDescTitle is the schema descriptor for title field.
	sectionDescTitle := sectionFields[1].Descriptor()
	// section.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	section.TitleValidator = sectionDescTitle.Validators[0].(func(string) error)
	// sectionDescNumber is the schema descriptor for number field.
	sectionDescNumber := sectionFields[3].Descriptor()
	// section.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	section.NumberValidator = sectionDescNumber.Validators[0].(func(int) error)
	stagedfileFields := schema.StagedFile{}.Fields()
	_ = stagedfileFields
	// stagedfileDescPath is the schema descriptor for path field.
	stagedfileDescPath := stagedfileFields[0].Descriptor()
	// stagedfile.PathValidator is a validator for the "path" field. It is called by the builders before save.
	stagedfile.PathValidator = stagedfileDescPath.Validators[0].(func(string) error)
}
