package modules

import "productzilla-academy/core-service/core"

const (
	MULTIPLE_CHOICE = "multiple_choise"
)

type Options struct {
	Label    string  `json:"label"`
	Value    string  `json:"value"`
	IsAnswer bool    `json:"is_answer"`
	Point    float64 `json:"point"`
}

type Question struct {
	core.Generic
	Type     string    `json:"type"`
	Question string    `json:"question"`
	Answer   string    `json:"answer,omitempty"`
	Options  []Options `json:"options,omitempty"`
}
