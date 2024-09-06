package server

import (
	"context"
	"easydemo/proto/hellopb"
	"easydemo/proto/raftpb"
	"fmt"
	"io"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Raft struct {
	raftpb.UnimplementedRaftServer
}

func (t *Raft) Consensus(serverStream raftpb.Raft_ConsensusServer) error {

	// 接收消息
	for {
		msg, err := serverStream.Recv()
		if err == io.EOF {
			fmt.Println("server1", err)
			return nil
		}
		if err != nil {
			fmt.Println("server2", err)
			return nil
		}

		fmt.Println("server", msg.From, msg.Term)
		msg.Term++
		serverStream.Send(msg)
		time.Sleep(5 * time.Second)
	}
}

type Hello struct {
	hellopb.UnimplementedHelloServer
}

func (t *Hello) Send(ctx context.Context, r *raftpb.RaftMessage) (*raftpb.RaftMessage, error) {
	return &raftpb.RaftMessage{
		MsgType: r.MsgType,
	}, nil
}

func Start() {

	lis, err := net.Listen("tcp", ":8088")

	if err != nil {

		fmt.Println("监听失败")
		return
	}

	var opts []grpc.ServerOption
	srv := grpc.NewServer(opts...)

	raftpb.RegisterRaftServer(srv, &Raft{})

	hellopb.RegisterHelloServer(srv, &Hello{})

	err = srv.Serve(lis)
	if err != nil {
		fmt.Println("启动失败")
		return
	}
}
