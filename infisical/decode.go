package infisical

import (
	"encoding/base64"
	"fmt"
)

func GetBase64DecodedSymmetricEncryptionDetails(key string, cipher string, IV string, tag string) (DecodedSymmetricEncryptionDetails, error) {
	cipherx, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return DecodedSymmetricEncryptionDetails{}, fmt.Errorf("Base64DecodeSymmetricEncryptionDetails: Unable to decode cipher text [err=%v]", err)
	}

	keyx, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return DecodedSymmetricEncryptionDetails{}, fmt.Errorf("Base64DecodeSymmetricEncryptionDetails: Unable to decode key [err=%v]", err)
	}

	IVx, err := base64.StdEncoding.DecodeString(IV)
	if err != nil {
		return DecodedSymmetricEncryptionDetails{}, fmt.Errorf("Base64DecodeSymmetricEncryptionDetails: Unable to decode IV [err=%v]", err)
	}

	tagx, err := base64.StdEncoding.DecodeString(tag)
	if err != nil {
		return DecodedSymmetricEncryptionDetails{}, fmt.Errorf("Base64DecodeSymmetricEncryptionDetails: Unable to decode tag [err=%v]", err)
	}

	return DecodedSymmetricEncryptionDetails{
		Key:    keyx,
		Cipher: cipherx,
		IV:     IVx,
		Tag:    tagx,
	}, nil
}

func GetPlainTextSecrets(key []byte, encryptedSecrets []Secret) ([]SingleEnvironmentVariable, error) {
	var plainTextSecrets []SingleEnvironmentVariable
	for _, secret := range encryptedSecrets {
		// Decrypt key
		keyIv, err := base64.StdEncoding.DecodeString(secret.SecretKeyIV)
		if err != nil {
			return nil, fmt.Errorf("unable to decode secret IV for secret key")
		}

		keyTag, err := base64.StdEncoding.DecodeString(secret.SecretKeyTag)
		if err != nil {
			return nil, fmt.Errorf("unable to decode secret authentication tag for secret key")
		}

		keyCiphertext, err := base64.StdEncoding.DecodeString(secret.SecretKeyCiphertext)
		if err != nil {
			return nil, fmt.Errorf("unable to decode secret cipher text for secret key")
		}

		plainTextKey, err := DecryptSymmetric(key, keyCiphertext, keyTag, keyIv)
		if err != nil {
			return nil, fmt.Errorf("unable to symmetrically decrypt secret key")
		}

		// Decrypt value
		valueIv, err := base64.StdEncoding.DecodeString(secret.SecretValueIV)
		if err != nil {
			return nil, fmt.Errorf("unable to decode secret IV for secret value")
		}

		valueTag, err := base64.StdEncoding.DecodeString(secret.SecretValueTag)
		if err != nil {
			return nil, fmt.Errorf("unable to decode secret authentication tag for secret value")
		}

		valueCiphertext, _ := base64.StdEncoding.DecodeString(secret.SecretValueCiphertext)
		if err != nil {
			return nil, fmt.Errorf("unable to decode secret cipher text for secret key")
		}

		plainTextValue, err := DecryptSymmetric(key, valueCiphertext, valueTag, valueIv)
		if err != nil {
			return nil, fmt.Errorf("unable to symmetrically decrypt secret value")
		}

		// Decrypt comment
		commentIv, err := base64.StdEncoding.DecodeString(secret.SecretCommentIV)
		if err != nil {
			return nil, fmt.Errorf("unable to decode secret IV for secret value")
		}

		commentTag, err := base64.StdEncoding.DecodeString(secret.SecretCommentTag)
		if err != nil {
			return nil, fmt.Errorf("unable to decode secret authentication tag for secret value")
		}

		commentCiphertext, _ := base64.StdEncoding.DecodeString(secret.SecretCommentCiphertext)
		if err != nil {
			return nil, fmt.Errorf("unable to decode secret cipher text for secret key")
		}

		plainTextComment, err := DecryptSymmetric(key, commentCiphertext, commentTag, commentIv)
		if err != nil {
			return nil, fmt.Errorf("unable to symmetrically decrypt secret comment")
		}

		plainTextSecret := SingleEnvironmentVariable{
			Key:     string(plainTextKey),
			Value:   string(plainTextValue),
			Type:    secret.Type,
			ID:      secret.ID,
			Tags:    secret.Tags,
			Comment: string(plainTextComment),
		}

		plainTextSecrets = append(plainTextSecrets, plainTextSecret)
	}

	return plainTextSecrets, nil
}
