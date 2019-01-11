package utils

import (
	"regexp"
	"strconv"
	"time"
)

/*
DateValidation - Validador de data de nascimento
*/
func DateValidation(date time.Time) int {
	loc, _ := time.LoadLocation("UTC")
	if time.Now().In(loc).Add(1*time.Hour).Sub(date) < 0 {
		return -1
	} else if time.Now().In(loc).Add(1*time.Hour).Sub(date) < 155520 {
		return 1
	}
	return 0
}

/*
EmailValidation - Validador de email
*/
func EmailValidation(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(email) == true {
		return true
	} else {
		return false
	}
}

/*
RgValidation - Validador de rg
*/
func RgValidation(rg string) bool {
	if len(rg) != 9 {
		return false
	}
	return true
}

/*
CpfValidation - validador de Cpf
*/
func CpfValidation(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}
	var eq bool
	var dig string
	for _, val := range cpf {
		if len(dig) == 0 {
			dig = string(val)
		}
		if string(val) == dig {
			eq = true
			continue
		}
		eq = false
		break
	}
	if eq {
		return false
	}
	i := 10
	sum := 0
	for index := 0; index < len(cpf)-2; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}
	prod := sum * 10
	mod := prod % 11
	if mod == 10 {
		mod = 0
	}
	digit1, _ := strconv.Atoi(string(cpf[9]))
	if mod != digit1 {
		return false
	}
	i = 11
	sum = 0
	for index := 0; index < len(cpf)-1; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}
	prod = sum * 10
	mod = prod % 11
	if mod == 10 {
		mod = 0
	}
	digit2, _ := strconv.Atoi(string(cpf[10]))
	if mod != digit2 {
		return false
	}

	return true
}
