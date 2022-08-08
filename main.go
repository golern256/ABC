package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Event struct {
	Name       string
	Priority   string
	Procent    int
	SumProcent int
}

func countPart(task []Event) int {
	sumPriority := 0
	for i := 0; i < len(task); i++ {
		prior, _ := strconv.Atoi(strings.TrimSpace(task[i].Priority))
		sumPriority += prior
	}
	return sumPriority
}

func countProcent(task []Event) []Event {

	wholePriority := countPart(task)
	for i := 0; i < len(task); i++ {
		tempPrior, _ := strconv.Atoi(strings.TrimSpace(task[i].Priority))
		task[i].Procent = tempPrior * 100 / wholePriority
	}

	return task
}

func countSumProcent(task []Event) []Event {

	sort.SliceStable(task, func(i, j int) bool {
		return task[i].Procent > task[j].Procent
	})
	task[0].SumProcent = task[0].Procent
	for i := 1; i < len(task); i++ {
		task[i].SumProcent = task[i].Procent + task[i-1].SumProcent
	}

	return task
}
func createGroup(task []Event) {

	tasks := countProcent(task)
	finalarr := countSumProcent(tasks)

	for i := 0; i < len(finalarr); i++ {

		if finalarr[i].SumProcent <= 80 {
			fmt.Println("Group A", finalarr[i].Name)
		}

		if finalarr[i].SumProcent > 80 && finalarr[i].SumProcent <= 95 {
			fmt.Println("Group B", finalarr[i].Name)
		}

		if finalarr[i].SumProcent > 95 {
			fmt.Println("Group C", finalarr[i].Name)
		}
	}

	fmt.Println(finalarr)
}

func readFile(fileName string) []Event {

	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var tasks []Event
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		tasks = append(tasks, Event{
			Name:     line[0],
			Priority: line[1],
		})
	}
	return tasks
}

func main() {

	var fileName string = "tasks.csv"
	tasks := readFile(fileName)
	createGroup(tasks)
}
