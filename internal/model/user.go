package model

import (
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	Role string `bson:"role"`
	DB   string `bson:"db"`
}

const (
	ROLE_DB_SEPERATOR = ":"
	ROLES_SEPERATOR   = ","
)

func (r *Role) String() string {
	return r.Role + ROLE_DB_SEPERATOR + r.DB
}

func ParseRoleFromStr(s string) (*Role, error) {
	s = strings.Trim(s, ROLE_DB_SEPERATOR)
	if len(s) == 0 {
		return nil, errors.New("could not parse role from empty string")
	}

	ss := strings.Split(s, ROLE_DB_SEPERATOR)
	if len(ss) != 2 {
		return nil, fmt.Errorf("could not parse role from invalid string: %s", s)
	}

	return &Role{ss[0], ss[1]}, nil
}

func ParseRolesFromStr(s string) ([]*Role, error) {
	if strings.Contains(s, " ") {
		return nil, errors.New("role string could not contains whitespace")
	}

	ss := strings.Split(s, ROLES_SEPERATOR)

	var roles []*Role

	for _, v := range ss {
		if v == "" {
			continue
		}

		role, err := ParseRoleFromStr(v)
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

type User struct {
	primitive.ObjectID `bson:"_id"`
	// UserID             string   `bson:"userId"`
	User       string   `bson:"user"`
	DB         string   `bson:"db"`
	Roles      []Role   `bson:"roles"`
	Mechanisms []string `bson:"mechanisms"`
}
