package sql

import _ "embed"

//go:embed schema.sql
var Schema []byte

//go:embed seed.sql
var Seed []byte
