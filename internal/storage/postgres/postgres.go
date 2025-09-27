package postgres

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewDB crea y retorna un pool de conexiones a la base de datos
// usando sql.Open() pero con el driver de 'pgx'. Se necesita el
// connection String (Data Source Name) creado anteriormente.
func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Se realiza la verificacion comprobando que la base anda y que las
	// credenciales son correctas.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
