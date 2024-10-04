package models

import (
	orm "backend/lib/orm/ORM"
)

func MigrateModels(MIGRATEION_DIR string, o orm.ORM) {
	models := []interface{}{
		User{},
		Groups{},
		Post{},
		Visibility{},
		Comment{},
		Like{},
		Chat{},
		Event{},
		GroupMember{},
		Notifications{},
		Followers{},
		Invite{},
		orm.Model{},
	}

	o.AutoMigrate(MIGRATEION_DIR, models...)
}
