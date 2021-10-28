package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joeyscat/mongoman/internal/model"
	"github.com/joeyscat/mongoman/internal/mongo"
	"github.com/olekukonko/tablewriter"
)

func ListUsers(uri, dbname string) {
	users, err := mongo.ListUsers(context.Background(), uri, dbname)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Roles", "DB"})

	for _, user := range users {
		var roles string
		for _, r := range user.Roles {
			roles = roles + model.ROLES_SEPERATOR + r.String()
		}
		roles = strings.Trim(roles, model.ROLES_SEPERATOR)

		table.Append([]string{
			user.User, roles, user.DB,
		})
	}

	table.Render()
}

func CreateUser(username, password, rolesStr, uri, dbname string) {

	var roles []*model.Role
	var err error
	if rolesStr != "" {
		roles, err = model.ParseRolesFromStr(rolesStr)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	var u mongo.UserToCreate
	u.Roles = roles
	u.Name = username
	u.Pwd = password

	err = mongo.CreateUser(context.Background(), u, uri, dbname)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func DeleteUser(username, uri, dbname string) {

	err := mongo.DeleteUser(context.Background(), username, uri, dbname)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
