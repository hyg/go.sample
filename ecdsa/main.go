// ecdsa project main.go
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"io"
	"log"
)

func main() {
	c := elliptic.P521()
	sec, _ := ecdsa.GenerateKey(c, rand.Reader)
	pub := &sec.PublicKey

	t := sha1.New()
	io.WriteString(t, "data") // when msg is a string
	//t.Write([]byte("data")) // when msg is []bye
	sum1 := t.Sum(nil)[:]

	r, s, _ := ecdsa.Sign(rand.Reader, sec, sum1)
	log.Printf("r=%d\ts=%d", r, s)

	b := ecdsa.Verify(pub, sum1, r, s)
	log.Printf("b=%v", b)

	b = ecdsa.Verify(pub, sum1, s, r)
	log.Printf("b=%v", b)

}
