package mongodbdriver

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	instance MongoDBClientPort
	once     sync.Once
	errInit  error
)

type MongoDBClientPort interface {
	DB() *mongo.Database
	Close()
}

// MongoDBClient representa um cliente para interagir com um banco de dados MongoDB
type MongoDBClient struct {
	db *mongo.Database
}

// InitializeMongoDBClient inicializa o cliente MongoDB usando o padrão singleton
func InitializeMongoDBClient(config MongoDBClientConfig) error {
	once.Do(func() {
		client := &MongoDBClient{}
		errInit = client.connect(config)
		if errInit != nil {
			instance = nil
		} else {
			instance = client
		}
	})
	return errInit
}

// GetMongoDBInstance retorna a instância do cliente MongoDB
func GetMongoDBInstance() (MongoDBClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("MongoDB client is not initialized")
	}
	return instance, nil
}

// connect estabelece a conexão com o banco de dados MongoDB utilizando a configuração fornecida
func (client *MongoDBClient) connect(config MongoDBClientConfig) error {
	dsn := config.dsn()
	clientOptions := options.Client().ApplyURI(dsn)

	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Verificar a conexão
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = conn.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	client.db = conn.Database(config.Database)
	return nil
}

// Close fecha a conexão com o banco de dados MongoDB
func (client *MongoDBClient) Close() {
	if client.db != nil {
		client.db.Client().Disconnect(context.TODO())
	}
}

// DB retorna a conexão com o banco de dados MongoDB
func (client *MongoDBClient) DB() *mongo.Database {
	return client.db
}