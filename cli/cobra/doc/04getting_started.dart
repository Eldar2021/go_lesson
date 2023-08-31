/*
Getting Started

While you are welcome to provide your own organization, 
typically a Cobra-based application will follow the following 
organizational structure:

▾ appName/
  ▾ cmd/
      add.go
      your.go
      commands.go
      here.go
    main.go

In a Cobra app, typically the main.go file is very bare. 
It serves one purpose: initializing Cobra.

```go
package main

import (
  "{pathToYourApp}/cmd"
)

func main() {
  cmd.Execute()
}
```
*/