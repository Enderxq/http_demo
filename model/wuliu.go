package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Exps struct {
	Id          uint   `gorm:"primary_key" json:"id"`
	Waybill     string `gorm:"type:varchar(100);not null" json:"waybill"`
	Company     string `gorm:"type:varchar(100);not null" json:"company"`
	Created_at  string `gorm:"type:varchar(100);not null" json:"created___at"`
	Origin      string `gorm:"type:varchar(100);not null" json:"origin"`
	Updated_at  string `gorm:"type:varchar(100);not null" json:"updated___at"`
	Destination string `gorm:"type:varchar(100);not null" json:"destination"`
	State       uint   `json:"state"`
	Routes      string `gorm:"type:text(1630);not null" json:"routes"`
}

func Report() {
	data := GetWuliu()
	for _, it:=range data {
		strJson, _ := json.Marshal(it)
		//fmt.Println(string(strJson), "post to web api ")
		HttpPostToServer("http://127.0.0.1:8080", strJson)
	}

}

func HttpPostToServer(addr string, jsonstr []byte) {
	req, err := http.NewRequest("POST", addr, bytes.NewBuffer(jsonstr))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	//fmt.Println("response status:", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Printf("response body:%s\n", body)
}


func GetWuliu() (data []*Exps){
	d1 := &Exps{
		Id: 1,
		Waybill: "111222",
		Company: "ems",
		Created_at: "2021.01.02",
		Origin: "",
		Updated_at: "2021.08.01",
		Destination: "",
		State: 0,
		Routes: "",
	}
	d2 := &Exps{
		Id: 2,
		Waybill: "222333",
		Company: "ems",
		Created_at: "2021.01.03",
		Origin: " ",
		Updated_at: "2021.08.01",
		Destination: " ",
		State: 1,
		Routes: " ",
	}
	data = append(data, d1, d2)
	//data = append(data, d)

	return
	/* //
	fmt.Println("get from sqlite ")
	db, err := sql.Open("sqlite3", "1111expresses.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, waybill, company, created_at, origin, updated_at, destination, state, routes from expresses order by id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	for rows.Next() {
		var id , state uint
		var waybill, company, created_at, origin, updated_at, destination, routes string
		err = rows.Scan( &id, &waybill, &company, &created_at, &origin, &updated_at, &destination, &state, &routes)
		if err != nil {
			log.Fatal(err)
		}

		d := &Exps{
			Id: id,
			Waybill: waybill,
			Company: company,
			Created_at: created_at,
			Origin: origin,
			Updated_at: updated_at,
			Destination: destination,
			State: state,
			Routes: routes,
		}
		data = append(data, d)
	}
	return

	*/
}

