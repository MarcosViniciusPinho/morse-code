package main

import (
	"log/slog"
	"os"

	"github.com/MarcosViniciusPinho/morse-code/internal/application/dto"
	"github.com/MarcosViniciusPinho/morse-code/internal/application/handler"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		slog.Error("invalid command or number of parameters, it should be used as follows: ./morse-code \".... . -.-- / .--- ..- -.. .\"")
		return
	}

	input := os.Args[1]

	inputDTO := dto.NewMorseCodeInputDTO(input)
	handler := handler.NewMorseCodeHandler(inputDTO)
	outputDTO, err := handler.Process()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info(outputDTO.GetOutput())
}
