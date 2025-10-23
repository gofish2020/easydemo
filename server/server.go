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

// Raft对象：这个就是演示双向stream的业务实现
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
		msg.Term++ // 将消息的 term+1 再发回给客户端
		serverStream.Send(msg)
		time.Sleep(5 * time.Second)
	}
}

// Hello对象：我们平时写的最多远程过程调用的业务实现
type Hello struct {
	hellopb.UnimplementedHelloServer
}

func (t *Hello) Send(ctx context.Context, r *raftpb.RaftMessage) (*raftpb.RaftMessage, error) {
	return &raftpb.RaftMessage{
		MsgType: r.MsgType,
	}, nil
}

// File对象演示：客户端通过stream发送文件到服务端，最后服务端接收完成以后，回复一个确认消息
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

	req.SendAndClose(&raftpb.FileInfoResp{ // 返回文件名，告知已经接收完成
		Isok: true,
		Name: fileName + ext,
	})
	return nil
}

func Start() {

	// 8088监听端口
	lis, err := net.Listen("tcp", ":8088")

	if err != nil {

		fmt.Println("监听失败")
		return
	}

	// 创建 gprc 服务
	var opts []grpc.ServerOption
	srv := grpc.NewServer(opts...)

	// 注册服务端业务处理对象
	raftpb.RegisterRaftServer(srv, &Raft{})
	hellopb.RegisterHelloServer(srv, &Hello{})
	raftpb.RegisterFileServer(srv, &File{})

	// 启动服务
	fmt.Println("启动服务...")
	err = srv.Serve(lis)
	if err != nil {
		fmt.Println("启动失败")
		return
	}
}
