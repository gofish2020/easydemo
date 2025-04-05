# 分开输出两个文件 *.pb.go 和 *_grpc.pb.go（和合并一起输出一样的效果）
protoc -I . --go_out=module=easydemo:.  --go-grpc_out=module=easydemo:.  proto/*.proto
