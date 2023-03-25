package task21

import "fmt"

type RussianSpeaker interface {
	SpeakRussian() string
}

type EnglishSpeaker interface {
	SpeakEnglish() string
}

type RussianPerson struct{}

func (p *RussianPerson) SpeakRussian() string {
	return "Здравствуй!"
}

func GreetEnglish(speaker EnglishSpeaker) {
	fmt.Println("Hello,", speaker.SpeakEnglish())
}

type RussianSpeakerAdapter struct {
	russianSpeaker RussianSpeaker
}

func (s *RussianSpeakerAdapter) SpeakEnglish() string {
	return s.russianSpeaker.SpeakRussian()
}

func Run() {
	fmt.Println("--- Task 21 ---")

	russianPerson := RussianPerson{}
	adapter := RussianSpeakerAdapter{russianSpeaker: &russianPerson}
	GreetEnglish(&adapter)
}
