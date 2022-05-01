package ctl

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func Health(w http.ResponseWriter, r *http.Request) {
	var (
		name string
		err  error
	)

	if name, err = os.Hostname(); err != nil {
		log.Printf("get hostname err:%+v", err)
		name = "未知"
	}

	data := map[string]string{}
	data["msg"] = "ok"
	data["host"] = name

	if err = json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
