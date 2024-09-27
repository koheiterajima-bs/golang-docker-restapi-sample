# どのDockerイメージを基にコンテナを作るかの指定
FROM --platform=linux/amd64 library/mysql:8.0.25

# コンテナ内で使用する環境変数を設定(todoという名前のデータベースを自動で作成)
ENV MYSQL_DATABASE todo

# ホストマシンにあるファイルをコンテナの中にコピーする命令(custom.cnfはホスト側にてMySQLの設定をカスタマイズするための設定ファイル)
# /etc/mysql/conf.dは、MySQLの設定ファイルが入るディレクトリで、ここにコピーすることでMySQLの設定を変更できる
COPY custom.cnf /etc/mysql/conf.d/

# ホストマシンのSQLファイルやフォルダを指している
# /docker-entrypoint-initdb.d/はMySQLコンテナが起動する際に自動的に実行されるスクリプトを置くディレクトリ、ここにSQLファイルを置くと、コンテナ起動時にこれらのファイルが実行され、データベースが初期化される
COPY sql /docker-entrypoint-init.db.d