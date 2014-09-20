# GoPasswordUtilities

Simple library for working with passwords in Go (Golang).

For more detail about the library and its features, reference your local godoc once installed.

## Documentation

```bash
godoc -http=:6060
```

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
	
	// Create a new password object
    pass := gpu.NewPassword("secret12")
    fmt.Println(*pass)
	
	// Generate 10,000 passwords
    // On the fly compile and execution.  Better once
    // statically compiled.
    // 0.19s user 0.07s system 84% cpu 0.303 total
    passChan := make(chan *gpu.Password, 10000)
    go func() {
        for i := 0; i < 10000; i++ {
            passChan <- gpu.GeneratePassword(8)
        }
        close(passChan)
    }()

    for pass := range passChan {
        fmt.Printf("%s\n", pass.Pass)
    }
    	
	// Get all the password info
    results, err := gpu.ProcessPassword(p)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("Has Rating: %s\n", results.ComplexityRating())
    
    // Generate a Very Strong password
    p := gpu.GenerateVeryStrongPassword(10)
    fmt.Println(p.Pass)
    
    // Hash a password that includes a salt
    p := gpu.GeneratePassword(10)
    fmt.Println(p)
    fmt.Printf("%x\n", p.MD5())
    fmt.Printf("%x\n", p.MD5(&gpu.SaltConf{Length: 32}))    
}
```