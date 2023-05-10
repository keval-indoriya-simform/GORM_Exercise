package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Employee struct {
	EmpID      int32      `gorm:"primaryKey" json:"emp_id,omitempty"`
	EmpName    string     `json:"emp_name,omitempty"`
	DeptName   string     `gorm:"type:varchar(50)" json:"dept_name,omitempty"`
	Salary     int32      `json:"salary,omitempty"`
	Department Department `gorm:"references:Dept_name;foreignKey:DeptName"`
}

type Department struct {
	Dept_name string `gorm:"primaryKey" json:"dept_name,omitempty"`
	Location  string `json:"location,omitempty"`
}

type EmployeeIndia struct {
	ID         int32
	Name       string
	Gender     string
	Department string
}

type EmployeeUK struct {
	ID         int32
	Name       string
	Gender     string
	Department string
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Day2 sslmode=disable password=Ke^@l081001 port=5432")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully Connected to Database")
	}

	defer func() {
		db, err := db.DB()
		if err != nil {
			fmt.Println(err)
		}
		err = db.Close()
		errorCheck(err)
		fmt.Println("Closing Database Connection")
	}()

	//db.Migrator().DropTable(&Employee{}, &Department{})
	err = db.AutoMigrate(&Employee{}, &Department{})
	errorCheck(err)

	//db.Create(&[]Employee{
	//	{
	//		EmpID:    101,
	//		EmpName:  "Mohan",
	//		DeptName: "Admin",
	//		Salary:   4000,
	//		Department: Department{
	//			Dept_name: "Admin",
	//			Location:  "Mumbai",
	//		},
	//	},
	//	{
	//		EmpID:    102,
	//		EmpName:  "Rajkumar",
	//		DeptName: "HR",
	//		Salary:   3000,
	//		Department: Department{
	//			Dept_name: "HR",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    103,
	//		EmpName:  "Akbar",
	//		DeptName: "IT",
	//		Salary:   4000,
	//		Department: Department{
	//			Dept_name: "IT",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    104,
	//		EmpName:  "Dorvin",
	//		DeptName: "Finance",
	//		Salary:   6500,
	//		Department: Department{
	//			Dept_name: "Finance",
	//			Location:  "Ahmedabad",
	//		},
	//	},
	//	{
	//		EmpID:    105,
	//		EmpName:  "Rohit",
	//		DeptName: "HR",
	//		Salary:   3000,
	//		Department: Department{
	//			Dept_name: "HR",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    106,
	//		EmpName:  "Rajesh",
	//		DeptName: "Finance",
	//		Salary:   5000,
	//		Department: Department{
	//			Dept_name: "Finance",
	//			Location:  "Ahmedabad",
	//		},
	//	},
	//	{
	//		EmpID:    107,
	//		EmpName:  "Preet",
	//		DeptName: "HR",
	//		Salary:   7000,
	//		Department: Department{
	//			Dept_name: "HR",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    108,
	//		EmpName:  "Maryam",
	//		DeptName: "Admin",
	//		Salary:   4000,
	//		Department: Department{
	//			Dept_name: "Admin",
	//			Location:  "Mumbai",
	//		},
	//	},
	//	{
	//		EmpID:    109,
	//		EmpName:  "Sanjay",
	//		DeptName: "IT",
	//		Salary:   6500,
	//		Department: Department{
	//			Dept_name: "IT",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    110,
	//		EmpName:  "Vasudha",
	//		DeptName: "IT",
	//		Salary:   7000,
	//		Department: Department{
	//			Dept_name: "IT",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    111,
	//		EmpName:  "Melinda",
	//		DeptName: "IT",
	//		Salary:   8000,
	//		Department: Department{
	//			Dept_name: "IT",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    112,
	//		EmpName:  "Komal",
	//		DeptName: "IT",
	//		Salary:   10000,
	//		Department: Department{
	//			Dept_name: "IT",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    113,
	//		EmpName:  "Gautham",
	//		DeptName: "Admin",
	//		Salary:   2000,
	//		Department: Department{
	//			Dept_name: "Admin",
	//			Location:  "Mumbai",
	//		},
	//	},
	//	{
	//		EmpID:    114,
	//		EmpName:  "Manisha",
	//		DeptName: "HR",
	//		Salary:   3000,
	//		Department: Department{
	//			Dept_name: "HR",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    115,
	//		EmpName:  "Chandni",
	//		DeptName: "IT",
	//		Salary:   4500,
	//		Department: Department{
	//			Dept_name: "IT",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    116,
	//		EmpName:  "Satya",
	//		DeptName: "Finance",
	//		Salary:   6500,
	//		Department: Department{
	//			Dept_name: "Finance",
	//			Location:  "Ahmedabad",
	//		},
	//	},
	//	{
	//		EmpID:    117,
	//		EmpName:  "Adarsh",
	//		DeptName: "HR",
	//		Salary:   3500,
	//		Department: Department{
	//			Dept_name: "HR",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    118,
	//		EmpName:  "Tejaswi",
	//		DeptName: "Finance",
	//		Salary:   5500,
	//		Department: Department{
	//			Dept_name: "Finance",
	//			Location:  "Ahmedabad",
	//		},
	//	},
	//	{
	//		EmpID:    119,
	//		EmpName:  "Cory",
	//		DeptName: "HR",
	//		Salary:   8000,
	//		Department: Department{
	//			Dept_name: "HR",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    120,
	//		EmpName:  "Monica",
	//		DeptName: "Admin",
	//		Salary:   5000,
	//		Department: Department{
	//			Dept_name: "Admin",
	//			Location:  "Mumbai",
	//		},
	//	},
	//	{
	//		EmpID:    121,
	//		EmpName:  "Rosalin",
	//		DeptName: "IT",
	//		Salary:   6000,
	//		Department: Department{
	//			Dept_name: "IT",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    122,
	//		EmpName:  "Ibrahim",
	//		DeptName: "IT",
	//		Salary:   8000,
	//		Department: Department{
	//			Dept_name: "IT",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    123,
	//		EmpName:  "Vikram",
	//		DeptName: "IT",
	//		Salary:   8000,
	//		Department: Department{
	//			Dept_name: "IT",
	//			Location:  "Banglore",
	//		},
	//	},
	//	{
	//		EmpID:    124,
	//		EmpName:  "Dheeraj",
	//		DeptName: "IT",
	//		Salary:   11000,
	//		Department: Department{
	//			Dept_name: "IT",
	//			Location:  "Banglore",
	//		},
	//	},
	//})
	//db.Create(&Department{Dept_name: "Marketing", Location: "Pune"})

	//1
	//  var value = map[string]interface{}{}
	//	db.Model(&Employee{}).Joins("inner join departments as dp on employees.dept_name=dp.dept_name").Select("MAX(salary),MIN(salary)").Where("dp.dept_name=? AND dp.location=?", "HR", "Banglore").Find(&value)
	//	fmt.Println(value)

	//2
	//var count int64
	//db.Debug().Model(&Employee{}).Where("dept_name IN ?", []string{"IT", "Finance"}).Count(&count)
	//fmt.Println(count)

	//3
	//var avgSalary = []map[string]interface{}{}
	//db.Debug().Model(&Employee{}).Select("dept_name, AVG(salary)").Group("dept_name").Find(&avgSalary)
	//for i := range avgSalary {
	//	fmt.Println(avgSalary[i])
	//}

	//4
	//var sum = []map[string]interface{}{}
	//db.Debug().Model(&Employee{}).Joins("Department", db.Select("location")).Select("SUM(salary)").Group("location").Find(&sum)
	//for i := range sum {
	//	fmt.Println(sum[i])
	//}

	//5
	//var dept = []map[string]interface{}{}
	//db.Debug().Model(&Employee{}).Select("dept_name, COUNT(*)").Group("dept_name").Having("COUNT(*) > 4").Find(&dept)
	//for i := range dept {
	//	fmt.Println(dept[i])
	//}

	//6
	//var emps []Employee
	//q1 := db.Model(&Employee{}).Joins("inner join departments as dp on employees.dept_name=dp.dept_name").Select("MAX(salary)").Where("dp.location=?", "Banglore")
	//q2 := db.Model(&Employee{}).Joins("inner join departments as dp on employees.dept_name=dp.dept_name").Select("MIN(salary)").Where("dp.location=?", "Mumbai")
	//db.Preload("Department").Where("salary IN ((?), (?))", q1, q2).Find(&emps)
	//fmt.Println(emps)

	//7
	//1st way
	//db.Debug().Raw("Create table employee_hr_admin AS ?", db.Model(&Employee{}).Joins("inner join departments as dp on employees.dept_name=dp.dept_name").Select("emp_id, emp_name, dp.dept_name, location").Where("dp.dept_name IN ?", []string{"HR", "Admin"})).Row()
	//db.Debug().Exec("Create table employee_hr_admin AS ?", db.Model(&Employee{}).Joins("inner join departments as dp on employees.dept_name=dp.dept_name").Select("emp_id, emp_name, dp.dept_name, location").Where("dp.dept_name IN ?", []string{"HR", "Admin"}))
	//2nd way
	//db.Raw("Create table employee_hr_admin AS SELECT emp_id, emp_name, dp.dept_name, location FROM employees inner join departments as dp on employees.dept_name=dp.dept_name WHERE dp.dept_name IN ('HR','Admin')").Row()

	//8
	//db.Exec("insert into employee_hr_admin ?", db.Debug().Model(&Employee{}).Joins("inner join departments as dp on employees.dept_name=dp.dept_name").Select("emp_id, emp_name, dp.dept_name, location").Where("dp.dept_name = ?", "IT"))

	//9
	//db.Debug().Preload("Department").Table("employees as e").Where("salary > (?)",
	//	db.Model(Employee{}).Select("AVG(salary)").Where("e.dept_name = dept_name").Group("dept_name"),
	//).Find(&emps)
	//for i := range emps {
	//	fmt.Println(emps[i])
	//}

	//10
	//db.Debug().Preload("Department").Model(Employee{}).Where("salary > (?)",
	//	db.Model(Employee{}).Select("salary").Where("emp_id = 109"),
	//).Find(&emps)
	//for i := range emps {
	//	fmt.Println(emps[i])
	//}

	//11
	//db.Debug().Model(Employee{}).Select("emp_name, dept_name, case when salary > (@avg) then 'Higher' when salary < (@avg) then 'Lower' when salary = (@avg) then 'Equal' END as relative_salary", sql.Named("avg", db.Select("avg(salary)").Table("employees"))).Find(&dept)
	//for i := range dept {
	//	fmt.Println(dept[i])
	//}

	//part 2
	db.AutoMigrate(&EmployeeIndia{}, &EmployeeUK{})

	//db.Create([]EmployeeIndia{
	//	{
	//		Name:       "Pranaya",
	//		Gender:     "Male",
	//		Department: "IT",
	//	},
	//	{
	//		Name:       "Priyanka",
	//		Gender:     "Female",
	//		Department: "IT",
	//	},
	//	{
	//		Name:       "Preety",
	//		Gender:     "Female",
	//		Department: "HR",
	//	},
	//	{
	//		Name:       "Subrat",
	//		Gender:     "Male",
	//		Department: "HR",
	//	},
	//	{
	//		Name:       "Anurag",
	//		Gender:     "Male",
	//		Department: "IT",
	//	},
	//})
	//
	//db.Create([]EmployeeUK{
	//	{
	//		Name:       "James",
	//		Gender:     "Male",
	//		Department: "IT",
	//	},
	//	{
	//		Name:       "Priyanka",
	//		Gender:     "Female",
	//		Department: "IT",
	//	},
	//	{
	//		Name:       "Sara",
	//		Gender:     "Female",
	//		Department: "HR",
	//	},
	//	{
	//		Name:       "Subrat",
	//		Gender:     "Male",
	//		Department: "HR",
	//	},
	//	{
	//		Name:       "Pam",
	//		Gender:     "Female",
	//		Department: "HR",
	//	},
	//})

	//1
	//db.Find(&dept)
	//db.Raw("(?) union (?)", db.Model(&EmployeeIndia{}).Select("employee_india.*, 'India' as country_name"), db.Model(&EmployeeUK{}).Select("employee_uks.*, 'UK' as country_name")).Find(&dept)
	//for i := range dept {
	//	fmt.Println(dept[i])
	//}

	//2
	//db.Debug().Model(&EmployeeUK{}).Not("name IN (?)", db.Model(&EmployeeIndia{}).Select("name")).Find(&dept)
	//for i := range dept {
	//	fmt.Println(dept[i])
	//}

	//3
	//db.Debug().Model(&EmployeeUK{}).Where("name IN (?)", db.Model(&EmployeeIndia{}).Select("name")).Find(&dept)
	//for i := range dept {
	//	fmt.Println(dept[i])
	//}
}
