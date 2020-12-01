package modules

import (
	"productzilla-academy/core-service/core"
	"productzilla-academy/core-service/core/courses"
)

type Module struct {
	core.General

	Course courses.Course `json:"course"`

	ParentModule *Module `json:"-"`

	SubModules []*Module `json:"sub_modules"`

	Questions []Question `json:"questions"`
}
