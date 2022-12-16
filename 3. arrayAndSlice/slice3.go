/*
Wrtie a function that counts the numbers of bits that are different in two SHA256 hashes
*/

package main

import (
	"crypto/sha256"
	"fmt"
)

func bigDiff(sha1, sha2 *[sha256.Size]byte) int {
	n := 0
	for i, _ := range sha1 {
		for b := sha1[i] ^ sha2[i]; b != 0; b &= b - 1 {
			n++
		}
	}
	return n
}

func main() {
	s1 := "understand"
	s2 := "UNDERSTAND"
	sha1 := sha256.Sum256([]byte(s1))
	sha2 := sha256.Sum256([]byte(s2))
	fmt.Printf("\n%x\n%x\n%t\n%T\n%T\n", sha1, sha2, sha1 == sha2, sha1, sha2)
	fmt.Println("Difference: ", bigDiff(&sha1, &sha2))
}
