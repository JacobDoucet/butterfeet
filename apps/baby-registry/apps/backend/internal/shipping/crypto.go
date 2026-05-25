package shipping

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// HashEmail returns sha256(lowercase(trim(email))) hex-encoded.
// Used as a stable lookup key; never reversible.
func HashEmail(email string) string {
	normalized := strings.ToLower(strings.TrimSpace(email))
	sum := sha256.Sum256([]byte(normalized))
	return hex.EncodeToString(sum[:])
}

// EncryptEmail returns an opaque blob that can be decoded back to the
// original email via DecryptEmail. Today this is identity-with-marker
// so the schema is stable for when AES-GCM is layered on in step 8.
func EncryptEmail(email string) string {
	return "v0:" + strings.ToLower(strings.TrimSpace(email))
}

// DecryptEmail reverses EncryptEmail. Returns "" if the blob is not a
// recognized version.
func DecryptEmail(blob string) string {
	if strings.HasPrefix(blob, "v0:") {
		return strings.TrimPrefix(blob, "v0:")
	}
	return ""
}
