// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EmailSubscribersColumns holds the columns for the "email_subscribers" table.
	EmailSubscribersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "email_address", Type: field.TypeString},
	}
	// EmailSubscribersTable holds the schema information for the "email_subscribers" table.
	EmailSubscribersTable = &schema.Table{
		Name:       "email_subscribers",
		Columns:    EmailSubscribersColumns,
		PrimaryKey: []*schema.Column{EmailSubscribersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "emailsubscriber_app_id_email_address",
				Unique:  true,
				Columns: []*schema.Column{EmailSubscribersColumns[4], EmailSubscribersColumns[5]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EmailSubscribersTable,
	}
)

func init() {
}
