package retenciones

import (
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/arriendamientosfideicomisos"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/dividendos"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/enajenaciondeacciones"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/fideicomisonoempresarial"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/intereses"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/intereseshipotecarios"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/operacionesderivados"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/pagosaextranjeros"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/planesretiro"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/premios"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/sectorfinanciero"
	"github.com/SaulEnriqueMR/kore-models/models/retenciones/complementos/serviciosplataformastecnologicas"
	"github.com/SaulEnriqueMR/kore-models/models/timbrefiscaldigital"
)

type Complemento struct {
	ArriendamientosFideicomisos      *arriendamientosfideicomisos.ArriendamientosFideicomisos           `xml:"Arrendamientoenfideicomiso" bson:"ArriendamientosFideicomisos,omitempty" json:"ArriendamientosFideicomisos,omitempty"`
	Dividendos                       *dividendos.Dividendos                                             `xml:"Dividendos" bson:"Dividendos,omitempty" json:"Dividendos,omitempty"`
	EnajenacionDeAcciones            *enajenaciondeacciones.EnajenacionDeAcciones                       `xml:"EnajenaciondeAcciones" bson:"EnajenacionDeAcciones,omitempty" json:"EnajenacionDeAcciones,omitempty"`
	FideicomisoEmpresarial           *fideicomisonoempresarial.FideicomisoEmpresarial                   `xml:"Fideicomisonoempresarial" bson:"FideicomisoEmpresarial,omitempty" json:"FideicomisoEmpresarial,omitempty"`
	Intereses                        *intereses.Intereses                                               `xml:"Intereses" bson:"Intereses,omitempty" json:"Intereses,omitempty"`
	InteresesHipotecarios            *intereseshipotecarios.InteresesHipotecarios                       `xml:"Intereseshipotecarios" bson:"InteresesHipotecarios,omitempty" json:"InteresesHipotecarios,omitempty"`
	OperacionesDerivados             *operacionesderivados.OperacionesDerivados                         `xml:"Operacionesconderivados" bson:"OperacionesDerivados,omitempty" json:"OperacionesDerivados,omitempty"`
	PagosAExtranjeros                *pagosaextranjeros.PagosAExtranjeros                               `xml:"Pagosaextranjeros" bson:"PagosAExtranjeros,omitempty" json:"PagosAExtranjeros,omitempty"`
	PlanesDeRetiro                   *planesretiro.PlanesDeRetiro                                       `xml:"Planesderetiro" bson:"PlanesDeRetiro,omitempty" json:"PlanesDeRetiro,omitempty"`
	Premios                          *premios.Premios                                                   `xml:"Premios" bson:"Premios,omitempty" json:"Premios,omitempty"`
	SectorFinanciero                 *sectorfinanciero.SectorFinanciero                                 `xml:"SectorFinanciero" bson:"SectorFinanciero,omitempty" json:"SectorFinanciero,omitempty"`
	ServiciosPlataformasTecnologicas *serviciosplataformastecnologicas.ServiciosPlataformasTecnologicas `xml:"ServiciosPlataformasTecnologicas" bson:"ServiciosPlataformasTecnologicas,omitempty" json:"ServiciosPlataformasTecnologicas,omitempty"`
	TimbreFiscalDigital              *timbrefiscaldigital.TimbreFiscalDigital                           `xml:"TimbreFiscalDigital" bson:"TimbreFiscalDigital,omitempty" json:"TimbreFiscalDigital,omitempty"`
}

type Addenda struct {
	Value string `xml:",chardata" bson:"Value" json:"Value"`
}
