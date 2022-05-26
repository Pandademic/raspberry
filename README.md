# raspberry
tasty cli framework for go

# Features
------

- simple
- fast
- small

# Example Program
---
This program can spit out a variety of greeetings.
``` go
package main

import (
  berry "github.com/pandademic/raspberry"
  "fmt"
)
func hiCmd(){
  fmt.Println("hi")
}
func helloCmd(){
  fmt.Println("hello")
}
func main(){
  cli := berry.Cli{AcceptedCommands:[]string{"-v","version","-h","help","hi","hello"},HelpMsg:"Available commands:\n-h\nhelp\n-v\nversion\nhi\nhello",Version:0.1}
  cli.Setup()
  cli.SetHandler("hi",hiCmd)
  cli.SetHandler("hello",helloCmd)
}
```
# Documentation
----
Check the [wiki](https://github.com/Pandademic/raspberry/wiki)

# Licence
----
This program is licenced under the [MIT licence](https://github.com/Pandademic/raspberry/blob/main/LICENSE)
