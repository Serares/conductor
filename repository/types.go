package repository

type PeerInfo struct {
	id       int64
	ip       int64
	hash     string
	fileHash string
}

type Repository interface {
	Store(p PeerInfo)
	Get(id int64)
}
