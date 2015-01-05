env
===

Go implementation for default values for environment variables.

~~~ go
package main

import "github.com/mikebeyer/env"

func main() {
  port := env.Get("PORT", "8080")
}
~~~
