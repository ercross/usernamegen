package usernamegen

import "context"

const (
	// defaultGeneratorBatchSize is the number of names the generator should generate in a generateNewBatch command
	defaultGeneratorBatchSize = 5000

	defaultMinAvailable = 1000

	appName = "usernamegen"

	defaultSeparator = "_"
)

var defaultConfig = Config{
	GeneratorBatchSize: defaultGeneratorBatchSize,
	MinAvailable:       defaultMinAvailable,
	Separator:          defaultSeparator,
}

type Config struct {
	Separator          string
	GeneratorBatchSize int64

	// MinAvailable is the minimum count of unused unique username
	// available in the Storage
	MinAvailable int64
	Logger       Logger
}

type Logger interface {
	LogInfoF(msg string, args ...interface{})
	LogErrorF(err error, msg string, args ...interface{})
}

type Storage interface {

	// SaveBatch saves a new batch of usernames into Storage.
	// It can not be ascertained that each element in names will be unique,
	// hence SaveBatch must be designed to ignore error on such conflict
	//
	// Errors are logged to Logger
	SaveBatch(ctx context.Context, names []Username) (savedCount int64)

	// FetchNewUsername from Storage
	// Note that the fetched username can not be reused
	FetchNewUsername(ctx context.Context) (string, error)
}
