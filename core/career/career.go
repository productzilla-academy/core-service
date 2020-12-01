package career

import "productzilla-academy/core-service/core"

type Career struct {
	core.General
	Name string `json:"name"`
}
