# 标准输出增强工具

## 介绍

将标准输出重定向到 o，o 将建立一个 http 服务器，将内容传输到任何 http 客户端

## 浏览器增强插件

在 chrome 客户中，你可以安装以下增强插件以配合 o 来使用：

- JSON Viewer
- Markdown Viewer
- Set Character Encoding 

## 示例

```shell
# 默认端口 8010
$ echo '{"name":"dueros", "age":3}' | o

# 在指定端口打开 o
$ echo '{"name":"dueros", "age":3}' | o 8888

# 指定文本类型
$ cat readme.md | o md
$ echo "<h1>hello world!</h1>" | o html
$ echo '{"name":"dueros", "age":3}' | o json

# 同时指定类型和端口
$ cat readme.md | o md 8888
$ cat readme.md | o 8888 md
```
