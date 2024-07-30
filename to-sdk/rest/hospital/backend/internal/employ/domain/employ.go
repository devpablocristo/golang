package domain

import "time"

type Employ struct {
	// `gorm:"-"` excluye el campo
	Person   `gorm:"-"`
	IDEmploy uint64 `json:"idEmploy" form:"idEmploy" gorm:"primary_key"`
	IDPerson uint64 `json:"idPerson" form:"idPerson gorm:"foreignkey:IDPerson"`

	CreadoEn      time.Time `gorm:"-"`
	ActualizadoEn time.Time `gorm:"-"`
	EliminadoEn   time.Time `gorm:"-"`

	PuestoEmploy       string `json:"puestoEmploy" form:"puestoEmploy" binding:"required"`
	MovilEmploy        string `json:"movilEmploy" form:"movilEmploy"`
	NumeroLegajoEmploy uint64 `json:"numeroLegajoEmploy" form:"numeroLegajoEmploy" binding:"required"`
	CelularEmploy      uint64 `json:"celularEmploy" form:"celularEmploy"`
}
