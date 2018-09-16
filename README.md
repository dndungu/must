# must
A tiny Golang package that will print to os.Stdout and exit(0) when certain conditions (mostly when err is not nil) are not met.

## must.BeNil
```go

result, err := DoSomething()

must.BeNil(err)

// is functionally equivalent to

import "log"

result, err := DoSomething()
if err != nil {
	log.Fatal(err.Error())
}
```

## must.GetEnv
```go
key := "ENV_VARIABLE"

value := must.GetEnv(key)

// is functionally equivalent to

import "log"

value, ok := os.LookupEnv(key)
if !ok {
	log.Fatal(fmt.Sprintf("Error, environment variable %s is not set.", key))
}
```

## must.ReadFile
```go
path := "/secrets/MY_TOKEN"

value := must.ReadFile(path)

// is functionally equivalent to

import (
	"io/ioutil"
	"log"
	"strings"
)

b, err := ioutil.ReadFile(path)
if err != nil {
	log.Fatal(err.Error())
}
value := strings.TrimSpace(string(b))
```
