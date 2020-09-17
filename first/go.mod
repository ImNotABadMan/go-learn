module first

go 1.15

// 用于构建程序

// 依赖包的路劲声明
// 包名 => 目录
replace first-moudle => ../first-moudle

// 依赖 包 版本
require first-moudle v1.0.0
