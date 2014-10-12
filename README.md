# GoPasswordUtilities

[![GoDoc](https://godoc.org/github.com/briandowns/GoPasswordUtilities?status.svg)](https://godoc.org/github.com/briandowns/GoPasswordUtilities) [![Build Status](https://travis-ci.org/briandowns/GoPasswordUtilities.svg?branch=master)](https://travis-ci.org/briandowns/GoPasswordUtilities)

Simple library for working with passwords in Go (Golang).  

Complexity check will check for upper case letters, lower case letters, numbers, special characters 
and also whether the password is dictionary based. 

For more detail about the library and its features, reference your local godoc once installed.

```bash
godoc -http=:6060
```

## Features 

- Generate a Password
- Run a Complexity Check
- Hash a Password
- Salt a Password

## Warning

This library is in alpha and will be subject to change.  Use with caution.  There's also a damn good chance it 
could be in a broken state at times as well.

## Installation

```bash
go get github.com/briandowns/GoPasswordUtilities
```

## Examples

### Import the package, generate a password and hash it.

```Go
package main

import (
	"fmt"
	gpu "github.com/briandowns/GoPasswordUtilities"
)

func main() {
	p := gpu.GeneratePassword(10)
	fmt.Println(p)
	fmt.Printf("%x\n", p.MD5())
}
```

### Create a new password object and get its information

```Go
    pass := gpu.New("secret12")
    fmt.Println(*pass)
    
    results, err := gpu.ProcessPassword(pass)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("Has Rating: %s\n", results.ComplexityRating())
    fmt.Printf("%t\n", results.DictionaryBased)
```

### Generate thousands of passwords.

```Go
    // On the fly compile and execution.  Better 
    ///once statically compiled.
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
```

### Generate a Very Strong password.
 
```Go
    p := gpu.GenerateVeryStrongPassword(10)
    fmt.Println(p.Pass)
```

### Hash a password that includes a generated salt.

```Go
    p := gpu.GeneratePassword(10)
    fmt.Println(p.Pass)
    hash1, _ := p.SHA256()
    hash2, salt := p.SHA256(&gpu.SaltConf{Length: 32})
    fmt.Printf("%x\n", hash1)
    fmt.Printf("%x - %x\n", hash2, salt)   
```
