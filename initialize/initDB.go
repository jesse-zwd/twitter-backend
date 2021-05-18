package initialize

import (
	"backend/global"
	"backend/util"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database init postgres connection
func Database(connString string) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	
	// Error
	if err != nil {
		util.Log().Panic("database connection failed", err)
	}

	//setting connection pool
	sqlDB, err := db.DB()
	if err != nil {
		util.Log().Panic("sqlDB failed", err)
	}

	//idle
	sqlDB.SetMaxIdleConns(50)
	//open connections
	sqlDB.SetMaxOpenConns(100)
	//timeout
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	global.GORM_DB = db

	migration()
}
