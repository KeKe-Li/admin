# 从我们自己的Go语言镜像开始，也可选择官方镜像仓库里的Go语言镜像
FROM centos:7

# 添加当前目录（应用根目录）内容，并设置工作目录
ADD centrifugo /usr/bin/centrifugo

RUN echo "centrifugo - nofile 65536" >> /etc/security/limits.d/centrifugo.nofiles.conf

RUN groupadd -r centrifugo && useradd -r -g centrifugo centrifugo

RUN mkdir /centrifugo && chown centrifugo:centrifugo /centrifugo && \
    mkdir /var/log/centrifugo && chown centrifugo:centrifugo /var/log/centrifugo

VOLUME ["/centrifugo", "/var/log/centrifugo"]

WORKDIR /centrifugo

USER centrifugo

EXPOSE 8000
