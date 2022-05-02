package services

import "strings"

func GetMessage(messages ...[]string) (msg string) {

	var resp []string

	if !validateParameters(messages) {
		return ""
	}

	for indexMensajes, mensajes := range messages {
		for indexMSG, msg := range mensajes {
			if indexMensajes == 0 {
				resp = append(resp, msg)
			} else {
				if resp[indexMSG] == "" && msg != "" {
					resp[indexMSG] = msg
				} else if resp[indexMSG] != msg {
					resp = append(resp, msg)
				}
			}
		}
	}
	return strings.Join(resp, " ")
}

func validateParameters(messages [][]string) bool {

	if messages == nil {
		return false
	} else if len(messages) < 3 {
		return false
	} else if messages[0] == nil && len(messages[0]) < 0 {
		return false
	} else if messages[1] == nil && len(messages[1]) < 0 {
		return false
	} else if messages[2] == nil && len(messages[2]) < 0 {
		return false
	}

	return true
}
