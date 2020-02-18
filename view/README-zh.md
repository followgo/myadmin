# myadmin-view

这是 myadmin 后台管理系统的前端，它基于 `vue-admin-template`。

## vue-admin-template 项目

- [线上地址](http://panjiachen.github.io/vue-admin-template)
- [国内访问](https://panjiachen.gitee.io/vue-admin-template)
- [使用文档](https://panjiachen.github.io/vue-element-admin-site/zh/)

## 启动项目（Win10）

提示

- 系统必须已经先安装 node v12+
- `--global` 参数要使用管理员运行的 powershell 中运行
- 建议不要直接使用 cnpm 安装，会有各种诡异的 bug。可以通过 `--registry` 参数解决下载速度慢的问题

1. 使用管理员运行 powershell，然后安装 node-sass 的构建工具

```bash
# 安装 node-gyp
npm install -g node-gyp --registry=https://registry.npm.taobao.org

# 安装 windows-build-tools
npm install -g --production windows-build-tools --registry=https://registry.npm.taobao.org
```

>此步骤会在系统中自动安装 Visual Studio Build Tools 和 Python 2.7，这是构建 node-sass 必要的工具。

2. 在项目目录下执行

```bash
# 安装依赖
npm install --registry=https://registry.npm.taobao.org

# 启动服务
npm run dev
```

浏览器访问 [http://localhost:9528](http://localhost:9528)

## 构建生产环境

```bash
npm run build:prod
```

## 其它

```bash
# 预览发布环境效果
npm run preview

# 预览发布环境效果 + 静态资源分析
npm run preview -- --report

# 代码格式检查
npm run lint

# 代码格式检查并自动修复
npm run lint -- --fix

# 优化 svg 图片文件
npm run svgo
```

## 浏览器支持

Modern browsers and Internet Explorer 10+.

| IE / Edge | Firefox | Chrome | Safari |
| --------------- | -------------- | -------------- | -------------- |
| IE10, IE11, Edge| last 2 versions| last 2 versions| last 2 versions

## License

[MIT](https://github.com/PanJiaChen/vue-element-admin/blob/master/LICENSE)

Copyright (c) 2017-present PanJiaChen