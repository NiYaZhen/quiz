# 使用 Golang 作為基礎映像
FROM golang:latest

# 設定工作目錄
WORKDIR /app

# 複製應用程式程式碼到容器中
COPY . .

# 編譯應用程式
RUN go build -o api

# 配置容器
EXPOSE 8080

# 設定容器執行的命令
CMD ["./api"]
