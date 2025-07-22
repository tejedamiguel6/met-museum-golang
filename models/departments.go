package models

type Department struct {
	DepartmentID int    `json:"departmentId"`
	DisplayName  string `json:"displayName"`
}

type DepartmentsResponse struct {
	Departments []Department `json:"departments"`
}
