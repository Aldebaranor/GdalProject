package entity

func (Earth_tif) TableName() string {
	return "earthTif"
}

type Earth_tif struct {
	Lon float64 `json:"lon" gorm:"column:lon"`
	Lat float64 `json:"lat" gorm:"column:lat"`
	Alt float64 `json:"alt" gorm:"column:alt"`
}
