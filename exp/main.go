package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/Kally95/Go_Web_App/hash"
)

func main() {
	toHash := []byte("this is my string to hash")
	h := hmac.New(sha256.New, []byte("my-secret-key"))
	h.Write(toHash)
	b := h.Sum(nil)
	h.Reset()

	fmt.Println(base64.URLEncoding.EncodeToString(b))

	hmac := hash.NewHMAC("my-secret-key")
	fmt.Println(hmac.Hash("this is my string to hash"))

}
