package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

const DefaultLimit = 10 // 默认每页10条

// PageOptions 。。。
type PageOptions struct {
	QuerySql   string        `json:"query_sql"`
	Conditions []interface{} `json:"conditions"`
	Page       int           `json:"page"`
	Limit      int           `json:"limit"`
}

// ResponseData ...
type ResponseData struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

//PageQuery 分页查询
func PageQuery(db *gorm.DB, p PageOptions) (ResponseData, error) {
	var Page, Limit int
	var r ResponseData
	if p.Page == 0 {
		Page = 1
	} else {
		Page = p.Page
	}
	if p.Limit == 0 {
		Limit = DefaultLimit
		p.Limit = Limit
	} else {
		Limit = p.Limit
	}
	var sum int
	var c []interface{}
	// 查询时，对时间参数进行格式化处理
	for _, info := range p.Conditions {
		var temp interface{}
		switch info.(type) {
		case time.Time:
			temp = info.(time.Time).Format("2006-01-02 15:04:05")
		default:
			temp = info
		}
		c = append(c, temp)
	}
	cRows, err := db.Raw(p.QuerySql, c...).Rows()
	if err != nil && err != gorm.ErrRecordNotFound {
		return r, err
	} else if err == gorm.ErrRecordNotFound {
		return r, nil
	}

	for cRows.Next() {  // 获取总条目
		sum++
	}
	r.Total = sum

	begin := (Page - 1) * Limit
	// 拼接sql limit 10 offset 0 或 limit 0,10
	pageSql := p.QuerySql + " limit " + strconv.Itoa(Limit) + " offset " + strconv.Itoa(begin)

	// 分页查询最终条目
	rows, err := db.Raw(pageSql, p.Conditions...).Rows()
	if err != nil && err != gorm.ErrRecordNotFound {
		return r, err
	} else if err == gorm.ErrRecordNotFound {
		return r, nil
	}

	defer rows.Close()

	var ret []interface{}
	columns, err := rows.Columns()
	if err != nil {
		return r, err
	}

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		rows.Scan(scanArgs...)
		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				switch col.(type) {
				case []byte:
					record[columns[i]] = string(col.([]byte))
				default:
					record[columns[i]] = col
				}
			}
		}
		ret = append(ret, record)
	}

	r.List = ret
	return r, nil
}
