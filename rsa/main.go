// rsa project main.go
package main

import (
	//"bytes"
	//"crypto/aes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/pem"
	//"encoding/binary"
	//"encoding/xml"
	"log"
	//"math/big"
	"crypto/x509"
	//"io/ioutil"
	//	"os"
	"fmt"
	"io"
)

//http://www.tuicool.com/articles/aMfIba

const pubstr = "<?xml version=\"1.0\"?> <RSAParameters xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\">   <Exponent>AQAB</Exponent>   <Modulus>202wRQxUPtO4COUkCVmSde6n+UqjtZroTNJWgL+MF910U7HxGl7Y2DBpon0iGDr2B82nxhJ88CWNetVT2bv6x7BweBxrhK605hAjNU3iKLWa5SeNS0YaGTDVCC2qthE80k7h3fCFw6jQ8MY2soXFPWRIROzdisLqmSjxhHDuOdc=</Modulus> </RSAParameters>"
const str = "DWiGR3Enw56In+qC2bGeWCueSSrSetqChUDkqRBGcSKZO4ojZfgwiW1z0rF9ed6OFRb/ABy+lvzjc3MeY3VWNvIQx1H+PCBQ3hy9qb24s9KnR2jiv73c5ROfXtwzyhMOeJcS3rHaLWwcv/bWjfKU6TU94p9sx1+FggJoqtB/0zc="

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

type Pubkey struct {
	Exponent string `xml:"Exponent"`
	Modulus  string `xml:"Modulus"`
}

func main() {

	block, _ := pem.Decode(pempub)
	if block == nil {
		log.Fatal("public key error")
	}
	//log.Print(block)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(pubInterface)
	}
	//log.Print(pubInterface)
	pub := pubInterface.(*rsa.PublicKey)
	//log.Print(pub)
	//rsa.EncryptPKCS1v15(rand.Reader, pub, origData)

	if true {

		msg := "data new"
		msg1 := []byte(msg)

		t := sha1.New()
		io.WriteString(t, msg)

		//sum1 := []byte(fmt.Sprint(t.Sum(nil)))
		sum1 := t.Sum(nil)[:]

		sum2 := []byte(fmt.Sprint(sha1.Sum(msg1)))
		//sig, _ := base64.StdEncoding.DecodeString("lPmi6ToXl5ddN7o6Y34cxpbb0rnjuUnT90I8YIzAaORIh7LJTUSz5V9FrcQPIbYatgMBcFulj9TVzQ3u1x2fjFPLzlPd/3nRib84bVHpuNFYTmLwJBYp6WvZnTwwJ2iBXoHogyLGgo8ioUsxDumyJycF7CNNX0hJJfK0XesWvj8=")
		sig, _ := base64.StdEncoding.DecodeString("J6UKX2f7mDLiwQgr8gNFlQDu6w43vfeBC7tmmEBu860zD6qyS1jDVPvo6QPiafegllBITC7E+MCCz5xQrVJjLsB7KZlXdEmHVAMVsjMKHl14OLs9I7RAToZwqimo+0NoJlgLevhpj/+WWWhbAyLzh2Agy61JU79dueY4tfeqbB8=")
		log.Print(sum1)
		log.Print(sum2)
		log.Print(sig)

		err := rsa.VerifyPKCS1v15(pub, crypto.SHA1, sum1, sig)
		log.Print(err)

		return
	}

	//b := pem.Block{"PUBLIC KEY", nil, []byte(pub)}
	b, _ := x509.MarshalPKIXPublicKey(pub)
	c := pem.Block{"PUBLIC KEY", nil, b}
	d := pem.EncodeToMemory(&c)
	log.Print(string(d))

	block, _ = pem.Decode(pemsec)
	//log.Print(block)
	if block == nil {
		log.Fatal("public key error")
	}
	sec, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	log.Print(sec)

	if true {

		msg := "data new"
		t := sha1.New()
		io.WriteString(t, msg)

		sum1 := t.Sum(nil)[:]

		s, err := rsa.SignPKCS1v15(rand.Reader, sec, crypto.SHA1, sum1)
		log.Print(base64.StdEncoding.EncodeToString(s))
		log.Print(err)

		return
	}

	l := x509.MarshalPKCS1PrivateKey(sec)
	m := pem.Block{"RSA PRIVATE KEY", nil, l}
	n := pem.EncodeToMemory(&m)
	log.Print(string(n))

	return

	sec, _ = rsa.GenerateKey(rand.Reader, 1024)
	pub = &sec.PublicKey

	//sha1 := sha1.New()

	msg := []byte("DecryptPKCS1v15SessionKey decrypts a session key using RSA and the padding scheme from PKCS#1 v1.5. If rand != nil, it uses RSA blinding to avoid timing side-channel attacks. It returns an error if the ciphertext is the wrong length or if the ciphertext is greater than the public modulus. Otherwise, no error is returned. If the padding is valid, the resulting plaintext message is copied into key. Otherwise, key is unchanged. These alternatives occur in constant time. It is intended that the user of this function generate a random session key beforehand and continue the protocol with the resulting value. This will remove any possibility that an attacker can learn any information about the plaintext. See “Chosen Ciphertext Attacks Against Protocols Based on the RSA Encryption Standard PKCS #1”, Daniel Bleichenbacher, Advances in Cryptology (Crypto '98).")
	//secmsg, _ := rsa.EncryptOAEP(sha1, rand.Reader, pub, msg, []byte(""))
	secmsg, err := rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
	//secmsg, _ = base64.StdEncoding.DecodeString("UK9CI8av+GXuw7vpxy/vwg1oy+3fyD0J5d+DXtmqxdfHO+v0ZR5aPHrWOrNsCodPn30VgQ1lTRmQ6BjeKgEbywC7WBMtjnvSaVYMDDCZpWmvPN7zaa1xh1Fo/4QH2hdqixyfU6YwmWBON6P34IBknoJYKdHEvrkf809r0Y+Peqc=")
	//secmsg, _ = base64.StdEncoding.DecodeString("HzEY/n/B2K78w9TMDe0ZcmGIjtvrshsLFJDrpdDEWILVGwUywDhPzCoxCAzJycujP9+PE4rO4rd/3R/IAiYFIS37wLzTjVl4nRTuiLDFI7ldqUr5DNp2GFRxd/c5vfYkBzNeCBlGhai5alnUwwrafyPtU6nJcVtdNYlXvmOQXYE=")
	//secmsg, _ = base64.StdEncoding.DecodeString("rB5OhZaTOTeTgI2WZ/L1orjBoC6djDCTGXOsTLq3ZLjRSKyzV7wpD67813ifo84EYaxrshr+rMcQ+mhaAYWiEhn4rYb76k86vXWkWUIGvwsprVgllN8T2v3uHhUkc6BMLUgXRypuJtvOwE37m9hQ4g9Qg6oDjxwp8MMtOm9VPfE=")
	//secmsg, _ = base64.StdEncoding.DecodeString("mWWpp3LU8hcNTvLBhY+XJ0WrEazrnBi7teXzNOCVQbnDT2dcKuA5/y1GVP21I16RUINHHDt4YIglrjSkKe/QULVuiXTRGmsjE4PDtT3CzDam7LiDW6RhCXvugOg0MP2+SgKhlD06+WQFppEDTspQKVWZQp+3EL4oNu5CwiRaZ5c=")

	//log.Print("secmsg\t", base64.StdEncoding.EncodeToString(secmsg))
	log.Print("secmsg\t", secmsg)
	log.Print(err)

	//demsg, _ := rsa.DecryptOAEP(sha1, rand.Reader, sec, secmsg, []byte(""))
	demsg, err := rsa.DecryptPKCS1v15(rand.Reader, sec, secmsg)
	log.Print("demsg\t", string(demsg))
	log.Print(err)

	/*
		var pubkey Pubkey
		err := xml.Unmarshal([]byte(pubstr), &pubkey)

		if err != nil {
			log.Fatal(err)
		}

		log.Print("pubkey\t", pubkey)
		e, _ := base64.StdEncoding.DecodeString(pubkey.Exponent)
		n, _ := base64.StdEncoding.DecodeString(pubkey.Modulus)
		var impkey rsa.PublicKey

		log.Print("n\t", n)

		b := []byte{0x00, 0x22, 0x33, 0x44}
		b[1] = e[0]
		b[2] = e[1]
		b[3] = e[2]
		b_buf := bytes.NewBuffer(b)
		var x int32
		binary.Read(b_buf, binary.BigEndian, &x)

		impkey.E = int(x)
		impkey.N = *big.Int(str)

		log.Print("impkey\t", impkey)
		log.Print("impkey.E\t", impkey.E)
		log.Print("impkey.N\t", impkey.N)

		sha1 := sha1.New()
		msg := []byte("EncryptPKCS1v15 encrypts the given message with RSA and the padding scheme from PKCS#1 v1.5. The message must be no longer than the length of the public modulus minus 11 bytes. WARNING: use of this function to encrypt plaintexts other than session keys is dangerous. Use RSA OAEP in new protocols.")
		secstr, _ := rsa.EncryptOAEP(sha1, rand.Reader, &impkey, msg, nil)

		log.Print(secstr)

		return
	*/

	/*
		seck, _ := rsa.GenerateKey(rand.Reader, 1024)

		pubstr := base64.StdEncoding.EncodeToString([]byte(seck.PublicKey.E))
		log.Print(pubstr)
		pubstr := base64.StdEncoding.EncodeToString([]byte(seck.PublicKey.N))
		log.Print(pubstr)
		//var b bytes.Buffer
		//xmlstr, _ := xml.Marshal(&seck)
		//log.Print("length\t", seck.PublicKey.N.BitLen())
		//log.Print("PublicKey.N\t", seck.PublicKey.N)
		/*
			for i := 0; i < 128; i++ {
				log.Print("for loop\t", i, seck.PublicKey.N[i])
			}

			//xmlstr := x509.MarshalPKCS1PrivateKey(seck)

			/*
				log.Print("xml\t", xmlstr)

				log.Print("seck\t", seck)
				log.Print("PublicKey\t", seck.PublicKey)
				log.Print("D\t", seck.D)
				log.Print("E\t", seck.E)
				log.Print("N\t", seck.N)
				log.Print("Primes\t", seck.Primes)
				log.Print("Precomputed\t", seck.Precomputed)
	*/
}
