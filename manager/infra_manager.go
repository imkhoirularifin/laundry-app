package manager

import (
	"database/sql"
	"fmt"
	"laundry-app/config"
)

type InfraManager interface {
	openConn() error
	Conn() *sql.DB
}

type infraManager struct {
	db  *sql.DB // Connection
	config *config.Config // Configuration
}

/*
	Open Database Connection based on existing Configuration
*/
func (i *infraManager) openConn() error {
	// dsn = data source name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", i.config.Host, i.config.Port, i.config.Username, i.config.Password, i.config.Database)

	db, err := sql.Open(i.config.Driver, dsn)
	if err != nil {
		return fmt.Errorf("failed to open connection %v", err.Error())
	}

	// Fill db on infraManager struct
	i.db = db
	return nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	// Fill config on infraManager struct
	conn := &infraManager{config: cfg}
	if err := conn.openConn(); err != nil {
		return nil, err
	}

	return conn, nil
}
