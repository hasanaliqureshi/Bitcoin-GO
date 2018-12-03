package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Private Key, Generated from 1. Private Key/main.go
	sum, _ := hex.DecodeString("e241ff7ab61a14ffe64bd04093a696d46c93cbb89df6f093f3f5e8bafee35a95")
	// Elliptic Curve Digital Signature Algorithm ECDSA
	k := new(big.Int)                                                    // creating a bigInt
	k.SetBytes(sum)                                                      // putting Private Key hex in it
	var priv *ecdsa.PrivateKey                                           // defining ecdsa.Privateakey
	priv = new(ecdsa.PrivateKey)                                         // allocating priv in ecdsa struct
	curve := crypto.S256()                                               // elliptic cure using ethereum-go library
	priv.PublicKey.Curve = curve                                         // defining PublicKey Curve
	priv.D = k                                                           // allocating private key in D *Int
	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarBaseMult(k.Bytes()) // performing Scalar Base
	pub := elliptic.Marshal(curve, priv.PublicKey.X, priv.PublicKey.Y)   // Generating Public Key from X and Y component
	// Generating Compressed Public Key
	ycomp := priv.PublicKey.Y.Bytes() // y component of Public Key
	var byteadd string
	if ycomp[binary.Size(ycomp)-1]%2 == 0 {
		byteadd = "02" // if last byte of y component of Public Key is Even, we'll add 0x02 byte to x component of Public Key
	} else {
		byteadd = "03" // if last byte of y component of Public Key is Odd, we'll add 0x03 byte to x component of Public Key
	}
	cpub := byteadd + fmt.Sprintf("%x", priv.PublicKey.X) // adding byte to x componeent of Public Key
	fmt.Println("Public Key : " + fmt.Sprintf("%x", pub))
	fmt.Println("Compressed Public Key : " + cpub)

}
