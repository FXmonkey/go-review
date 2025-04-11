package main

import (
	"fmt"
	"go-review/Company"
	"go-review/Idol"
	"strings"
)

func main() {
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
	company := Company.NewIdolCompany("凌晨十二点")
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

}
