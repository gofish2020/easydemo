package client

import (
	"bufio"
	"context"
	"easydemo/proto/hellopb"
	"easydemo/proto/raftpb"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
}

// 创建 grpc 客户端
func NewClient() *Client {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials())) // 不安全证书

	// 客户端拨号
	conn, err := grpc.NewClient(":8088", opts...)
	if err != nil {
		return nil
	}

	return &Client{
		conn: conn,
	}
}

func (t *Client) Send(ctx context.Context, req *raftpb.RaftMessage) (*raftpb.RaftMessage, error) {

	// 创建 grpc Hello 客户端
	client := hellopb.NewHelloClient(t.conn)

	// 发送 req 返回res结果
	res, err := client.Send(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 打印结果
	msgType := res.MsgType
	fmt.Println(msgType.Number(), msgType.String())

	// 再发送一次  req 返回res结果
	req.MsgType = raftpb.MessageType_INSTALL_SNAPSHOT
	res, err = client.Send(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 打印结果
	msgType1 := res.MsgType
	fmt.Println(msgType1.Number(), msgType1.String())

	return res, nil

}

func (t *Client) Consensus() {

	// 创建 grpc Raft 客户端
	raftClient := raftpb.NewRaftClient(t.conn)

	// 可以看到这里和我们普通的 rpc调用直接返回结果不一样，而是返回了一个新对象
	raftClientStream, err := raftClient.Consensus(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	// 通过 raftClientStream 我们可以持续的与服务端【发送和接收数据】
	from := uint64(time.Now().UnixMilli())
	msg1 := &raftpb.RaftMessage{}
	msg1.From = from
	msg1.Term = 1
	raftClientStream.Send(msg1) // 发送一个数据包

	for {
		msg, err := raftClientStream.Recv() // 接收
		if err == io.EOF {
			raftClientStream.CloseSend()
			raftClientStream = nil
			return
		}

		if err != nil {
			raftClientStream.CloseSend()
			raftClientStream = nil
			return
		}
		fmt.Println("client", from, msg.Term) // 这个 Term值被服务端不断的 +1
		raftClientStream.Send(msg)            // 发送
		time.Sleep(5 * time.Second)
	}

}

func (t *Client) SendFile(fileNanme string) error {

	// 创建 grpc File 客户端
	fileClient := raftpb.NewFileClient(t.conn)
	// 返回一个 fileStream 对象
	fileStream, err := fileClient.Sendfile(context.Background())
	if err != nil {
		return err
	}

	// 打开本地文件
	file, err := os.Open(fileNanme)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建一个 Reader
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024) // 每次读取 1024 字节

	for {
		// 读取数据到缓冲区，每次读取 1024 字节
		_, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF { // 文件读取完毕
				fileinfo := &raftpb.FileContext{}
				fileinfo.Islastframe = true // 最后一帧标记
				fileinfo.Context = buffer
				pos := strings.LastIndex(file.Name(), ".")
				if pos != -1 {
					fileinfo.Ext = file.Name()[pos:]
				}
				fileStream.Send(fileinfo) // 发送最后一帧文件数据
			}
			break
		}

		// 持续的发送文件数据
		fileinfo := &raftpb.FileContext{}
		fileinfo.Islastframe = false //不是最后一帧
		fileinfo.Context = buffer
		fileStream.Send(fileinfo)

	}
	// 发送完毕，等待【接收服务端】的返回
	res, err := fileStream.CloseAndRecv()
	if err != nil {
		fmt.Println(res, err)
		return err
	}
	fmt.Println(res)
	return nil
}
