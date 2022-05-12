# MyTail
- Goのosパッケージのみを使用して`tail`コマンドを実装しました
## 機能
- 引数にファイル名を指定するとファイルの最後の10行を標準出力にプリントする
- `-c`オプションのあとに数値をいれると、ファイルの最後からcバイトを標準出力にプリントする
- 複数のファイルを指定可能
## How to Use
### インストール
```
git clone git@github.com:JimpeiYamamoto/My_tail.git && cd MyTail
```
### ビルド
```
go mod init ztail
```
### 実行
```
./ztail filename
./ztail -c 100 filename
```
## こだわり
- osパッケージのみ使用
- ファイルの読み込み回数を最適化