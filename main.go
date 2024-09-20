package main

import (
	"fmt"
	"github.com/fengzhiyuqiao/student/student"
)

func main() {
	fmt.Println("hello world")
	stu := student.New("feng", 18)
	AddStudentScore(stu, map[string]int{"math": 100})
	fmt.Println(stu.GetScore("math"))
}

// AddStudentScore 功能: 添加分数
//
//	参  数:
//	   	student	: *student.Student
//	返回值:
func AddStudentScore(student *student.Student, score map[string]int) {
	for k, v := range score {
		student.AddScore(k, v)
	}
}
