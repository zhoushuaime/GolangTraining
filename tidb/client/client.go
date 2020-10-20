package client

import (
	"fmt"
	"os"
	"strings"
	"zhoushuai.com/GolangTraining/tidb/tool"
)

var (
	sql1Path = "/sql/sql1.sql"
	sql2Path = "/sql/sql2.sql"
)

func init() {
	runMode := os.Getenv("RUN_MODE")
	sql1Path = fmt.Sprint(tool.GetCurrentDir(), sql1Path)
	sql2Path = fmt.Sprint(tool.GetCurrentDir(), sql2Path)
	// 设置环境变量来区分运行的是test还是go build，以免执行test时，生成的相对路径找不到
	if runMode == "test" {
		sql1Path = "../sql/sql1.sql"
		sql2Path = "../sql/sql2.sql"
	}

}

// TransClient ...
type TransClient struct {
	TaskId   int64
	ClientId int
}

// TransDetail ...
type TransDetail struct {
	Trans    string
	ClientId int
	RandNum  int64
}

// NewTransClient ...
func NewTransClient(clientIds ...int) *TransClient {
	clientId := 0
	if len(clientIds) > 0 {
		clientId = clientIds[0]
	}
	return &TransClient{
		TaskId:   tool.RandNum(),
		ClientId: clientId,
	}
}

// GetAllTransDetail ...
func (trans *TransClient) GetAllTransDetail() ([]TransDetail, error) {
	result := make([]TransDetail, 0)
	for i := 1; i < 3; i++ {
		trans.ClientId = i
		res, err := trans.GetTransDetail()
		if err != nil {
			return nil, err
		}
		result = append(result, res...)
	}
	return result, nil
}

// GetTransDetail ...
func (trans *TransClient) GetTransDetail() ([]TransDetail, error) {

	result := make([]TransDetail, 0)
	fileName := ""
	switch trans.ClientId {
	case 1:
		fileName = sql1Path
	case 2:
		fileName = sql2Path
	default:
		return nil, fmt.Errorf("unsupported transId:%v", trans.TaskId)
	}
	res, err := tool.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	transSql := strings.Split(res, "\n")

	for _, sql := range transSql {
		transDetail := TransDetail{}
		transDetail.Trans = sql
		transDetail.ClientId = trans.ClientId
		transDetail.RandNum = tool.RandNum()
		result = append(result, transDetail)
	}

	return result, nil
}
