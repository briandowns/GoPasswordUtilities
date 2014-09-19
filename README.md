# GoPasswordUtilities

Simple library for working with passwords in Go.

## Warning

This library is in alpha and will be subject to change.  Use with caution.  There's also a damn good chance it 
could be in a broken state at times as well.

## Installation

```bash
go get github.com/bdowns328/GoPasswordUtilities
```

## Example

```Go
package main

import (
	"fmt"
	gpu "github.com/bdowns328/GoPasswordUtilities"
)

func main() {
    // Generate a password and hash it.
	p := gpu.GeneratePassword(10)
	fmt.Println(p)
	fmt.Printf("%x\n", p.MD5())
}
```