package main

import (
	"encoding/csv"
	"eod/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

	//set configurations
	stage1MaxThreads := 5
	stage2MaxThreads := 5
	stage3MaxThreads := 8
	inputPath := "csv/Before Eod.csv"
	outputPath := "csv/After Eod.csv"
	//freeCount is how many customers get bonus
	freeCount := 100
	//bonus is how much bonus will customer get
	bonus := float64(10)

	//read contents from csv
	csvContents, err := readCsv(inputPath)

	if err != nil {
		log.Fatal(err)
	}

	//prepare data, convert string to eod struct
	eodData := prepareData(csvContents)

	stage1(stage1MaxThreads, &eodData)
	stage2(stage2MaxThreads, &eodData)
	stage3(stage3MaxThreads, &eodData, freeCount, bonus)

	writeCsv(&eodData, outputPath)
}

//readCsv reads from *path* and return slice of slice of string and *error*
func readCsv(path string) ([][]string, error) {
	// open file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// remember to close the file at the end of the program
	defer f.Close()

	csvReader := csv.NewReader(f)

	//skip first row since it's a header
	csvReader.Read()

	csvContents, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return csvContents, nil
}

//prepareDate converts to a slice of eod model
func prepareData(csvContents [][]string) []models.EOD {
	eodData := make([]models.EOD, 0)

	for _, v := range csvContents {
		cols := strings.Split(v[0], ";")

		//convert row to eod
		var eod models.EOD

		index := 0
		inc := func(i *int) int { *i++; return *i }

		eod.ID, _ = strconv.Atoi(cols[index])
		eod.Name = cols[inc(&index)]
		eod.Age, _ = strconv.Atoi(cols[inc(&index)])
		eod.Balanced, _ = strconv.ParseFloat(cols[inc(&index)], 64)
		eod.PreviousBalanced, _ = strconv.ParseFloat(cols[inc(&index)], 64)
		eod.AveragedBalanced, _ = strconv.ParseFloat(cols[inc(&index)], 64)
		eod.FreeTransfer, _ = strconv.Atoi(cols[inc(&index)])

		eodData = append(eodData, eod)
	}

	return eodData
}

//stage1 averages the balanced & previos balanced and the result is set to averaged balanced
func stage1(maxThreads int, eodData *[]models.EOD) {
	wg := sync.WaitGroup{}
	wg.Add(maxThreads)

	lengthRows := len(*eodData)

	for i := 0; i < maxThreads; i++ {
		n := lengthRows / maxThreads

		go func(i int) {
			defer wg.Done()

			start := i * n
			end := n * (i + 1)

			if i == maxThreads-1 {
				end = lengthRows
			}

			for j := start; j < end; j++ {
				(*eodData)[j].AveragedBalanced = ((*eodData)[j].Balanced + (*eodData)[j].PreviousBalanced) / 2
				(*eodData)[j].No1ThreadNo = fmt.Sprintf("%d", i)
			}
		}(i)
	}
	wg.Wait()
}

//stage2 sets free transfer per user based on the balanced
func stage2(maxThreads int, eodData *[]models.EOD) {
	wg := sync.WaitGroup{}
	wg.Add(maxThreads)

	lengthRows := len(*eodData)

	for i := 0; i < maxThreads; i++ {
		n := lengthRows / maxThreads

		go func(i int) {
			defer wg.Done()

			start := i * n
			end := n * (i + 1)

			if i == maxThreads-1 {
				end = lengthRows
			}

			for j := start; j < end; j++ {
				// I think di antara 100-150 is excluding 100 and 150, so the extreme value is 101 and 149, cmiiw
				if (*eodData)[j].Balanced > 100 && (*eodData)[j].Balanced < 150 {
					(*eodData)[j].FreeTransfer = 5
					(*eodData)[j].No2aThreadNo = fmt.Sprintf("%d", i)
				} else if (*eodData)[j].Balanced > 150 {
					// (*eodData)[j].FreeTransfer = 25
					// my bad, get it wrong,
					// the right is to add the balance, not the free transfer
					(*eodData)[j].Balanced += 25
					(*eodData)[j].No2bThreadNo = fmt.Sprintf("%d", i)
				}
			}
		}(i)
	}
	wg.Wait()
}

//stage3 increase the balance of first *freeCount* users by bonus int
func stage3(maxThreads int, eodData *[]models.EOD, freeCount int, bonus float64) {
	wg := sync.WaitGroup{}
	wg.Add(maxThreads)

	for i := 0; i < maxThreads; i++ {
		lengthRows := len(*eodData)

		n := lengthRows / maxThreads

		go func(i int) {
			defer wg.Done()
			start := i * n
			end := n * (i + 1)

			if i == maxThreads-1 {
				end = lengthRows
			}

			for j := start; j < end; j++ {
				if (*eodData)[j].ID >= 0 && (*eodData)[j].ID <= freeCount {
					(*eodData)[j].Balanced += bonus
				}
				(*eodData)[j].No3ThreadNo = fmt.Sprintf("%d", i)
			}
		}(i)
	}
	wg.Wait()
}

//writeCsv writes the data to a csv file
func writeCsv(eodData *[]models.EOD, path string) error {
	fOut, err := os.Create(path)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return err
	}

	defer fOut.Close()

	headerAfter := "id;Nama;Age;Balanced;No 2b Thread-No;No 3 Thread-No;Previous Balanced;Average Balanced;No 1 Thread-No;Free Transfer;No 2a Thread-No"

	csvWriter := csv.NewWriter(fOut)
	csvWriter.Comma = ';'
	csvWriter.Write([]string{headerAfter})

	for _, v := range *eodData {
		csvWriter.Write([]string{v.ToString()})
	}

	csvWriter.Flush()

	return csvWriter.Error()
}
