package core

import "os"

func GetKaiheilaBotVerifyToken() string {
	k := os.Getenv("KAIHEILA_BOT_VARIFY_TOKEN")
	if k == "" {
		return "xxx"
	}
	return k
}

func GetKaiheilaBotToken() string {
	k := os.Getenv("KAIHEILA_BOT_TOKEN")
	if k == "" {
		return "xxx"
	}
	return k
}

func GetKaiheilaBotEncryptKey() string {
	k := os.Getenv("KAIHEILA_BOT_ENCRYPT_KEY")
	if k == "" {
		return "xxxx"
	}
	return k
}
