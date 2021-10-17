FROM golang:1.16-stretch
WORKDIR /opt

# 构建完清除代码，直接使用二进制文件，保护源码
COPY . /opt/code
RUN cd /opt/code; go build main.go; cp main /usr/local/bin/httpserver; cp entrypoint.sh /usr/local/bin/; \
    chmod +x /usr/local/bin/httpserver; chmod +x /usr/local/bin/entrypoint.sh; \
    rm -rf /opt/code

# 不要用root用户
RUN useradd user
USER user

# entrypoint.sh为扩展脚本, 方便启动的初始化配置
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD httpserver