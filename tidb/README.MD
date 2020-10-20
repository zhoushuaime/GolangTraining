# Go  build

 - go build
 - GOOS="linux" go build -o tidb-linux  main.go (on macOS)
 
# 说明
 - Golang的相对路径是相对于执行命令时的目录，所以内部用了相对路径方法。
 - 直接go build即可，go run 得到的是临时路径，会报找不到路径错误。
 - 若要运行test文件，请设置环境变量RUN_MODE=test。