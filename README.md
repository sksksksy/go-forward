# go-forward，不同局域网通讯
# 该代码主要用于交换信息，部署在公共服务器上
#### 起源：
> 解决两个不同局域网之中的设备通讯问题
> 简要概述：
> 局域网中的设备通常是不能直连外部网络的，这种网络环境下，局域网设备中的ip地址为该局域网中的私有地址，当局域网需要和外部通讯时，会利用局域网中的可对外通讯的网关设备进行转发。通常为Nat映射
> ，nat映射也分好几种类型。所以想要不同局域网中的设备通讯，通过在nat上动手脚也是一种方法。
> 但该程序不是通过nat的形式通讯
> 试想，若两个局域网中的设备都能和同一台公网设备通讯，然后局域网中的设备和公网设备进行一些列沟通，让公网设备转发局域网设备的信息，是否也能实现通讯?
> go-forward代码主要就是基于该方法实现。
> 在go-forward中保存一张巨大的表，不同的sid对应不同设备的网络连接，然后根据一些规定查找对应连接，进行通讯
