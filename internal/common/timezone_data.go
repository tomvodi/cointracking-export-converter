package common

type TimezoneData struct {
	Value         string `json:"value"`
	Title         string `json:"title"`
	OffsetSeconds int    `json:"-"`
}
