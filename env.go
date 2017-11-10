package must

import (
	"fmt"
	"log"
	"os"
)

// GetEnv checks and returns an environment variable's value as a string if it
// exists or logs and exits the Go program.
func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal(fmt.Sprintf("Error, environment variable %s is not set.", key))
	}
	return value
}
