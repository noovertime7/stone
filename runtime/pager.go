package runtime

import "gorm.io/gorm"

type DataBasePager interface {
	GetPage() int
	GetPageSize() int
	Fitter
}

type DefaultPager interface {
	GetPage() int
	GetPageSize() int
	GetKeyWord() string
}

type Fitter interface {
	IsFitter() bool
	Do(tx *gorm.DB)
}
