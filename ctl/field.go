package ctl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Field struct {
	Id string `json:"id"`
	// 名称
	Title string `json:"title"`
	// 类型
	Type string `json:"type"`
	// 表达式
	Exps string `json:"exps"`
}

var fields = []*Field{
	{Id: "name", Title: "姓名", Type: "string", Exps: ""},
	{Id: "dept", Title: "部门", Type: "string", Exps: ""},
	{Id: "baseSalary", Title: "基本工资", Type: "number", Exps: ""},
	{Id: "perf", Title: "绩效", Type: "number", Exps: ""},
	{Id: "bonus", Title: "奖金", Type: "number", Exps: "IF(perf<=3.5,500,1000)"},
	{Id: "salary", Title: "工资", Type: "number", Exps: "SUM(baseSalary,bonus)"},
}

func Fields(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(fields); err != nil {
		log.Fatal(err)
	}
}

func FieldUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		field *Field
		err   error
		body  []byte
	)

	field = &Field{}
	if body, err = ioutil.ReadAll(r.Body); err != nil {
		log.Printf("request body err:%+v \n", err)
	} else if err = json.Unmarshal(body, field); err != nil {
		log.Printf("json unmarshal err:%+v \n", err)
	}

	for i, f := range fields {
		if f.Id == field.Id {
			fields[i] = field
		}
	}
	res := map[string]string{}
	res["msg"] = "ok"
	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal(err)
	}
}
