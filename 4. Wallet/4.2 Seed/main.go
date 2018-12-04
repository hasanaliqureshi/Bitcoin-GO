package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	mnemonicWords := []byte("into green spirit close pool estate cinnamon unaware consider silent siege fatal") // Generated from 4. Wallets/ 4.1 Mnemonics/ main.go
	passPhrase := "asdjlaskjdklsajds"                                                                           // BIP39 : There are no “wrong” passphrases in BIP-39. Every passphrase leads to some wallet, which unless previously used will be empty.

	s := pbkdf2.Key([]byte(mnemonicWords), []byte("mnemonic"+passPhrase), 2048, 64, sha512.New)
	seed := hex.EncodeToString(s)
	fmt.Println(seed)
}
