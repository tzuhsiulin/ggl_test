package utils

import "os"

func IsProdEnv() bool {
	return os.Getenv("ENV") == "prod"
}
