package models

type Wallet struct{
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
	HDpath     string `json:"hdpath"`
}