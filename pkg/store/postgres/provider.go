package postgres

import (
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"log"
	"service_history/internal/app/contract"
	"service_history/pkg/config"
	"time"
)

type provider struct {
	db        *sql.DB
	cs        string
	idlConns  int
	openConns int
	lifetime  time.Duration
	cfg       *config.SQL
}

// NewProvider return contract
func NewProvider(cfg *config.SQL) contract.IProvider {
	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable",
		cfg.Host, cfg.Port, cfg.UserID, cfg.Password, cfg.Database)

	return &provider{
		cs:        info,
		idlConns:  cfg.MaxIdleCons,
		openConns: cfg.MaxOpenCons,
		lifetime:  time.Duration(cfg.ConnMaxLifetime),
		cfg:       cfg,
	}
}

// Open Pull connection
func (p *provider) Open() error {
	var err error
	p.db, err = sql.Open("postgres", p.cs)
	if err != nil {
		return err
	}
	p.db.SetMaxIdleConns(p.idlConns)
	p.db.SetMaxOpenConns(p.openConns)
	p.db.SetConnMaxLifetime(p.lifetime * time.Minute)

	err = p.db.Ping()
	if err != nil {
		return err
	}
	log.Println("PG_DB connection open")

	return nil
}

// Close connection
func (p *provider) Close() error {
	return p.db.Close()
}

// GetConn return sql database
func (p *provider) GetConn() *sql.DB {
	return p.db
}
