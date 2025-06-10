package handler

import (
	"fmt"

	"github.com/MarcosViniciusPinho/morse-code/internal/application/dto"
	"github.com/MarcosViniciusPinho/morse-code/internal/core/domain"
)

type MorseCodeHandler struct {
	inputDTO dto.MorseCodeInputDTO
}

func NewMorseCodeHandler(inputDTO dto.MorseCodeInputDTO) MorseCodeHandler {
	return MorseCodeHandler{
		inputDTO: inputDTO,
	}
}

func (mch MorseCodeHandler) Process() (dto.MorseCodeOutputDTO, error) {
	domain := domain.NewMorseCodeDomain(mch.inputDTO.GetInput())
	err := domain.Process()
	if err != nil {
		return dto.MorseCodeOutputDTO{}, fmt.Errorf("error trying to process morse code to text: %v", err)
	}
	return dto.NewMorseCodeOutputDTO(domain.GetOutput()), nil
}
