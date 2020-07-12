package main

import (
	"GoWeb/src/web03_actions/model"
	"html/template"
	"net/http"
)

//测试range
func testRange(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("range.html"))
	var emps []*model.Employee
	emp := &model.Employee{
		ID:       1,
		LastName: "李小璐",
		Email:    "lxl@jnl.com",
	}
	emps = append(emps, emp)
	emp2 := &model.Employee{
		ID:       2,
		LastName: "白百何",
		Email:    "bbh@cyf.com",
	}
	emps = append(emps, emp2)
	emp3 := &model.Employee{
		ID:       3,
		LastName: "马蓉",
		Email:    "mr@wbq.com",
	}
	emps = append(emps, emp3)

	//执行
	t.Execute(w, emps)
}

//测试if
func testIf(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("if.html"))
	age := 17
	//执行
	t.Execute(w, age > 18)
}

//测试with
func testWith(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("with.html"))
	//执行
	t.Execute(w, "狸猫")
}

//测试template
func testTemplate(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("template1.html", "template2.html"))
	//执行
	t.Execute(w, "我能在两个文件中显示吗？")
}

//测试define
func testDefine(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("define.html"))
	//执行
	t.ExecuteTemplate(w, "model", "")
}

//测试testDefine2
func testDefine2(w http.ResponseWriter, r *http.Request) {
	age := 17
	var t *template.Template
	if age < 18 {
		//解析模板文件
		t = template.Must(template.ParseFiles("define2.html"))
	} else {
		//解析模板文件
		t = template.Must(template.ParseFiles("define2.html", "content1.html"))
	}
	//执行
	t.ExecuteTemplate(w, "model", "")
}

func main() {
	http.HandleFunc("/testIf", testIf)
	http.HandleFunc("/testRange", testRange)
	http.HandleFunc("/testWith", testWith)
	http.HandleFunc("/testTemplate", testTemplate)
	http.HandleFunc("/testDefine", testDefine)
	http.HandleFunc("/testDefine2", testDefine2)

	http.ListenAndServe(":8080", nil)
}
