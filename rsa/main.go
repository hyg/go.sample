// rsa project main.go
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	//"encoding/base64"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

//http://www.tuicool.com/articles/aMfIba

var pempub = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC/W8HgftMmG+nhg8pOWuWz/bLZ
SlIvGlDbNx2LQ1qkbBXObEfFVUUAAXnEb5NpUoVJKEj7RJzLOonB/x75ElG/QEiB
6Ct3cadbwB7aGFwqqBLp9wfceMR9gIzElepbwH3STUvavUjCgn9yr53AaipGvFlQ
0yrGWNSn5ZDttuk9XQIDAQAB
-----END PUBLIC KEY-----`)

var pemsec = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQC/W8HgftMmG+nhg8pOWuWz/bLZSlIvGlDbNx2LQ1qkbBXObEfF
VUUAAXnEb5NpUoVJKEj7RJzLOonB/x75ElG/QEiB6Ct3cadbwB7aGFwqqBLp9wfc
eMR9gIzElepbwH3STUvavUjCgn9yr53AaipGvFlQ0yrGWNSn5ZDttuk9XQIDAQAB
AoGASI7wWsF8Ks0Wx84DHebVhoRCFqZZt0aRNi4V48JsUkAxnI3uQOLuQOxOUa/F
A4CozW6bDbucgGr35TlpcyQBHo2Bg0kQHxA5qXhWa4sfVzpDUCt3nospnHaFKlzO
/lfEoaU/KxjMlViRirVJLeSgz4JEI0es+OQNkIlVca60LLkCQQDrpiHvRF1nUnh1
YkQL6oqFKBCtiha/eDpybsVvtRUC8S8SgD/UgZCr+T1rxTEml8tMQ+yDhjAaT3/1
F18tY5KTAkEAz+JsY+yPdPKnJmScIaLIPDFEesxFpMZstt1aoS1aSREqiO+9hhHH
Plt7zaTk6koZqZzNs1db0ogodz5KUSw2TwJBAJBjLw/IN+MDKUPjfgY/I7kLH4z1
u5J+PHG5ZchYkBNJbKpNYs72xIpbIUNThBY9lBea1uSP6BF2/NRUCcFp7XkCQQC3
9GvnzGhxm1vP/I2wsgQwR4SKiYJDOhbvhlbxc1mGeLtD66mxHsBJ7NhT9EthC2tE
DO51eaNWXIg6ZJOM2uu/AkEAvFvGU9gQMCMbOh7/l66zXnWg3casU3LG7Zl8HitT
DoLJA4hNI+WBcmM3zH+UZBCzBuTX7cjaKwtNjJsp15zvvA==
-----END RSA PRIVATE KEY-----
`)

// sig1 ==> "data"
// sig2 ==> "data new"
// use base64.StdEncoding.DecodeString() transfer to []byte
var sig1 = "lPmi6ToXl5ddN7o6Y34cxpbb0rnjuUnT90I8YIzAaORIh7LJTUSz5V9FrcQPIbYatgMBcFulj9TVzQ3u1x2fjFPLzlPd/3nRib84bVHpuNFYTmLwJBYp6WvZnTwwJ2iBXoHogyLGgo8ioUsxDumyJycF7CNNX0hJJfK0XesWvj8="
var sig2 = "J6UKX2f7mDLiwQgr8gNFlQDu6w43vfeBC7tmmEBu860zD6qyS1jDVPvo6QPiafegllBITC7E+MCCz5xQrVJjLsB7KZlXdEmHVAMVsjMKHl14OLs9I7RAToZwqimo+0NoJlgLevhpj/+WWWhbAyLzh2Agy61JU79dueY4tfeqbB8="

var msgv2utf8 = "UK9CI8av+GXuw7vpxy/vwg1oy+3fyD0J5d+DXtmqxdfHO+v0ZR5aPHrWOrNsCodPn30VgQ1lTRmQ6BjeKgEbywC7WBMtjnvSaVYMDDCZpWmvPN7zaa1xh1Fo/4QH2hdqixyfU6YwmWBON6P34IBknoJYKdHEvrkf809r0Y+Peqc="
var msgv15utf8 = "HzEY/n/B2K78w9TMDe0ZcmGIjtvrshsLFJDrpdDEWILVGwUywDhPzCoxCAzJycujP9+PE4rO4rd/3R/IAiYFIS37wLzTjVl4nRTuiLDFI7ldqUr5DNp2GFRxd/c5vfYkBzNeCBlGhai5alnUwwrafyPtU6nJcVtdNYlXvmOQXYE="
var msgv2asc = "rB5OhZaTOTeTgI2WZ/L1orjBoC6djDCTGXOsTLq3ZLjRSKyzV7wpD67813ifo84EYaxrshr+rMcQ+mhaAYWiEhn4rYb76k86vXWkWUIGvwsprVgllN8T2v3uHhUkc6BMLUgXRypuJtvOwE37m9hQ4g9Qg6oDjxwp8MMtOm9VPfE="
var msgv15asc = "mWWpp3LU8hcNTvLBhY+XJ0WrEazrnBi7teXzNOCVQbnDT2dcKuA5/y1GVP21I16RUINHHDt4YIglrjSkKe/QULVuiXTRGmsjE4PDtT3CzDam7LiDW6RhCXvugOg0MP2+SgKhlD06+WQFppEDTspQKVWZQp+3EL4oNu5CwiRaZ5c="

// import public key from pem format
func importPublicKeyfromPEM(pempub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pempub)
	//log.Print(block)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	//log.Print(pubInterface)
	pub := pubInterface.(*rsa.PublicKey)
	//log.Print(pub)
	return pub
}

// export public key to pem format
func exportPublicKeytoPEM(pub *rsa.PublicKey) []byte {
	b, _ := x509.MarshalPKIXPublicKey(pub)
	c := pem.Block{"PUBLIC KEY", nil, b}
	d := pem.EncodeToMemory(&c)
	//log.Print(string(d))

	return d
}

// import private key from pem format
func importPrivateKeyfromPEM(pemsec []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(pemsec)
	//log.Print(block)
	sec, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	//log.Print(sec)
	return sec
}

// export private key to pem format
func exportPrivateKeytoPEM(sec *rsa.PrivateKey) []byte {
	l := x509.MarshalPKCS1PrivateKey(sec)
	m := pem.Block{"RSA PRIVATE KEY", nil, l}
	n := pem.EncodeToMemory(&m)
	//log.Print(string(n))

	return n
}

// import private key from pem format
func importPrivateKeyfromEncryptedPEM(pemsec, password []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(pemsec)
	//log.Print(block)
	buf, _ := x509.DecryptPEMBlock(block, password)
	sec, _ := x509.ParsePKCS1PrivateKey(buf)
	//log.Print(sec)
	return sec
}

// export private key to pem format
func exportPrivateKeytoEncryptedPEM(sec *rsa.PrivateKey, password []byte) []byte {
	l := x509.MarshalPKCS1PrivateKey(sec)
	m, _ := x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", l, password, x509.PEMCipherAES256)
	n := pem.EncodeToMemory(m)
	//log.Print(string(n))

	return n
}

func VerifySign(pub *rsa.PublicKey, msg, sig []byte) error {
	t := sha1.New()
	//io.WriteString(t, msg)  // when msg is a string
	t.Write(msg)
	sum := t.Sum(nil)[:]
	//log.Print(sum)

	err := rsa.VerifyPKCS1v15(pub, crypto.SHA1, sum, sig)
	//log.Print(err)

	return err
}

func Sign(sec *rsa.PrivateKey, msg []byte) []byte {
	t := sha1.New()
	//io.WriteString(t, msg) // when msg is a string
	t.Write(msg)
	sum1 := t.Sum(nil)[:]

	sig, _ := rsa.SignPKCS1v15(rand.Reader, sec, crypto.SHA1, sum1)
	//log.Print(base64.StdEncoding.EncodeToString(sig))

	return sig
}

// msg <= 86 byte
// label should be empty string if send to C#
func EncryptV2(pub *rsa.PublicKey, msg, label []byte) ([]byte, error) {
	sha1 := sha1.New()
	return rsa.EncryptOAEP(sha1, rand.Reader, pub, msg, label)
}

// msg <= 117 byte
func EncryptV15(pub *rsa.PublicKey, msg []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
}

func DecryptV2(sec *rsa.PrivateKey, msg, label []byte) ([]byte, error) {
	sha1 := sha1.New()
	return rsa.DecryptOAEP(sha1, rand.Reader, sec, msg, label)
}

func DecryptV15(sec *rsa.PrivateKey, msg []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, sec, msg)
}

func main() {

	sec, _ := rsa.GenerateKey(rand.Reader, 1024)
	pub := &sec.PublicKey

	pempub = exportPublicKeytoPEM(pub)
	//pemsec = exportPrivateKeytoPEM(sec)
	pemsec = exportPrivateKeytoEncryptedPEM(sec, []byte("asdfgh"))
	fmt.Printf("%s", pemsec)

	pub = importPublicKeyfromPEM(pempub)
	//sec = importPrivateKeyfromPEM(pemsec)
	sec = importPrivateKeyfromEncryptedPEM(pemsec, []byte("asdfgh"))

	sig := Sign(sec, []byte("data"))

	err := VerifySign(pub, []byte("data"), sig)
	log.Print("err is nil ==> good sig: \t", err)

	msg := []byte("http://www.weibo.com/huangyg")
	secmsg, _ := EncryptV2(pub, msg, []byte(""))
	demsg, _ := DecryptV2(sec, secmsg, []byte(""))
	//secmsg, err := EncryptV15(pub, msg)
	//demsg, err := DecryptV15(sec, secmsg)

	log.Printf("Plaintext 1 :\t %s", msg)
	log.Printf("Plaintext 2 :\t %s", demsg)

}
