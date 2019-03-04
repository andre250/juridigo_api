package utils

import (
	"math/rand"
	"strings"
)

/*
GenerateRecoveryToken - Responsável por gerar um token de validação de
recuperação de senha.
*/
func GenerateRecoveryToken(size int) string {
	var recovery []string

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@!&%")

	for x := 0; x < size; x++ {
		recovery = append(recovery, string(letterRunes[rand.Intn(len(letterRunes))]))
	}

	finaltoken := strings.Join(recovery, "")
	return finaltoken
}
