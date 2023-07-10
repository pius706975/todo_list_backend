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
	MigrateCMD.Flags().BoolVarP(&migrateUP, "dbUP", "u", false, "run migration up")

	MigrateCMD.Flags().BoolVarP(&migrateDOWN, "dbDOWN", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	DBInit()
	defer db.Close()

	up := DoMigrateUp(migrateUP)
	if up != nil {
		return up
	}

	down := DoMigrateDown(migrateDOWN)
	if down != nil {
		return down
	}

	return nil
}

func DoMigrateUp(up bool) error {
	
	switch up {
	case true:
		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS activities (
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
			CREATE TABLE IF NOT EXISTS todos (
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

		log.Println("Migration up done")
	}
	
	return nil
}

func DoMigrateDown(down bool) error {
	switch down {
	case true:
		_, err := db.Exec("DROP TABLE todos")
		if err != nil {
			return err
		}

		_, err = db.Exec("DROP TABLE activities")
		if err != nil {
			return err
		}

		log.Println("Migration down done")
	}

	return nil
}
