package main

import "bytes"

func spamMaker(message string) string {
	linkPrefix := "http://"
	lenLinkPrefix := len(linkPrefix)
	if len(message) <= lenLinkPrefix {
		return message
	}
	lenMessage := len(message)
	isThisLink := false
	result := bytes.NewBuffer(make([]byte, 0, lenMessage))

	for i := 0; i < lenMessage; i++ {
		if isThisLink {
			if message[i] == ' ' {
				isThisLink = false
				result.WriteByte(message[i])
			} else {
				result.WriteByte('*')
			}
		} else if i+lenLinkPrefix > lenMessage {
			result.WriteByte(message[i])
		} else if linkPrefix == message[i:i+lenLinkPrefix] {
			result.WriteString(linkPrefix)
			i += lenLinkPrefix - 1
			isThisLink = true
		} else {
			result.WriteByte(message[i])
		}
	}

	return result.String()
}
