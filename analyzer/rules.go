package analyzer

import (
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func checkMessage(pass *analysis.Pass, pos token.Pos, msg string) {
	checkLowercase(pass, pos, msg)
	checkEnglish(pass, pos, msg)
	checkSpecial(pass, pos, msg)
	checkSensitive(pass, pos, msg)
}

func checkLowercase(pass *analysis.Pass, pos token.Pos, msg string) {
	r := []rune(msg)[0]
	if unicode.IsUpper(r) {
		pass.Reportf(pos, "log message should start with lowercase letter: %s", msg)
	}
}

func checkEnglish(pass *analysis.Pass, pos token.Pos, msg string) {
	for _, r := range msg {
		if r > unicode.MaxASCII {
			pass.Reportf(pos, "log message must be in english: %s", msg)
			return
		}
	}
}

func checkSpecial(pass *analysis.Pass, pos token.Pos, msg string) {
	for _, r := range msg {
		if !(unicode.IsSpace(r) || unicode.IsNumber(r) || unicode.IsLetter(r)) {
			pass.Reportf(pos, "log message contains forbidden special characters: %s", msg)
			return
		}
	}
}

func checkSensitive(pass *analysis.Pass, pos token.Pos, msg string) {
	msg = strings.ToLower(msg)
	for _, keyword := range sensitive {
		if strings.Contains(msg, keyword) {
			pass.Reportf(pos, "log message contains sensitive data: %s", msg)
			return
		}
	}
}

var sensitive = []string{"password", "token", "api_key"}
