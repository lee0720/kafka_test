package main

import (
	"fmt"
	"log"
	"main/kafka/product"
	"main/model"
	"strconv"

	"github.com/Shopify/sarama"
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

	producer, err := sarama.NewAsyncProducer([]string{"192.168.89.71:9092", "192.168.89.71:9093", "192.168.89.71:9094"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// entity_id := "1058552983"
	// entity_type := model.EntityType_ORGANIZATION

	// e := model.Entity{
	// 	EntityID:   &entity_id,
	// 	EntityType: &entity_type,
	// 	EntityName: "潍坊市长松柴油机有限责任公司",
	// }

	/*
	   "company_id" : "1029985986",
	             "negative_sentiment_level" : "1.0",
	             "company_name" : "新洋丰农业科技股份有限公司",
	             "link" : "http://www.hb.chinanews.com/news/2020/0326/336816.html",
	             "tag_id" : "2050200119",
	             "abstract" : "按照疫情防控要求及疫情形势，新洋丰农业科技股份有限公司1月27日全线停产，2月19日逐步复工复产，期间面临了值守人员交通不便、员工返岗受阻、原材料及产成品运输不畅等问题",
	             "id" : "41440",
	             "source" : "中国新闻网-财经新闻",
	             "title" : "湖北新洋丰农业科技加快生产助力春耕生产",
	             "publish_date" : "2020-03-26 10:58:00",
	             "news_id" : "1953735182"

	*/

	// eb := barn.UpdateEntityBucketRequest{
	// 	Op:            barn.Operator_INCREASE,
	// 	EntityBuckets: make([]*barn.EntityBucket, 0),
	// }
	// for i := 0; i < 10; i++ {
	// 	eb.EntityBuckets = append(eb.EntityBuckets, &barn.EntityBucket{
	// 		BucketUrl:  fmt.Sprintf(string(essentials.BucketURLRiskEventTag), fmt.Sprintf("%d", i)),
	// 		EntityId:   fmt.Sprintf("%d", i),
	// 		EntityType: common.EntityType_ORGANIZATION,
	// 		Count:      int32(i),
	// 		Date:       "20201030",
	// 	})
	// }
	// product.Productor("update_entity_bucket_count", "UpdateEntityBucketRequest", &eb, producer)

	// ub := barn.UpdateUserBucketRequest{
	// 	Op:          barn.Operator_INCREASE,
	// 	UserBuckets: make([]*barn.UserBucket, 0),
	// }
	// for i := 0; i < 10; i++ {
	// 	ub.UserBuckets = append(ub.UserBuckets, &barn.UserBucket{
	// 		BucketUrl: string(essentials.BucketURLDeal),
	// 		UserId:    fmt.Sprintf("%d", i),
	// 		Count:     int32(i),
	// 		Date:      "20201030",
	// 	})
	// }
	// product.Productor("update_user_bucket_count", "UpdateUserBucketRequest", &ub, producer)

	// gb := barn.UpdateGlobalBucketRequest{
	// 	Op:            barn.Operator_INCREASE,
	// 	GlobalBuckets: make([]*barn.GlobalBucket, 0),
	// }
	// for i := 0; i < 10; i++ {
	// 	gb.GlobalBuckets = append(gb.GlobalBuckets, &barn.GlobalBucket{
	// 		BucketUrl: string(essentials.BucketURLDeal),
	// 		Count:     int32(i),
	// 		Date:      "20201030",
	// 	})
	// }
	// product.Productor("update_global_bucket_count", "UpdateGlobalBucketRequest", &gb, producer)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			id := fmt.Sprintf("%d", i-100)
			riskPublicOpinion := model.RiskPublicOpinion{
				ID:          strconv.FormatInt(int64(i), 10),
				NewsID:      "1953735182",
				PublishDate: "2020-03-26 14:52:00",
				Title:       "湖北新洋丰农业科技加快生产助力春耕生产",
				Link:        "www.baidu.com",
				Source:      "新闻",
				Entities: []model.Entity{
					model.Entity{
						EntityID:   &id,
						EntityType: model.EntityType_ORGANIZATION,
						EntityName: "字节跳动",
					},
				},
				TagID:                  "2050200119",
				NegativeSentimentLevel: 1.0,
			}

			product.Productor("risk.sentiment_v1", riskPublicOpinion.ID, &riskPublicOpinion, producer)
		} else {
			entityID := fmt.Sprintf("%d", i)
			entityIDT := fmt.Sprintf("%d", (i-100)+200)

			riskEvent := model.Risk{
				ID:          strconv.FormatInt(int64(i), 10),
				TagID:       "10100004004",
				Content:     `{"10100004004":{"publishDate":"20200619","patentName":"一种组合式实木床"}}`,
				PublishDate: "2020-03-26 14:52:00",
				Entities: []model.Entity{
					model.Entity{
						EntityID:   &entityID,
						EntityType: model.EntityType_ORGANIZATION,
						EntityName: "华为",
					},
					model.Entity{
						EntityID:   &entityIDT,
						EntityType: model.EntityType_ORGANIZATION,
						EntityName: "腾讯",
					},
				},
			}

			product.Productor("risk.event_v1", riskEvent.ID, &riskEvent, producer)
		}
	}

	for i := 100; i < 200; i++ {
		if i%2 == 0 {
			id := fmt.Sprintf("%d", i)
			riskPublicOpinion := model.RiskPublicOpinion{
				ID:          strconv.FormatInt(int64(i), 10),
				NewsID:      "1953735182",
				PublishDate: "2020-03-26 14:52:00",
				Title:       "湖北新洋丰农业科技加快生产助力春耕生产",
				Link:        "www.baidu.com",
				Source:      "新闻",
				Entities: []model.Entity{
					model.Entity{
						EntityID:   &id,
						EntityType: model.EntityType_ORGANIZATION,
						EntityName: "字节跳动",
					},
				},
				TagID:                  "2050200119",
				NegativeSentimentLevel: 1.0,
			}

			product.Productor("risk.sentiment_v1", riskPublicOpinion.ID, &riskPublicOpinion, producer)
		} else {
			entityID := fmt.Sprintf("%d", i)
			entityIDT := fmt.Sprintf("%d", i+200)

			riskEvent := model.Risk{
				ID:          strconv.FormatInt(int64(i), 10),
				TagID:       "10100004004",
				Content:     `{"10100004004":{"publishDate":"20200619","patentName":"一种组合式实木床"}}`,
				PublishDate: "2020-03-26 14:52:00",
				Entities: []model.Entity{
					model.Entity{
						EntityID:   &entityID,
						EntityType: model.EntityType_ORGANIZATION,
						EntityName: "华为",
					},
					model.Entity{
						EntityID:   &entityIDT,
						EntityType: model.EntityType_ORGANIZATION,
						EntityName: "腾讯",
					},
				},
			}

			product.Productor("risk.event_v1", riskEvent.ID, &riskEvent, producer)
		}
	}

	/*
	   "tag_name" : "专利到期风险",
	            "tag_id" : "10100004004",
	            "company" : [
	              {
	                "name" : "青岛百瑞家居有限公司",
	                "id" : "1090201214"
	              }
	            ],
	            "id" : "d89e1e39-d38c-4df0-b269-83cdb8047d09",
	            "publish_date" : "20290915",
	            "content" : """{"10100004004":{"publishDate":"20200619","patentName":"一种组合式实木床"}}"""

	*/

	// for i := 100; i < 200; i++ {
	// 	entityID := "1024736719"
	// 	entityIDT := "1024736720"

	// 	riskEvent := model.Risk{
	// 		ID:          strconv.FormatInt(int64(i), 10),
	// 		TagID:       "10100004004",
	// 		Content:     `{"10100004004":{"publishDate":"20200619","patentName":"一种组合式实木床"}}`,
	// 		PublishDate: "2020-03-26 14:52:00",
	// 		Entities: []model.Entity{
	// 			model.Entity{
	// 				EntityID:   &entityID,
	// 				EntityType: model.EntityType_ORGANIZATION,
	// 				EntityName: "华为",
	// 			},
	// 			model.Entity{
	// 				EntityID:   &entityIDT,
	// 				EntityType: model.EntityType_ORGANIZATION,
	// 				EntityName: "腾讯",
	// 			},
	// 		},
	// 	}

	// 	product.Productor("risk.event_v1", riskEvent.ID, &riskEvent, producer)
	// }
}
