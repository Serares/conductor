package repository

type PeerInfo struct {
	Id              int64
	Ip              string
	TransactionHash string
	FileHash        string
	PubKey          string // used to encrypt the files before sending them
}

type Repository interface {
	Store(p PeerInfo) (int64, error)
	Get(id int64) (PeerInfo, error)
}
