TODO: 書く
# GolangWantedlyHomeWork

## Wantedly の宿題リポジトリ
宿題として課された課題に取り組んだリポジトリです。

## 動作環境

* OS
   * CentOS7.4
* docker
	* 17.09.0.ce-1.el7.centos


## 実装内容

* WebAPIサーバーを作成する
	* gin + PostgreSQL

* docker-compose で環境が構築されるようにする

### API

| path | method | description |
|:-----------:|:------------:|:------------:|
| / | GET | Hello World!! と json で返す
| /users | GET | user の一覧を表示
| /users/:id | GET | 指定した id の user を表示
| /users | POST | user を追加
| /users/:id | PUT | 指定した id の user を更新
| /users/:id | DELETE | 指定した id の user を削除
