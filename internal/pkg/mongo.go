package pkg

import (
	"errors"

	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

// mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb][?options]]
func GetDBNameFromURI(uri string) (string, error) {
	cs, err := connstring.ParseAndValidate(uri)
	if err != nil {
		return "", err
	}

	if cs.Database == "" {
		return "", errors.New("dbname is not included in the uri flag")
	}

	return cs.Database, nil
}
