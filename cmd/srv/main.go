package main

import (
	"demo/project_trans/model"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:password@(localhost)/tablename?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()

	db.AutoMigrate(&model.Exps{})
}

func UnJson(b []byte) (data *model.Exps) {
	var n model.Exps
	err1 := json.Unmarshal(b, &n)
	if err1 != nil {
		log.Fatalln(err1)
	}
	return &n
}

func main() {
	handler := func(w http.ResponseWriter, req *http.Request) {
		b, e := ioutil.ReadAll(req.Body)

		if e != nil {
			fmt.Printf("%v\n", e)
		} else {
			fmt.Printf("%s\n", string(b))
		}

		t := UnJson(b)


		DbTo(t)

		io.WriteString(w, "OK\n")
	}

	http.HandleFunc("/", handler)

	log.Println("Starting http server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func DbTo(m *model.Exps) {
	db.Debug().Model(m).Create(m)
}
