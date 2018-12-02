package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

// ---- Base58 Function from https://github.com/btcsuite/btcutil/blob/master/base58/base58.go
var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

const (
	// alphabet is the modified base58 alphabet used by Bitcoin.
	alphabet     = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	alphabetIdx0 = '1'
)

func Encode(b []byte) string {
	x := new(big.Int)
	x.SetBytes(b)

	answer := make([]byte, 0, len(b)*136/100)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, alphabet[mod.Int64()])
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabetIdx0)
	}

	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}

// ---- Base58 Function from https://github.com/btcsuite/btcutil/blob/master/base58/base58.go

func main() {
	// 1. Generating 32 bytes of cryptographically generated random integers
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	// 2. Adding 0x80 bytes ([128]) to the randomly generated private key
	ek := "80" + fmt.Sprintf("%x", b) // i.e Extended Private Key
	hash, _ := hex.DecodeString(ek)   // Converting string to hex
	// 3. Taking SHA256 hash of the extended key
	h := sha256.Sum256(hash)
	// 4. Taking again SHA256 hash of the result of first SHA256 hash
	hashtwo, _ := hex.DecodeString(fmt.Sprintf("%x", h))
	h2 := sha256.Sum256(hashtwo)
	// 5. Taking first 4 bytes of the hash
	byteSplit := h2[:4] // i.e checksum
	// fmt.Println(fmt.Sprintf("%x", byteSplit))
	// 6. Adding checksum to the extended private key
	ekc := ek + fmt.Sprintf("%x", byteSplit)
	ekch, _ := hex.DecodeString(ekc) // string to hex
	// 7. Base58 Encoding
	wifKey := Encode([]byte(ekch))
	println(wifKey) // Wallet Import Format Private Key
}
