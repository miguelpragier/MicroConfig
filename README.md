# MicroConfig
Simple Configuration File Engine

The goal is to load and set application config constants from a json file that respects our specific format.

Installation:
```bash
go get github.com/miguelpragier/microconfig
```

## Here's the config file format, with some dummy sample pairs

```json
{
  "pairs":[
    {
      "key": "Environment",
      "value": "dev"
    },
    {
      "key": "SomeKey",
      "value": "Some Value"
    },
    {
      "key": "MyAge",
      "value": "43"
    },
    {
      "key": "RandomFloatValue",
      "value": "78.910"
    },
    {
      "value": "DatabaseConnectionString",
      "key":  "databaseuser:a3232323bc4343d5454e656576876899f97@tcp(dbs.databaseserv.er)/mydatabase",
      "osEnv":true
    }
  ]
}
```

When osEnv is present, the constant is loaded in OS ENVIRONMENT variables as well.

## Gross useCase example considering the above example config file

```golang
package main

import (
	"github.com/miguelpragier/microconfig"
	"fmt"
	"os"
	"strings"
)

func main(){
  const myConfigFile = "./config.dev.json"
  
  if err:=microconfig.Load(myConfigFile); err!=nil {
    panic(err)
  }
  
  databaseConnectionString := os.GetEnv("DatabaseConnectionString")
  
  if databaseConnectionString=="" {
    panic("OMG! We have no database directions...")
  }
  
  appName, err := microconfig.GetString("SomeKey")
  
  if err!=nil {
    panic(err)
  }
  
  fmt.Printf("Expected: %v, Checked: %v", true, microconfig.Exists("Environment",false))
  
  // Note that Exists will search with caseInsensitive==false, resulting in a notFound/negative return
  fmt.Printf("Expected: %v, Checked: %v", false, microconfig.Exists("ENVIRONMENT",false))
  
  myIntegerConfigVariable,_ := microconfig.GetInt("MyAge")
  
  myFloat64ConfigVariable,_ := microconfig.GetFloat("RandomFloatValue")
}
```
