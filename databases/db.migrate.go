package databases

import (
	"log"

	"github.com/spf13/cobra"
)

var MigrateCMD = &cobra.Command{
	Use: "migrate",
	Short: "db migration",
	RunE: dbMigrate,
}

var migrateUP bool
var migrateDOWN bool

func init()  {
	MigrateCMD.Flags().BoolVarP(&migrateUP, "dbUP", "u", true, "run migration up")

	MigrateCMD.Flags().BoolVarP(&migrateDOWN, "dbDOWN", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	DBInit()
	defer db.Close()
	
	if migrateUP {
		err := DoMigrate("up")
		if err != nil {
			return err
		}
	}

	return nil
}

func DoMigrate(direction string) error {
	switch direction {
	case "up":
		_, err := db.Exec(`
			CREATE TABLE activities (
				activity_id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
				title VARCHAR(255),
				email VARCHAR(255),
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
			)`)
		if err != nil {
			return err
		}

		_, err = db.Exec(`
			CREATE TABLE todos (
				todo_id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
				activity_group_id BIGINT UNSIGNED DEFAULT 0,
				title VARCHAR(255),
				priority VARCHAR(255),
				is_active BOOLEAN DEFAULT true,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
				FOREIGN KEY (activity_group_id) 
				REFERENCES activities(activity_id) 
				ON DELETE CASCADE
			)`)
		if err != nil {
			return err
		}
	}

	log.Println("Migration up done")
	
	return nil
}