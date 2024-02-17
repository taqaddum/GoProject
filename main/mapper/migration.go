package mapper

import (
	"GoProject/main/enum"
	"GoProject/main/model"
	"GoProject/main/util"
	"fmt"
	"github.com/go-faker/faker/v4"
	"log/slog"
	"xorm.io/xorm"
)

func Migration(db *xorm.Engine) error {
	var beans = []any{
		new(model.User),
	}

	if err := db.Sync2(beans...); err != nil {
		slog.Error("数据迁移发生错误", err.Error())
		return err
	}

	admin := &model.User{
		Username:  "admin@aurora.org",
		Authority: enum.Admin,
	}

	if has, _ := db.Exist(admin); !has {
		password := faker.Password()
		admin.Passwd = util.BcryptPasswd(password)

		if _, err := db.Omit("store", "group_id").Insert(admin); err != nil {
			slog.Error("初始化管理员发生错误", err.Error())
			return err
		}

		fmt.Println("管理员账号:", admin.Username)
		fmt.Println("管理员密码:", password)
	}

	return nil
}
