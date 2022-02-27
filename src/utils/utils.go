package utils

import (
	"math/rand"
	"net/http"
	"os/exec"
	"strconv"
)

const randRange int = 1e10

func GenProxyCert(scriptPath, scriptName, host, savePath string) error {
	genCmd := exec.Command(scriptPath+scriptName, host, strconv.Itoa(rand.Intn(randRange)), scriptPath, savePath)
	_, err := genCmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func CopyHeaders(from, to http.Header) {
	for header, values := range from {
		for _, value := range values {
			to.Add(header, value)
		}
	}
}
