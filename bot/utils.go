package bot

import "strings"

func fetchCommand(msg string) (string, bool) {
	if len(msg) == 0 {
		return "", false
	}
	if msg[0] == '/' {
		ss := strings.Split(msg, " ")
		return ss[0], true
	}
	return "", false
}
