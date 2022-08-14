package models

type ScoreForm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	UID  string `json:"uid"`
	Ans  []struct {
		ID  int   `json:"id"`
		Key []int `json:"key"`
	} `json:"ans"`
}
