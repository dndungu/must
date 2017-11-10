# must
A tiny Golang package that will print to os.Stdout and exit(0) when certain conditions (mostly when err is not nil) are not met.

## must.BeNil
```go

result, err := DoSomething(value)
must.BeNil(err)

// would be equivalent to

result, err := DoSomething(value)
if err != nil {
  log.Fatal(err.Error())
}
```

## must.GetEnv
```go
 value := must.GetEnv("ENV_VARIABLE")

 // would be equivalent to

 value, ok := os.LookupEnv("ENV_VARIABLE")
 if !ok {
   log.Fatal(fmt.Sprintf("Error, environment variable %s is not set.", key))
 }
```
