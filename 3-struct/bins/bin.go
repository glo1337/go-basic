package bins

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	List []Bin `json:"list"`
}

func NewBin() *Bin {
	return &Bin{
		Id:        "first",
		Private:   true,
		CreatedAt: time.Now(),
		Name:      "firstBin",
	}
}

func NewBinList() *BinList {
	return &BinList{}
}
