package models

type Employee_Reg struct {
	Name          string `db:"emp_name,omitempty"`
	LastName      string `db:"emp_lastname,omitempty"`
	Email         string `db:"emp_email,omitempty"`
	Mobile_Number string `db:"mobile_number,omitempty"`
	Job_Title     string `db:"emp_job_title,omitempty"`
	Password      string `db:"emp_password,omitempty"`
}

type ResponseDB struct {
	Id int64 `db:"emp_id"`
	// Id int64 `db:"user_id"`

}
