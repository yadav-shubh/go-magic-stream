package models

type Ranking struct {
	RankingValue int64  `bson:"ranking_value" json:"ranking_value"`
	RankingName  string `bson:"ranking_name" json:"ranking_name"`
}

type RankingDTO struct {
	RankingValue int64  `bson:"ranking_value" json:"ranking_value" validate:"required,min=1"`
	RankingName  string `bson:"ranking_name" json:"ranking_name" validate:"required,min=2"`
}
