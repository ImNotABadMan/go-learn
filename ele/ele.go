package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//a1 := 2 * 21 / 43
	var (
		nameStr  string
		str      string
		extraStr string
		yh       string
		zj       float64
		zyh      float64
		zExtra   float64
		zjAll    float64 // 总原有订单价格
		zjYhAll  float64 // 最终订单价格(优惠)
	)

	nameStrSlice := getStdStrSlice("输入名字顺序, -分开：", &nameStr, "-")
	strArr := getStdStrSlice("输入每个人金额,-分开", &str, "-")
	extraArr := getStdStrSlice("输入额外金额, -分开", &extraStr, "-")
	yhArr := getStdStrSlice("输入额外金额, -分开", &yh, "-")

	// 费用切片
	zjFlArr, zj := stringSliceToFloatSlice(strArr)

	// 额外费用切片
	_, zExtra = stringSliceToFloatSlice(extraArr)

	// 优惠切片
	_, zyh = stringSliceToFloatSlice(yhArr)

	// 总原有订单价格
	zjAll = zj + zExtra
	// 最终订单价格
	zjYhAll = zjAll - zyh

	fmt.Println("结果：")
	// 总金额
	fmt.Println("总金额：")
	fmt.Println("\t", str)
	zjStr := strconv.FormatFloat(zj, 'f', 2, 64)
	fmt.Println("\t", strings.Join(strArr, " + ")+" = "+zjStr)

	fmt.Println("额外费用：")
	fmt.Println("\t", extraStr)
	zExtraStr := strconv.FormatFloat(zExtra, 'f', 2, 64)
	fmt.Println("\t", strings.Join(extraArr, " + ")+" = "+zExtraStr)

	// 最终订单金额
	fmt.Println("最终订单原有金额: " + fmt.Sprintf("%.2f", zjAll))

	// 总优惠
	fmt.Println("总优惠：")
	fmt.Println("\t", yh)
	zyhStr := strconv.FormatFloat(zyh, 'f', 2, 64)
	fmt.Println("\t", strings.Join(yhArr, " + ")+" = "+zyhStr)

	// 最终优惠订单金额
	fmt.Println("最终优惠订单金额: " + fmt.Sprintf("%.2f", zjYhAll))

	// 第一种算法，最终优惠价格/原总金额 * 每个人原来价格
	sf1(zjFlArr, zjYhAll, zj, nameStrSlice)

	// 第二种算法，每个人原来的价格 - 订单最终优惠价格/订单原总金额
	sf2(zjFlArr, zyh, zExtra, zj, nameStrSlice)

	fmt.Scanln()

}

func getStdStrSlice(tip string, str *string, sep string) []string {
	fmt.Println(tip)
	fmt.Scanln(str)

	if sep == "" {
		sep = "-"
	}

	return strings.Split(*str, sep)
}

func stringSliceToFloatSlice(strArr []string) (flSlice []float64, sum float64) {
	for _, value := range strArr {
		fl, _ := strconv.ParseFloat(value, 64)
		flSlice = append(flSlice, fl)
		sum += fl
	}

	return flSlice, sum
}

func sf1(zjFlArr []float64, zjYhAll float64, zj float64, nameStrSlice []string) {
	var (
		everyPayArr    []float64 // 最终每个人需要
		everyPayStrArr []string  //  最终每个人需要
		jiaoyan        float64   // 校验总和
	)
	fmt.Println("第一种算法，订单最终优惠价格/订单原总金额(不包含额外费用) * 每个人原来价格：")
	for index, value := range zjFlArr {
		//每个人分担包装费用
		bzfFl := float64(0)
		// 每个人优惠费用
		tmp := value * zjYhAll / zj
		// 最终费用
		tmpFinally := tmp + bzfFl

		everyPayArr = append(everyPayArr, tmpFinally)
		everyPayStrArr = append(everyPayStrArr, fmt.Sprintf("%.2f", tmpFinally))

		fmt.Println("\t", nameStrSlice[index])
		fmt.Println("\t    分担额外费用：", fmt.Sprintf("%.2f", bzfFl))
		fmt.Println("\t    使用优惠费用：", value, " * ", zjYhAll, " / ", zj,
			" = ", fmt.Sprintf("%.2f", tmp))
		fmt.Println("\t    最终需要费用：", fmt.Sprintf("%.2f", tmp), " + ",
			fmt.Sprintf("%.2f", bzfFl), " = ", fmt.Sprintf("%.2f", tmpFinally))
	}

	for _, value := range everyPayArr {
		jiaoyan += value
	}

	fmt.Println("\t检验总和: " + strings.Join(everyPayStrArr, " + ") + " = " +
		fmt.Sprintf("%.2f", jiaoyan))
}

func sf2(zjFlArr []float64, zyh float64, zExtra float64, zj float64, nameStrSlice []string) {
	var (
		everyPayArr    []float64 // 最终每个人需要
		everyPayStrArr []string  //  最终每个人需要
		jiaoyan        float64   // 校验总和
	)
	fmt.Println("第二种算法，每个人原来的价格 - 每个人原来的价格/订单原总金额(不包含额外费用) * 优惠总额 " +
		"+ 分担额外费用: ")
	for index, value := range zjFlArr {
		// 每个人分担包装费用
		bzfFl := value / zj * zExtra
		// 每个人优惠费用
		tmp := value - value/zj*zyh
		// 最终费用
		tmpFinally := tmp + bzfFl
		everyPayArr = append(everyPayArr, tmpFinally)
		everyPayStrArr = append(everyPayStrArr, fmt.Sprintf("%.2f", tmpFinally))
		fmt.Println("\t", nameStrSlice[index])
		fmt.Println("\t    分担额外费用：", value, "/", zj, "*", zExtra,
			" = ", fmt.Sprintf("%.2f", bzfFl))

		fmt.Println("\t    使用优惠费用：", value, " - ", value, " / ", zj, " * ", zyh,
			" = ", fmt.Sprintf("%.2f", tmp))

		fmt.Println("\t    最终需要费用：", fmt.Sprintf("%.2f", tmp), " + ",
			fmt.Sprintf("%.2f", bzfFl), " = ", fmt.Sprintf("%.2f", tmpFinally))
	}

	for _, value := range everyPayArr {
		jiaoyan += value
	}

	fmt.Println("\t检验总和: " + strings.Join(everyPayStrArr, " + ") + " = " +
		fmt.Sprintf("%.2f", jiaoyan))
}
