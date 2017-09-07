# gin-ueditor

在gin框架下使用UEditor富文本编辑器，实现各种上传功能<br>
主要参考了[jimmykuu/Go-UEditor](https://github.com/jimmykuu/Go-UEditor)和[zxysilent/ueditor](https://github.com/zxysilent/ueditor)这两个项目<br>

**1.下载最新版本的 UEditor**<br>
[下载UEditor](http://ueditor.baidu.com/website/download.html)<br>
选择一个版本下载（选择哪个版本不重要，都需要稍微修改）。这里选择最新的(1.4.3.3) PHP 版本, UTF-8<br>
**2.解压**<br>
解压到statics目录。解压后的目录结构<br>
```
| + conf<
| + controllers
| + models
| - static
    | - ueditor
  	   | + dialogs
       | + lang
       | - php
       	   config.json
       | + themes
       | + third-party
       index.html
       ueditor.all.js
       ueditor.all.min.js
       ueditor.config.job.js
       ueditor.config.js
       ueditor.parse.js
       ueditor.parse.min.js
| + routes
| + upload
| + views
```
删除static/ueditor/php目录下所有扩展名为 .php的文件，并且将static/ueditor/php改名成static/ueditor/conf<br>
**3.配置ueditor**<br>
修改ueditor.config.js,第 33行<br>
``, serverUrl: "php/controller.php"`` <br>
改成<br>
``, serverUrl: "/ueditor/controller"`` <br>
//这个url必须和我们在routes.go里配置的 ``router.Any("/ueditor/controller", controllers.Action)`` 对应<br>

其它的看代码里的注释吧
