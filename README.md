# nba-live
`nba-live`是一个在终端下观看`NBA&CBA`文字直播的命令行工具.

所有数据来自[手机版直播吧](https://m.zhibo8.cc/).

## 安装运行
#### Go
```shell
go get -u github.com/xwjdsh/nba-live
nba-live
```
#### Homebrew
```shell
brew tap xwjdsh/tap
brew install xwjdsh/tap/nba-live
nba-live
```
#### Docker
```shell
docker pull wendellsun/nba-live
docker run -it --rm wendellsun/nba-live
```
#### Manual
从[releases](https://github.com/xwjdsh/nba-live/releases)下载可执行文件并将其放到`PATH`环境变量中，然后在终端输入`nba-live`运行。
