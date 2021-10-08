package controllers

import (
	"bpk-dx/controllers/render"
	"bpk-dx/models"
	"bpk-dx/views"
	"fmt"
	"net/http"
)

type Emp struct{}

func (e Emp) EmpMain(resp http.ResponseWriter, r *http.Request) {
	emps, err := models.GetAllEmployees()
	if err != nil {
		fmt.Fprintln(resp, "GetAllEmployees error!", err)
		return
	}

	new := r.URL.Query().Get("new")

	viewData := &views.EmpMainIndexView{
		New:         new,
		Emp_list:    emps,
		ShowDelete:  false,
		Delete_list: []*models.Employee{},
	}
	if err := render.View(resp, "emp_main", viewData); err != nil {
		fmt.Fprintln(resp, "EmpMain error!", err)
	}
}

func (e Emp) EmpAdd(resp http.ResponseWriter, r *http.Request) {
	if err := render.View(resp, "emp_add", nil); err != nil {
		fmt.Fprintln(resp, "EmpAdd error!", err)
	}
}

func (e Emp) EmpCreate(resp http.ResponseWriter, r *http.Request) {
	emp, err := models.CreateEmployee(&models.Employee{
		Id:           r.PostFormValue("id"),
		Surname:      r.PostFormValue("surname"),
		Lastname:     r.PostFormValue("lastname"),
		Department:   r.PostFormValue("department"),
		Section_name: r.PostFormValue("section_name"),
		Section_code: r.PostFormValue("section_code"),
	})
	if err != nil {
		fmt.Fprintln(resp, "EmpCreate error", err)
	} else {
		http.Redirect(resp, r, "/employee?new="+emp.Id, http.StatusFound)
	}
}

func (e Emp) DeleteEmployee(resp http.ResponseWriter, r *http.Request) {
	sel := r.URL.Query().Get("id")
	//fmt.Fprintln(resp, "Select", sel)
	del, err := models.DeleteEmployee(sel)
	if err != nil {
		fmt.Fprintln(resp, "DeleteEmployee error", err)
	}
	emps, err := models.GetAllEmployees()
	if err != nil {
		fmt.Fprintln(resp, "GetAllEmployees error!", err)
		return
	}

	viewData := &views.EmpMainIndexView{
		New:         "",
		Emp_list:    emps,
		ShowDelete:  true,
		Delete_list: del,
	}
	fmt.Println(sel)
	if err := render.View(resp, "emp_main", viewData); err != nil {
		fmt.Fprintln(resp, "DeletEmployee error!", err)
	}
	//http.Redirect(resp, r, "/employee/", http.StatusFound)
}
