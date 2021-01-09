package config

type JWT struct {
	SigningKey   string `json:"signing_key"`
	ExpireSecond int64  `json:"expire_second"`
}
