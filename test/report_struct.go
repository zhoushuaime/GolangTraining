package test

import (
	"fmt"
	"sync"
	"time"
)

// IdList ...
type IdList struct {
	IdList []string `json:"id_list"`
}

// ReportList ...
type ReportList struct {
	ReportList []struct {
		Vin   string `json:"vin"`
		Theme string `json:"theme"`
	} `json:"report_list"`
}

// GetIDParams ...
type GetIDParams struct {
	Province  string `json:"province"`
	City      string `json:"city"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Vin       string `json:"vin"`
	MaxRecord int    `json:"max_record"`
}

// GetReportParams ...
type GetReportParams struct {
	IdList []string `json:"id_list"`
}

// Vin ...
type Vin struct {
	Vin []string `json:"vin"`
}

// Report  插入数据库的结构体...
type Report struct {
	ID               string       `json:"id" bson:"id"`                   // 诊断报告id
	Resource         string       `json:"resource" bson:"resource"`       // 数据来源
	FaultCodes       []FaultCodes `json:"fault_codes" bson:"fault_codes"` // 故障码
	FaultN           int          `json:"fault_n" bson:"fault_n"`         // 故障码个数
	RecStatus        int          `json:"rec_status" bson:"rec_status"`
	FlagCustomized   int          `json:"flag_customized" bson:"flag_customized"`
	Latitude         string       `json:"latitude" bson:"latitude"`
	Longitude        string       `json:"longitude" bson:"longitude"`
	FlowN            int          `json:"flow_n" bson:"flow_n"`                         // 数据流个数
	DataFlow         []DataFlow   `json:"data_flow" bson:"data_flow"`                   //数据流
	DiagnosticUserId string       `json:"diagnostic_user_id" bson:"diagnostic_user_id"` // 诊断用户ID
	Language         string       `json:"language" bson:"language"`
	ProSerialNo      string       `json:"pro_serial_no" bson:"pro_serial_no"`   // 诊断设备ID
	WalletAddress    string       `json:"wallet_address" bson:"wallet_address"` // 钱包地址
	UpdateTime       string       `json:"update_time" bson:"update_time"`       // 诊断报告生成时间，格式:YYYY-MM-DD hh:mm:ss
	Theme            string       `json:"theme" bson:"theme"`                   // 报告主题
	Vin              string       `json:"vin" bson:"vin"`                       // 汽车Vin码
	MileAge          int          `json:"mileage" bson:"mileage"`
	SoftwareVersion  string       `json:"software_version" bson:"software_version"`
	IsBussinessReads int          `json:"is_bussiness_reads" bson:"is_bussiness_reads"`
	IsOwnerReads     int          `json:"is_owner_reads" bson:"is_owner_reads"`
	IsPay            int          `json:"is_pay" bson:"is_pay"`
}

// FaultCodes ..
type FaultCodes struct {
	Code        string `json:"code" bson:"code"`
	Description string `json:"description" bson:"description"`
	FaultId     string `json:"fault_id" bson:"fault_id"`
	Status      string `json:"status" bson:"status"`
}

// DataFlow DataFlow
type DataFlow struct {
	Unit   string `json:"unit" bson:"unit"` // 百分比
	FlowID string `json:"flow_id" bson:"flow_id"`
	Name   string `json:"name" bson:"name"`
	Value  string `json:"value" bson:"value"`
}

// ReportLists ...
type ReportLists struct {
	ReportList []Report `json:"report_list"`
}

//==============================================================================
const reportCollection = "diagnose_report_tmp"

// CName ...
func (*Report) CName() string {
	return reportCollection
}

type Config struct {
	URL string
}

var config *Config
var oSingle sync.Once
var TT = time.Now()
//func init() {
//	config = new(Config)
//	config.URL = "test"
//}

/**
单例模式 懒汉式
 */
// NewConfig ...
func NewConfig() *Config {
	if config == nil {
		fmt.Println("config is nil")
		config =  new(Config)
		return config
	}
	fmt.Println("config is not nil")
	return config
}

func GetInstance() *Config {
	oSingle.Do(
		func() {
			config = new(Config)
		})
	return config
}
