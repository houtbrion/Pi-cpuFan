# テスト用ツール
Raspberry PiのCPU温度の測定と，ファン制御用ピンの電圧を強制的にHigh,Low切り替えを行うshellスクリプトで，
ハード作成時の動作確認などに利用するものです．


各ツールの役割は名前からわかると思います．
* getTemperature.sh
* fanOff.sh
* fanOn.sh


## 1. 各プログラムの説明
### getTemperature.sh
カーネルからCPU温度を読み出しています．
```
    #!/bin/sh
    cat /sys/class/thermal/thermal_zone0/temp
```

### fanOff.sh
ピン番号を直接指定しているので，ピン番号を変更する場合はソースを編集してください．
```
    #!/bin/sh
    
    echo 18 >/sys/class/gpio/export
    echo out >/sys/class/gpio/gpio18/direction
    echo 0 > /sys/class/gpio/gpio18/value
```

### fanOn.sh
こちらもピン番号の変更には，編集が必要です．
```
    #!/bin/sh
    
    echo 18 >/sys/class/gpio/export
    echo out >/sys/class/gpio/gpio18/direction
    echo 1 > /sys/class/gpio/gpio18/value
```




