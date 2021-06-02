package initialize

import (
	"backend/cache"
	"backend/util"
	"os"

	"github.com/joho/godotenv"
)

// Init
func Init() {
	// load .env
	godotenv.Load()

	// set log level
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// connect database
	Database(os.Getenv("PSQL_DSN"))
	cache.Redis()
}
