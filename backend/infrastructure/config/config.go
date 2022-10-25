package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Database          Database
	Grobid            Grobid
	IntBycrptPassword int    `env:"INT_BYCRPT_PASSWORD" validate:"required"`
	JWTTokenSecret    string `env:"JWT_TOKEN_SECRET" validate:"required"`
	S3                S3     `env:"S3" validate:"required"`
}

type Grobid struct {
	GrobidUrlPdfToTei string `env:"GROBID_URL_PDF_TO_TEI" validate:"required"`
}

type S3 struct {
	AccessKeyIdS3     string `env:"ACCESS_KEY_ID_S3" validate:"required"`
	SecretAccessKeyS3 string `env:"SECRET_ACCESS_KEY_S3" validate:"required"`
	BucketName        string `env:"BUCKET_NAME" validate:"required"`
	Endpoint          string `env:"ENDPOINT" validate:"required"`
}

type Database struct {
	DBName          string        `env:"MYSQL_DBNAME" default:"root" validate:"required"`
	DBUser          string        `env:"MYSQL_DBUSER" default:"root" validate:"required"`
	DBPass          string        `env:"MYSQL_DBPASS" default:"root"`
	Host            string        `env:"MYSQL_HOST" default:"localhost" validate:"required"`
	Port            string        `env:"MYSQL_PORT" default:"3306" validate:"required"`
	MaxOpenConns    int           `env:"MYSQL_MAX_OPEN_CONNS" default:"30" validate:"required"`
	MaxIdleConns    int           `env:"MYSQL_MAX_IDLE_CONNS" default:"6" validate:"required"`
	ConnMaxLifetime time.Duration `env:"MYSQL_CONN_MAX_LIFETIME" default:"30m" validate:"required"`
	MaxIdleTime     time.Duration `env:"MYSQL_MAX_IDLE_TIME" default:"0"`
}

func New() (Config, error) {
	Config := Config{}
	// build config from env
	log.Println("Mapping Env...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//mysql config
	Config.Database.DBName = os.Getenv("MYSQL_DBNAME")
	Config.Database.DBUser = os.Getenv("MYSQL_DBUSER")
	Config.Database.DBPass = os.Getenv("MYSQL_DBPASS")
	Config.Database.Host = os.Getenv("MYSQL_HOST")
	Config.Database.Port = os.Getenv("MYSQL_PORT")

	// bcrypt
	Int, err := strconv.Atoi(os.Getenv("INT_BYCRPT_PASSWORD"))
	if err != nil {
		log.Fatal("WRONG ENV: INT_BYCRPT_PASSWORD")
	}
	Config.IntBycrptPassword = Int

	// jwt token secret
	Config.JWTTokenSecret = os.Getenv("JWT_TOKEN")

	// Grobid Service
	Config.Grobid.GrobidUrlPdfToTei = os.Getenv("GROBID_URL_PDF_TO_TEI")

	// S3 Storage
	Config.S3.AccessKeyIdS3 = os.Getenv("ACCESS_KEY_ID_S3")
	Config.S3.SecretAccessKeyS3 = os.Getenv("SECRET_ACCESS_KEY_S3")
	Config.S3.BucketName = os.Getenv("BUCKET_NAME")
	Config.S3.Endpoint = os.Getenv("ENDPOINT")

	return Config, nil
}
