package main

import (
	"fmt"
	"go-review/Company"
	"go-review/Idol"
	"go-review/Option"
	"go-review/Test"
	"strings"
)

func main() {
	defer func() {
		fmt.Print("语法基础Done\n")
	}()
	manager := Idol.NewIdolManager()
	idol1 := Idol.NewIdol("略略", 20, "采梦", "凌晨十二点")
	idol2 := Idol.NewIdol("小鱼", 20, "采梦", "凌晨十二点")
	manager.AddIdol(idol1)
	manager.AddIdol(idol2)
	// fmt.Printf("name:%s,age:%v,team:%s", idol1.GetName(), idol1.GetAge(), idol1.GetTeam())
	var checkName string
	fmt.Print("请输入要查看的idol的名字:")
	fmt.Scanln(&checkName)
	checkIdol := manager.GetIdol(checkName)
	fmt.Printf("查到的信息,名字:%s, 年龄:%d, 队伍:%s, 公司:%s\n",
		checkIdol.GetName(), checkIdol.GetAge(), checkIdol.GetTeam(), checkIdol.GetCompany())
	if age := checkIdol.GetAge(); age > 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("未成年")
	}

	// 公司信息
	// company := Company.NewIdolCompany("凌晨十二点")
	company := Company.NewIdolCompany(
		Option.WithCompanyName("凌晨十二点"),
	)
	builder := strings.Builder{}
	builder.WriteString("2019-05-23")
	builder.WriteString(" 18:00:00")
	company.UpdateEstablishDate(builder.String())
	company.AddLocal("上海")
	company.AddTeam("采梦")
	company.AddTeam("梦花火")
	fmt.Printf("公司信息,名字:%s, 成立时间:%s, 本地:%v, 团队:%v\n", company.GetName(), company.GetEstablishDate(), company.GetLocal(), company.GetTeam())
	for index, team := range company.GetTeam() {
		fmt.Printf("%d  团队:%v\n", index, team)
	}
	companyMg := Company.NewCompanyManager()
	companyMg.AddCompany(company)
	compangAll := companyMg.GetAllCompanies()
	teamDone := make([][]string, 0, 1024)
	for _, comp := range compangAll {
		teamDone = append(teamDone, comp.GetTeam())
	}
	fmt.Println(teamDone)
	teamDoneList := make([]string, 0)
	for _, team := range teamDone {
		// ... 展开切片，将多个切片合并为一个切片
		teamDoneList = append(teamDoneList, team...)
	}
	fmt.Println(teamDoneList)
	teamDoneList = teamDoneList[:1]
	fmt.Println(teamDoneList)
	Test.Test_chan()
	Test.Test_string()

	// 泛型
	q_int := Test.NewQueue[int]()
	q_int.Push(2)
	q_int.Push(1)
	fmt.Printf("Pop 队列值为%v\n", q_int.Pop())
	fmt.Printf("现在队列为\n")
	for _, s := range *q_int.ShowQueue() {
		fmt.Printf("value = %v ", s)
	}
}
