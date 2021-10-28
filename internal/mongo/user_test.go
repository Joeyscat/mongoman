package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/joeyscat/mongoman/internal/model"
	"github.com/stretchr/testify/assert"
)

const (
	testUser02 = "test_02"
	testUser03 = "test_03"
	testUser04 = "test_04"
)

func TestListUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	type args struct {
		ctx    context.Context
		uri    string
		dbname string
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.User
		wantErr bool
	}{
		{
			"OK",
			args{ctx, "mongodb://test_01:123456@127.0.0.1:27017/test", ""},
			[]*model.User{{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListUsers(tt.args.ctx, tt.args.uri, tt.args.dbname)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(tt.want) != 0 {
				assert.NotEmpty(t, got)
			}

			for _, v := range got {
				t.Logf("%#+v\n", v)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	type args struct {
		ctx    context.Context
		u      UserToCreate
		uri    string
		dbname string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"OK",
			args{ctx, UserToCreate{
				testUser02, "12345",
				[]*model.Role{{Role: "readWrite", DB: "test"}},
			}, "mongodb://test_01:123456@127.0.0.1:27017/test", ""},
			false,
		},
		{
			"OK",
			args{ctx, UserToCreate{
				testUser03, "12345",
				[]*model.Role{{Role: "readWrite", DB: ""}, {Role: "userAdmin", DB: "test"}},
			}, "mongodb://test_01:123456@127.0.0.1:27017/test", ""},
			false,
		},
		{
			"OK",
			args{ctx, UserToCreate{
				testUser04, "12345",
				[]*model.Role{{Role: "userAdmin", DB: "test_01"}},
			}, "mongodb://test_01:123456@127.0.0.1:27017/test", ""},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.ctx, tt.args.u, tt.args.uri, tt.args.dbname); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	type args struct {
		ctx      context.Context
		username string
		uri      string
		dbname   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"OK",
			args{ctx, testUser02, "mongodb://test_01:123456@127.0.0.1:27017/test", ""},
			false,
		},
		{
			"OK",
			args{ctx, testUser03, "mongodb://test_01:123456@127.0.0.1:27017/test", ""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUser(tt.args.ctx, tt.args.username, tt.args.uri, tt.args.dbname); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_deleteUser(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		uri      string
		dbname   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := deleteUser(tt.args.ctx, tt.args.username, tt.args.uri, tt.args.dbname); (err != nil) != tt.wantErr {
				t.Errorf("deleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
