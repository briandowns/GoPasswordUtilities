// Copyright 2014 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//
// Simple library for working with passwords in Go.
//

package GoPasswordUtilities

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"log"
	"regexp"
)

var (
	characters     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]~"
	passwordScores = map[int]string{
		0: "Horrible",
		1: "Weak",
		2: "Medium",
		3: "Strong",
		4: "Very Strong"}
)

type Password struct {
	Pass   string
	Length int
}

type PasswordComplexity struct {
	Length          int
	Score           int
	ContainsUpper   bool
	ContainsLower   bool
	ContainsNumber  bool
	ContainsSpecial bool
}

// Use this if you're not generating a new password.
func NewPassword(password string) *Password {
	p := Password{Pass: password, Length: len(password)}
	return &p
}

// Generate and return a password as a string and as a
// byte slice of the given length.
func GeneratePassword(length int) *Password {
	passwordBuffer := new(bytes.Buffer)
	randBytes := make([]byte, length)
	if _, err := rand.Read(randBytes); err == nil {
		for j := 0; j < length; j++ {
			tmpIndex := int(randBytes[j]) % len(characters)
			char := characters[tmpIndex]
			passwordBuffer.WriteString(string(char))
		}
	}
	p := &Password{Pass: passwordBuffer.String(),
		Length: len(passwordBuffer.String()),
	}
	return p
}

// Generate a MD5 sum for the given password.
func (p *Password) MD5() [16]byte {
	return md5.Sum([]byte(p.Pass))
}

// Generate a SHA256 sum for the given password.
func (p *Password) SHA256() [32]byte {
	return sha256.Sum256([]byte(p.Pass))
}

// Generate a SHA512 sum for the given password.
func (p *Password) SHA512() [64]byte {
	return sha512.Sum512([]byte(p.Pass))
}

// Get the length of the password.  This method is being put on the
// password struct in case someone decides not to do a complexity
// check.
func (p *Password) GetLength() int {
	return p.Length
}

// Parse the password and note its attributes.
func ProcessPassword(p *Password) (nil, error) {
	c := &PasswordComplexity{}
	matchLower := regexp.MustCompile(`[a-z]`)
	matchUpper := regexp.MustCompile(`[A-Z]`)
	matchNumber := regexp.MustCompile(`[0-9]`)
	matchSpecial := regexp.MustCompile(`[\!@\#\$\%\^\&\*\(\\\)\-_\=\+,\.\?\/\:\;{}\[\]~]`)

	if p.Length < 8 {
		log.Println("Password isn't long enough for evaluation.")
		return nil, "Password isn't long enough for evaluation."
	} else {
		c.Length = p.Length
	}

	if matchLower.MatchString(p.Pass) {
		c.ContainsLower = true
		c.Score += 1
	}
	if matchUpper.MatchString(p.Pass) {
		c.ContainsUpper = true
		c.Score += 1
	}
	if matchNumber.MatchString(p.Pass) {
		c.ContainsNumber = true
		c.Score += 1
	}
	if matchSpecial.MatchString(p.Pass) {
		c.ContainsSpecial = true
		c.Score += 1
	}

	return nil, nil
}

// Get the score of the password.
func (c *PasswordComplexity) GetScore() int {
	return c.Score
}

// Get whether the password contains an upper case letter.
func (c *PasswordComplexity) HasUpper() bool {
	return c.ContainsUpper
}

// Get whether the password contains a lower case letter.
func (c *PasswordComplexity) HasLower() bool {
	return c.ContainsLower
}

// Get whether the password contains a number.
func (c *PasswordComplexity) HasNumber() bool {
	return c.ContainsNumber
}

// Get whether the password contains a special character.
func (c *PasswordComplexity) HasSpecial() bool {
	return c.ContainsSpecial
}

// Return the rating for the password.
func (c *PasswordComplexity) ComplexityRating() string {
	return passwordScores[c.Score]
}
