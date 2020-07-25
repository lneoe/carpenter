package tool

import (
	"context"
	"fmt"

	"carpenter/service"
)

type UserTool struct{}

func (tool UserTool) Create(username, password string, quota int64) error {
	d := service.NewUserDao(context.TODO(), service.DB)
	password = service.PasswordSHA224(password)
	user, err := d.Create(username, password, quota)
	if err != nil {
		return err
	}

	fmt.Printf("User Created\n%+v\n", user)
	return nil
}

func (tool UserTool) List() error {
	d := service.NewUserDao(context.TODO(), service.DB)

	users, err := d.Fetch(0, 100)
	if err != nil {
		return err
	}

	// header: id	username	download	upload	quota
	format := "%[1]d\t%[2]s\t%.3[3]fGB\t%.3[4]fGB\t%.3[5]fGB\n"

	bytesToGB := func(in int64 /* bytes */) float64 {
		return float64(in) / 1024 / 1024 / 1024
	}

	printUser := func(user service.User) {
		fmt.Printf(format, user.ID, user.Username, bytesToGB(user.Download), bytesToGB(user.Upload), bytesToGB(user.Quota))
	}

	fmt.Printf("id\tusername\tdownload\tupload\tquota\n")
	for _, user := range users {
		printUser(user)
	}

	return nil
}

func (tool UserTool) Delete(id int64) error {
	d := service.NewUserDao(context.TODO(), service.DB)
	return d.Delete(id)
}
