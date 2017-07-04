package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type NecesidadRechazada struct {
	Id            int        `orm:"column(id);pk;auto"`
	Necesidad     *Necesidad `orm:"column(necesidad);rel(fk)"`
	Justificacion string     `orm:"column(justificacion)"`
	Fecha         time.Time  `orm:"column(fecha);type(date)"`
}

func (t *NecesidadRechazada) TableName() string {
	return "necesidad_rechazada"
}

func init() {
	orm.RegisterModel(new(NecesidadRechazada))
}

// AddNecesidadRechazada insert a new NecesidadRechazada into database and returns
// last inserted Id on success.
func AddNecesidadRechazada(m *NecesidadRechazada) (id int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	m.Fecha = time.Now()
	if _, err = o.Insert(m); err!=nil{
		o.Rollback()
		return
	} else {
		m.Necesidad.Estado.Id=3;
		if _,err = o.Update(m.Necesidad); err!=nil{
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}

// GetNecesidadRechazadaById retrieves NecesidadRechazada by Id. Returns error if
// Id doesn't exist
func GetNecesidadRechazadaById(id int) (v *NecesidadRechazada, err error) {
	o := orm.NewOrm()
	v = &NecesidadRechazada{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNecesidadRechazada retrieves all NecesidadRechazada matches certain condition. Returns empty list if
// no records exist
func GetAllNecesidadRechazada(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NecesidadRechazada)).RelatedSel(5)
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

	var l []NecesidadRechazada
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

// UpdateNecesidadRechazada updates NecesidadRechazada by Id and returns error if
// the record to be updated doesn't exist
func UpdateNecesidadRechazadaById(m *NecesidadRechazada) (err error) {
	o := orm.NewOrm()
	v := NecesidadRechazada{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNecesidadRechazada deletes NecesidadRechazada by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNecesidadRechazada(id int) (err error) {
	o := orm.NewOrm()
	v := NecesidadRechazada{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NecesidadRechazada{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}