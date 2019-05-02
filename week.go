package vk_bmstu_schedule_status

import "time"

type WeekType int

const (
	TypeNumerator WeekType = iota
	TypeDenominator
)

func (weekType WeekType) FullWeekName() string {
	switch weekType {
	case TypeNumerator:
		return "числитель"
	case TypeDenominator:
		return "знаменатель"
	}
	return ""
}

func (weekType WeekType) ShortWeekName() string {
	switch weekType {
	case TypeNumerator:
		return "ЧС"
	case TypeDenominator:
		return "ЗН"
	}
	return ""
}

func WeekTypeFromNumber(weekNumber int) WeekType {
	if weekNumber%2 == 0 {
		return TypeDenominator
	}
	return TypeNumerator
}

func CurrentWeek() int {
	_, week := time.Now().ISOWeek()
	return week - 5
}
