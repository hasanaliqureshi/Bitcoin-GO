package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"golang.org/x/crypto/ripemd160"
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
	// Public Key Generated from 2. Public Key/main.go
	pub, _ := hex.DecodeString("03cc9a549a0abd655d0518c562e72c31253e6bd9192dd4f369af3dba173ef1575e")
	// Taking SHA256 hash of Public Key
	sha := sha256.Sum256(pub)
	shaHex, _ := hex.DecodeString(fmt.Sprintf("%x", sha))
	// Taking RipeMD160 hash of the result of SHA256
	hasher := ripemd160.New()
	hasher.Write(shaHex)
	hashBytes := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	// Generating Mainnet Address
	madd := "00" + hashString                                // Adding 0x00 byte to the result of RipeMD160 hash
	maddh, _ := hex.DecodeString(madd)                       // Converting string to hex
	smaddo := sha256.Sum256(maddh)                           // Taking SHA256 hash of hex
	smaddh, _ := hex.DecodeString(fmt.Sprintf("%x", smaddo)) // SHA256 result to hex
	smaddt := sha256.Sum256(smaddh)                          // Taking SHA256 again of hex
	bmadd := smaddt[:4]                                      // checksum of first 4 bytes of 2nd SHA256 hash
	maddf := madd + fmt.Sprintf("%x", bmadd)                 // adding checksum to RipeMD160 hash
	maddfh, _ := hex.DecodeString(maddf)                     // Converting string to hex
	fmt.Println("Mainnet Address: " + Encode(maddfh))        // main net address
	// Generating Testnet Address
	tadd := "6f" + hashString                                // Adding 0x00 byte to the result of RipeMD160 hash
	taddh, _ := hex.DecodeString(tadd)                       // Converting string to hex
	staddo := sha256.Sum256(taddh)                           // Taking SHA256 hash of hex
	staddh, _ := hex.DecodeString(fmt.Sprintf("%x", staddo)) // SHA256 result to hex
	staddt := sha256.Sum256(staddh)                          // Taking SHA256 again of hex
	btadd := staddt[:4]                                      // checksum of first 4 bytes of 2nd SHA256 hash
	taddf := tadd + fmt.Sprintf("%x", btadd)                 // adding checksum to RipeMD160 hash
	taddfh, _ := hex.DecodeString(taddf)                     // Converting string to hex
	fmt.Println("Testnet Address: " + Encode(taddfh))        // test net address
}
