package mongo

type Options struct {
	URI      string
	Database string
}

func NewOptions(uri, database string) *Options {
	return &Options{
		URI:      uri,
		Database: database,
	}
}
