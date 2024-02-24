package migrations

import (
	"fmt"
	"os"

	"github.com/RajaSunrise/hotel/database"
	"github.com/RajaSunrise/hotel/models/entity"
)

func Migrate() {
	err := database.DB.AutoMigrate(
		&entity.User{},
		&entity.Room{},
		&entity.Review{},
		&entity.Promotion{},
		&entity.Facility{},
		&entity.Location{},
		&entity.Employee{},
		&entity.Content{},
		&entity.Booking{},
		&entity.Payment{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, []any{"succes to migrate"}...)
}