package input

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func parseHeaders(headers string) map[string]string {
	if len(headers) == 0 {
		return nil
	}

	headerMap := make(map[string]string)
	for _, header := range strings.Split(headers, ",") {
		keyVal := strings.Split(header, ":")
		if len(keyVal) != 2 {
			panic("Error: Invalid set of response headers in command line arg. See help.")
		}
		headerMap[keyVal[0]] = keyVal[1]
	}
	return headerMap
}

func parseTimeStringIntoDuration(durStr string) time.Duration {
	if durStr == "" {
		return 0
	}

	if len(durStr) < 2 || durStr[len(durStr)-1] != 's' {
		panic("Invalid sleep time input. See help.")
	}

	if unicode.IsLetter(rune(durStr[len(durStr)-2])) {
		if durStr[len(durStr)-2] == 'm' {
			count, err := strconv.Atoi(durStr[0 : len(durStr)-2])
			if err != nil {
				panic(fmt.Sprintf("%v", err))
			}
			return time.Duration(count) * time.Millisecond
		} else {
			panic("Invalid sleep time input. See help.")
		}
	}

	count, err := strconv.Atoi(durStr[0 : len(durStr)-1])
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	return time.Duration(count) * time.Second
}
