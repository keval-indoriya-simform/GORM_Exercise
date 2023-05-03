package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Employee struct {
	EmpID      int32      `gorm:"primaryKey" json:"emp_id,omitempty"`
	EmpName    string     `json:"emp_name,omitempty"`
	EmpNo      int32      `json:"emp_no,omitempty"`
	Employment Employment `gorm:"foreignKey:EmpID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"employment,omitempty"`
}

type Employment struct {
	ID          int32  `json:"id,omitempty"`
	EmpID       int32  `json:"emp_id,omitempty"`
	EmpProfile  string `json:"emp_profile,omitempty"`
	EmpCountry  string `json:"emp_country,omitempty"`
	EmpJoinDate string `gorm:"type:date" json:"emp_join_date,omitempty"`
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	emps    []Employee
	country []string
)

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Day1 sslmode=disable password=Ke^@l081001 port=5432")

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

	//db.Migrator().DropTable(&Employee{}, &Employment{})
	err = db.AutoMigrate(&Employee{}, &Employment{})
	errorCheck(err)

	//db.Create(&[]Employee{
	//	{
	//		EmpID:   101,
	//		EmpName: "Keval",
	//		EmpNo:   7123,
	//		Employment: Employment{
	//			EmpProfile:  "trainee",
	//			EmpCountry:  "India",
	//			EmpJoinDate: "2023/02/07",
	//		},
	//	},
	//	{
	//		EmpID:   102,
	//		EmpName: "Juhi",
	//		EmpNo:   1234,
	//		Employment: Employment{
	//			EmpProfile:  "trainee",
	//			EmpCountry:  "US",
	//			EmpJoinDate: "2022/05/07",
	//		},
	//	},
	//	{
	//		EmpID:   103,
	//		EmpName: "Shaswat",
	//		EmpNo:   4657,
	//		Employment: Employment{
	//			EmpProfile:  "trainee",
	//			EmpCountry:  "Italy",
	//			EmpJoinDate: "2021/10/07",
	//		},
	//	},
	//	{
	//		EmpID:   104,
	//		EmpName: "Abhishek",
	//		EmpNo:   7777,
	//		Employment: Employment{
	//			EmpProfile:  "software developer",
	//			EmpCountry:  "venes",
	//			EmpJoinDate: "2022/05/07",
	//		},
	//	},
	//	{
	//		EmpID:   105,
	//		EmpName: "Kishan",
	//		EmpNo:   8888,
	//		Employment: Employment{
	//			EmpProfile:  "software developer",
	//			EmpCountry:  "Nepal",
	//			EmpJoinDate: "2021/05/07",
	//		},
	//	},
	//	{
	//		EmpID:   106,
	//		EmpName: "Heema",
	//		EmpNo:   9999,
	//		Employment: Employment{
	//			EmpProfile:  "software developer",
	//			EmpCountry:  "India",
	//			EmpJoinDate: "2022/05/07",
	//		},
	//	},
	//	{
	//		EmpID:   107,
	//		EmpName: "Hari",
	//		EmpNo:   1111,
	//		Employment: Employment{
	//			EmpProfile:  "manager",
	//			EmpCountry:  "Germany",
	//			EmpJoinDate: "2021/05/07",
	//		},
	//	},
	//	{
	//		EmpID:   108,
	//		EmpName: "alana",
	//		EmpNo:   2222,
	//		Employment: Employment{
	//			EmpProfile:  "software developer",
	//			EmpCountry:  "US",
	//			EmpJoinDate: "2019/07/07",
	//		},
	//	},
	//	{
	//		EmpID:   109,
	//		EmpName: "Manish",
	//		EmpNo:   3333,
	//		Employment: Employment{
	//			EmpProfile:  "software developer",
	//			EmpCountry:  "London",
	//			EmpJoinDate: "2021/05/07",
	//		},
	//	},
	//	{
	//		EmpID:   110,
	//		EmpName: "krisha",
	//		EmpNo:   4444,
	//		Employment: Employment{
	//			EmpProfile:  "HR",
	//			EmpCountry:  "canada",
	//			EmpJoinDate: "2019/07/07",
	//		},
	//	},
	//})

	//Q1. Display employee details and their profile where country = India
	//q1(db)

	//Q2. Display name of 5 countries order by country name
	//q2(db)

	//Q3. Display employee who joined in Year 2022 and are not from India
	//q3(db)

	//Q4. Display employee who joined in Between January and June 2022
	//q4(db)

	//Q5. Display employee details where country = Germany and has joined in 2021
	//q5(db)

	//Q6. Display employee name whose name has 'a' as 3rd letter and 'n' as 4th letter
	//q6(db)

	//Q7. Update employee name as Arjun where emp_id = 101
	//q7(db)

	//Q8. Delete details of emp_id 105
	//q8(db)

	//Q9. Display details of employees from India, US. (using IN)
	//q9(db)

	//Q10. Display details of all employees including their profile and country. (Include employees who are not in employment table)
	//q10(db)
	//db.Create(&Employee{
	//	EmpID:      111,
	//	EmpName:    "parth",
	//	EmpNo:      5555,
	//	Employment: Employment{},
	//})

}

func display(emps []Employee) {
	for i := range emps {
		data, _ := json.MarshalIndent(emps[i], "", "\t")
		fmt.Println(string(data))
	}
}

func q1(db *gorm.DB) {
	db.Debug().InnerJoins("Employment", db.Select("emp_profile").Where(&Employment{EmpCountry: "India"})).Find(&emps)
	display(emps)
}

func q2(db *gorm.DB) {
	db.Model(&Employment{}).Distinct("emp_country").Order("emp_country").Limit(5).Find(&country)
	fmt.Println(country)
}

func q3(db *gorm.DB) {
	db.Debug().Joins("Employment").Where("emp_country != ? and emp_join_date BETWEEN ? AND ?", "India", "2022-01-01", "2022-12-31").Find(&emps)
	display(emps)
}

func q4(db *gorm.DB) {
	db.Debug().Joins("Employment").Where("emp_join_date BETWEEN ? AND ?", "2022-01-01", "2022-06-30").Find(&emps)
	display(emps)
}

func q5(db *gorm.DB) {
	db.Debug().Joins("Employment").Where("emp_country = ? and DATE_PART('Year', emp_join_date) = ?", "Germany", "2021").Find(&emps)
	display(emps)
}

func q6(db *gorm.DB) {
	db.Debug().Joins("Employment").Where("emp_name LIKE ?", "__an%").Find(&emps)
	display(emps)
}

func q7(db *gorm.DB) {
	db.Model(&Employee{}).Where("emp_id = ?", 101).Update("emp_name", "Arjun")
}

func q8(db *gorm.DB) {
	db.Where("emp_id = ?", 105).Delete(&Employee{})
}

func q9(db *gorm.DB) {
	db.Debug().Joins("Employment").Where("emp_country IN ?", []string{"India", "US"}).Order("emp_id").Find(&emps)
	display(emps)
}

func q10(db *gorm.DB) {
	//1st way
	db.Debug().Joins("Employment").Find(&emps)
	//2nd way
	//db.Preload("Employment").Order("emp_id").Find(&emps)
	display(emps)
}
