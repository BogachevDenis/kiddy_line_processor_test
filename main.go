package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kiddy_line_processor_test.git/database"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)
type setting struct {
	ServerHost		string
	HTTPServerPort	string
	GRPCServerPort	string
	N_football  	int
	N_baseball		int
	N_soccer		int
	PgHost     		string
	PgPort     		string
	PgUser     		string
	PgPass     		string
	PgBase     		string
}
type Inf struct {
	mu      sync.Mutex
	data map[string]map[string]string
	storage map[string]int
}
func NewCounters() *Inf  {
	return &Inf{
		data: make(map[string]map[string]string),
		storage: make(map[string]int),
	}
}
var cfg setting

var i *Inf



func init()  {
	i = NewCounters()

	file, e := os.Open("setting.cfg")
	if e != nil{
		log.WithFields(log.Fields{
			"file" : "setting.cfg",
			"error" : e,
		}).Fatal("File can`t be opened")
	}
	log.WithFields(log.Fields{
		"file" : "setting.cfg",
	}).Info("Config file was opened")

	defer file.Close()
	stat, e := file.Stat()
	if e != nil{
		log.WithFields(log.Fields{
			"file" : "setting.cfg",
			"error" : e,
		}).Fatal("File can`t be stat")
	}

	readByte := make([]byte, stat.Size())
	_, error := file.Read(readByte)
	if error != nil{
		log.WithFields(log.Fields{
			"file" : "setting.cfg",
			"error" : error,
		}).Fatal("File can`t be read")
	}


	err := json.Unmarshal(readByte, &cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"file" : "cfg",
			"error" : err,
		}).Fatal("Config data can`t be unmarshal")
	}

	error_db := database.Connect(cfg.PgUser, cfg.PgPass, cfg.PgBase)
	if error_db  != nil {
		log.WithFields(log.Fields{
			"database" : "error",
			"error" : error_db,
		}).Fatal("Connection error")
	} else {
		i.storage["DB"] = 200
		log.WithFields(log.Fields{
			"database" : "connect",
		}).Info("Connection ok")
	}

	go worker_footbal(i,cfg.N_football)
	go worker_baseball(i,cfg.N_baseball)
	go worker_soccer(i,cfg.N_soccer)
}

func main() {
	r := gin.Default()
	r.GET("/ready", ready )
	r.Run(":"+cfg.HTTPServerPort) // listen and serve on 0.0.0.0:8080

}


func worker_footbal(i *Inf, n_footbal int) {

	for {
		time.Sleep(time.Duration(n_footbal) * time.Second)
		resp, err := http.Get("http://kiddy_line_processor_testgit_lines-provider_1:8000/api/v1/lines/football")

		if err != nil {
			log.WithFields(log.Fields{
				"lines-provider" : "error",
				"error" : err,
			}).Fatal("lines-provider GET Error")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.WithFields(log.Fields{
				"lines-provider" : "error",
				"error" : err,
			}).Fatal("lines-provider GET Error")
		}

		i.mu.Lock()
		err = json.Unmarshal(body, &i.data)
		i.mu.Unlock()

		if err != nil {
			log.WithFields(log.Fields{
				"lines-provider body" : "error",
				"error" : err,
			}).Fatal("lines-provider data can`t be unmarshal")
		}

		i.mu.Lock()
		i.storage["FOOTBALL"] = resp.StatusCode
		i.mu.Unlock()

		i.mu.Lock()
		transform, _ := strconv.ParseFloat(i.data["lines"]["FOOTBALL"], 32)
		i.mu.Unlock()

		transformFloat := float32(transform)

		i.mu.Lock()
		error := database.Insert_football(transformFloat)
		i.mu.Unlock()
		if error != nil {
			log.WithFields(log.Fields{
				"Insert_football" : "error",
				"error" : error,
			}).Fatal("Insert Database error")
		}

	}
}
func worker_baseball(i *Inf, n_baseball int) {
	for {
		time.Sleep(time.Duration(n_baseball) * time.Second)

		resp, err := http.Get("http://kiddy_line_processor_testgit_lines-provider_1:8000/api/v1/lines/baseball")
		if err != nil {
			log.WithFields(log.Fields{
				"lines-provider" : "error",
				"error" : err,
			}).Fatal("lines-provider GET Error")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.WithFields(log.Fields{
				"lines-provider" : "error",
				"error" : err,
			}).Fatal("lines-provider GET Error")
		}

		i.mu.Lock()
		err = json.Unmarshal(body, &i.data)
		i.mu.Unlock()

		if err != nil {
			log.WithFields(log.Fields{
				"lines-provider body" : "error",
				"error" : err,
			}).Fatal("lines-provider data can`t be unmarshal")
		}

		i.mu.Lock()
		i.storage["BASEBALL"] = resp.StatusCode
		i.mu.Unlock()

		i.mu.Lock()
		transform, _ := strconv.ParseFloat(i.data["lines"]["BASEBALL"],32)
		i.mu.Unlock()

		transformFloat := float32(transform)

		i.mu.Lock()
		error := database.Insert_baseball(transformFloat)
		i.mu.Unlock()

		if error != nil {
			log.WithFields(log.Fields{
				"Insert_football" : "error",
				"error" : error,
			}).Fatal("Insert Database error")
		}
	}
}
func worker_soccer(i *Inf, n_soccer int) {
	for {
		time.Sleep(time.Duration(n_soccer) * time.Second)

		resp, err := http.Get("http://kiddy_line_processor_testgit_lines-provider_1:8000/api/v1/lines/soccer")

		if err != nil {
			log.WithFields(log.Fields{
				"lines-provider" : "error",
				"error" : err,
			}).Fatal("lines-provider GET Error")
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.WithFields(log.Fields{
				"lines-provider" : "error",
				"error" : err,
			}).Fatal("lines-provider GET Error")
		}

		i.mu.Lock()
		err = json.Unmarshal(body, &i.data)
		i.mu.Unlock()

		if err != nil {
			log.WithFields(log.Fields{
				"lines-provider body" : "error",
				"error" : err,
			}).Fatal("lines-provider data can`t be unmarshal")
		}

		i.mu.Lock()
		i.storage["SOCCER"] = resp.StatusCode
		i.mu.Unlock()

		i.mu.Lock()
		transform, _ := strconv.ParseFloat(i.data["lines"]["SOCCER"],64)
		i.mu.Unlock()

		transformFloat := float32(transform)

		i.mu.Lock()
		error := database.Insert_soccer(transformFloat)
		i.mu.Unlock()

		if error != nil {
			log.WithFields(log.Fields{
				"Insert_football" : "error",
				"error" : error,
			}).Fatal("Insert Database error")
		}
	}
}

func ready(c *gin.Context) {

	i.mu.Lock()
	if (i.storage["FOOTBALL"] == 200)&&(i.storage["BASEBALL"] == 200)&&(i.storage["SOCCER"] == 200)&&(i.storage["DB"] == 200) {
		c.JSON(200, gin.H{
			"db_connection": "ok",
			"data_status": "ok",
		})
		i.mu.Unlock()
	} else {
		c.JSON(500, gin.H{
			"db_connection": "error",
			"data_status": "error",
		})
	}
}
