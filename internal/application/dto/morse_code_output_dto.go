package dto

type MorseCodeOutputDTO struct {
	output string
}

func NewMorseCodeOutputDTO(output string) MorseCodeOutputDTO {
	return MorseCodeOutputDTO{
		output: output,
	}
}

func (mcrd MorseCodeOutputDTO) GetOutput() string {
	return mcrd.output
}
