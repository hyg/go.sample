// openpgp project main.go
package main

import (
	"crypto/openpgp"
	"crypto/openpgp/armor"
	"fmt"
	"github.com/howeyc/gopass"
	"log"
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

func main() {
	fmt.Println("OpenPGP sample!")

	pubringFile, _ := os.Open("E:\\huangyg\\PGP\\p2club.pkr")
	defer pubringFile.Close()
	pubring, _ := openpgp.ReadKeyRing(pubringFile)
	theirPublicKey := getKeyByEmail(pubring, "huangyg@p2club.org")

	privringFile, _ := os.Open("E:\\huangyg\\PGP\\p2club.skr")
	defer privringFile.Close()
	privring, _ := openpgp.ReadKeyRing(privringFile)
	myPrivateKey := getKeyByEmail(privring, "huangyg@p2club.org")

	//theirPublicKey.Serialize(os.Stdout)
	//myPrivateKey.Serialize(os.Stdout)
	//myPrivateKey.SerializePrivate(os.Stdout, nil)

	log.Println(theirPublicKey)
	log.Println(myPrivateKey)

	myPrivateKey.PrivateKey.Decrypt([]byte("abcde"))
	for myPrivateKey.PrivateKey.Encrypted {
		fmt.Print("PGP passphrase: ")
		pgppass := gopass.GetPasswd()

		myPrivateKey.PrivateKey.Decrypt([]byte(pgppass))
		if myPrivateKey.PrivateKey.Encrypted {
			fmt.Println("Incorrect. Try again or press ctrl+c to exit.")
		}
	}

	var hint openpgp.FileHints
	hint.IsBinary = false
	hint.FileName = "_CONSOLE"
	hint.ModTime = time.Now()

	w, _ := armor.Encode(os.Stdout, "PGP MESSAGE", nil)
	plaintext, _ := openpgp.Encrypt(w, []*openpgp.Entity{theirPublicKey}, myPrivateKey, &hint, nil)
	fmt.Fprintf(plaintext, "黄勇刚在熟悉OpenPGP代码\n")
	plaintext.Close()
	w.Close()
	fmt.Printf("\n")
}
