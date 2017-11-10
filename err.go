package must

import (
	"log"
)

// BeNil checks if an err value is nil and does nothing else it exits the
// Go program's process.
func BeNil(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
