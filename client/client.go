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

	raftClient raftpb.RaftClient
}

func NewClient() *Client {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials())) // 不安全证书

	// 拨号
	conn, err := grpc.NewClient(":8088", opts...)
	if err != nil {
		return nil
	}

	return &Client{
		conn:       conn,
		raftClient: raftpb.NewRaftClient(conn),
	}
}

func (t *Client) Send(ctx context.Context, req *raftpb.RaftMessage) (*raftpb.RaftMessage, error) {

	client := hellopb.NewHelloClient(t.conn)
	res, err := client.Send(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	msgType := res.MsgType
	fmt.Println(msgType.Number(), msgType.String())

	req.MsgType = raftpb.MessageType_INSTALL_SNAPSHOT
	res, err = client.Send(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	msgType1 := res.MsgType
	fmt.Println(msgType1.Number(), msgType1.String())

	return res, nil

}

func (t *Client) Consensus() {

	raftClientStream, err := t.raftClient.Consensus(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	from := uint64(time.Now().UnixMilli())
	msg1 := &raftpb.RaftMessage{}
	msg1.From = from
	msg1.Term = 1
	raftClientStream.Send(msg1)

	for {
		msg, err := raftClientStream.Recv() // 收
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
		fmt.Println("client", from, msg.Term)
		msg.Term++
		raftClientStream.Send(msg) // 发
		time.Sleep(5 * time.Second)
	}

}

func (t *Client) SendFile(fileNanme string) error {

	fileClient := raftpb.NewFileClient(t.conn)
	fileStream, err := fileClient.Sendfile(context.Background())
	if err != nil {
		return err
	}

	file, err := os.Open(fileNanme)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建一个 Reader
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024) // 每次读取 1024 字节

	for {
		// 读取数据到缓冲区
		_, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fileinfo := &raftpb.FileContext{}
				fileinfo.Islastframe = true
				fileinfo.Context = buffer
				pos := strings.LastIndex(file.Name(), ".")
				if pos != -1 {
					fileinfo.Ext = file.Name()[pos:]
				}
				fileStream.Send(fileinfo)
			}
			break
		}

		fileinfo := &raftpb.FileContext{}
		fileinfo.Islastframe = false
		fileinfo.Context = buffer
		fileStream.Send(fileinfo)

	}
	res, err := fileStream.CloseAndRecv()
	if err != nil {
		fmt.Println(res, err)
		return err
	}
	fmt.Println(res)
	return nil
}
