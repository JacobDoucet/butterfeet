package shipping

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// NewToken returns a fresh URL-safe random token (raw) and its sha256 hex hash.
// Only the hash is persisted; the raw token only ever lives inside the share URL.
func NewToken() (raw, hash string, err error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", "", err
	}
	raw = base64.RawURLEncoding.EncodeToString(buf)
	hash = HashToken(raw)
	return raw, hash, nil
}

func HashToken(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}
