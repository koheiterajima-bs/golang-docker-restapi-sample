FROM golang:1.17.1-alpine as builder

# コンテナ内で作業を行うディレクトリを/buildに設定する
WORKDIR /build

# ローカルのディレクトリの内容をコンテナ内のカレントディレクトリにコピーする
COPY ../../ ./

# CGO:C言語とGoを一緒に使う機能だが、Alpineなどの軽量OS環境では通常無効
ARG CGO_ENABLE=0

# GOOSは、GoでそのOS用のバイナリを作成するかを指定するための変数
ARG GOOS=linux

# GOARCHはGoでどのCPUアーキテクチャ用のバイナリを作成するかを指定するための変数
ARG GOARCH=amd64

# -ldflags '-s -w'はバイナリのサイズを小さくするためのオプション
RUN go build -ldflags '-s -w' ./cmd/sample-alpine

# 軽量な実行環境を作成する(ここからがマルチステージビルドの2つ目のステージ)
FROM alpine

# マルチステージビルドの特徴で、前のステージ(builder)でビルドした成果物を現在のステージにコピーする
COPY --from=builder /build/sample-api /opt/app/

# コンテナが起動されたときに実行されるコマンドを指定している
ENTRYPOINT ["/opt/app/sample-api"]

# 2つのステージを使って、まずアプリケーションをビルドし、その後そのアプリケーションだけを含んだ軽量なコンテナを作成して実行する