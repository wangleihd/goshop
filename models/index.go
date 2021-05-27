package models

var (
	Indexs map[string]*Index
)

type Index struct {
	IndexId    string
	Score      int64
	PlayerName string
}

type IndexReturnJSON struct {
	Banners []NideshopAd `json:"banners"`
}

func init() {
	Indexs = make(map[string]*Index)
	Indexs["hjkhsbnmn123"] = &Index{"hjkhsbnmn123", 100, "astaxie"}
	Indexs["mjjkxsxsaa23"] = &Index{"mjjkxsxsaa23", 101, "someone"}
}

func GetIndex() map[string]*Index {
	return Indexs
}
