package config

import (
	"os"
	"strconv"
)

var (
	HostZync     = getOrDefaultS("INDEXER_ZYNC_HOST", "http://localhost:4080/")
	UserZync     = getOrDefaultS("INDEXER_ZYNC_HOST", "admin")
	PassZync     = getOrDefaultS("INDEXER_ZYNC_PASS", "admin123")
	ChunkReader  = getOrDefaultI("INDEXER_READER_SIZECHUNK", 500)
	DefaultIndex = getOrDefaultS("INDEXER_ZYNC_NAME_INDEX", "indexEmail")
)

func getOrDefaultS(osEnv, defaultValue string) string {
	value := os.Getenv(osEnv)
	if value == "" {
		value = defaultValue
	}
	return value
}

func getOrDefaultI(osEnv string, defaultValue int) int {
	value, error := strconv.Atoi(os.Getenv(osEnv))
	if error != nil {
		value = defaultValue
	}
	return value
}
