## risk & news kafka 消息结构
- entity_type 枚举，整数类型
	- EntityType_ENTITY_TYPE_UNSPECIFIED： 0
	- EntityType_PERSON：                  109006023
	- EntityType_ORGANIZATION：            109006022
	- EntityType_DEAL：                    109006024
	- EntityType_VERTICAL：                109006025

### risk kafka 消息结构
> 事件和舆情 只是content 内容不一样，结构是一样的

- 事件
``` json
{
    "id":"f617360b-3172-4f59-8af8-0973f55a971d",     // 事件 ID
    "entity":[     // 事件关联的 实体
        {
            "entity_id":"1058552983",   // 实体 id
            "entity_type":109006022,    // 实体类型，枚举类型
            "entity_name":"潍坊市长松柴油机有限责任公司"   // 实体名称
        }
    ],
    "publish_date":"20220618",    // 公布时间
    "tag_id":"10100004004",    // 事件类型
    //如果是事件，content 存储原来的事件content
    //如果是舆情，content 存储下面json结构
    "content":"{"10100004004":{"publishDate":"20130102","patentName":"柴油机水箱散热装置"}}"
}
```

- 舆情
``` json
{
    "id":"2642161",    // 舆情 id
    "entity":[   // 舆情 关联的 实体
        {
            "entity_id":"1201585340",   // 实体id
            "entity_type":109006022,    // 实体类型，枚举类型
            "entity_name":"潍坊市长松柴油机有限责任公司"   // 实体名称
        }
    ],
    "publish_date":"20220618",   // 公布时间
    "tag_id":"10100004004",    // 舆情类型
    //如果是事件，content 存储原来事件content
    //如果是舆情，content 存储下面json结构
    "content":"{"news_id":"2179421888","title":"24000亿特斯拉，怎么和“韭菜”杠上了？","link":"https://www.toutiao.com/i6876007738221003267/","negative_sentiment_level":"4.0","source":"头条-德林社","abstract":"3月初，特斯拉Model 3遭消费者投诉，称《环保随车清单》显示的自动驾驶硬件HW3.0版本，被“减配”为性能相差21倍的HW2.5版本，车主维权事件爆发"}"
}
```

### news kafka 消息结构

``` json
{
    "id":"2165203995",             // 新闻 id
    "title":"生物制药研发商凯瑞康宁完成4000万美元C轮融资",   // 新闻标题
    "abstract":"abstract",          // 新闻摘要
    "publish_time":"2019-10-10 11:22:33",    // 新闻的发布时间
    "link":"http://www.ebrun.com/20200918/402712.shtml",    // 新闻链接
    "picture_url":"PictureURL",    // 新闻图片链接
    "related_entities":[    // 新闻相关的 实体
        {
            "entity_id":"1058552983",    // 实体 id
            "entity_type":109006022,    // 实体类型，枚举类型
            "entity_name":"潍坊市长松柴油机有限责任公司"    // 实体名称
        }
    ],
    "content":"content"    // 新闻的内容
}
```


## 附录
- go 模型
``` go
package main

type Risk struct {
	ID          string    `json:"id,omitempty"`
	Entity      []*Entity `json:"entity,omitempty"`
	PublishDate string    `json:"publish_date,omitempty"`
	TagID       string    `json:"tag_id,omitempty"`
	Content     string    `json:"content,omitempty"`
}
type Entity struct {
	EntityID   *string     `json:"entity_id,omitempty"`
	EntityType *EntityType `json:"entity_type,omitempty"`
	EntityName string      `json:"entity_name,omitempty"`
}

//如果是事件，content 存储原来南京给的事件content
//如果是舆情，content 存储下面json结构
type PublicOpinion struct {
	NewsID                 string `json:"news_id,omitempty"`
	Title                  string `json:"title,omitempty"`
	Link                   string `json:"link,omitempty"`
	NegativeSentimentLevel string `json:"negative_sentiment_level,omitempty"`
	Source                 string `json:"source,omitempty"`
	Abstract               string `json:"abstract,omitempty"`
}

type News struct {
	ID              string    `json:"id,omitempty"`
	Title           string    `json:"title,omitempty"`
	Abstract        string    `json:"abstract,omitempty"`
	PublishTime     string    `json:"publish_time,omitempty"`
	Link            string    `json:"link,omitempty"`
	PictureURL      string    `json:"picture_url,omitempty"`
	RelatedEntities []*Entity `json:"related_entities,omitempty"`
	Source          string    `json:"source,omitempty"`
	Content         string    `json:"content,omitempty"`
}

// EntityType that Cornerstone supports. All implementations
// can be found in themes/kernel.
type EntityType int32

const (
	EntityType_ENTITY_TYPE_UNSPECIFIED EntityType = 0
	EntityType_PERSON                  EntityType = 109006023
	EntityType_ORGANIZATION            EntityType = 109006022
	EntityType_DEAL                    EntityType = 109006024
	EntityType_VERTICAL                EntityType = 109006025
)

var EntityType_name = map[int32]string{
	0:         "ENTITY_TYPE_UNSPECIFIED",
	109006023: "PERSON",
	109006022: "ORGANIZATION",
	109006024: "DEAL",
	109006025: "VERTICAL",
}

var EntityType_value = map[string]int32{
	"ENTITY_TYPE_UNSPECIFIED": 0,
	"PERSON":                  109006023,
	"ORGANIZATION":            109006022,
	"DEAL":                    109006024,
	"VERTICAL":                109006025,
}

```