package model

type GetRecommendationListResp struct {
	CommonRes
	Data GetRecommendationListRespData `json:"data"`
}

type GetRecommendationListRespData struct {
	Total              int64                `json:"total"`
	RecommendationList []RecommendationList `json:"recommendation_list"`
}

type RecommendationList struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
