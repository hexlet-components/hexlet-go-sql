package migrations

import "embed"

// FS contains all goose migration files.
//
//go:embed *.sql
var FS embed.FS
