package monitoring

import (
	"context"

	"github.com/devpablocristo/golang/sdk/internal/core/monitoring/ports"
	portsmysql "github.com/devpablocristo/golang/sdk/pkg/databases/sql/mysql/go-sql-driver/ports"
)

type mysqlRepository struct {
	mysql portsmysql.Repository
}

func NewMySqlRepository(db portsmysql.Repository) ports.Repository {
	return &mysqlRepository{
		mysql: db,
	}
}

func (r *mysqlRepository) CheckDbConn(ctx context.Context) error {
	return r.mysql.DB().PingContext(ctx)
}
