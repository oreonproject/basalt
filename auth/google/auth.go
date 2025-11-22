package google

import (
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"oreonproject/basalt/utils"
)

func CodeVerifierKeyGen() []byte {
	// We will Initialise the logs
	utils.LogInit("logs/auth.log")
	codeVerifierKey := make([]byte, 32) // creates a new byte slice
	log.Print("PreSaltKey Initialised")

	// Use the Read Function from crypto/rand to populate the codeVerifier with cryptographically secure random values
	crand.Read(codeVerifierKey)
	log.Print("Populated the PreSaltKey")
	log.Println("Code Verification Key Generated")

	return codeVerifierKey
}

func CodeChallengeGen(codeVerifier []byte) {
	var codeChallenge string

	hasher := sha256.New()
	hasher.Write(codeVerifier) // Hashes the codeVerifier

	codeChallenge = base64.RawURLEncoding.EncodeToString(codeVerifier) // Encodes the base64 URL to a string
	fmt.Println(codeChallenge)
}
