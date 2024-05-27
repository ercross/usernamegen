package usernamegen

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const usernamesCollection = "usernames_gen"

type mongoStorage struct {
	db  *mongo.Database
	cfg Config
}

func UseMongoDB(ctx context.Context, dsn string, cfg Config) Storage {
	m := mustConnect(ctx, dsn, cfg.Logger)
	m.cfg = cfg

	m.ensureUniqueUsernameIndex(ctx)

	// check that collection contains documents
	count, err := m.db.Collection(usernamesCollection).EstimatedDocumentCount(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to ensure usernames collection availability"))
	}
	if count <= m.cfg.GeneratorBatchSize {
		return m
	}

	generated := generateNewBatch(m.cfg.GeneratorBatchSize, m.cfg.Separator)
	m.SaveBatch(ctx, generated)
	return m
}

func mustConnect(ctx context.Context, dsn string, logger Logger) *mongoStorage {

	opt := options.Client().SetConnectTimeout(time.Second * 10).ApplyURI(dsn)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(fmt.Errorf("%s: failed to connect to mongo: %w", appName, err))
	}

	// ping the connection to be sure the connection is properly configured
	if err = client.Ping(ctx, nil); err != nil {
		panic(fmt.Errorf("%s: failed to ping mongo: %w", appName, err))
	}

	db := client.Database(usernamesCollection, &options.DatabaseOptions{})

	logger.LogInfoF("%s: successfully connected to MongoDB", appName)

	return &mongoStorage{
		db: db,
	}
}

// ensureUniqueUsernameIndex is idempotent.
// See https://www.mongodb.com/docs/manual/reference/method/db.collection.createIndex/#behaviors
func (m *mongoStorage) ensureUniqueUsernameIndex(ctx context.Context) {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetName("username_unique").SetUnique(true),
	}
	_, err := m.db.Collection(usernamesCollection).Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		panic(fmt.Errorf("%s: failed to create unique username index: %w", appName, err))
	}
}

func (m *mongoStorage) SaveBatch(ctx context.Context, names []Username) (savedCount int64) {
	opts := options.InsertMany().SetOrdered(false)
	var docs []interface{}
	for _, name := range names {
		docs = append(docs, bson.M{
			"username": name,
			"used":     false,
		})
	}
	result, err := m.db.Collection(usernamesCollection).InsertMany(ctx, docs, opts)
	if err != nil {
		m.cfg.Logger.LogErrorF(err, "%s: failed to save generated usernames batch", appName)
		return
	}
	return int64(len(result.InsertedIDs))
}

func (m *mongoStorage) FetchNewUsername(ctx context.Context) (string, error) {
	filter := bson.M{"used": false}
	result := m.db.Collection(usernamesCollection).FindOne(ctx, filter)
	if result.Err() != nil {
		return "", fmt.Errorf("failed to fetch new username: %w", result.Err())
	}
	var username Username
	if err := result.Decode(&username); err != nil {
		return "", fmt.Errorf("failed to decode username: %w", err)
	}
	go m.maybeGenerateNewBatch(context.Background())

	return username.Username, nil
}

func (m *mongoStorage) maybeGenerateNewBatch(ctx context.Context) {
	count, err := m.db.Collection(usernamesCollection).CountDocuments(ctx, bson.M{"used": false})
	if err != nil {
		m.cfg.Logger.LogErrorF(err, "%s: failed to obtain available usernames count", appName)
		return
	}
	if count > m.cfg.MinAvailable {
		return
	}

	generated := generateNewBatch(m.cfg.GeneratorBatchSize, m.cfg.Separator)
	savedCount := m.SaveBatch(ctx, generated)

	if savedCount < m.cfg.GeneratorBatchSize {
		m.maybeGenerateNewBatch(context.Background())
	}
}
