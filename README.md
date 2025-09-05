# env

dotenv loader the chad way.

## why `env` over x?

* 50~ lines of code
* 0-dep
* relies on go 1.1
* based code

## example

[source](/example)
```go
package main

import (
	"fmt"
	"os"

	"github.com/nxtgo/env"
)

func init() {
	err := env.Load("")
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(".env file is loaded :D")
	fmt.Println(os.Getenv("BASED"))
}
```

# license

under CC0 1.0 (public domain) + ip waiver.
