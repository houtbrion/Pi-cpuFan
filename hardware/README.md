# ハードウェア

そのディレクトリには，Raspberry Pi用のユニバーサル基板を利用してファン制御回路のシールドを作成する
場合のfritzingの回路図をだけでなく，「コンソールなど用Piシールド」では，コンソール用USBシリアル，
I2C接続のRTC端子を追加．このシールドで利用したRTCモジュールは
Spark FunのBOB-12708です．あと，秋月電子通商でも入手が容易な8564NBの場合の回路図も
付けてあります．

# 1. ここで紹介する試作ハードについて
ファン制御，RTCなどの回路を使って，Raspberry Piによるシリアルをコンソールとして利用する小型のサーバ(ディスプレイ，キーボードなど無し)
を作成し，ケースに収める試作の実例を紹介します．

ただし，コンソール用の端子は通常使われているRS-232などではなく，USBシリアルコンバータを
Raspberry PiのHATとして実装しているため，メンテナンスでネットワークが利用できない場合は
USB接続でログインすることができます．

# 2. コンソールやRTCの設定
## 2.1 Raspberry Pi3のシリアルコンソールの問題 (4は未検証)
UARTの速度がCPUの動作周波数と連動してしまうため，シリアルコンソールでは
文字化け，入力不可能な状態が続く．

[参考URL](https://github.com/RPi-Distro/repo/issues/22)

この試作では，ディスプレイやUSBキーボードを繋がず，
シリアルコンソール使うため，Raspberry Pi OSを
マイクロSDに書き込んだ際に，ブート用コンフィグファイル(/boot/config.txt)に以下の行を追加しておく．

### Pi3の場合
```
enable_uart=1
core_freq=250
```

### Pi2, Pi zeroやzeroWの場合
```
enable_uart=1
```

## 2.2 RTC
Raspberry PiはRTCが搭載されていないため，NTPなどで時刻を同期しないと人手で時刻をメンテナンス必要が
出てしまい，非常に面倒である．
この面倒を避けるため，今回の試作では，シリアルコンソール用のシールドにRTCも搭載している．
これを利用するため，以下のURLを参照し，パッケージのインスールや設定を行う．

なお，今回試作したHATは8564NBを搭載しているため，まったく同じ回路を利用する場合は後者の方の
対処をお願いします．
### Spark Fun BOB-12708
- [DS1307関係参考URL1](<https://mynotebook.h2np.net/post/1085> "Raspberry Pi 2 にRTCモジュールをつける") Raspberry Pi 2 にRTCモジュールをつける
- [DS1307関係参考URL2](<https://learn.adafruit.com/adding-a-real-time-clock-to-raspberry-pi?view=all> "Adding a Real Time Clock to Raspberry Pi") Adding a Real Time Clock to Raspberry Pi

### (秋月)8564NB
- [8564NB関係参考URL1](<http://news.mynavi.jp/articles/2014/08/21/raspberry-pi4/002.html> "超小型PC「Raspberry Pi」で夏休み自由課題・第4回 - Raspberry Piの屋外モバイル、高層エレベーターの気圧変化を調べる") 超小型PC「Raspberry Pi」で夏休み自由課題・第4回 - Raspberry Piの屋外モバイル、高層エレベーターの気圧変化を調べる
- [8564NB関係参考URL2](<http://tomosoft.jp/design/?p=5812> "Raspberry PIへリアルタイムクロックモジュールのI2C接続") Raspberry PIへリアルタイムクロックモジュールのI2C接続

# 3. 用意した部品
## 3.1 共通
### 3.1.1 本体など
* 本体 : [Raspberry Pi 2 Model B][pi2]
* ヒートシンク : [Seeed Studio 800035001 Raspberry Pi用アルミニウム・ヒートシンク][sink]
* USBシリアル変換 : [スイッチサイエンスSSCI-010320 FTDI USBシリアル変換アダプター(5V/3.3V切り替え機能付き)][ftdi]

### 3.1.2 ファンの制御回路用
* 抵抗 : [カーボン抵抗(炭素皮膜抵抗) 1/4W 3kΩ (100本)][3kohm]
* ダイオード : [ショットキーダイオード][diode]
* トランジスタ : [トランジスタ (2SC1815GR 60V 150mA) 20個][2SC1815]

ただし，今回旧型Pi用のヒートシンクを利用したのは，旧型用は小型のヒートシンクが余分についている代わりに，熱伝導両面テープが付いているので，特別な材料とかなしに，チップに貼り付けることができるためです．

## 3.2 千石電商の箱をそのまま利用
* [Seeed Studio 114990129/141107004【ファン付き】Raspberry Pi B+専用ケース(クリア/ロゴなし/組立式)][case1]
* [Seeed Studio 114990130 Raspberry Pi B+専用ケース(クリア/ロゴなし/組立式)][case2]

GPIOの横の部品にスリットがあり，そこが折れたため，スリットがないファンなしの箱を追加で買って，部品交換しましまた．

![箱の壊れた部分][brokenCase]

![千石のファン付き箱に収めた写真][system2]

## 3.3 改造版
シリアルを簡単に繋げられるようにしたかったので，GPIO端子周りに余裕がある箱を改造してファンを付けた．
* [MultiComp MC-RP002-CLR Raspberry Pi B+専用ケース(クリア/ロゴあり)][case3]
* [SHICOH F4006AP-05QCW (0406-5) DCファン(5V)][dcFan]
* [カモンSRS-40 ファン防振シリコンシート][si-sheet]
* [40mm角ファンガード][FanGuard]

![自作箱に収めた写真][system1]

![自作箱を開けた写真][openCase]

# 4. 回路図など
![ブレッドボード利用時の配線イメージ][breadboard]

![回路図][circuit]

# 5. ハードテスト用プログラムの準備
本番運用の前に，試作したハードがきちんと動作するか確認する必要があるため，
[toolsディレクトリ](../tools)内の動作確認用プログラムを利用します．


- getTemperature.sh : CPU温度を取得するプログラム
- fanOff.sh : ファン制御用回路につながるGPIOピンをLowにするプログラム
- fanOn.sh : ファン制御用回路につながるGPIOピンをHighにするプログラム

ただし，fanOff.shとfanOn.shはピン番号を試作したハードに合わせて修正する
必要があるため，以下の参考にしてください．

### fanOff.sh
下の「18」と「gpio18」を自分が用意した回路に合わせて修正．
```
    #!/bin/sh
    
    echo 18 >/sys/class/gpio/export
    echo out >/sys/class/gpio/gpio18/direction
    echo 0 > /sys/class/gpio/gpio18/value
```

### fanOn.sh
こちらも同じくGPIOの番号を修正．
```
    #!/bin/sh
    
    echo 18 >/sys/class/gpio/export
    echo out >/sys/class/gpio/gpio18/direction
    echo 1 > /sys/class/gpio/gpio18/value
```

# 6. 動作確認
上のgetTemperature.shでCPU温度が取得できることと，fanOn.sh, fanOff.shで
ファンが回ることを確認してください．


# 7. 製作記事
なお，RTC無しで小さめの箱に収めた試作は以下の記事で紹介していますので
そちらを参照してください．

## 市販の箱に収めた場合
- [市販の箱をそのまま使った場合の制作記事][sengoku]
- [試作したハード][sengoku-pic]

## 自分で箱に穴を開けて，ファンを付けた場合
- [改造箱版の制作記事][original]
- [試作したハード][original-pic]とその[内部][inside]


<!--以下はリンクの定義-->
<!--参考文献-->
[sengoku]: <http://hautbrion.blogspot.jp/2015/05/raspberry-pi-2usb-1.html> "千石の箱版"
[original]: <http://hautbrion.blogspot.jp/2014/10/raspberry-pi-2.html> "改造箱版"

<!--ハード関連-->

[pi2]: <http://akizukidenshi.com/catalog/g/gM-09024/> "Raspberry Pi 2 Model B"
[case1]: <http://www.sengoku.co.jp/mod/sgk_cart/detail.php?code=EEHD-4N4C> "Seeed Studio 114990129/141107004【ファン付き】Raspberry Pi B+専用ケース(クリア/ロゴなし/組立式)"
[case2]: <http://www.sengoku.co.jp/mod/sgk_cart/detail.php?code=EEHD-4R3F> "Seeed Studio 114990130 Raspberry Pi B+専用ケース(クリア/ロゴなし/組立式)"
[sink]: <https://www.sengoku.co.jp/mod/sgk_cart/detail.php?code=EEHD-4FBE> "Seeed Studio 800035001 Raspberry Pi用アルミニウム・ヒートシンク"
[ftdi]: <http://www.sengoku.co.jp/mod/sgk_cart/detail.php?code=EEHD-0SK8> "スイッチサイエンスSSCI-010320 FTDI USBシリアル変換アダプター(5V/3.3V切り替え機能付き)"
[3kohm]: <http://akizukidenshi.com/catalog/g/gR-25302/> "カーボン抵抗(炭素皮膜抵抗) 1/4W 3kΩ (100本)"
[diode]: <http://akizukidenshi.com/catalog/g/gI-00881/> "ショットキーダイオード"
[2SC1815]: <http://akizukidenshi.com/catalog/g/gI-00881/> "トランジスタ (2SC1815GR 60V 150mA) 20個"
[dcFan]: <https://www.sengoku.co.jp/mod/sgk_cart/detail.php?code=557S-43JV> "SHICOH F4006AP-05QCW (0406-5) DCファン(5V)"
[si-sheet]: <https://www.sengoku.co.jp/mod/sgk_cart/detail.php?code=4AZ6-GFHU> "カモンSRS-40 ファン防振シリコンシート"
[FanGuard]: <https://www.sengoku.co.jp/mod/sgk_cart/detail.php?code=328B-2CE8> "40mm角ファンガード"
[case3]: <http://www.sengoku.co.jp/mod/sgk_cart/detail.php?code=EEHD-4KRC> "MultiComp MC-RP002-CLR Raspberry Pi B+専用ケース(クリア/ロゴあり)"


<!--イメージファイル-->
[system1]: system1.jpg "自作箱に収めた写真"
[system2]: system2.jpg "千石のファン付き箱に収めた写真"
[breadboard]: breadboard.jpg "ブレッドボード利用時の配線イメージ"
[circuit]: circuit.jpg "回路図"
[openCase]: openCase.jpg "自作箱を開けた写真"
[brokenCase]: brokenCase.jpg "箱の壊れた部分"
[sengoku-pic]: system2.jpg "千石でかった箱"
[priginal-pic]: system1.jpg "穴を開けた箱"
[inside]: openCase.jpg "内部"



