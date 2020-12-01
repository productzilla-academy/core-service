package courses

import (
	"productzilla-academy/core-service/core"
	"productzilla-academy/core-service/core/career"
)

type Course struct {
	core.General
	Career career.Career `json:"career"`
}
