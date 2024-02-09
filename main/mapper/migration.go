package mapper

import (
	"GoProject/main/model"
	"log/slog"
	"xorm.io/xorm"
)

func Migration(db *xorm.Engine) error {
	var beans = []any{
		new(model.Users),
		new(model.Files),
		new(model.Folders),
		new(model.Shares),
		new(model.Folders),
		new(model.Groups),
		new(model.Downloads),
		new(model.Trash),
		new(model.Transfers),
	}

	if err := db.Sync2(beans...); err != nil {
		slog.Error("数据迁移发生错误", err.Error())
		return err
	}
	return nil
}
