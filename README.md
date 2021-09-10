# GOGOOJ 
&nbsp;&nbsp;&nbsp;&nbsp;个人娱乐项目 </br>
&nbsp;&nbsp;&nbsp;&nbsp;一个玩具在线代码测评系统，类似leetcode </br>
&nbsp;&nbsp;&nbsp;&nbsp;分为前端，后端和判题端 </br>
&nbsp;&nbsp;&nbsp;&nbsp;前端为Vue+ElementUI+echarts,UI设计参考(抄)[QingdaoU/OnlineJudge](https://github.com/QingdaoU/OnlineJudge) </br>
&nbsp;&nbsp;&nbsp;&nbsp;后端为Go+iris+gosql </br>
&nbsp;&nbsp;&nbsp;&nbsp;判题端用C配合ulimit和seccomp </br>
&nbsp;&nbsp;&nbsp;&nbsp;代码全手敲,净代码超过3W行，历时8个月 </br>

### 一些无关紧要的特性
* 三端分离，支持多判题端
* 简单，Go语言特性少，不管多久过去回来看还是那么直白(指搬砖代码)
* 快！在1C2G1M带宽的服务器上进行业务实测，带宽最先到达瓶颈，CPU反而不到10%
* 支持ACM/OI规则，支持SPJ
                                                   
### 架构
    GOGOOJ
    |--Frontend       --Vue
    |--WebServer      --Go
    |--JudgeServer    --Go and C
    |--utils          --some useful utils, such as session, log
    |--config         --the config of Webserver
    
### 预览
&nbsp;用户端[点击跳转](http://49.234.91.99/) </br>
&nbsp;后台
![img1](https://i.bmp.ovh/imgs/2021/09/50b3193cc12db2b4.png)
![img2](https://i.bmp.ovh/imgs/2021/09/566a065a01d72df5.png)
![img3](https://i.bmp.ovh/imgs/2021/09/11f4367493fa2fb3.png)
![img3](https://i.bmp.ovh/imgs/2021/09/bd1973c0a5e27dbd.png)