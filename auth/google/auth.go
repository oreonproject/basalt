package google

import (
	crand "crypto/rand"
	"crypto/sha256"
	"log"
	"oreonproject/basalt/utils"
)

func CodeVerifyKeyGen() []byte {
	// We will Initialise the logs
	utils.LogInit("logs/auth.log")

	key := make([]byte, 32) // creates a new byte slice

	// Use the Read Function from crypto/rand to populate the key with cryptographically secure random values
	crand.Read(key)
	log.Print("Populated the Key")

	// Initialise a hasher
	hasher := sha256.New()
	// Write to the key using the hash digest for key as an extra safeguard
	hasher.Write(key)
	log.Println("Code Verification Key Generated")
	return key
}
