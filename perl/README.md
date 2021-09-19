# perl版
shellスクリプト版のソフトウェアは，ピン番号を変更する場合や，ファンを制御する
温度を変更する場合はソースを直接編集する必要があります．



## 1. インストール

現状は，systemctl用の定義ファイル等を作っていないので，fanctrlを適当な
ディレクトリにコピーした上で，実行手順を/etc/rc.localに追加
してください．

あと，perlとperlのSwitchモジュールを使っているので，
インストールしていない場合はインストールしてください．

## 2. パラメータ調整

以下はプログラムそのもののとなっています．
```
    #!/usr/bin/perl
    
    use Switch;
    
    open(IN, "/sys/class/thermal/thermal_zone0/temp");
    $cputemp = <IN>/1000.0; #numbers
    close(IN);
    $cputempavg = $cputemp;
    
    system("echo 18 > /sys/class/gpio/export");
    system("echo out > /sys/class/gpio/gpio18/direction");
    system("echo 1 > /sys/class/gpio/gpio18/value");
    $fan_run = 1;
    
    while(1)
    {
    	open(IN, "/sys/class/thermal/thermal_zone0/temp");
    	$cputemp = <IN>/1000.0; #numbers
    	close(IN);
    	
    	$cputempavg = $cputempavg*0.9 + $cputemp*0.1;
    
    	if($cputempavg > 45)
    	{
    		$fan_run = 1;
    	}
    	if($cputempavg < 35)
    	{
    		$fan_run = 0;
    	}
    	if($cputempavg > 65)
    	{
    		$fan_run = 2;
    	}
    	switch ($fan_run) {
    		case 1 { system("echo 1 > /sys/class/gpio/gpio18/value");}
    		case 0 { system("echo 0 > /sys/class/gpio/gpio18/value");}
    		else   { system("shutdown -h now");}
    	}
    	sleep(10);
    }
```


### 2.1. ファン制御用のピン番号
ファンを制御するピンは18番に決め打ちになっているので，ピンを変更する場合は，「fanctrl」の以下の「gpio18」をファンの制御信号のピンに変更してください．
```
    	switch ($fan_run) {
    		case 1 { system("echo 1 > /sys/class/gpio/gpio18/value");}
    		case 0 { system("echo 0 > /sys/class/gpio/gpio18/value");}
    		else   { system("shutdown -h now");}
    	}
```

ピン番号の参考文献は「[ここ][gpio_map]」です．

### 2.2. CPU温度定義

現状の「fanctrl」は45度でファンがまわり，35度でファンが止まります．ファンでは温度が下がらず65度以上になると強制シャットダウンするようになっているので，これらのしきい値は使う環境に合わせて変更してください．
```
    	if($cputempavg > 45)
    	{
    		$fan_run = 1;
    	}
    	if($cputempavg < 35)
    	{
    		$fan_run = 0;
    	}
    	if($cputempavg > 65)
    	{
    		$fan_run = 2;
    	}
```

[gpio_map]: https://www.raspberrypi.org/documentation/computers/os.html#gpio-and-the-40-pin-header "ピン配置"




