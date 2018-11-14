# faithtop
Desktop GUI framework for Go ( using GTK-2 Lib )

# Windows 32bit

1. Install [msys2](https://www.msys2.org)

  （如果你在中国，请在安装之后给msys2添加中国源 [链接](https://blog.csdn.net/liyuanbhu/article/details/56496501)

2. Run those command in MSYS2 Mingw64 terminal :

`pacman -Syy`

`pacman -S vim git mingw-w64-i686-gtk2 pkg-config mingw-w64-i686-toolchain base-devel glib2-devel`

3. Install Go into C:\msys2\usr\local

4. Add those 2 lines to C:\msys2\etc\profile

```shell
export GOROOT=/usr/local/go
export GOPATH=/home/username/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

5.Install faithtop

（如果你在中国，请先在msys2里面运行以下命令：

```shell
mkdir -p ~/go/src/golang.org/x
cd ~/go/src/golang.org/x/
git clone --depth=1 https://github.com/golang/sys
```
然后再往下看
）

Run `go get github.com/gofaith/faithtop
