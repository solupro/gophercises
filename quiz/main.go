package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/labstack/gommon/log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var questions [][]string
var timeout int

func initQuestions(csvPath string)  {
	f, err := os.Open(csvPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	questions, err = r.ReadAll()
	if err != nil {
		panic(err)
	}
}

func main()  {
	var csvPath string
	flag.StringVar(&csvPath, "csv", "problems.csv", "question csv file path")
	flag.IntVar(&timeout, "t", 30, "timeout for the game")
	flag.Parse()

	initQuestions(csvPath)

	timer := time.NewTimer(time.Duration(timeout) * time.Second)

	total := 0
	correct := 0
	wrong := 0
	ch := make(chan string)

GAME_OVER:
	for {
		q, a := getQuestion()
		fmt.Printf("what %s? ", q)
		total += 1

		go func() {
			var input string
			fmt.Scanf("%s", &input)
			ch <- input
		}()

		select {
		case <-timer.C:
			fmt.Println("game over...")
			break GAME_OVER

		case answer := <-ch:
			num, err := strconv.Atoi(answer);
			if err != nil || num != a {
				fmt.Println("wrong!")
				wrong += 1
			} else {
				fmt.Println("correct!")
				correct += 1
			}
		}
	}

	fmt.Printf("total:%d, correct:%d, wrong:%d\n", total, correct, wrong)
}

func getQuestion() (string, int) {
	for {
		len := len(questions)
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		idx := r.Intn(len)

		question := questions[idx]
		a, err := strconv.Atoi(question[1])
		if err != nil {
			log.Error(err)
			questions = append(questions[:idx], questions[idx+1:]...)
			continue
		}
		q := question[0]
		return q, a
	}
}