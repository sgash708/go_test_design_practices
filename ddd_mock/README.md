# 目的
* DDDでのプロジェクト作成
* mockを使用して依存関係を気にせずテストすること

# ユースケース
* Userを一人登録する
    * userを登録する場合には、User(model)の構造体を渡す。
    * 返り値は、「Userモデル」「エラー」
* UserGroupを作成する

# 作成手順
## 1. domain/model
TableDrivenで、実際に保存されるカラムを「構造体」として記述する。
## 2. domain
* <code>Serviceクラス</code>で、使用する関数を<code>interface</code>として定義する。
* 実際の登録処理は、<code>Repositoryクラス</code>で行う。
