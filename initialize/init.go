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

	// load translation file
	if err := LoadLocales("initialize/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("translation file loading failed!", err)
	}

	// connect database
	Database(os.Getenv("PSQL_DSN"))
	cache.Redis()
}
