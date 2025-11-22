package authGoogle

import (
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"oreonproject/basalt/utils"
)

func CodeVerifierKeyGen() []byte {
	// We will Initialise the logs
	utils.LogInit("logs/oauth.log")
	codeVerifierKey := make([]byte, 32) // creates a new byte slice

	// Use the Read Function from crypto/rand to populate the codeVerifier with cryptographically secure random values
	crand.Read(codeVerifierKey)
	log.Println("Code Verification Key Generated")

	return codeVerifierKey
}

func CodeChallengeGen(codeVerifier []byte) string {
	utils.LogInit("logs/oauth.log")
	var codeChallenge string

	hasher := sha256.New()
	hasher.Write(codeVerifier) // Hashes the codeVerifier

	codeChallenge = base64.RawURLEncoding.EncodeToString(codeVerifier) // Encodes the base64 URL to a string
	log.Print("codeChallenge Generated")
	return codeChallenge
}

func StateTokGen() string {
	stateTok := make([]byte, 16)
	crand.Read(stateTok) // Populates State Token

	state := base64.RawURLEncoding.EncodeToString(stateTok) // Encodes state to base64 for added protections
	return state
}
