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
		"admin.login": "Login",
		"admin.login.username": "Username",
		"admin.login.password": "Password",
		"admin.login.submit": "Login",
	},
	"ko_KR": {
		"admin.login": "로그인",
		"admin.login.username": "유저이름",
		"admin.login.password": "비밀번호",
		"admin.login.submit": "로그인",
	},
}
