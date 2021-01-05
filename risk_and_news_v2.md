## ES模型

- risk
```json
{
	"mappings": {
		"properties": {
			"id": {
				"type": "keyword"
			},
			"entities": {
                "type":"nested",
				"properties":{  
					"entity_id":{
						"type":"keyword"
					},
                    "entity_type":{
						"type":"keyword"
					},
					"entity_name":{ 
						"index":"false",
						"type":"keyword"
					}
				}
			},
			"publish_date": {
				"type": "date",
				"format":"yyyy-MM-dd"
			},
			"tag_id": {
				"type": "keyword"
			},
			"content": {
                "index": "false",
                "type": "text"
            }			
		}
	}
}


//如果是事件，content 存储原来南京给的事件content
//如果是舆情，content 存储下面json结构
{
    "news_id": "",
    "title": "",
    "link": "",
    "negative_sentiment_level": "",
    "source": "",
    "abstract": ""
}


// 原来 news_id 建立了索引，使用json之后没有了，是否需要？ 
{
    "news_id": {
        "type": "keyword"
    },
    "title": {
        "index": false,
        "type": "keyword"
    },
    "link": {
        "index":false,
        "type": "keyword"
    },
    "negative_sentiment_level": {
        "index":false,
        "type": "keyword"
    },
    "source": {
        "index":false,
        "type": "keyword"
    },
    "abstract": {
        "index": false,
        "type": "text"
    }
}

```
- news

``` go

type News struct {
	ID string     `json:"id,omitempty"`
	Title string  `json:"title,omitempty"`
	Abstract string `json:"abstract,omitempty"`
	PublishTime time.Time `json:"publish_time,omitempty"`
	Link string `json:"link,omitempty"`
	PictureURL string `json:"picture_url,omitempty"`
	Entity `json:"entity,omitempty"`
	Source string `json:"source,omitempty"`
	Content string `json:"content,omitempty"`
}


```



```json
{
	"mappings": {
		"properties": {
			"id": {
				"type": "keyword"
			},
			"title": {
				"index": "true",
				"type": "text"
			},
			"abstract": {
				"index": "true",
				"type": "text"
			},
			"publish_time": {
				"index": "true",
				"type": "date",
				"format": "yyyy-MM-dd HH:mm:ss"
			},
			"link": {
				"index": "false",
				"type": "keyword"
			},
			"picture_url": {
				"index": "false",
				"type": "keyword"
			},
			"related_entities": {
				"type":"nested",
				"properties":{  
					"entity_id":{
						"type":"keyword"
					},
					"entity_type":{
						"type":"keyword"
					},
					"entity_name":{ 
						"index":"false",
						"type":"keyword"
					}
				}
			},
			"source": {
				"index": "false",
				"type": "keyword"
			},
			"content": {
				"index": "true",
				"type": "text"
			}
		}
	}
}

```

// TODO: 改成存到 mysql，性能问题通过添加 cache 解决
## 赛道 -> 关键词 存储方案
- 存放 redis
- 数据类型： Set集合
- key:vertical_id value:关键词组
- 启用rdb持久化


## 接口设计

- risk

```protoc
service RiskManagementTwirp {  
    rpc GetRisks (RiskRequest) returns (RiskConnection);
}


message RiskRequest {
    google.protobuf.Int32Value first =1;
    google.protobuf.StringValue after =2;
    google.protobuf.Int32Value last = 3;
    google.protobuf.StringValue before =4;
    RiskFilter filter = 5;
}
message RiskFilter {
    string entity_id = 1;
	string entity_type = 2;
	google.protobuf.StringValue start_time = 3;
    google.protobuf.StringValue end_time = 4;
	repeated string tag_ids = 5;
}

message RiskConnection{
    int32 total_count = 1;
    repeated RiskEdge edges = 2;
    repeated Risk nodes = 3;
    io.mvalley.common.PageInfo page_info = 4;
}

message RiskEdge {
    Risk node = 1;
    string cursor = 2;
}

// 配合前端组件，待定
message Risk {
    string publish_date = 1;
    repeated Entity entities = 2;
    string tag_name = 3; 
    string abstract = 4;
	
	string tag_code = 6;
    string content = 7;
}

message Entity {
    google.protobuf.StringValue entity_id = 1;
    google.protobuf.StringValue entity_type = 2;
    string entity_name = 3;
}

```

- news

```protoc
service NewsTwirp {  
    rpc GetNews (NewsRequest) returns (NewsConnection);
}

message NewsRequest {
    google.protobuf.Int32Value first =1;
    google.protobuf.StringValue after =2;
    google.protobuf.Int32Value last = 3;
    google.protobuf.StringValue before =4;
    repeated EntityFilter entity = 5;
	repeated keywords = 6;
}


message EntityFilter {
	string entity_id = 1;
	string entity_type = 2;
}


message NewsConnection{
    int32 total_count = 1;
    repeated NewsEdge edges = 2;
    repeated News nodes = 3;
    io.mvalley.common.PageInfo page_info = 4;
}

message NewsEdge {
    News node = 1;
    string cursor = 2;
}

// 配合前端组件，待定
message News {
	string title = 1; 
	string abstract = 2;
    string publish_date = 3;
	string like = 4;
	string picture_url = 5;
	repeated Entity entities = 6;
}

message Entity {
    google.protobuf.StringValue entity_id = 1;
    google.protobuf.StringValue entity_type = 2;
    string entity_name = 3;
}

```