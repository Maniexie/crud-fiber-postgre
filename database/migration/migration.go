package migration

import (
	"fmt"

	"github.com/Maniexie/crud-fiber-postgres/database"
	"github.com/Maniexie/crud-fiber-postgres/models/entities"
)

func RunMigrate() {
	err := database.DB.AutoMigrate(&entities.Mahasiswa{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Migrate Berhasil")
}
