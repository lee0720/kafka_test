package main

import (
	"main/kafka/consumer"
)

/*
 "_source" : {
          "link" : "http://www.ebrun.com/20200918/402712.shtml",
          "id" : "2165203995",
          "title" : "生物制药研发商凯瑞康宁完成4000万美元C轮融资",
          "publish_date" : "2020-09-18 15:22:56"
        }
*/

func main() {

	go consumer.Consumer("risk.event_v1")
	consumer.Consumer("risk.sentiment_v1")
	// go consumer.Consumer("update_entity_bucket_count")
	// go consumer.Consumer("update_global_bucket_count")
	// consumer.Consumer("user_bucket")

}
