package main

import (
    "os"
    "fmt"
    "encoding/csv"
    "bufio"
    "log"
    "strconv"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("Usage: %s FILENAME\n", os.Args[0])
        os.Exit(0)
    }
    var filename = os.Args[1]
    //fmt.Println(filename)
    fh, _ := os.Open(filename)
    reader := bufio.NewReader(fh)
    r := csv.NewReader(reader)
	r.Comma = ';'

    records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%T\n", records)
    var sum = 0
    for _, row := range records {
        val, _ := strconv.Atoi(row[2])
        sum += val
    }
    fmt.Println(sum)
}
