package postgres

import (
	"RestAPIv2/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func InitDB(cfg *config.Config) (*Storage, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("db connection error: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("couldn't connect to the DB: %v", err)
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave, alias string) (int64, error) {
	var id int64
	err := s.db.QueryRow("INSERT INTO url (url, alias) VALUES ($1, $2) RETURNING id", urlToSave, alias).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("couldn't insert URL: %v", err)
	}
	return id, nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	var url string
	err := s.db.QueryRow("SELECT url FROM url WHERE alias = $1", alias).Scan(&url)
	if err != nil {
		return "", fmt.Errorf("couldn't get URL: %v", err)
	}

	return url, nil
}

func (s *Storage) DeleteURL(alias string) error {
	_, err := s.db.Exec("DELETE FROM url WHERE alias = $1", alias)
	if err != nil {
		return fmt.Errorf("couldn't delete URL: %v", err)
	}

	return nil
}
