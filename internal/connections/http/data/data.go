package data

type AddToStore struct {
	Key any `json:"key"`
	Val any `json:"val"`
}

type GetFromStore struct {
	Val any `json:"val"`
}
