package setup

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/sqlite"
	"github.com/esceer/vault/cmd/config"
	dbmigration "github.com/esceer/vault/db-migration"
)

func ConnectToDB(cfg *config.Config) (*sql.DB, error) {
	return sql.Open(cfg.DatabaseDriver, cfg.DataSource)
}

func RunMigrationScripts(cfg *config.Config) error {
	connStr := fmt.Sprintf("sqlite:%v", cfg.DataSource)
	u, _ := url.Parse(connStr)
	db := dbmate.New(u)
	db.AutoDumpSchema = false
	db.FS = dbmigration.MigrationsFs
	return db.CreateAndMigrate()
}
