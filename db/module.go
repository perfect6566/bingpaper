package db

import "time"

type Papers struct {
	Image
	Createtime time.Time `bson:"createtime" json:"createtime"`
}

type Image struct {
	Startdate     string        `bson:"startdate" json:"startdate"`
	Fullstartdate string        `bson:"fullstartdate" json:"fullstartdate"`
	Enddate       string        `bson:"enddate" json:"enddate"`
	URL           string        `bson:"url" json:"url"`
	Urlbase       string        `bson:"urlbase" json:"urlbase"`
	Copyright     string        `bson:"copyright" json:"copyright"`
	Copyrightlink string        `bson:"copyrightlink" json:"copyrightlink"`
	Title         string        `bson:"title" json:"title"`
	Quiz          string        `bson:"quiz" json:"quiz"`
	Wp            bool          `bson:"wp" json:"wp"`
	Hsh           string        `bson:"hsh" json:"hsh"`
	Drk           int           `bson:"drk" json:"drk"`
	Top           int           `bson:"top" json:"top"`
	Bot           int           `bson:"bot" json:"bot"`
	Hs            []interface{} `bson:"hs" json:"hs"`
}
type Bingresp struct {
	Images []Image `json:"images"`
	Tooltips struct {
		Loading  string `json:"loading"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Walle    string `json:"walle"`
		Walls    string `json:"walls"`
	} `json:"tooltips"`
}