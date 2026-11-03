package standardlibrary

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Example_crypto() {
	sum := sha256.Sum256([]byte("hello"))
	fmt.Println(hex.EncodeToString(sum[:])[:16])

	mac := hmac.New(sha256.New, []byte("secret")) // sign webhooks, tokens
	mac.Write([]byte("msg"))
	sig := mac.Sum(nil)
	fmt.Println(len(sig), hmac.Equal(sig, sig)) // constant-time compare

	token := make([]byte, 16)
	rand.Read(token) // crypto/rand: safe for secrets
	fmt.Println(len(hex.EncodeToString(token)))
	// Output:
	// 2cf24dba5fb0a30e
	// 32 true
	// 32
}
