# Pi-cpuFan

Raspberry Pi用の冷却ファン制御プログラムで，以前作った[簡易版(python)][Pi-CoolingFan]をgoで作り直して，機能強化したもの．

## 1. 動作環境

### 1.1 動作を確認した各種ソフトのバージョン
- OS : Raspberry Pi bullseye
- go言語 : go version go1.17 linux/arm
- periph
  - periph.io/x/conn/v3 : v3.6.8
  - periph.io/x/d2xx    : v0.0.1
  - periph.io/x/host/v3 : v3.7.0

### 1.2 ビルドに必要なその他のツール
- make


## 2. ハードウェア構成
GPIO16番にファンを制御する回路をつけることが前提になっていますが，
ピン番号を変えたい場合は，設定ファイル中で定義している以下の部分を書き換えてください．

```
"fanpin":"GPIO16"
```

以下の回路図やブレッドボードでの配線図も貼っておきますが，[CADソフト(Fritzing)][fritzing]の配線図も同封してあります．

### 2.1 配線
![配線イメージ][breadboard]


### 2.2 デバイスを接続するピン番号の変更

GPIOのピン番号は，設定ファイル(本リポジトリのbuttonAndShutdown.cfg)で変更可能です．

以下は例であるが，ファン制御回路のGPIOの番号は「fanpin」の行で設定しており，GPIO番号16番(ピン番号36番:[参考資料][gpio_map])にしています．
```
{
	"use_syslog":true,
	"use_stdout":false,
	"log_file_name":"",
	"low_threshold":30,
	"high_threshold":40,
	"fanpin":"GPIO16"
}
```



## 3. インストール
### 3.1 準備
1章のリストを見て，go言語とmakeだけはインストールしておいてください．goやmakeをインストールする際に，その他の開発ツールを入れる必要があると思いますが，それについては，go言語の[オフィシャルサイト][golang]の指示に従ってください．

### 3.2 調整
#### 3.2.1 設定ファイルの内容を回路に合わせる

2.2節の「デバイスを接続するピン番号の変更」でも説明したように，回路に合わせて使っているデバイスを接続したGPIOの番号に合わせて同封の「cpufan.cfg」の「fanpin」の行を変更してください．


#### 3.2.2 ログ出力先の選択
再度の掲示になりますが，設定ファイルの中身は以下のようになっています．ここで，「usesyslog」はsyslogを使うか否か，「usestdout」は標準出力にログを垂れ流す場合(回路の動作確認などに利用)，「logfilename」にログファイルのフルパスを指定すると，指定したファイルにログを追加書き込みしていきます．
```
$ cat cpufan.cfg
{
	"use_syslog":true,
	"use_stdout":false,
	"log_file_name":"",
	"low_threshold":30,
	"high_threshold":40,
	"fanpin":"GPIO16"
}
$
```
複数のログ出力先を有効にすることも可能ですし，全部ON(もしくはOFF)にするものOKです．

#### 3.2.3 CPU温度の設定
前節では，設定ファイルでGPIOのピン番号を設定しましたが，冷却用ファンを回す温度のしきい値も設定できます．このプログラムはCPUの温度が上がるとファンを回し，CPUの温度が下がるとファンを止めるため，CPU温度が「high_threshold」を超えるとファンを回し，温度が下がってき「low_threshold」を下回るとファンを止めます．ファンが止まったり動作したりを繰り替えすと音が気になるなどの影響があるため，ファンが回り始めるしきい値「high_threshold」とファンが止まるしきい値「low_threshold」をある程度差をあけてください．もちろん「high_threshold」を大きくします．



#### 3.2.4 インストール先ディレクトリの修正
現状インストール先のディレクトリは/usr/localになっています．
もし，修正したい場合は以下の項目を修正してください．

- Makefile : BASE_DIRの設定行(例:BASE_DIR=/usr/local)
- プログラム本体(cpufan.go) : 設定ファイルのフルパス(例 : const defaultConfigFileName string = "/usr/local/etc/cpufan.cfg")
- systemdの設定(cpufan.service) : バイナリパス名(例 : ExecStart=/usr/local/bin/cpufan)


### 3.3 インストール
以下の2ステップです．
- ```make all```
- ```sudo make install```

rebootで動作しますが，もし，即座に動かしたい場合は，以下のコマンドを実行してください．
```
# systemctl restart cpufan.service
```
うまく動作していれば，systemctlコマンドで「Active: active (running) since 時刻」となるはずです．
```
# systemctl status cpufan.service
● cpufan.service - cpufan
     Loaded: loaded (/etc/systemd/system/cpufan.service; enabled; vendor preset>
     Active: active (running) since Mon 2021-09-06 14:09:13 JST; 47min ago
   Main PID: 10691 (cpufan)
      Tasks: 9 (limit: 2059)
        CPU: 1.706s
     CGroup: /system.slice/cpufan.service
             mq10691 /usr/local/bin/cpufan

Sep 06 14:09:13 raspberrypi systemd[1]: Started cpufan.
#
```




[breadboard]: breadboard.jpg "ブレッドボード図面"
[Pi-CoolingFan]: <https://github.com/houtbrion/Pi-CoolingFan> "Pi-CoolingFan repository"
[golang]: https://go.dev/ "go公式サイト"
