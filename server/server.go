package server

import (
	"context"
	"easydemo/proto/hellopb"
	"easydemo/proto/raftpb"
	"fmt"
	"io"
	"net"
	"os"
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

type File struct {
	raftpb.UnimplementedFileServer
}

func (t *File) Sendfile(req raftpb.File_SendfileServer) error {

	fileName := fmt.Sprintf("%d", time.Now().UnixMilli())
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm) // 创建临时文件
	if err != nil {
		return err
	}

	ext := ""
	for {
		fileContext, err := req.Recv()
		if err == io.EOF {
			file.Close()
			os.Remove(file.Name())
			fmt.Println("eof", err)
			return err
		}
		if err != nil {
			file.Close()
			os.Remove(file.Name())
			fmt.Println("err", err)
			return err
		}

		file.Write(fileContext.Context) // 接收文件数据

		if fileContext.Islastframe { // 最后一帧数据
			ext = fileContext.Ext // 记录文件后缀
			break
		}
	}

	file.Close()
	if ext != "" { //存在后缀，重命名文件
		os.Rename(fileName, fileName+ext)
	}

	req.SendAndClose(&raftpb.FileInfoResp{ // 返回文件名
		Isok: true,
		Name: fileName + ext,
	})
	return nil
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

	raftpb.RegisterFileServer(srv, &File{})

	err = srv.Serve(lis)
	if err != nil {
		fmt.Println("启动失败")
		return
	}
}
