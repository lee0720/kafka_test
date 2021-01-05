package model

type Entity struct {
	EntityID   *string    `json:"entity_id,omitempty"`
	EntityType EntityType `json:"entity_type,omitempty"`
	EntityName string     `json:"entity_name,omitempty"`
}

//如果是事件，content 存储原来南京给的事件content
//如果是舆情，content 存储下面json结构

type News struct {
	ID              string   `json:"id,omitempty"`
	Title           string   `json:"title,omitempty"`
	Abstract        string   `json:"abstract,omitempty"`
	PublishTime     string   `json:"publish_time,omitempty"`
	Link            string   `json:"link,omitempty"`
	PictureURL      string   `json:"picture_url,omitempty"`
	RelatedEntities []Entity `json:"related_entities,omitempty"`
	Source          string   `json:"source,omitempty"`
	Content         string   `json:"content,omitempty"`
}

// Risk ...
type Risk struct {
	ID          string   `json:"id,omitempty"`
	Entities    []Entity `json:"entities,omitempty"`
	PublishDate string   `json:"publish_date,omitempty"`
	TagID       string   `json:"tag_id,omitempty"`
	Content     string   `json:"content,omitempty"`
}

// RiskPublicOpinion ...
type RiskPublicOpinion struct {
	Entities               []Entity `json:"entities,omitempty"`
	ID                     string   `json:"id,omitempty"`
	NewsID                 string   `json:"news_id,omitempty"`
	NegativeSentimentLevel int32    `json:"negative_sentiment_level,omitempty"`
	Link                   string   `json:"link,omitempty"`
	TagID                  string   `json:"tag_id,omitempty"`
	Source                 string   `json:"source,omitempty"`
	Title                  string   `json:"title,omitempty"`
	PublishDate            string   `json:"publish_date,omitempty"`
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
