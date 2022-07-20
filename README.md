# 远控平台

## 需求分析

### 目标操作系统

大部分为Windows操作系统。~~但是根据经验来看，不能排除其他操作系统可能。~~

### 信息收集程序

开发一个可执行程序（exe），这个exe安装到目标主机上。

这个程序的主要功能是收集信息。

**收集信息表**

|     功能     | 目前支持的系统  |     备注     |
| :----------: | :-------------: | :----------: |
| 系统基本信息 | Mac OS、Windows | 主要硬件信息 |
|     截图     | Mac OS、Windows | Mac需要权限  |
|              |                 |              |

### 测试服务端

一个简单的服务端，用于模拟测试显示客户端传送过来的数据。

**另外一个重要功能就是发送命令给客户端。**



## 协议约定

|       功能       |       客户端发送        | 服务器回复 |
| :--------------: | :---------------------: | :--------: |
| 检测服务器存活性 |          HELLO          |     OK     |
|     发送截图     | screen_文件名\_文件大小 |     -      |
|                  |                         |            |

