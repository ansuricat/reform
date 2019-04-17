// Package dialects implements reform.Dialect selector.
package dialects

import (
	"github.com/ansuricat/reform"
	"github.com/ansuricat/reform/dialects/mssql"
	"github.com/ansuricat/reform/dialects/mysql"
	"github.com/ansuricat/reform/dialects/postgresql"
	"github.com/ansuricat/reform/dialects/sqlite3"
	"github.com/ansuricat/reform/dialects/sqlserver"
)

// ForDriver returns reform Dialect for given driver string, or nil.
func ForDriver(driver string) reform.Dialect {
	switch driver {
	case "postgres", "pgx":
		return postgresql.Dialect
	case "mysql":
		return mysql.Dialect
	case "sqlite3":
		return sqlite3.Dialect
	case "mssql":
		return mssql.Dialect
	case "sqlserver":
		return sqlserver.Dialect
	default:
		return nil
	}
}
