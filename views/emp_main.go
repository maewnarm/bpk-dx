package views

import "bpk-dx/models"

type EmpMainIndexView struct {
	New         string
	Emp_list    []*models.Employee
	ShowDelete  bool
	Delete_list []*models.Employee
}
