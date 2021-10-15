package main

import (
	"context"
	"encoding/json"
	"github.com/qianlnk/pgbar"
	"log"
	"os"
	"strconv"
	"time"

	v2 "gopkg.in/olivere/elastic.v2"

	v7 "github.com/olivere/elastic/v7"
)

var loger *log.Logger

var size = 5000


var v7client *v7.Client

var oldIndexTypeMap  = map[string]string{
	"category": "pro_category",
	"brand":"pro_brand",
}

func init() {
	file := "./" + time.Now().Format("2006") + "_log" + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	loger = log.New(logFile, "[elastic_sync]", log.LstdFlags|log.Lshortfile|log.LUTC) // 将文件设置为loger作为输出


	v7client, err = v7.NewClient(
		v7.SetURL("http://10.10.10.53:9200"),
		v7.SetBasicAuth("elastic", "iceasy2021"),
	)

	if err != nil {
		panic(err)
	}
}

func main() {
	sync()
}

func sync() {

	client, err := v2.NewClient(
		v2.SetURL("http://10.10.10.161:9200"),
		v2.SetBasicAuth("iceasyadm", "m7M0Z%hbRsDQirISAu"),
	)

	v2.SetErrorLog(loger)
	v2.SetInfoLog(loger)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(getValues(oldIndexTypeMap))
	c, _ := client.Count(getValues(oldIndexTypeMap)...).Types(getKeys(oldIndexTypeMap)...).Do()
	log.Println("本次统计到需要更新总数: ", c)
	pgb := pgbar.New("")
	pgbar.Println("ES 同步")
	b := pgb.NewBar("ES 1st", int(c))

	query := v2.MatchAllQuery{}
	s, err := client.Scroll(getValues(oldIndexTypeMap)...).Types(getKeys(oldIndexTypeMap)...).Query(query).Scroll("3m").Size(size).Do()
	if err != nil {
		loger.Println(err)
	}
	for {
		r, err := client.Scroll().ScrollId(s.ScrollId).Scroll("3m").Do()
		if err != nil {
			loger.Println(err)
		}
		if r != nil && r.Hits != nil {
			handlerTargetData(r.Hits, b)
		}
	}
}
// 处理数据
func handlerTargetData(hits *v2.SearchHits, b *pgbar.Bar) {
	bulkRequest := v7client.Bulk()
	for i := 0; i < len(hits.Hits); i++ {
		doc, _ := hits.Hits[i].Source.MarshalJSON()
		bulkRequest.Add(v7.NewBulkIndexRequest().Index(hits.Hits[i].Type).Id(getId(doc)).Doc(string(doc)))
	}
	_, err := bulkRequest.Do(context.Background())
	if err != nil {
		loger.Println("bulk 错误 %s", err)
	}
	b.Add(len(hits.Hits))
}

func getId(doc []byte) string {
	bean := make(map[string]interface{})
	err := json.Unmarshal(doc, &bean)
	if err != nil {
		loger.Println("json转译失败 %s", err)
	}
	return strconv.Itoa(int(bean["id"].(float64)))
}



func getKeys(m map[string]string) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func getValues(m map[string]string) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	values := make([]string, len(m))
	for _, v := range m {
		values[j] = v
		j++
	}
	return values
}