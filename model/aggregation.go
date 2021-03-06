package model

type WorkTime struct {
	Tag     string `json:"tag"`
	Time    string `json:"time"`
	Percent string `json:"percent"`
}

type WorkTimes []*WorkTime

type Response struct {
	WorkTimes WorkTimes
	ClockDatas []ClockData
}
