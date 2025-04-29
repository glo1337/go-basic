package bins

import (
	"time"

	"github.com/google/uuid"
)

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
		Id:        uuid.New().String(),
		Private:   true,
		CreatedAt: time.Now(),
		Name:      "firstBin",
	}
}

func NewBinList() *BinList {
	return &BinList{}
}
