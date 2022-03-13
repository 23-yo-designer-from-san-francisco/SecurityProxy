package repeatReqHandlers

import (
	"Proxy/db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

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
