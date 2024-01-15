package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"

	cr "crypto/rand"
)

func justInt() {
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// Always the same value.
	// A limitation of the go.dev/play sandbox.
	fmt.Println(time.Now().UnixNano())
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}

func shuffleStrings() {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	rand.Shuffle(len(answers), func(i, j int) {
		answers[i], answers[j] = answers[j], answers[i]
	})
	fmt.Println(answers)

}

func crypto() {

	rcn, err := cr.Int(cr.Reader, big.NewInt(27))
	if err != nil {
		log.Fatalf("Should be able to get a random number: %v\n", err)
	}
	fmt.Println(rcn.Int64())

	pn, err := cr.Prime(cr.Reader, 10)
	if err != nil {
		log.Fatalf("Should be able to get a prime number: %v\n", err)
	}
	fmt.Println(pn.Int64())

	// token
	t, err := token()
	if err != nil {
		log.Fatalf("cannot get token: %v\n", err)
	}
	fmt.Println(t)

	// hex
	src := []byte("Hello Gopher!")
	dst := make([]byte, hex.EncodedLen(len(src)))
	fmt.Println("dst len", len(dst))
	hex.Encode(dst, src)
	fmt.Printf("%s\n", dst)
}

func token() (string, error) {
	b := make([]byte, 32)
	_, err := cr.Read(b)
	if err != nil {
		return "", err
	}
	fmt.Printf("token in hex: %x\n", b)
	return base64.StdEncoding.EncodeToString(b), nil
}

func main() {
	justInt()
	shuffleStrings()
	fmt.Println(rand.Perm(10))
	crypto()
}
