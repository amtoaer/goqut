<h1 align="center">- 马原毛概刷题工具-GoLang版 -</h1>
<p align="center">
<img src="./img/screenshot.svg" width="800">
</p>
<p align="center">

</p>

## 介绍
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Famtoaer%2Fgoqut.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Famtoaer%2Fgoqut?ref=badge_shield)

本项目为马原、毛概的选择题刷题工具，题目大体按照章节顺序排布。如遇题目缺字、错字等问题，或者有功能上的建议，欢迎提`issue`反馈。

## 说明

比起 python 版本的优点：

1. 题库内置，使用方便
2. 使用编译型语言，生成二进制文件，运行速度快
3. 可双击直接运行，不依赖本机环境
4. 程序重写，源码更加清晰，更易维护
5. 改良存储机制，节省空间
6. 加入额外功能（一键查看答案等）

## 目录
```bash
.
├── build.sh # 构建脚本
├── file
│   ├── data.go # 读取题库文件并解析
│   └── history.go # 读取错题本和顺序刷题记录并解析
├── go.mod
├── go.sum
├── img
│   └── screenshot.svg
├── LICENSE
├── main.go # 程序入口
├── question
│   ├── bind.go # 实体和接口的绑定
│   ├── entity.go # 题库实体和接口定义
│   └── helper.go # 题库的辅助函数
├── README.md
├── statik
│   └── statik.go # 由题库转换而成的代码文件
├── text
│   └── plain.go # 错误提示的文本
└── utils
    └── utils.go # 工具函数

6 directories, 15 files
```

## 功能

+ [x] 顺序刷题
+ [x] 随机刷题
+ [x] 考试模拟
+ [x] 错题本

## 行为

> 顺序刷题的题目进度/错题本会被保存在~/.config/goqut/${subject}.archive内。

运行程序后，程序会根据科目选项初始化题库，如果存档文件不存在会自动写入初始配置。接着选择需要的功能进入刷题界面。

刷题时程序的行为：

1. 每一道题目开始覆盖写入存档。
2. 进入逻辑判断：
   + 非错题本：
     1. 第一次即选到正确选项，跳转到下一道题目。
     2. 多次才选到正确选项，计入错题本并跳转到下一道题目。
     3. 输入`ans`得到题目答案，防止某些多选题卡太久。处理逻辑同2。
     4. 输入`quit`退出程序（与使用“关闭”按钮效果相同）
   + 错题本：
     1. 第一次即选到正确选项，移出错题本，跳转到下一道题目。
     2. 否则只跳转到下一道题目。

## 用法

> 无特殊情况请使用最新版。
> 
> **windows 系统请使用 1.0.8 及之后的版本。**

下载[release](https://github.com/amtoaer/goqut/releases)页面的相应系统版本，双击运行。

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Famtoaer%2Fgoqut.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Famtoaer%2Fgoqut?ref=badge_large)