本项目为基于fabric开发的防医疗骗保平台。
主要功能有：
		1、参保人员管理
        2、药品数据审计
        3、医疗费用结算
	    4、审计结果反馈
 
服务端部署方法：
	在配置好fabric环境拥有docker容器的服务器端，在~/go/src/github.com目录下解压该压缩包，将解压文件改名为ChainblockProject，
    进入文件夹删除main.go文件，将main2.go文件改名为main.go。在当前目录下运行./start.sh。
客户端部署方法：
	解压文件到 (GO安装路径)\src\github.com，将ChainblockProject2.6文件夹放置github.com目录下，把文件名改成“ChainblockProject”
    进入文件夹删除main2.go文件，用goland运行该文件,打开浏览器输入网址:localhost://9002。