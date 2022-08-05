package main

import (
	"fmt"

	structs "github.com/oceanm7n/planes/common"
	c "github.com/oceanm7n/planes/connection"
	"github.com/ssgreg/logf"
)

func logger() (*logf.Logger, logf.ChannelWriterCloseFunc) {
	writer, writerClose := logf.NewChannelWriter.Default()

	return logf.NewLogger(logf.LevelInfo, writer), writerClose
}

func main() {
	logger, writerClose := logger()
	defer writerClose()

	logger.Info("Application started")

	query := structs.Query{
		From: "TLV",
		To:   "SVO",
		Date: "2022-09-01",
	}

	eth := c.NewConnection("Etihad", "ETH", query)

	res := eth.SendRequest()

	response_string := eth.ParseRequest(res)

	fmt.Println(response_string)
}
