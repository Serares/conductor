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
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", options.Hostname, options.Username, options.Database, options.Password))
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	return &dbRepo{
		db: db,
	}, nil

}

func (r *dbRepo) Get(id int64) (PeerInfo, error) {
	row := r.db.QueryRow("SELECT * FROM peers where id=$1", id)
	newPeerInfo := PeerInfo{}
	err := row.Scan(&newPeerInfo.Id, &newPeerInfo.Ip, &newPeerInfo.TransactionHash)

	return newPeerInfo, err
}

func (r *dbRepo) Store(p PeerInfo) (string, error) {
	var lastInsertedId string
	err := r.db.QueryRow("INSERT INTO peers VALUES($1, $2, $3) RETURNING id", p.Id, p.Ip, p.TransactionHash).Scan(&lastInsertedId)
	if err != nil {
		return "", err
	}

	return lastInsertedId, nil
}
