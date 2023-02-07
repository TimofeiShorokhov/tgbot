package data

type SignState int

const (
	 StateName SignState = iota
	 StateNumber
)

const (
	Classic = "20"
	D2 = "25"
	D3 = "30"
	D4 = "35"
)

type Client struct {
	State SignState
	ID int
	Name string
	Number string
}

