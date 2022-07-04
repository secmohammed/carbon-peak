package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)
func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    csvReader.Comma = ';'
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}
type TimeslotPeak struct {
    LowestTime int
    HighestTime int
    Points int
}
func main() {
   flag.Parse()
   filePath := flag.Arg(0)
    if filePath == "" {
        log.Fatal("No input file specified")
    }
    records := readCsvFile(filePath)
    var timeslotPeaks []TimeslotPeak

    for _, record := range records {

        start, err := strconv.Atoi(record[0])
        if err != nil {
            log.Fatal("Unable to parse start time for " + record[0], err)
        }
        end, err := strconv.Atoi(record[1])
        if err != nil {
            log.Fatal("Unable to parse end time for " + record[1], err)
        }
        points, err := strconv.Atoi(record[2])
        if err != nil {
            log.Fatal("Unable to parse points for " + record[2], err)
        }
        if len(timeslotPeaks) == 0 {
            timeslotPeaks = append(timeslotPeaks, TimeslotPeak{start, end, points})
        } else {
            var found = false

            for i, timeslotPeak := range timeslotPeaks {
                if start == timeslotPeak.LowestTime  && end == timeslotPeak.HighestTime {
                    timeslotPeaks[i].Points += points
                    found = true
                    break
                }
                if end > timeslotPeak.LowestTime && timeslotPeak.HighestTime >= end  {
                    timeslotPeaks[i].Points += points
                    found = true
                    break
                }
            }
            if !found {
                timeslotPeaks = append(timeslotPeaks, TimeslotPeak{start, end, points})
            }
        }


    }
    sort.Slice(timeslotPeaks, func(i, j int) bool {
        return timeslotPeaks[i].Points > timeslotPeaks[j].Points
    })
    fmt.Println(timeslotPeaks[0].Points)

}
