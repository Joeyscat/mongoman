package mongo

import (
	"context"
	"errors"

	"github.com/joeyscat/mongoman/internal/model"
	"github.com/joeyscat/mongoman/internal/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserInfoResult struct {
	Users []*model.User `bson:"users"`
	Ok    int           `bson:"ok"`
}

// ListUsers List all users for this db
//
// https://docs.mongodb.com/manual/reference/method/db.getUsers/
func ListUsers(ctx context.Context, uri, dbname string) ([]*model.User, error) {
	if dbname == "" {
		var err error

		dbname, err = pkg.GetDBNameFromURI(uri)
		if err != nil {
			return nil, errors.New("dbname is not specified, and dbname is not included in the uri")
		}
	}
	return listUsers(ctx, uri, dbname)
}

func listUsers(ctx context.Context, uri, dbname string) ([]*model.User, error) {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	var r UserInfoResult
	err = client.Database(dbname).RunCommand(ctx, bson.D{primitive.E{Key: "usersInfo", Value: 1}}).Decode(&r)
	if err != nil {
		return nil, err
	}

	if r.Ok != 1 {
		return nil, errors.New("command [usersInfo] return not ok")
	}

	return r.Users, nil
}

type UserToCreate struct {
	Name  string // not empty
	Pwd   string // not empty
	Roles []*model.Role
}

type SimpleResult struct {
	Ok int `bson:"ok"`
}

// CreateUser Create a user
//
// https://docs.mongodb.com/manual/reference/command/createUser/#mongodb-dbcommand-dbcmd.createUser
func CreateUser(ctx context.Context, u UserToCreate, uri, dbname string) error {
	if dbname == "" {
		var err error

		dbname, err = pkg.GetDBNameFromURI(uri)
		if err != nil {
			return errors.New("dbname is not specified, and dbname is not included in the uri")
		}
	}
	return createUser(ctx, u, uri, dbname)
}

func createUser(ctx context.Context, u UserToCreate, uri, dbname string) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	var roles bson.A
	for _, v := range u.Roles {
		roles = append(roles, bson.M{"role": v.Role, "db": v.DB})
	}

	var r SimpleResult
	err = client.Database(dbname).RunCommand(ctx,
		bson.D{
			primitive.E{Key: "createUser", Value: u.Name},
			primitive.E{Key: "pwd", Value: u.Pwd},
			primitive.E{Key: "roles", Value: roles},
		},
	).Decode(&r)
	if err != nil {
		return err
	}

	if r.Ok != 1 {
		return errors.New("command [usersInfo] return not ok")
	}
	return nil
}

// DeleteUser Delete a user
//
// https://docs.mongodb.com/manual/reference/command/dropUser/#mongodb-dbcommand-dbcmd.dropUser
func DeleteUser(ctx context.Context, username, uri, dbname string) error {
	if dbname == "" {
		var err error

		dbname, err = pkg.GetDBNameFromURI(uri)
		if err != nil {
			return errors.New("dbname is not specified, and dbname is not included in the uri")
		}
	}
	return deleteUser(ctx, username, uri, dbname)
}

func deleteUser(ctx context.Context, username, uri, dbname string) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	var r SimpleResult
	err = client.Database(dbname).RunCommand(ctx,
		bson.D{
			primitive.E{Key: "dropUser", Value: username},
		},
	).Decode(&r)
	if err != nil {
		return err
	}

	if r.Ok != 1 {
		return errors.New("command [usersInfo] return not ok")
	}
	return nil
}
