# AnimeManager
一个基于go的简单程序，用于将动漫字幕组的命名格式更改为Plex/Emby可以识别的格式

## 用法
### 编译
```shell
git clone https://github.com/WiDayn/AnimeManager
cd AnimeManager
go build AnimeManager/cmd
```
### 使用
config可参照`/configs/default.toml`
```shell
  -b    批量处理Dir
  -c string
        用于匹配的config文件名
  -d string
        匹配的Dir
```