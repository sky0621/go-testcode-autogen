# go-testcode-autogen

## 概要

##### go のテストコードひな形を自動生成するツール

##### 指定ディレクトリ配下の go ファイルについて、公開functionに対するテストコードを持つテストファイルを自動生成します。

##### 対象の go ファイルに対するテストファイルが既に存在する場合は何もしません。

##### 公開functionがない go ファイルに対してはテストファイルを作成しません。

##### 「vendor」配下は無視するようにしています。

## 前提

##### ・go の実行環境（v1.7）が存在すること

##### ・vendoringツールとして glide がインストール済みであること

## 使い方

##### 1) cd $GOPATH/src/github.com/sky0621/go-testcode-autogen

##### 2) glide up

##### 3) cd $GOPATH/src/github.com/sky0621/go-testcode-autogen/cmd

##### 4) go run main.go -target=../example/sampleproject

###### ※「-target」にテストコードのひな形を作りたいプロジェクトのディレクトリフルパスを指定

## TODO
 
##### 機能実現スピード最優先での実装なので要リファクタ
