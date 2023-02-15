# wechatbot
最近chatGPT异常火爆，想到将其接入到个人微信是件比较有趣的事，所以有了这个项目。项目基于[openwechat](https://github.com/eatmoreapple/openwechat)
开发

### 目前实现了以下功能
 + 群聊@回复
 + 私聊回复
 + 自动通过回复

# 安装使用
````
# 获取项目
git clone git@github.com:JustZyx/wechat-assistant.git

# 进入项目目录
cd wechat-assistant

# 复制配置文件
copy config.dev.json config.json

# 启动项目
go run main.go

启动前需替换config中的api_key
