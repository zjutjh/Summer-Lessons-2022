# 前端第二节课环境配置

## 下载安装 NodeJs

> **NodeJs** 使用和浏览器一样的引擎，使得我们可以在浏览器以外的地方运行 JavaScript 代码，本课程使用 nodejs 是为了搭建一个组件化的开发环境。安装了 nodejs 的同时，还会自带一个包管理工具 **npm**

检测是否安装成功：在终端(powreshell / cmd / bash )中任意目录下输入

```bash
node -v
```

正确显示版本号则安装成功

## 搭建开发环境

首先将默认的包管理工具 npm 替换成 yarn / tyarn

```sh
npm i -g yarn --registry=https://registry.npmmirror.com
npm i -g tyarn --registry=https://registry.npmmirror.com
```

> yarn 比 npm 安装依赖更快
>
> tyarn 是一个 yarn 的国内镜像版本，下载依赖包的速度更快

## 使用 npm 创建一个前端项目

在代码文件夹（或者是桌面）打开 终端

```bash
yarn create vite
```

![create](https://s3.bmp.ovh/imgs/2022/08/06/3d860d97c423b169.png)

> Project name 自己随意定一个
>
> framework 和 variant 按照图中选择

> [vite](https://cn.vitejs.dev) 是一个脚手架，负责在本地搭建起服务器，并且能调动各种依赖，将我们的代码(`.vue`, `.scss`, `.ts`)编译成 `html`，`js`，`css` 这些浏览器能运行的文件

## 启动项目

```bash
cd course-preview # 这里输入自定义的 Project name
tyarn # 用 镜像源 安装项目的依赖包
yarn dev # 在本地搭建一个服务器，自动部署项目到本地
# 在浏览器打开生成的 ip 地址
```

这样就可以在浏览器上看到默认的项目了

## 观察项目目录

```tree
.
├── README.md
├── index.html
├── package.json         # 依赖列表
├── public               # 公共资源文件夹，存放一些图片，字体等静态资源
│   └── vite.svg
├── src                  # 源代码文件夹，我们的代码都是在这个文件夹下写
│   ├── App.vue
│   ├── assets
│   │   └── vue.svg
│   ├── components
│   │   └── HelloWorld.vue
│   ├── main.ts
│   ├── style.css
│   └── vite-env.d.ts
├── tsconfig.json         # Typescript 的配置文件
├── tsconfig.node.json    # 同上
└── vite.config.ts        # 脚手架 vite 的配置文件
```
