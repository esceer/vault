package setup

import (
	"fmt"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/sqlite"
	"github.com/esceer/vault/backend/cmd/config"
	dbmigration "github.com/esceer/vault/backend/db-migration"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDB(cfg *config.Config) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(cfg.DataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func RunMigrationScripts(cfg *config.Config) error {
	connStr := fmt.Sprintf("sqlite:%v", cfg.DataSource)
	u, _ := url.Parse(connStr)
	db := dbmate.New(u)
	db.AutoDumpSchema = false
	db.FS = dbmigration.MigrationsFs
	return db.CreateAndMigrate()
}
