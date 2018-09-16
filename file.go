package must

import (
	"io/ioutil"
	"strings"
)

// ReadFile will read a file using ioutil.ReadFile and return the contents
// as a string.
func ReadFile(path string) string {
	b, err := ioutil.ReadFile(path)
	BeNil(err)
	return strings.TrimSpace(string(b))
}
