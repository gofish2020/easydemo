package main

import (
	"context"
	"easydemo/client"
	"easydemo/proto/raftpb"
	"easydemo/server"
	"os"
	"time"
)

func main() {

	len := len(os.Args)
	if len > 1 { // server

		switch os.Args[1] { //运行指令 go run main.go 1
		case "1":
			server.Start()
		}
	} else { // client
		// 运行指令 go run main.go
		cli := client.NewClient()
		cli.Send(context.Background(), &raftpb.RaftMessage{
			MsgType: raftpb.MessageType_HEARTBEAT,
		})

		go cli.Consensus() // 在tcp上开一个client stream

		time.Sleep(1 * time.Second)
		go cli.Consensus() // 在tcp上开另外一个client stream

		time.Sleep(1 * time.Minute)

	}
}