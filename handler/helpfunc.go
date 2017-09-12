package handler

import (
	"regexp"
	"strconv"
	"strings"
	"log"
	"net/http"
)

//Split is good
func Split(str string) (string, string) {
	arg := ""
	s := strings.Split(str, " ")
	command := s[0]
	if len(s) > 1 {
		arg = s[1]
	}
	return command, arg
}

//ClearString make me happy
func ClearString(str string) string {
	str = str[1 : len(str)-1]
	return str
}

//ClearID make me happy
func ClearID(str string) string {
	floatID, _ := strconv.ParseFloat(str, 64)
	id := int(floatID)
	s := strconv.Itoa(id)
	return s
}

//CheckArg test args
func CheckArg(arg string) bool {
	var validID = regexp.MustCompile(`[0-9]{1}[:-]{1}[0-9]{1}`)

	if len(arg) == 3 {
		return validID.MatchString(arg)
	}

	return false
}

//WhoWin return who win or draw
func WhoWin(a int, b int) int {
	var k int
	if a == b {
		k = 0
	}

	if a > b {
		k = 1
	}

	if a < b {
		k = 2
	}
	return k
}

//SendContent allow send video or gif or photo
func SendContent(id, message, content string) {
	log.Println(content)
	message = strings.Replace(message, " ", "%20", -1)
	_, err := http.Get("https://api.vk.com/method/messages.send?access_token=449c941ace6703fd55cfaee6b5863da12ecd29ad56168d842b73809a747bd2987f49a0ad4f3e64ab81178" + "&user_ids=" + id + "&message=" + message + "&attachment=" + content)
	if err != nil {
		log.Println(err)
	}
}
