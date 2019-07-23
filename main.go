package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
)

const ver = "1.0.0"

var dir = flag.String("d", ".", "Migration dir")
var pgpassword = flag.String("p", "./.pgpass", "PGPASSWORD file path")
var table = flag.String("t", "ishare_migrations", "Migration table")
var version = flag.Bool("v", false, "Print version")

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("dbmig version %s\n", ver)
		return
	}

	// 1. read password file
	line, err := ioutil.ReadFile(*pgpassword)
	if err != nil {
		log.Panicf("can not read PGPASSWORD file: %s", err.Error())
	}

	content := strings.TrimSpace(string(line))
	comps := strings.Split(content, ":")

	// 2. connect db
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		comps[3], comps[4], comps[0], comps[1], comps[2])
	db, err := sql.Open("postgres", connStr)

	// Loop over dir
	files, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Panicf("Failed to read migration dir: %s", err.Error())
	}

	re := regexp.MustCompile("([0-9a-zA-Z_]+).sql$")
	for _, f := range files {
		matches := re.FindStringSubmatch(f.Name())
		if f.IsDir() || len(matches) != 2 {
			continue
		}

		sql := fmt.Sprintf("SELECT 1 FROM %s WHERE migration_name = $1 LIMIT 1", *table)
		rows, err := db.Query(sql, matches[1])
		if err != nil {
			log.Panicf("Failed to select migrations table %s: %s\n", *table, err.Error())
		}

		log.Printf("Checking %s...", matches[1])
		if !rows.Next() {
			// Not yet loaded
			path := path.Join(*dir, f.Name())
			data, err := ioutil.ReadFile(path)
			if err != nil {
				log.Panicf("Failed to read %s: %s\n", f.Name(), err.Error())
			}

			sql = string(data)
			_, err = db.Exec(sql)
			if err != nil {
				log.Panicf("Failed to apply migration %s: %s\n", f.Name(), err.Error())
			}

			// Insert migration record
			sql = fmt.Sprintf("INSERT INTO %s (migration_name) VALUES ($1)", *table)
			_, err = db.Exec(sql, matches[1])
			if err != nil {
				log.Panicf("Failed to insert migration record: %s\n", err.Error())
			}

			log.Printf("Applied\n")
		} else {
			// Already Load it
			log.Printf("Already loaded\n")
		}
	}
}
