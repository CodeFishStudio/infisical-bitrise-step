package infisical

import (
	"crypto/aes"
	"crypto/cipher"
)

func DecryptSymmetric(key []byte, cipherText []byte, tag []byte, iv []byte) ([]byte, error) {
	// Case: empty string
	if len(cipherText) == 0 && len(tag) == 0 && len(iv) == 0 {
		return []byte{}, nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, len(iv))
	if err != nil {
		return nil, err
	}

	var nonce = iv
	var ciphertext = append(cipherText, tag...) // the aesgcm open method expects auth tag at the end of the cipher text

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
