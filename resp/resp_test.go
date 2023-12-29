package resp

import (
	"testing"
	"github.com/stretchr/testify/assert"
	// "fmt"
)

func TestNullMsg(t *testing.T){
	respMsg := "$-1\r\n"

	actualResp, err := DeserializeResp(respMsg)

	expResp := "NULL"

	assert.Nil(t, err)
	assert.Equal(t, expResp, actualResp)
}

func TestParseSimpleString(t *testing.T) {
	respMsgs := []string{
		"+OK\r\n",
		"+hello world\r\n",
	}
	expectedMsgs := []string{
		"OK", 
		"hello world",
	}

	testRespMsgParser(t, parseSimpleString, respMsgs, expectedMsgs)
}

func TestParseError(t *testing.T) {
	respMsgs := []string{
		"-Error message\r\n",
	}

	expectedMsgs := []string{
		"Error message",
	}

	testRespMsgParser(t, parseError, respMsgs, expectedMsgs)
}

func TestParseBulkString(t *testing.T) {
	respMsgs := []string{
		"$0\r\n\r\n",
		"$-1\r\n",
		"$5\r\nhello\r\n",
		"$-1\r\n", // null bulk string
	}

	expectedMsgs := []string{
		"",
		""
	}

	testRespMsgParser(t, parseBulkString, respMsgs, expectedMsgs)
}

func TestParseInteger(t *testing.T) {
	respMsgs := []string{
		
	}

	expectedMsgs := []string{
		"",
		""
	}

	testRespMsgParser(t, parseBulkString, respMsgs, expectedMsgs)
}

func TestParseArray(t *testing.T) {
	respMsgs := []string{
		"*0\r\n",
		"*-1\r\n", // null array
		"*1\r\n$4\r\nping\r\n",
		"*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n",
		"*2\r\n$3\r\nget\r\n$3\r\nkey\r\n",
		"*3\r\n:1\r\n:2\r\n:3\r\n",
		"*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Hello\r\n-World\r\n",
		"*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$5\r\nhello\r",
		"*3\r\n$5\r\nhello\r\n$-1\r\n$5\r\nworld\r\n", // null element in array
		}

	expectedMsgs := []string{
		"",
		""
	}

	testRespMsgParser(t, parseBulkString, respMsgs, expectedMsgs)
}

func testRespMsgParser(t *testing.T, parser func(string)string, respMsgs []string, expectedMsgs []string) {
	for idx, respMsg := range respMsgs {
		actualMsg := parser(respMsg)
		assert.Equal(t, expectedMsgs[idx], actualMsg)
	}
}