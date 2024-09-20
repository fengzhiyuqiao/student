package main

import (
	"fmt"
	"github.com/fengzhiyuqiao/student/student"
)

func main() {
	fmt.Println("hello world")
	stu := student.New("feng", 18)
	stu.AddScore("math", 100)
	fmt.Println(stu.GetScore("math"))
}

// AddStudentScore 功能: 添加分数
//
//	参  数:
//	   	student	: *student.Student
//	返回值:
func AddStudentScore(student *student.Student) {
	student.AddScore("math", 100)
}
