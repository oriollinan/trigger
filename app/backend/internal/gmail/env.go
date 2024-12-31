package gmail

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
** This defines the expected enviroment variables for the Gmail Service
 */

var (
	errEnvNotFound      string     = "Enviroment argument %s not found"
	enviromentArguments [17]string = [...]string{
		"MONGO_INITDB_ROOT_USERNAME",
		"MONGO_INITDB_ROOT_PASSWORD",
		"MONGO_PORT",
		"MONGO_HOST",
		"MONGO_DB",

		"AUTH_SERVICE_BASE_URL",
		"USER_SERVICE_BASE_URL",
		"SESSION_SERVICE_BASE_URL",
		"SYNC_SERVICE_BASE_URL",
		"ACTION_SERVICE_BASE_URL",
		"GMAIL_SERVICE_BASE_URL",
		"SETTINGS_SERVICE_BASE_URL",

		"ADMIN_TOKEN",

		"GMAIL_PORT",
		"AUTH_PORT",
		"GOOGLE_CLIENT_ID",
		"GOOGLE_CLIENT_SECRET",
	}
)

func Env(envPath string) error {
	err := godotenv.Load(envPath)

	if err != nil {
		return err
	}
	for _, envArg := range enviromentArguments {
		v := os.Getenv(envArg)

		if v != "" {
			continue
		}
		return fmt.Errorf(errEnvNotFound, envArg)
	}
	return nil
}
