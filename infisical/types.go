package infisical

import "time"

type ServiceTokenResponse struct {
	Id           string  `json:"_id"`
	Name         string  `json:"name"`
	Workspace    string  `json:"workspace"`
	Scopes       []Scope `json:"scopes"`
	User         `json:"user"`
	LastUsed     time.Time `json:"lastUsed"`
	ExpiresAt    time.Time `json:"expiresAt"`
	EncryptedKey string    `json:"encryptedKey"`
	Iv           string    `json:"iv"`
	Tag          string    `json:"tag"`
	Permissions  []string  `json:"permissions"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	V            int       `json:"__v"`
}

type Scope struct {
	Environment string `json:"environment"`
	SecretPath  string `json:"secretPath"`
	Id          string `json:"_id"`
}

type User struct {
	Id           string        `json:"_id"`
	AuthMethods  []string      `json:"authMethods"`
	Email        string        `json:"email"`
	IsMfaEnabled bool          `json:"isMfaEnabled"`
	MfaMethods   []interface{} `json:"mfaMethods"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	V            int           `json:"__v"`
	FirstName    string        `json:"firstName"`
	LastName     string        `json:"lastName"`
}

type SecretsResponse struct {
	Secrets []Secret `json:"secrets"`
}

type Secret struct {
	ID        string `json:"_id"`
	Version   int    `json:"version"`
	Workspace string `json:"workspace"`
	Type      string `json:"type"`
	Tags      []struct {
		ID        string `json:"_id"`
		Name      string `json:"name"`
		Slug      string `json:"slug"`
		Workspace string `json:"workspace"`
	} `json:"tags"`
	Environment             string    `json:"environment"`
	SecretKeyCiphertext     string    `json:"secretKeyCiphertext"`
	SecretKeyIV             string    `json:"secretKeyIV"`
	SecretKeyTag            string    `json:"secretKeyTag"`
	SecretValueCiphertext   string    `json:"secretValueCiphertext"`
	SecretValueIV           string    `json:"secretValueIV"`
	SecretValueTag          string    `json:"secretValueTag"`
	SecretCommentCiphertext string    `json:"secretCommentCiphertext"`
	SecretCommentIV         string    `json:"secretCommentIV"`
	SecretCommentTag        string    `json:"secretCommentTag"`
	Algorithm               string    `json:"algorithm"`
	KeyEncoding             string    `json:"keyEncoding"`
	Folder                  string    `json:"folder"`
	V                       int       `json:"__v"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}

type DecodedSymmetricEncryptionDetails = struct {
	Cipher []byte
	IV     []byte
	Tag    []byte
	Key    []byte
}

type SingleEnvironmentVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
	ID    string `json:"_id"`
	Tags  []struct {
		ID        string `json:"_id"`
		Name      string `json:"name"`
		Slug      string `json:"slug"`
		Workspace string `json:"workspace"`
	} `json:"tags"`
	Comment string `json:"comment"`
}
