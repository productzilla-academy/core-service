package assignment

import (
	"productzilla-academy/core-service/core/courses"
	"productzilla-academy/core-service/core/courses/modules"
)

type Assignment struct {
	Module    *modules.Module    `json:"module"`
	Course    *courses.Course    `json:"course"`
	Questions []modules.Question `json:"questions"`
}
