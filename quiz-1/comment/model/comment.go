package model

// Comment - 評論
type Comment struct {
	// 唯一UUID
	UUID string `json:"-" bson:"uuid"`
	// 父ID
	ParentID string `json:"parentid" bson:"parentid"`
	// 評論內容
	Comment string `json:"comment" bson:"comment"`
	// 評論作者
	Author string `json:"author" bson:"author"`
	// 更新時間
	Update string `json:"update" bson:"update"`
	// 是否為收藏評論
	Favorite bool `json:"favorite" bson:"favorite"`
}
