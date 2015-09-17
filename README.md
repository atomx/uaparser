
# uaparser

### [![GoDoc](https://godoc.org/github.com/atomx/uaparser?status.png)](https://godoc.org/github.com/atomx/uaparser)

### Example usage

```go
package main

import (
  "fmt"
  "github.com/atomx/uaparser"
)

func main() {
  id, version := uaparser.Browser("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:40.0) Gecko/20100101 Firefox/40.0")

  major, minor := uaparser.Unversion(version)

  fmt.Printf("%s %d.%d", uaparser.Browsers[id], major, minor)
}
```
