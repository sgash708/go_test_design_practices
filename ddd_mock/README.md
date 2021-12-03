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
### 特徴
モデルはどの層にも依存しない。

### 詳細
TableDrivenで、実際に保存されるカラムを「構造体」として記述する。

## 2. domain/repository
### 特徴
<code>Interface</code>を定義する。
当層は、modelに依存する。

### 詳細
<code>Serviceクラス</code>で、使用する関数を<code>interface</code>として定義する。
(重要) 実際の登録処理は、<code>Infra層</code>で行う。

## 3. domain/service
### 特徴
ビジネスロジックを記述する

### 詳細
まず、<code>User</code>構造体を定義する。
<code>domain/repository</code>を受け取り、


# 参考
* [Golangで書くレイヤードアーキテクチャ　簡易APIサーバーを構築](https://qiita.com/hmarf/items/7f4d39c48775c205b99b)
