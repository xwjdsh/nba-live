```text
          __                ___          
   ____  / /_  ____ _      / (_)   _____ 
  / __ \/ __ \/ __ `/_____/ / / | / / _ \
 / / / / /_/ / /_/ /_____/ / /| |/ /  __/
/_/ /_/_.___/\__,_/     /_/_/ |___/\___/ 

```
[![Release](http://github-release-version.herokuapp.com/github/xwjdsh/nba-live/release.svg?style=flat)](https://github.com/xwjdsh/nba-live/releases/latest)
[![Build Status](https://travis-ci.org/xwjdsh/nba-live.svg?branch=master)](https://travis-ci.org/xwjdsh/nba-live)
[![Go Report Card](https://goreportcard.com/badge/github.com/xwjdsh/nba-live)](https://goreportcard.com/report/github.com/xwjdsh/nba-live)
[![](https://images.microbadger.com/badges/image/wendellsun/nba-live.svg)](https://microbadger.com/images/wendellsun/nba-live)
[![DUB](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/xwjdsh/manssh/blob/master/LICENSE)

`nba-live`是一个在终端下观看`NBA&CBA`文字直播的命令行工具.

所有数据来自[手机版直播吧](https://m.zhibo8.cc/).

## 截图
![](https://raw.githubusercontent.com/xwjdsh/nba-live/master/screenshot/nba-live-select.png)
![](https://raw.githubusercontent.com/xwjdsh/nba-live/master/screenshot/nba-live-cui.png)

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
从[releases](https://github.com/xwjdsh/nba-live/releases)下载可执行文件并将其放到`PATH`环境变量对应的路径中，然后在终端输入`nba-live`运行。

## 致谢
* [manifoldco/promptui](https://github.com/manifoldco/promptui)
* [marcusolsson/tui-go](https://github.com/marcusolsson/tui-go)
* [基于Python命令行的NBA文字直播小工具](http://www.jianshu.com/p/b4077b8810bd)

## 协议
[MIT License](https://github.com/xwjdsh/nba-live/blob/master/LICENSE)
