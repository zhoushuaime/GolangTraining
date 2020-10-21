package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
	"sync"
	"zhoushuai.com/GolangTraining/tidb/client"
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

type TransInfo struct {
	ClientId       int
	CurrentVersion int
	CurrentCount   int
	CurrentData    client.TransDetail
}

func main() {
	db, err := GetDBConn()
	if err != nil {
		panic(fmt.Sprintf("get db error:%v", err))
	}

	client1 := client.NewTransClient(1)
	trans1, err := client1.GetTransDetail()
	if err != nil {
		panic(fmt.Errorf("GetAllTransDetail failed,err:%v", err))
	}

	client2 := client.NewTransClient(2)
	trans2, err := client2.GetTransDetail()
	if err != nil {
		panic(fmt.Errorf("GetAllTransDetail failed,err:%v", err))
	}

	var mux sync.Mutex
	ClientMap := make(map[int]TransInfo)

	clientChan1 := make(chan client.TransDetail, len(trans1))
	clientChan2 := make(chan client.TransDetail, len(trans2))

	go func() {
		for i := 0; i < len(trans1); i++ {
			clientChan1 <- trans1[i]
		}
		//close(clientChan1)
	}()

	go func() {
		for i := 0; i < len(trans2); i++ {
			clientChan2 <- trans2[i]
		}
		//close(clientChan2)
	}()
	var executeUpdate func(sql string)
	executeUpdate = func(sql string) {
		if err = db.Update(sql); err == nil {
			fmt.Printf("execute success,sql:%s\n", sql)
		} else {
			fmt.Printf("execute failed,sql:%s\n", sql)
		}
	}

	var handle func(minVersion, l int, clientChan chan client.TransDetail)
	handle = func(minVersion, l int, clientChan chan client.TransDetail) {
		for {
			select {
			case res := <-clientChan:
				go func(res client.TransDetail) {
					mux.Lock()
					defer mux.Unlock()
					key := res.ClientId
					value, exist := ClientMap[key]
					if exist {
						var temp TransInfo
						temp.CurrentData = res
						temp.CurrentVersion = res.TransNum
						if value.CurrentVersion > res.TransNum {
							fmt.Printf("back to queque,currentVersion:%v,res.TransNum:%v,currentCount:%v\n", value.CurrentVersion, res.TransNum, value.CurrentCount)
							clientChan <- value.CurrentData // 塞回去
							temp.CurrentCount = value.CurrentCount
							ClientMap[key] = temp
							return
						}

						temp.CurrentCount = value.CurrentCount + 1
						ClientMap[key] = temp
						fmt.Printf("%+v\n", value.CurrentData.Trans)
						executeUpdate(value.CurrentData.Trans)
						if temp.CurrentCount == l {
							fmt.Printf("%+v\n", temp.CurrentData.Trans)
							executeUpdate(temp.CurrentData.Trans)

						}

						return
					}

					if l == 1 {
						fmt.Printf("%+v\n", res.Trans)
						executeUpdate(res.Trans)
						return
					}
					var temp TransInfo
					temp.CurrentData = res
					temp.CurrentVersion = res.TransNum
					temp.CurrentCount = value.CurrentCount + 1

					if res.TransNum > minVersion {
						temp.CurrentVersion = minVersion
						clientChan <- res
						return
					}
					ClientMap[key] = temp

				}(res)
			}

		}
	}

	// 模拟两个客户端，根据clientId区分
	for i := 0; i < 2; i++ {
		go func(clientId int) {
			switch clientId {
			case 0:
				handle(trans1[0].TransNum, len(trans1), clientChan1)
			case 1:
				handle(trans1[0].TransNum, len(trans2), clientChan2)
			}

		}(i)
	}

	s := strings.Repeat("=", 50)
	fmt.Println(fmt.Sprintf("%s finished %s", s, s))
	select {}

}

/*
type Tt struct {
	Data    string
	Num     int
	ClintId int
}

type ClientInfo struct {
	ClientId       int
	CurrentVersion int
	CurrentCount   int
	CurrentData    Tt
}*/

/*func main1() {
	a := []Tt{
		{
			Data:    "A1",
			Num:     1,
			ClintId: 1,
		},
		{
			Data:    "A2",
			Num:     2,
			ClintId: 1,
		},
		{
			Data:    "A3",
			Num:     3,
			ClintId: 1,
		},
	}
	b := []Tt{
		{
			Data:    "C1",
			Num:     1,
			ClintId: 2,
		},
	}

	var mux sync.Mutex
	ClientMap := make(map[int]ClientInfo)

	clientChan1 := make(chan Tt, len(a))
	clientChan2 := make(chan Tt, len(b))

	go func() {
		for i := 0; i < len(a); i++ {
			clientChan1 <- a[i]
		}
		//close(clientChan1)
	}()

	go func() {
		for i := 0; i < len(b); i++ {
			clientChan2 <- b[i]
		}
		//close(clientChan2)
	}()

	var handle func(minVersion, l int, clientChan chan Tt)
	handle = func(minVersion, l int, clientChan chan Tt) {
		for {
			select {
			case res := <-clientChan:
				go func(res Tt) {
					mux.Lock()
					defer mux.Unlock()
					key := res.ClintId
					value, exist := ClientMap[key]
					if exist {
						var temp ClientInfo
						temp.CurrentData = res
						temp.CurrentVersion = res.Num
						if value.CurrentVersion > res.Num {
							fmt.Printf("back,currentVersion:%v,res.Num:%v,currentCount:%v\n", value.CurrentVersion, res.Num, value.CurrentCount)
							clientChan <- value.CurrentData // 塞回去
							temp.CurrentCount = value.CurrentCount
							ClientMap[key] = temp
							return
						}

						temp.CurrentCount = value.CurrentCount + 1
						ClientMap[key] = temp
						fmt.Printf("%+v\n", value.CurrentData)

						if temp.CurrentCount == l {
							fmt.Printf("%+v\n", temp.CurrentData)
						}

						return
					}

					if l == 1 {
						fmt.Printf("%+v\n", res)
						return
					}
					var temp ClientInfo
					temp.CurrentData = res
					temp.CurrentVersion = res.Num
					temp.CurrentCount = value.CurrentCount + 1

					if res.Num > minVersion {
						temp.CurrentVersion = minVersion
						clientChan <- res
						return
					}
					ClientMap[key] = temp

				}(res)
			}

		}
	}

	for i := 0; i < 2; i++ {
		go func(i int) {
			switch i {
			case 0:
				handle(a[0].Num, len(a), clientChan1)
			case 1:
				handle(b[0].Num, len(b), clientChan2)
			}

		}(i)
	}

	fmt.Println(strings.Repeat("=", 50))

	select {}
}*/
