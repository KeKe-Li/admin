# 从我们自己的Go语言镜像开始，也可选择官方镜像仓库里的Go语言镜像
FROM centos:7

# 添加当前目录（应用根目录）内容，并设置工作目录
ADD centrifugo /usr/bin/centrifugo

RUN echo "centrifugo - nofile 65536" >> /etc/security/limits.d/centrifugo.nofiles.conf

RUN groupadd -r centrifugo && useradd -r -g centrifugo centrifugo
# 编译应用，内容详见后面
RUN mkdir /centrifugo && chown centrifugo:centrifugo /centrifugo && \
    mkdir /var/log/centrifugo && chown centrifugo:centrifugo /var/log/centrifugo
# 创建一个存放运行时数据的磁盘卷
VOLUME ["/centrifugo", "/var/log/centrifugo"]

WORKDIR /centrifugo

USER centrifugo
# 对外暴露应用服务监听端口
EXPOSE 8000
