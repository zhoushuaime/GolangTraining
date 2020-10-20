package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
	"sync"
	"zhoushuai.com/GolangTraining/tidb/client"
	"zhoushuai.com/GolangTraining/tidb/tool"
)

const (
	User       = "root"
	Pass       = ""
	Port       = 4000
	Host       = "127.0.0.1"
	Database   = "test"
	queryParam = "charset=utf8mb4&parseTime=True&loc=Local"
)

// SQLDao ...
type SQLDao struct {
	DB *gorm.DB
}

// GetDBConn ...
func GetDBConn() (*SQLDao, error) {

	sqlDao := new(SQLDao)
	var err error
	sqlDao.DB, err = gorm.Open("mysql", //tidb兼容
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", User, Pass, Host, Port, Database, queryParam))
	return sqlDao, err
}

var sqlStr = `select * from test`

// Exec ...
func (dao *SQLDao) Find(out interface{}, cond ...interface{}) error {
	return dao.DB.Raw(sqlStr).Scan(out).Error
}

// Update ...
func (dao *SQLDao) Update(sqlStr string) error {
	return dao.DB.Exec(sqlStr).Error

}

// Out ...
type Out struct {
	Id int64 `json:"id" gorm:"column:id"`
	A  int64 `json:"a" gorm:"column:a"`
	B  int64 `json:"b" gorm:"column:b"`
}

func main() {
	db, err := GetDBConn()
	if err != nil {
		panic(fmt.Sprintf("get db error:%v", err))
	}

	fmt.Println(strings.Repeat("=", 50))

	cli := client.NewTransClient()
	transDetail, err := cli.GetAllTransDetail()
	if err != nil {
		panic(fmt.Errorf("GetAllTransDetail failed,err:%v", err))
	}

	input := make([]interface{}, 0)
	for i := 0; i < len(transDetail); i++ {
		input = append(input, transDetail[i])
	}

	// 全排列数据
	outPermute := tool.Permutation(input)

	updateSqls := make([]string, 0)

	for _, values := range outPermute {
		for _, value := range values {
			if v, ok := value.(client.TransDetail); ok {
				updateSqls = append(updateSqls, v.Trans)
			}
		}
	}

	// 模拟2个client
	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for _, sql := range updateSqls {
				if err = db.Update(sql); err == nil {
					fmt.Printf("execute success,sql:%s\n", sql)
				} else {
					fmt.Printf("execute failed,sql:%s\n", sql)
				}
			}
		}()
	}

	wg.Wait()

	s := strings.Repeat("=", 50)
	fmt.Println(fmt.Sprintf("%s finished %s", s, s))

}
