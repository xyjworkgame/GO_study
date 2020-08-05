
# 生成proto
protoc -I . --go_out=plugins=grpc:. ./hello.proto