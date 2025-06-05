package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectGORM() *gorm.DB {
	dsn := "host=aws-0-us-east-1.pooler.supabase.com user=postgres.vekkygjkpcymoaccgzlc password='CxCYfn*7z2$%S?x' dbname=postgres port=5432 sslmode=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar no banco Supabase: %v", err)
	}
	return db
}
