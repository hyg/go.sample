// openpgp project main.go
package main

import (
	"crypto/openpgp"
	"crypto/openpgp/armor"
	"fmt"
	//"github.com/howeyc/gopass"

	"os"
	"time"
)

func getKeyByEmail(keyring openpgp.EntityList, email string) *openpgp.Entity {
	for _, entity := range keyring {
		for _, ident := range entity.Identities {
			if ident.UserId.Email == email {
				return entity
			}
		}
	}

	return nil
}

// NewEntity creates a new entity. It doesn't provide an option for comments.
func NewEntity(name, email, outFile string) (ne *openpgp.Entity, err error) {
	ne, err = openpgp.NewEntity(name, "", email, nil)
	if err != nil {
		return
	}

	out, err := os.Create(outFile)
	if err != nil {
		ne = nil
		return
	}

	hdr := map[string]string{
		"Version": "Keybase Go client (OpenPGP version v0.1)",
	}

	keyOut, err := armor.Encode(out, openpgp.PrivateKeyType, hdr)
	if err != nil {
		ne = nil
		return
	}

	defer func() {
		keyOut.Close()
		out.Close()
	}()

	err = ne.SerializePrivate(keyOut, nil)
	if err != nil {
		ne = nil
		return
	}

	return
}

func main() {
	//NewEntity("hyg", "huangyg@mars22.com", "secfile.key")
	//return

	pubringFile, _ := os.Open("path to public keyring")
	defer pubringFile.Close()
	pubring, _ := openpgp.ReadKeyRing(pubringFile)
	//theirPublicKey := getKeyByEmail(pubring, "someone@xuemen.com")
	theirPublicKey := getKeyByEmail(pubring, "someone@their.org")

	secringFile, _ := os.Open("path to private keyring")
	defer secringFile.Close()
	sevring, _ := openpgp.ReadKeyRing(secringFile)
	myPrivateKey := getKeyByEmail(sevring, "huangyg@xuemen.com")

	//theirPublicKey.Serialize(os.Stdout)
	//myPrivateKey.Serialize(os.Stdout)
	//myPrivateKey.SerializePrivate(os.Stdout, nil)

	myPrivateKey.PrivateKey.Decrypt([]byte("passphrase"))
	/*
		// bug: have to input the correct passphrase at the first time

		for myPrivateKey.PrivateKey.Encrypted {
			fmt.Print("PGP passphrase: ")
			pgppass := gopass.GetPasswd()

			myPrivateKey.PrivateKey.Decrypt([]byte(pgppass))
			if myPrivateKey.PrivateKey.Encrypted {
				fmt.Println("Incorrect. Try again or press ctrl+c to exit.")
			}
		}
	*/

	var hint openpgp.FileHints
	hint.IsBinary = false
	hint.FileName = "_CONSOLE"
	hint.ModTime = time.Now()

	w, _ := armor.Encode(os.Stdout, "PGP MESSAGE", nil)
	defer w.Close()
	plaintext, _ := openpgp.Encrypt(w, []*openpgp.Entity{theirPublicKey}, myPrivateKey, &hint, nil)
	defer plaintext.Close()

	fmt.Fprintf(plaintext, "黄勇刚在熟悉OpenPGP代码\n")
}
