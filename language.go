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
		"test.test": "Test {0}",
		"admin.login": "Login",
	},
	"ko_KR": {
		"admin.login": "로그인",
	},
}
