package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Informacion_persona_natural struct {
	Id                                int		 `orm:"column(num_documento_persona)"`
	TipoDocumento                     int		 `orm:"column(tipo_documento)"`
	DigitoVerificacion                int			`orm:"column(digito_verificacion)"`
	PrimerApellido                    string `orm:"size(128)"`
	SegundoApellido                   string `orm:"size(128)"`
	PrimerNombre                      string `orm:"size(128)"`
	SegundoNombre                     string `orm:"size(128)"`
	Cargo                             string `orm:"size(128)"`
	IdPaisNacimiento                  string `orm:"size(128)"`
	Perfil                            int
	Profesion                         string `orm:"size(128)"`
	Especialidad                      string `orm:"size(128)"`
	MontoCapitalAutorizado            int
	Genero                            string `orm:"size(128)"`
	GrupoEtnico                       string `orm:"size(128)"`
	ComunidadLgbt                     bool
	CabezaFamilia                     bool
	PersonasACargo                    bool
	NumeroPersonasACargo              int
	EstadoCivil                       string `orm:"size(128)"`
	Discapacitado                     bool
	TipoDiscapacidad                  string `orm:"size(128)"`
	DeclaranteRenta                   bool
	MedicinaPrepagada                 bool
	ValorUvtPrepagada                 int
	CuentaAhorroAfc                   bool
	NumCuentaBancariaAfc              string `orm:"size(128)"`
	IdEntidadBancariaAfc              int
	InteresViviendaAfc                int
	DependienteHijoMenorEdad          bool
	DependienteHijoMenos23Estudiando  bool
	DependienteHijoMas23Discapacitado bool
	DependienteConyuge                bool
	DependientePadreOHermano          bool
	IdNucleoBasico                    int
	IdArl                             int
	IdEps                             int
	IdFondoPension                    int
	IdCajaCompension                  int
	IdNitArl                          int
	IdNitEps                          int
	IdNitFondoPension                 int
	IdNitCajaCompensacion             int
}

func init() {
	orm.RegisterModel(new(Informacion_persona_natural))
}

// AddInformacion_persona_natural insert a new Informacion_persona_natural into database and returns
// last inserted Id on success.
func AddInformacion_persona_natural(m *Informacion_persona_natural) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInformacion_persona_naturalById retrieves Informacion_persona_natural by Id. Returns error if
// Id doesn't exist
func GetInformacion_persona_naturalById(id int) (v *Informacion_persona_natural, err error) {
	o := orm.NewOrm()
	v = &Informacion_persona_natural{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInformacion_persona_natural retrieves all Informacion_persona_natural matches certain condition. Returns empty list if
// no records exist
func GetAllInformacion_persona_natural(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Informacion_persona_natural))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Informacion_persona_natural
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateInformacion_persona_natural updates Informacion_persona_natural by Id and returns error if
// the record to be updated doesn't exist
func UpdateInformacion_persona_naturalById(m *Informacion_persona_natural) (err error) {
	o := orm.NewOrm()
	v := Informacion_persona_natural{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInformacion_persona_natural deletes Informacion_persona_natural by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInformacion_persona_natural(id int) (err error) {
	o := orm.NewOrm()
	v := Informacion_persona_natural{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Informacion_persona_natural{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
