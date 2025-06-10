package dto

type MorseCodeInputDTO struct {
	input string
}

func NewMorseCodeInputDTO(input string) MorseCodeInputDTO {
	return MorseCodeInputDTO{
		input: input,
	}
}

func (mcrd MorseCodeInputDTO) GetInput() string {
	return mcrd.input
}
