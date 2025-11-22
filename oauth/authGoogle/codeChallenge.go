package authGoogle

import (
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"oreonproject/basalt/utils"
	"strings"
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

func CodeChallengeGen() string {

	CodeVerifier := CodeVerifierKeyGen()

	hasher := sha256.New()
	hasher.Write(CodeVerifier)

	challenge := base64.RawURLEncoding.EncodeToString(CodeVerifier)

	return challenge
}
func StateTokGen() string {
	stateTok := make([]byte, 16)
	crand.Read(stateTok) // Populates State Token

	state := base64.RawURLEncoding.EncodeToString(stateTok) // Encodes state to base64 for added protections
	state = strings.ReplaceAll(state, "%", "")
	fmt.Println(state)
	return state
}
