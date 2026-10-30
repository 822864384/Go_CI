# ---------- 构建阶段 ----------
FROM golang:1.26-alpine AS builder

WORKDIR /src

# 静态编译，去掉 CGO 依赖，才能在精简基础镜像里跑
ENV CGO_ENABLED=0 \
    GOOS=linux

# 先只复制依赖清单，让 go mod download 单独占一层，充分利用构建缓存
COPY go.mod ./
# 本 demo 无第三方依赖，所以没有 go.sum。
# 日后引入依赖时改成下面两行：
# COPY go.mod go.sum ./
# RUN go mod download

COPY . .

RUN go build -trimpath -ldflags="-s -w" -o /denom .

# ---------- 运行阶段 ----------
FROM alpine:3

RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -H -u 10001 appuser

COPY --from=builder /denom /usr/local/bin/denom

# 用非 root 用户运行，安全基线
USER appuser

ENTRYPOINT ["/usr/local/bin/denom"]
CMD ["-amount", "386"]