package service

import "quiz-game/models"

type DataParser interface {
	Parse([][]string) []models.Problem
}

type DataParserImpl struct {
}

func NewDataParser() DataParser {
	return new(DataParserImpl)
}

func (i *DataParserImpl) Parse(data [][]string) []models.Problem {
	problems := make([]models.Problem, len(data))
	for i, value := range data {
		problems[i] = models.Problem{
			Question: value[0],
			Answer:   value[1],
		}
	}
	return problems
}
