# GoKey

## Package based on [Unkey's Article](https://unkey.dev/blog/uuid-ux) around key generation

### Installation
--- 
```bash
go get -u github.com/stneto1/gokey
```

### Usage
```go
package main

import (
    "log"

    "github.com/stneto1/gokey"
)

func main() {
    key := gokey.MustGenerateKey("sk", 10)
    log.Println("Generated key", key) // Generated key sk_550e8400e2
}
```