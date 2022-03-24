# Winrar Password Discovery  WinRAR密码探索器

#### 介绍
WinRAR密码探索器，通过字典探索你丢失的rar密码，支持RAR5,RAR4加密后文件

本程序主要解决您在忘记自己压缩包密码情况下,通过字典方式进行winrar解压

注意,本软件基于MIT开源,仅供学习交流,切勿用于非法用途!
#### 软件架构
Golang 


#### 使用说明
generate-password.go (generate-password.exe)  密码字典生成工具,打开后输入密码生成规则,生成1-∞  位长度密码,可根据模式选择不同规则密码。

![生成密码](https://gitee.com/aliu/discoveryRAR/raw/develop/jpg/generate-password.png "生成密码")


discoveryRAR.go （discoveryRAR.exe） RAR探索器根据生成的password.txt字典进行查找探索。

![探索RAR](https://gitee.com/aliu/discoveryRAR/raw/develop/jpg/discovery.png "探索RAR")

