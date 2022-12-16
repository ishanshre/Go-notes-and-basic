package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("y"))
	fmt.Printf("%x\n%x\n%t\n%T\n%T\n", c1, c2, c1 == c2, c1, c2)
}

/*
This is a example
The function of sha256 in the crypto/sha256 package produces the cryptographic hash or digest of a message stored in
an arbitary byte slice. The digest has 256 bits, so its type is [32]byte ---> [32]unit8>
If the two digest is same then the two messages are highly like to be same.
If the digests are different then the message are different

*/
