package oauth

import (
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"oreonproject/basalt/utils"
)

func CodeVerifierKeyGen() string {
	log := utils.LogInit("oauth.log")

	verifier := make([]byte, 32)
	crand.Read(verifier)

	verifierStr := base64.RawURLEncoding.EncodeToString(verifier)
	log.Println("Code Verification Key Generated")
	return verifierStr
}

// Generate code challenge from a verifier
func CodeChallengeGen(verifier string) string {
	hash := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(hash[:])
	return challenge
}

func StateTokGen() string {
	log := utils.LogInit("oauth.log")
	stateTok := make([]byte, 16)
	crand.Read(stateTok)
	state := base64.RawURLEncoding.EncodeToString(stateTok)
	log.Println("State Token Generated")
	return state
}
