package repeatReqHandlers

import (
	"Proxy/db"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func SendRequestList(respWriter http.ResponseWriter, _ *http.Request) {
	dbConn, err := db.CreateNewDatabaseConnection()
	if err != nil {
		logrus.Warn("Can't connect to database")
		logrus.Error(err)
	}

	defer dbConn.Close()

	reqList, err := dbConn.GetRequestList()
	if err != nil {
		logrus.Warn("Can't get data from DB")
		_, _ = fmt.Fprintf(respWriter, "Can't get request info\n")
		return
	}

	if len(reqList) == 0 {
		_, _ = fmt.Fprintf(respWriter, "No requests saved\n")
		return
	}

	_, _ = fmt.Fprintf(respWriter,
		`
				<html>
				<head>
				<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.0.0/dist/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
				</head>
				<table class="table">
					<thead>
						<td>ID</td>
						<td>Host</td>
						<td>Request</td>
						<td>Execute</td>
						<td>Vulnerability</td>
					</thead>
				<tbody>`,
	)
	for i, req := range reqList {
		_, _ = fmt.Fprintf(respWriter,
			`<tr>
					<td>%d</td>
					<td>%s</td>
					<td>%s</td>
					<td>
						<a href="/req?id=%d">
							Execute!
						</a>
					</td>
					<td>
						<a href="/scan?id=%d">
							Exploit vulnerability
						</a>
					</td>
			</tr>
			`,
			i+1,
			req.Host,
			req.Request,
			req.Id,
			req.Id)
	}
	_, _ = fmt.Fprintf(respWriter,
		"</tbody></table></html>",
	)
}
