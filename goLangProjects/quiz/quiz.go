package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	var quizTime int
	fmt.Println("Enter the time for the quiz")
	fmt.Scanf("%d", &quizTime)
	questions, err := getQuestions() //reads csv file and returns the questions
	if err != nil {
		fmt.Println("Error in Parsing file " + err.Error())
	}
	//fmt.Println(questions)
	score := quiz(questions, quizTime)
	fmt.Printf("Your score is %d", score)
}

func quiz(questions [][]string, quiztime int) int {
	var score int
	go getAnswers(questions, &score)                  //passing the score as a pointer
	time.Sleep(time.Second * time.Duration(quiztime)) //Sleeping for the quizTime
	return score
}

func getAnswers(questions [][]string, score *int) {
	var ans string

	for i := 0; i < len(questions); i++ {
		fmt.Print("Problem #" + strconv.Itoa(i) + ":" + questions[i][0] + "=")
		fmt.Scan(&ans)
		if ans == questions[i][1] {
			*score++
		}

	}

}

func getQuestions() ([][]string, error) {
	questionsByte, err := ioutil.ReadFile("problems.csv") // Read File returns a byte array
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	allQuestions := string(questionsByte) //Converting the bytes array into string
	questionsReader := strings.NewReader(allQuestions)
	csvReader := csv.NewReader(questionsReader)
	questions, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return questions, err //questions is a string 2D array where each row element contains each row of csv File

}
