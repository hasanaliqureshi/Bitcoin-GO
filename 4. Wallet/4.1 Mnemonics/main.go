package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// 1. Create 128 bits entropy
	e := make([]byte, 16)
	_, err := rand.Read(e)
	if err != nil {
		panic(err)
	}

	ent := ""
	for _, b := range e {
		ent = ent + fmt.Sprintf("%08b", b)
	}

	es := sha256.Sum256(e)
	esh, _ := hex.DecodeString(fmt.Sprintf("%x", es))

	hashbits := ""
	for _, b := range esh {
		hashbits = hashbits + fmt.Sprintf("%08b", b)
	}

	cs := ent + hashbits[:4]
	mc := make([]int, 12)
	var str string
	for i := 0; i < 12; i++ {
		startIndex := i * 11
		endIndex := startIndex + 11
		if endIndex >= len(cs)-1 {
			str = cs[startIndex:]
		} else {
			str = cs[startIndex:endIndex]
		}
		asInt, err := strconv.ParseInt(str, 2, 64)
		if err != nil {
			panic(err)
		}
		mc[i] = int(asInt)
	}

	size := int(math.Pow(2, 11))

	dict := make([]string, size)
	reverseDict := make(map[string]int, size)

	file, err := os.Open("dictionary.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		word := scanner.Text()
		dict[i] = word
		reverseDict[word] = i
		i++
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Print(dict)
	m := make([]string, 12)
	for k, code := range mc {
		m[k] = dict[code]
	}
	fmt.Println(m)
}
