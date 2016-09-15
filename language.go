package main

import (
	"strings"
	"strconv"
)

const FallBackLanguage = "en_US"

type Language struct {
	Map map[string]string
}

func (l *Language) Init(lang string) bool {
	if rsc, ok := langList[lang]; ok {
		l.Map = rsc
		return true
	}
	return false
}

func (l *Language) Get(key string, val ...string) (ret string) {
	if msg, ok := l.Map[key]; ok {
		for i := 0; i < len(msg); i++ {
			if msg[i] == '{' {
				next := strings.IndexByte(msg[i:], '}')
				if next > 0 {
					num := msg[i:i + next]
					index, _ := strconv.Atoi(num)

					if len(val) > index {
						ret += val[index]

						i += next + 1
					}
				}
			}
			if len(msg) > i {
				ret += string(msg[i])
			}
		}
	}

	return
}

var langList = map[string]map[string]string{
	"en_US": {
		"account.login": "Login",
		"account.login.username": "Username",
		"account.login.password": "Password",
		"account.login.submit": "Login",
		"account.login.noaccount": "Don't have account?",

		"account.register": "Register",
		"account.register.username": "Username",
		"account.register.password": "Password",
		"account.register.passwordConfirm": "Confirm password",
		"account.register.email": "E-mail",
		"account.register.submit": "Register",
		"account.register.haveaccount": "I have my account.",
	},
	"ko_KR": {
		"account.login": "로그인",
		"account.login.username": "유저이름",
		"account.login.password": "비밀번호",
		"account.login.submit": "로그인",
		"account.login.noaccount": "계정이 없습니까?",

		"account.register": "가입",
		"account.register.username": "유저이름",
		"account.register.password": "비밀번호",
		"account.register.passwordConfirm": "비밀번호 확인",
		"account.register.email": "이메일",
		"account.register.submit": "가입",
		"account.register.haveaccount": "계정이 있습니다",
	},
}
