package migrations

import (
	"github.com/go-rel/rel"
)

// MigrateCreateSamples definition
func MigrateCreateSamples(schema *rel.Schema) {
	schema.CreateTable("todos", func(t *rel.Table) {
		t.ID("id")
	})
}

// RollbackCreateSamples definition
func RollbackCreateSamples(schema *rel.Schema) {
	schema.DropTable("todos")
}
