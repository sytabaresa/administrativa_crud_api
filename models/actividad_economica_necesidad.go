package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ActividadEconomicaNecesidad struct {
	Id                 int           `orm:"column(id);pk;auto"`
	SolicitudNecesidad *Necesidad    `orm:"column(solicitud_necesidad);rel(fk)"`
	ActividadEconomica string `orm:"column(actividad_economica)"`
}

func (t *ActividadEconomicaNecesidad) TableName() string {
	return "actividad_economica_necesidad"
}

func init() {
	orm.RegisterModel(new(ActividadEconomicaNecesidad))
}

// AddActividadEconomicaNecesidad insert a new ActividadEconomicaNecesidad into database and returns
// last inserted Id on success.
func AddActividadEconomicaNecesidad(m *ActividadEconomicaNecesidad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetActividadEconomicaNecesidadById retrieves ActividadEconomicaNecesidad by Id. Returns error if
// Id doesn't exist
func GetActividadEconomicaNecesidadById(id int) (v *ActividadEconomicaNecesidad, err error) {
	o := orm.NewOrm()
	v = &ActividadEconomicaNecesidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllActividadEconomicaNecesidad retrieves all ActividadEconomicaNecesidad matches certain condition. Returns empty list if
// no records exist
func GetAllActividadEconomicaNecesidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ActividadEconomicaNecesidad)).RelatedSel(5)
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

	var l []ActividadEconomicaNecesidad
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

// UpdateActividadEconomicaNecesidad updates ActividadEconomicaNecesidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateActividadEconomicaNecesidadById(m *ActividadEconomicaNecesidad) (err error) {
	o := orm.NewOrm()
	v := ActividadEconomicaNecesidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteActividadEconomicaNecesidad deletes ActividadEconomicaNecesidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteActividadEconomicaNecesidad(id int) (err error) {
	o := orm.NewOrm()
	v := ActividadEconomicaNecesidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ActividadEconomicaNecesidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}