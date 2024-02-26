package mapper

import (
	"GoProject/main/enum"
	"GoProject/main/model"
	"GoProject/main/util"
	"fmt"
	"github.com/go-faker/faker/v4"
	"log"
	"xorm.io/xorm"
)

func Migration(db *xorm.Engine) {
	var beans = []any{
		new(model.User),
	}

	if err := db.Sync2(beans...); err != nil {
		log.Fatal("数据迁移发生错误", err.Error())
	}

	admin := &model.User{
		Username:  "admin@aurora.org",
		Authority: enum.Admin,
	}

	if has, _ := db.Exist(admin); !has {
		password := faker.Password()
		admin.Passwd = util.BcryptPasswd(password)

		if _, err := db.Omit("store", "group_id").Insert(admin); err != nil {
			log.Fatal("初始化管理员发生错误", err.Error())
		}

		fmt.Println("管理员账号:", admin.Username)
		fmt.Println("管理员密码:", password)
	}
}
