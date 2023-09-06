package repository

type PeerInfo struct {
	Id              string
	Ip              string
	TransactionHash string
}

type Repository interface {
	Store(p PeerInfo) (string, error)
	Get(id int64) (PeerInfo, error)
}
