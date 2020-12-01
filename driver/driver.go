package driver

import "productzilla-academy/core-service/driver/config"

type Driver interface {
	Configuration() config.ConfigProvider
	Registry() Registry
	CallRegistry() Driver
}
