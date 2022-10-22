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
	scanner := bufio.NewScanner(os.Stdin)

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

}
