package models

import "crypto/rsa"

type Key struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	Id         string
}
