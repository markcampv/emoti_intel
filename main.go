package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql",
		"user:password/journaldb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var emotion string
	var response string
	var journal string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hey there, what would you like to do today? You can either write or read")
	fmt.Scanln(&journal)
	switch journal {
	case "write":
		fmt.Println("Choose your emotions wisely. How are you feeling right now?")
		fmt.Scanln(&emotion)
		switch emotion {
		case "happy":
			fmt.Println("Looks like you're in a good mood today! Tell me more :)")
			//let user input how they feel
			if scanner.Scan() {
				response = scanner.Text()
			}
		case "excited":
			fmt.Println("Must be something happening in your life today. Spill the beans!")
			//let user input how they feel
			if scanner.Scan() {
				response = scanner.Text()
			}
		case "anger":
			fmt.Println("Hey, it's ok. we've all been there. Let's count to 3 together...")
			//let user input how they feel
			if scanner.Scan() {
				response = scanner.Text()
			}
		case "sad":
			//let user input how they feel
			fmt.Println("One step at a time. Tell me how you feel")
			if scanner.Scan() {
				response = scanner.Text()
			}
		case "dissapointment":
			fmt.Println("I get it, things happen to the best of us. What happened?")
			//let user input how they feel
			if scanner.Scan() {
				response = scanner.Text()
			}
		default:
			fmt.Println("Take your time! It's ok not to know how you feel. Come back when you're in the mood to jot something down.")

		}

		//write code to save data in database so that user can see history

		date := time.Now()
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		stmt, err := tx.Prepare("insert into emoti(Emotion, Response, JournalDate) values(?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(emotion, response, date)
		if err != nil {
			log.Fatal(err)
		}

		tx.Commit()
	case "read":
		fmt.Println("Fetching journal entries")
		rows, err := db.Query("select * from emoti")
		if err != nil {
			log.Fatal(err)
		}
		//Get column names
		columns, err := rows.Columns()
		if err != nil {
			log.Fatal(err)
		}
		//Make a slice for the values
		values := make([]interface{}, len(columns))

		// rows.Scan wants '[]interface{}' as an argument, so we must copy the
		// references into such a slice
		scanArgs := make([]interface{}, len(values))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		// Fetch rows
		for rows.Next() {
			err = rows.Scan(scanArgs...)
			if err != nil {
				log.Fatal(err)
			}

			// Print data
			for i, value := range values {
				switch value.(type) {
				case nil:
					fmt.Println(columns[i], ":NULL")
				case []byte:
					fmt.Println(columns[i], ": ", string(value.([]byte)))
				default:
					fmt.Println(columns[i], ": ", value)
				}
			}
			fmt.Println("-----------------------------------")
		}
	}

}
