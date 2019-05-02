package swagger

import "time"

type Secret struct {
	Hash           string    `json:"hash,omitempty" xml:"hash"`
	SecretText     string    `json:"secretText,omitempty" xml:"secretText"`
	CreatedAt      time.Time `json:"createdAt,omitempty" xml:"createdAt"`
	ExpiresAt      time.Time `json:"expiresAt,omitempty" xml:"expiresAt"`
	RemainingViews int64     `json:"remainingViews,omitempty" xml:"remainingViews"`
}

func (secret Secret) save() {
	saveSecret(secret)
}

func (secret Secret) delete() {
	removeSecret(secret.Hash)
}

func (secret Secret) expired() bool {
	return secret.RemainingViews < 0 || (!secret.CreatedAt.Equal(secret.ExpiresAt) && time.Now().After(secret.ExpiresAt))
}
