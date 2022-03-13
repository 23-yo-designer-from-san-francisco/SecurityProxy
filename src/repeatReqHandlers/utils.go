package repeatReqHandlers

import (
	"Proxy/db"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Param struct {
	key        string
	value      string
	vulnerable bool
}

func (p *Param) serialize() string {
	return fmt.Sprintf("%s=%s", p.key, p.value)
}

func (p *Param) fakeReplaceValue(val string) string {
	return fmt.Sprintf("%s=%s", p.key, url.QueryEscape(val))
}

const CookieRegex = `[Cc]ookie:.+`
const GETParamsRegex = `\?[a-zA-Z0-9~\-_.!*'(),%=&]+`
const POSTParamsRegex = `\n\r\n(.+)`

func getReqFromParam(respWriter http.ResponseWriter, request *http.Request) db.Request {
	dbConn, err := db.CreateNewDatabaseConnection()
	if err != nil {
		logrus.Warn("Can't connect to database")
		logrus.Error(err)
	}

	defer dbConn.Close()

	info := request.URL.Query()["id"]
	if len(info) < 1 {
		_, _ = fmt.Fprintf(respWriter,
			"Set id param to query in URL to repeat request\nVisit http://localhost/ for more info\n")
		return db.Request{}
	}

	if len(info) > 1 {
		_, _ = fmt.Fprintf(respWriter, "WARN: using first ID\n")
	}

	id, err := strconv.Atoi(info[0])
	if err != nil {
		logrus.Error(err)
	}

	return dbConn.GetReqById(id)
}

func findGETParams(req string) []Param {
	r := regexp.MustCompile(GETParamsRegex)
	matches := r.FindAllString(req, -1)
	return splitKeyAndValue(matches[0][1:])
}

func findPOSTParams(req string) []Param {
	r := regexp.MustCompile(POSTParamsRegex)
	matches := r.FindAllStringSubmatch(req, -1)
	return splitKeyAndValue(matches[0][1])
}

func splitKeyAndValue(matches string) []Param {
	str := strings.Split(matches, "&")
	params := make([]Param, 0)
	for _, paramStr := range str {
		var param Param
		parts := strings.Split(paramStr, "=")
		param.key = parts[0]
		param.value = parts[1]
		params = append(params, param)
	}
	return params
}

func tryCookie(req string) []Param {
	r := regexp.MustCompile(CookieRegex)
	matches := r.FindAllString(req, -1)

	cookieString := matches[0][8:]
	cookies := strings.Split(cookieString, "; ")

	params := make([]Param, 0)
	for _, cookie := range cookies {
		var param Param
		parts := strings.Split(cookie, "=")
		param.key = parts[0]
		param.value = parts[1]
		params = append(params, param)
	}

	return params
}
