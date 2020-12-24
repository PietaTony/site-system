package step

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //mysql lib
)

func dbConnect() *sql.DB {
	db, err := sql.Open(sqlType, sqlLogin)
	panicError(err)
	return db
}

func printResult(result sql.Result) {
	affected, err := result.RowsAffected()
	log.Printf("Rows Affected: %d", affected)
	printError(err)

	insertID, err := result.LastInsertId()
	log.Printf("Insert ID: %d", insertID)
	printError(err)
}

//Create use Step to create into "step" table without ID
func Create(step Step) (Step, error) {
	db := dbConnect()
	defer db.Close()

	query, err := db.Prepare(`
		INSERT INTO site_system.step (
			reader_num,
			tag_num,
			user_num,
			style,
			start,
			end,
			is_remove
		) Value (?,?,?,?,?,?,?)`)
	printError(err)
	defer query.Close()

	result, _ := query.Exec(
		step.ReaderNum,
		step.TagNum,
		step.UserNum,
		step.Style,
		step.Start,
		step.End,
		step.IsRemove)
	printResult(result)

	insertID, err := result.LastInsertId()
	step.ID = uint64(insertID)

	return step, err
}

//SelectAll SelectALL into "step" table
func SelectAll() (steps []Step, err error) {
	db := dbConnect()
	defer db.Close()

	result, err := db.Query(`SELECT
		id,
		reader_num,
		tag_num,
		user_num,
		style,
		start,
		end,
		is_remove
	from step`)
	printError(err)
	defer result.Close()

	steps = []Step{}
	for result.Next() {
		step := New()
		err = result.Scan(
			&step.ID,
			&step.ReaderNum,
			&step.TagNum,
			&step.UserNum,
			&step.Style,
			&step.Start,
			&step.End,
			&step.IsRemove)
		printError(err)
		if err != nil {
			return steps, err
		}
		log.Println(step)
		steps = append(steps, step)
	}
	return steps, err
}
