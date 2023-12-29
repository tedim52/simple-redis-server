package resp

import (
	"errors"
	"strings"
	// "fmt"
)

const (
	nullValue = "$-1\r\n"
	crlf = "\r\n"

	simpleStringType = string("+")
	errorType = string("-")
	intType = string(":")
	bulkStringType = string("$")
	arrayType = string("*")
)

func DeserializeResp(respMsg string) (string, error) {
	if len(respMsg) < 0 {
		return "", errors.New("empty string received")
	}
	if nullValue == respMsg || len(respMsg) < 0 {
		return "NULL", nil
	}

	dataType := respMsg[0]
	var parsedMsg string
	switch string(dataType) {
	case simpleStringType:
		parsedMsg = parseSimpleString(respMsg)
	case bulkStringType:
		parsedMsg = parseBulkString(respMsg)
	case arrayType:
		parsedMsg = parseArray(respMsg)
	case intType:
		parsedMsg = parseInt(respMsg)
	case errorType:
		parsedMsg = parseError(respMsg)
	default:
		return "", errors.New("unsupported data type")
	}

	return parsedMsg, nil
}

// format: +<data><CRLF>
func parseSimpleString(respMsg string) string {
	if len(respMsg) < 1 + len(crlf) {
		// is this an okay behavior for this case? figure it out later
		return ""
	}

	msg := respMsg[1:]
	msgWoCrlf := strings.TrimRight(msg, crlf)
	return msgWoCrlf
}


func parseBulkString(respMsg string) string {
	// TODO: implement
	return ""
}

func parseInt(respMsg string) string {
	// TODO: implement

	return ""
}


func parseArray(respMsg string) string {
	// TODO: implement

	return ""
}

func parseError(respMsg string) string {
	if len(respMsg) < 1 + len(crlf) {
		// is this an okay behavior for this case? figure it out later
		return ""
	}

	errorMsg := respMsg[1:]
	errorMsgWoCrlf := strings.TrimRight(errorMsg, crlf)
	return errorMsgWoCrlf

	return errorMsgWoCrlf
}

