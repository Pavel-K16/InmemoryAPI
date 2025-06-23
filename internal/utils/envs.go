package utils

import (
	"os"
	"taskapi/internal/logger"
)

var (
	log = logger.LoggerInit()
)

func DefaultEnvs() map[string]string {
	return map[string]string{
		"TASKAPI_LISTEN_PORT":            "8080",
		"TASKAPI_HTTP_READ_TIMEOUT_SEC":  "15",
		"TASKAPI_HTTP_WRITE_TIMEOUT_SEC": "15",
		"SYNC_TASKSTATUS_INTERVAL_SEC":   "3",
	}
}

func GetEnvs() map[string]string {
	res := make(map[string]string)

	envs := DefaultEnvs()

	for key, val := range envs {
		if incValue := os.Getenv(key); incValue != "" {
			res[key] = incValue
		} else {
			res[key] = val
		}
	}

	return res
}

func SetEnvs(envs map[string]string) error {
	for key, val := range envs {
		if err := os.Setenv(key, val); err != nil {
			log.Errorf("Error setting env %s: %s", key, err)

			return err
		}
	}

	return nil
}
