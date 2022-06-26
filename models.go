package main

import (
	"os"
	"fmt"
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

type Transcation struct{
	gorm.Model
	TranscationID string
	BillAmount int64 
	Place string 
 	Date string
	Npeople int64 
	Receiver string 
	Rnumber string 
	Payer string 
	Pnumber string 
	Share int64 
	Status string 
}



func InitialMigration(){
	val := os.Getenv("DBSTR")
	db, err := gorm.Open("postgres", val)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	
	}
	defer db.Close()
	db.AutoMigrate(&Transcation{})
	// db.CreateTable(&Transcation{})
}

func CreateConnection() *gorm.DB{
	val := os.Getenv("DBSTR")
	db, err := gorm.Open("postgres", val)
	if err != nil {

		panic("failed to connect database")
	
	}
	return db

}

var DBcon *gorm.DB = CreateConnection()

func Save(payload string) string{
	jsobj := gjson.Parse(payload)
	tid := uuid.New().String()
	
	go jsobj.Get("split").ForEach(func(_,value gjson.Result) bool{
		col := Transcation{
			TranscationID:tid,
			BillAmount:jsobj.Get("billAmont").Int(),
			Place:jsobj.Get("place").String(),
			Date:jsobj.Get("date").String(),
			Npeople:jsobj.Get("nPeople").Int(),
			Receiver:jsobj.Get("spentBy.name").String(),
			Rnumber:jsobj.Get("spentBy.number").String(),
			Payer:"",
			Pnumber:"",
			Share:0,
			Status:"pending",
		}
		valoj := gjson.Parse(value.String())
		col.Payer = valoj.Get("name").String()
		col.Pnumber = valoj.Get("number").String()
		col.Share = valoj.Get("shareAmount").Int()
		DBcon.Create(&col)
		return true
	})
	return `{"transcation_id":"`+tid+`"}`

}

func GetAll(number string,intent string) string {
	var transactions []Transcation
	col := "rnumber = ?"
	payload := ""
	resp := "[]" 
	paid := make(map[string]int64)
	
	if intent != "recivable"{
		col = "pnumber = ?"
	}
	DBcon.Where(col,number).Find(&transactions)
	result , err := json.Marshal(transactions)
	if err != nil {
        fmt.Println(err)
        return ""
    }
	if intent == "recivable"{
		data,_ := sjson.SetRaw("","data",string(result))
		gjson.Get(data,"data").ForEach(func(_,value gjson.Result) bool{
			jsobj := gjson.Parse(value.String())
			tid := jsobj.Get("TranscationID").String()
			if !gjson.Get(payload,tid).Exists(){
				payload ,_ = sjson.Set(payload,tid+".place",value.Get("Place").String())
				payload ,_ = sjson.Set(payload,tid+".transcationId",tid)
				payload ,_ = sjson.Set(payload,tid+".date",value.Get("Date").String())
				payload ,_ = sjson.Set(payload,tid+".billAmount",value.Get("BillAmount").Int())
				payload ,_ = sjson.Set(payload,tid+".nPeople",value.Get("Npeople").Int())
			}
			acc,_ := sjson.Set("{}","number",value.Get("Pnumber").String())
			acc,_ = sjson.Set(acc,"name",value.Get("Payer").String())
			acc,_ = sjson.Set(acc,"share",value.Get("Share").Int())
			acc,_ = sjson.Set(acc,"status",value.Get("Status").String())
			if value.Get("Status").String() == "paid"{
				paid[tid]++
			}
			payload ,_ = sjson.SetRaw(payload,tid+".split.-1",acc)
			return true
		})
		data,_ = sjson.SetRaw("","data",payload)
		gjson.Get(data,"data").ForEach(func(key,value gjson.Result) bool{
			bill := ""
			if paid[key.String()] == value.Get("nPeople").Int(){
				bill,_ = sjson.Set(value.String(), "OverallPaymentStatus","Paid") 
			}else if paid[key.String()]>0{
				bill,_ = sjson.Set(value.String(), "OverallPaymentStatus","Partial") 
			}else{
				bill,_ = sjson.Set(value.String(), "OverallPaymentStatus","Pending") 
			}
			resp,_ = sjson.SetRaw(resp,"-1",bill)
			return true
		})
		
	}else{
		data,_ := sjson.SetRaw("","data",string(result))
		gjson.Get(data,"data").ForEach(func(_,value gjson.Result) bool{
			payload,_ = sjson.Set("","billAmount",value.Get("BillAmount").Int())
			payload,_ = sjson.Set(payload,"place",value.Get("Place").String())
			payload,_ = sjson.Set(payload,"date",value.Get("Date").String())
			payload,_ = sjson.Set(payload,"spentBy.name",value.Get("Receiver").String())
			payload,_ = sjson.Set(payload,"spentBy.number",value.Get("Rnumber").String())
			payload,_ = sjson.Set(payload,"nPeople",value.Get("Npeople").String())
			payload,_ = sjson.Set(payload,"myShare",value.Get("Share").Int())
			payload,_ = sjson.Set(payload,"status",value.Get("Status").String())
			payload,_ = sjson.Set(payload,"Paylink","/pay/"+value.Get("Pnumber").String()+"/"+value.Get("TranscationID").String())
			resp,_ = sjson.SetRaw(resp,".-1",payload)
			return true
		})

	}
	return resp
}

func UpdateStatus(tid string,number string) {
	DBcon.Model(&Transcation{}).Where("transcation_id = ? AND pnumber = ?",tid,number).Update("status","paid")
}