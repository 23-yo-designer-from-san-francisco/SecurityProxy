package repeatReqHandlers

import (
	"Proxy/utils"
	"bufio"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func ExecRepReq(respWriter http.ResponseWriter, request *http.Request) {
	req := getReqFromParam(respWriter, request)

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	reqReader := bufio.NewReader(strings.NewReader(req.Request))
	buffer, err := http.ReadRequest(reqReader)
	if err != nil {
		logrus.Error(err)
	}

	httpReq, err := http.NewRequest(buffer.Method, req.Host, buffer.Body)
	if err != nil {
		logrus.Error(err)
	}

	utils.CopyHeaders(buffer.Header, httpReq.Header)

	resp, err := client.Do(httpReq)
	if err != nil {
		logrus.Error(err)
	}

	utils.CopyHeaders(resp.Header, respWriter.Header())
	respWriter.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(respWriter, resp.Body)
	_ = resp.Body.Close()

}
