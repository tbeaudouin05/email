package main

import (
	"time"

	"github.com/thomas-bamilo/csv/gosqlcsv"
	"github.com/thomas-bamilo/email/goemail"
)

func main() {

	gosqlcsv.GoSQLCsv()
	time.Sleep(5 * time.Second)
	goemail.GoEmail()

}
