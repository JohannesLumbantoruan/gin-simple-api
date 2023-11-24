package main

import (
	"os"
	"crypto/rand"
	"fmt"
)

func InitEnv() {
	atSecret := make([]byte, 64)
	rtSecret := make([]byte, 64)

	rand.Read(atSecret)
	rand.Read(rtSecret)

	os.Setenv("ACCESS_TOKEN_SECRET", fmt.Sprintf("%x", atSecret))
	os.Setenv("REFRESH_TOKEN_SECRET", fmt.Sprintf("%x", rtSecret))
}