package config

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

// Config - основная структура конфигурации, содержащая параметры сервера и базы данных.
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

// ServerConfig - структура для параметров конфигурации сервера.
type ServerConfig struct {
	Host string
	Port string
}

// PostgresConfig - структура для параметров конфигурации PostgreSQL.
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
	SslMode  string
}

// Inj - интерфейс для инъекции зависимостей
type Inj interface {
	DB() *pgxpool.Pool
}

// inj - структура, содержащая зависимости
type inj struct {
	db *pgxpool.Pool
}

// DB - возвращает объект подключения к базе данных
func (i *inj) DB() *pgxpool.Pool {
	return i.db
}

var I Inj

// MustLoad - функция загрузки конфигурации из .env файла
func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file: %v", err)
	}

	cfg := &Config{
		Server: ServerConfig{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("PORT"),
		},
		Postgres: PostgresConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Db:       os.Getenv("DB_NAME"),
			SslMode:  os.Getenv("DB_SSLMODE"),
		},
	}

	log.Println("Configuration successfully loaded")
	return cfg
}

// Init - инициализация конфигурации и создание подключения к базе данных
func (cfg *Config) Init() *Config {
	var (
		i   = &inj{}
		err error
	)

	// Инициализация клиента базы данных
	dbURL := "postgres://" + cfg.Postgres.User + ":" + cfg.Postgres.Password + "@" +
		cfg.Postgres.Host + ":" + cfg.Postgres.Port + "/" + cfg.Postgres.Db + "?sslmode=" + cfg.Postgres.SslMode

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if i.db, err = pgxpool.New(ctx, dbURL); err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("Successfully connected to PostgreSQL")

	I = i
	return cfg
}
