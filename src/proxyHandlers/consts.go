package proxyHandlers

const reqDumpErr = "Request dump error: "
const dbConnectErr = "Can't connect to database"

const httpOkResponse = "HTTP/1.1 200 Connection established\r\n\r\n"

const scriptsDir = "/scripts"
const certsDir = "/certs/"
const genCertScript = "/gen_cert.sh"
const certKey = "/cert.key"

var skipHeaderList = []string{
	"proxy-connection",
}
