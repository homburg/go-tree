# Tree

[![Build Status](https://travis-ci.org/homburg/go-tree.svg?branch=master)](https://travis-ci.org/homburg/go-tree)

Create an text-based tree (like `$ tree`) from any text
using a custom delimiter.

## CLI

https://github.com/homburg/tree/

## USAGE

```go
package main

import (
	"fmt"
	"github.com/homburg/tree"
)

func main() {

	t := tree.New("-")

	t.EatLines([]string{
		"root"
		"root-first",
		"root-first-second.txt",
		"third",
	})

	fmt.Println(t.Format())
	// Output:
	// root
	// └─── first
	//      └── second.txt
	// third.txt
}
```

## LICENSE(S)

MIT 2016 Thomas B Homburg

[LICENSE.md](LICENSE.md)
[LICENSES.md](LICENSES.md)

