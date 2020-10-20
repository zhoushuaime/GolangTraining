# Go  build

 - go build
 - GOOS="linux" go build -o tidb-linux  main.go (on macOS)
 
# 说明
 - Golang的相对路径是相对于执行命令时的目录，所以当把文件放到其他地方时会有找不到文件，此处用了相对路径的方法去解决。
 - 直接go build即可，因为go run 得到的是临时路径，仍然会报找不到路径错误。
 - 若要运行test文件，请设置环境变量RUN_MODE=test。