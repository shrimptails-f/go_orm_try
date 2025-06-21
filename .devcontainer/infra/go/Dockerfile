FROM golang:1.24.2-bookworm

# Task Runner のインストール
RUN curl -sL "https://github.com/go-task/task/releases/download/v3.9.0/task_linux_amd64.deb" -o task.deb && dpkg -i task.deb && rm task.deb

ENV TZ=Asia/Tokyo

# 環境変数を設定して、日本語UTF-8ロケールを使用
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
# 必要なパッケージのインストールとロケールの設定
# commitのコメントで日本語を扱いたいので追加。
RUN apt-get update && \
    apt-get install -y locales && \
    localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8

# ブランチ名がbashで表示されるようにカスタマイズ
COPY ../.bashrc /root/.bashrc

ENV APP_ROOT /data
WORKDIR $APP_ROOT

RUN apt-get update && \
    apt-get install -y vim && \
    rm -rf /var/lib/apt/lists/*

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/rinchsan/gosimports/cmd/gosimports@latest


RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# migrate バイナリのインストール
ENV MIGRATE_VERSION=4.16.2

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz \
    -o migrate.linux-amd64.tar.gz && \
    tar -xvzf migrate.linux-amd64.tar.gz && \
    mv migrate /usr/local/bin/ && \
    chmod +x /usr/local/bin/migrate && \
    rm migrate.linux-amd64.tar.gz