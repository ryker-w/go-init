# go Application Shutdown

Shutdown the application with a signal.

## Install & Build

```bash
go get github.com/lishimeng/go-app-shutdown.git
```

## Usage
```go
package main

import (
 "fmt"
 "github.com/lishimeng/go-app-shutdown"
 "time"
)

func main() {
    go work()
    shutdown.WaitExit(&shutdown.Configuration{
        BeforeExit: func(s string) {
            fmt.Printf("Shutdown %s\n", s)
        },
    })
}

func work() {
    for{
        time.Sleep(time.Second * 3)
        fmt.Println("work....")
    }
}

```
