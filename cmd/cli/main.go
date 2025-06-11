package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/MarcosViniciusPinho/morse-code/internal/application/dto"
	"github.com/MarcosViniciusPinho/morse-code/internal/application/handler"
)

func main() {
	input := strings.Join(os.Args[1:], " ")

	inputDTO := dto.NewMorseCodeInputDTO(input)
	handler := handler.NewMorseCodeHandler(inputDTO)
	outputDTO, err := handler.Process()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info(outputDTO.GetOutput())
}
