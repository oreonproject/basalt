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

	CodeVerifierKey := make([]byte, 32) // creates a new byte slice

	// Use the Read Function from crypto/rand to populate the CodeVerifierKey with cryptographically secure random values
	crand.Read(CodeVerifierKey)
	log.Print("Populated the Key")

	// Initialise a hasher
	hasher := sha256.New()
	// Write to the CodeVerifierKey using the hash digest for CodeVerifierKey as an extra safeguard
	hasher.Write(CodeVerifierKey)
	log.Println("Code Verification Key Generated")
	return CodeVerifierKey
}
