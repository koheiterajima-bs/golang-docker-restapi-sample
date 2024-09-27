/*
MySQLで新しいユーザーの作成
@'%'は、任意のIPアドレスから接続可能を意味する
IDENTIFIED BY 'todo-password'はそのユーザーのパスワードを設定する部分
*/
CREATE USER 'todo-app'@'%' IDENTIFIED BY 'todo-password';

/*
GRANTは、ユーザーに権限を付与するためのSQL文
ON todo.*は、どのデータベースとテーブルに対して権限を与えるかを指定する
TO 'todo-app'@'%'は、どのユーザーに権限を付与するかを指定する
*/
GRANT SELECT, INSERT, UPDATE, DELETE ON todo.* TO 'todo-app'@'%';