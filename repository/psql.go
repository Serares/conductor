package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type dbRepo struct {
	db *sql.DB
}

type ConnectionOptions struct {
	Username string
	Password string
	Hostname string
	Database string
}

func NewPsqlRepo(options ConnectionOptions) (*dbRepo, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s db=%s sslmode=disabled password=%s", options.Username, options.Database, options.Password))
	if err != nil {
		return nil, fmt.Errorf("error connectiong to db: %v", err)
	}
	return &dbRepo{
		db: db,
	}, nil

}

func (r *dbRepo) Get(id int64) (PeerInfo, error) {
	row := r.db.QueryRow("SELECT * FROM conductor where id=?", id)
	newPeerInfo := PeerInfo{}
	err := row.Scan(&newPeerInfo.Id, &newPeerInfo.Ip, &newPeerInfo.TransactionHash, &newPeerInfo.FileHash)

	return newPeerInfo, err
}

func (r *dbRepo) Store(p PeerInfo) (int64, error) {
	ins, err := r.db.Prepare("INSERT INTO conductor VALUES(NULL, ?,?,?)")
	if err != nil {
		return 0, err
	}

	defer ins.Close()

	res, err := ins.Exec(p.Ip, p.TransactionHash, p.FileHash)
	if err != nil {
		return 0, err
	}
	var lastId int64
	if lastId, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return lastId, nil
}
