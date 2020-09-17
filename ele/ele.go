package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//a1 := 2 * 21 / 43
	var (
		nameStr        string
		str            string
		extraStr       string
		yh             string
		zj             float64
		zyh            float64
		zExtra         float64
		zjAll          float64   // 总原有订单价格
		zjYhAll        float64   // 最终订单价格(优惠)
		everyPayArr    []float64 // 最终每个人需要
		everyPayStrArr []string  //  最终每个人需要
		jiaoyan        float64   // 校验总和
	)
	fmt.Println("输入名字顺序, -分开：")
	fmt.Scanln(&nameStr)

	fmt.Println("输入每个人金额,-分开")
	fmt.Scanln(&str)

	fmt.Println("输入额外金额, -分开")
	fmt.Scanln(&extraStr)

	fmt.Println("输入优惠, -分开")
	fmt.Scanln(&yh)

	strArr := strings.Split(str, "-")
	yhArr := strings.Split(yh, "-")
	extraArr := strings.Split(extraStr, "-")
	nameStrArr := strings.Split(nameStr, "-")

	// 名字切片
	nameStrSlice := make([]string, len(nameStrArr))
	for index, value := range nameStrArr {
		nameStrSlice[index] = value
	}

	// 费用切片
	var zjFlArr []float64
	for _, value := range strArr {
		fl, _ := strconv.ParseFloat(value, 64)
		zjFlArr = append(zjFlArr, fl)
		zj += fl
	}

	// 额外费用切片
	var extraFlArr []float64
	for _, value := range extraArr {
		fl, _ := strconv.ParseFloat(value, 64)
		extraFlArr = append(extraFlArr, fl)
		zExtra += fl
	}

	// 优惠切片
	var yhFlArr []float64
	for _, value := range yhArr {
		// string 转成float
		yhFl, _ := strconv.ParseFloat(value, 64)
		yhFlArr = append(yhFlArr, yhFl)
		zyh += yhFl
	}
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

	jiaoyan = 0
	for _, value := range everyPayArr {
		jiaoyan += value
	}

	fmt.Println("\t检验总和: " + strings.Join(everyPayStrArr, " + ") + " = " +
		fmt.Sprintf("%.2f", jiaoyan))

	everyPayArr = []float64{}
	everyPayStrArr = []string{}

	// 第二种算法，每个人原来的价格 - 订单最终优惠价格/订单原总金额
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

	jiaoyan = 0
	for _, value := range everyPayArr {
		jiaoyan += value
	}

	fmt.Println("\t检验总和: " + strings.Join(everyPayStrArr, " + ") + " = " +
		fmt.Sprintf("%.2f", jiaoyan))

	fmt.Scanln()

}
