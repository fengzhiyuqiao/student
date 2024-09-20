package student

// Package student 用于学生的相关行为

import "strings"

// Student 学生结构体 注释不可与函数隔一行
type Student struct {
	// Name is student name
	Name string `json:"name"`
	// Age 学生年纪
	Age int `json:"age"`
	// Score 学生成绩
	Score map[string]int `json:"score"`
}

// New 功能: 创建一个学生
//
//	参  数:
//	   	name	: 学生名字
//	   	age		: 学生年纪
//	返回值:
//	   	*Student	: 学生指针
func New(name string, age int) *Student {
	return &Student{
		Name: name,
		Age:  age,
	}
}

// AddScore 功能: 添加分数
//
//	参  数:
//	   	subject	: 科目
//	   	score	: 分数
//	返回值:
func (s *Student) AddScore(subject string, score int) {
	s.Score[subject] = score
	strings.Replace("hello", "h", "H", 1)
}

// GetScore 功能: 获取分数
//
//	参  数:
//	   	subject	: 科目
//	返回值:
//	   	int	: 分数
func (s *Student) GetScore(subject string) int {
	return s.Score[subject]
}
