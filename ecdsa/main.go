// ecdsa project main.go
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"io"
	"log"
	"os"
)

// import public key from pem format
func importPublicKeyfromPEM(pempub []byte) *ecdsa.PublicKey {
	block, _ := pem.Decode(pempub)
	//log.Print(block)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	//log.Print(pubInterface)
	pub := pubInterface.(*ecdsa.PublicKey)
	//log.Print(pub)
	return pub
}

// export public key to pem format
func exportPublicKeytoPEM(pub *ecdsa.PublicKey) []byte {
	b, _ := x509.MarshalPKIXPublicKey(pub)
	c := pem.Block{"EC PUBLIC KEY", nil, b}
	d := pem.EncodeToMemory(&c)
	//log.Print(string(d))

	return d
}

// import private key from pem format
func importPrivateKeyfromPEM(pemsec []byte) *ecdsa.PrivateKey {
	block, _ := pem.Decode(pemsec)
	//log.Print(block)
	sec, _ := x509.ParseECPrivateKey(block.Bytes)
	//log.Print(sec)
	return sec
}

// export private key to pem format
func exportPrivateKeytoPEM(sec *ecdsa.PrivateKey) []byte {
	l, _ := x509.MarshalECPrivateKey(sec)
	m := pem.Block{"EC PRIVATE KEY", nil, l}
	n := pem.EncodeToMemory(&m)

	keypem, _ := os.OpenFile("sec.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(keypem, &pem.Block{Type: "EC PRIVATE KEY", Bytes: l})
	//log.Print(string(n))

	return n
}

// import private key from pem format
func importPrivateKeyfromEncryptedPEM(pemsec, password []byte) *ecdsa.PrivateKey {
	block, _ := pem.Decode(pemsec)
	//log.Print(block)
	buf, _ := x509.DecryptPEMBlock(block, password)
	sec, _ := x509.ParseECPrivateKey(buf)
	//log.Print(sec)
	return sec
}

// export private key to pem format
func exportPrivateKeytoEncryptedPEM(sec *ecdsa.PrivateKey, password []byte) []byte {
	l, _ := x509.MarshalECPrivateKey(sec)
	m, _ := x509.EncryptPEMBlock(rand.Reader, "EC PRIVATE KEY", l, password, x509.PEMCipherAES256)
	n := pem.EncodeToMemory(m)
	//log.Print(string(n))

	keypem, _ := os.OpenFile("sec.Encrypted.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(keypem, &pem.Block{Type: "EC PRIVATE KEY", Bytes: l})

	return n
}

func main() {
	c := elliptic.P521()
	sec, _ := ecdsa.GenerateKey(c, rand.Reader)
	pub := &sec.PublicKey
	log.Print("pub", pub)
	log.Print("sec", sec)

	pempub := exportPublicKeytoPEM(pub)
	pemsec := exportPrivateKeytoEncryptedPEM(sec, []byte("asdfgh"))
	log.Print("pempub", pempub)
	log.Print("pemsec", pemsec)

	pub = importPublicKeyfromPEM(pempub)
	//sec = importPrivateKeyfromPEM(pemsec)
	sec = importPrivateKeyfromEncryptedPEM(pemsec, []byte("asdfgh"))
	log.Print("pub", pub)
	log.Print("sec", sec)

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
