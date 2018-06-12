# go-testcode-autogen

## 概要

##### go のテストコードひな形を自動生成するツール

##### 指定ディレクトリ配下の go ファイルについて、公開functionに対するテストコードを持つテストファイルを自動生成します。

##### 対象の go ファイルに対するテストファイルが既に存在する場合は何もしません。

##### 公開functionがない go ファイルに対してはテストファイルを作成しません。

##### 「vendor」配下は無視するようにしています。

## 使い方

##### 1) 以下より、対象モジュールをダウンロード

https://github.com/sky0621/go-testcode-autogen/releases

##### 2) sh go_testcode_autogen_linux_amd64 --target golang/src/github.com/sky0621/go-testcode-autogen/_example/sampleproject/
 -target=../example/sampleproject

###### ※「-target」にテストコードのひな形を作りたいプロジェクトのディレクトリフルパスを指定

## 留意点

##### ・同一パッケージに同一関数名（所持する構造体が異なる）がある場合、テストコードも同一関数名が生成され、コンパイルエラーが発生

###### ※この場合は自動生成後に手動でテストコードのテスト関数名修正が必要

## TODO
 
##### 機能実現スピード最優先での実装なので要リファクタ
