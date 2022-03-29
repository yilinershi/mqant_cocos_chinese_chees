# 自动编译
* 需要先初始化文件配置，在文件夹目录下执行命令：`tsc --init`,此时，会在目录下生成一个tsconfig.json文件
* 在vscode中选择 `终端-执行任务-选择当前目录下的tsconfig.json文件`

# 问题
* 今天自己遇到一个很SB的问题，自己的vscode没有拖到application里，导致item2里的命令可用，而在vscode里命令不可用，在vscode中一直报错`command not found:xxxx`

# 注意
* 在选择执行任务时，一定要区分`构建`和`监视`的区别，构建是一次性将ts翻译成js,而监视是只要项目中ts内容有变化，就会立马将ts翻译成js
