package main

import (
	"log"

	"github.com/race-classification/race"
)

func main() {
	source := race.NewFileSource()
	repository := race.NewRepository(source)
	classifier := race.NewClassifier()
	service := race.NewService(repository, classifier)

	err := service.GenerateResults("data/race.csv", "data/result.csv")
	if err != nil {
		log.Fatal("Falha ao gerar o resultado: ", err)
	}
	log.Print("Arquivo gerado com sucesso em: ./data/result.csv")
}
