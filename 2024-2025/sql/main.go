package main

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	_ "github.com/kataras/tablewriter"
	"log"
	"os"
	"sql-demo/entities"
	"sql-demo/utils"
	"time"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	_ = godotenv.Load()
	dburl := os.Getenv("DBURL")
	db, err := sql.Open("pgx", dburl)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, _ = db.Exec("SET search_path TO course")
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Minute * 3)

	// Ping and PingContext may be used to determine if communication with
	// the database server is still possible.
	//
	// When used in a command line application Ping may be used to establish
	// that further queries are possible; that the provided DSN is valid.
	//
	// When used in long running service Ping may be part of the health
	// checking system.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	status := "up"
	if err := db.PingContext(ctx); err != nil {
		status = "down"
	}
	log.Println(status)

	// A *DB is a pool of connections. Call Conn to reserve a connection for exclusive use.
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // Return the connection to the pool.

	// Print projects before update
	projects, err := GetProjects(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	utils.PrintProjects(projects)
	//for i, proj := range projects {
	//	log.Printf("%d: %v\n", i+1, proj)
	//}

	// Update project budgets by 10% increase for project after 2020
	loc, _ := time.LoadLocation("Europe/Sofia")
	const shortForm = "2006-Jan-02"
	startDate, _ := time.ParseInLocation(shortForm, "2020-Jan-01", loc)
	result, err := conn.ExecContext(ctx, `UPDATE projects SET budget = ROUND(budget * 1.1) WHERE start_date > $1;`, startDate)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Toatal budgets updated: %d\n", rows)

	// Print projects after update
	projects, err = GetProjects(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	utils.PrintProjects(projects)

	stmt, _ := db.Prepare("INSERT INTO projects(name,description,budge,start_date,finished,company_id) VALUES ($1,$2,$3,$4,$5);")

}

func GetProjects(ctx context.Context, conn *sql.Conn) ([]entities.Project, error) {
	rows, err := conn.QueryContext(ctx, "SELECT * FROM projects")
	if err != nil {
		log.Fatal(err)
	}

	projects := []entities.Project{}
	for rows.Next() {
		p := entities.Project{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Budget, &p.Finished, &p.StartDate, &p.CompanyID); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var userId uint
	for i, _ := range projects {
		userRows, err := conn.QueryContext(ctx, "SELECT user_id FROM projects_users WHERE project_id = $1", projects[i].ID)
		if err != nil {
			return nil, err
		}
		for userRows.Next() {
			if err := userRows.Scan(&userId); err != nil {
				userRows.Close()
				return nil, err
			}
			projects[i].UserID = append(projects[i].UserID, userId)
		}
		userRows.Close()
		if userRows.Err() != nil {
			return nil, userRows.Err()
		}
	}
	return projects, nil
}
