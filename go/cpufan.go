package main

import (
    "fmt"
    "log"
    "log/syslog"
    "time"
    "os"
    "encoding/json"
    "io/ioutil"
    "strconv"
    "strings"

    "periph.io/x/conn/v3/gpio"
    "periph.io/x/conn/v3/gpio/gpioreg"
    "periph.io/x/host/v3"
)

const defaultConfigFileName string = "/usr/local/etc/cpufan.cfg"
//const thresholdDiff int = 5
const cpuTemperatureFile string = "/sys/class/thermal/thermal_zone0/temp"

type Config struct {
    UseSyslog     bool   `json:"use_syslog"`
    UseStdout     bool   `json:"use_stdout"`
    LogFileName   string `json:"log_file_name"`
    LowThreshold  int    `json:"low_threshold"`
    HighThreshold int    `json:"high_threshold"`
    FanPin        string `json:"fanpin"`
}

type Logger struct {
    syslogFlag bool
    sysLog *syslog.Writer
    stdoutFlag bool
    logFile *os.File
}

func (out *Logger) Init(sFlag bool, s *syslog.Writer, stdFlag bool, fp *os.File) {
    out.syslogFlag = sFlag
    out.sysLog = s
    out.stdoutFlag = stdFlag
    out.logFile = fp
}

func (logger *Logger) Log(msg string) {
    if logger.syslogFlag {
        logger.sysLog.Err(msg)
    }
    if logger.stdoutFlag {
        log.Println(msg)
    }
    if logger.logFile != nil {
        logger.logFile.WriteString(time.Now().String()+" : "+ msg+ "\n")
    }
}

type Fan struct {
    state bool
    onThresh int
    offThresh int
}

func (fan *Fan) Init(on int, off int){
    fan.state = false
    fan.onThresh = on
    fan.offThresh = off
}

func (fan *Fan) setState(s bool) {
    fan.state = s
}

func getCpuTemp() int{
    data, _ := ioutil.ReadFile( cpuTemperatureFile )
    var tempStr string = strings.TrimRight(string(data),"\n")
    var temperature int
    temperature, err := strconv.Atoi(tempStr)
    if err != nil {
        os.Exit(1)
    }
    temperature=int(temperature/1000)
    return temperature
}

func (fan *Fan) checkCpuTemperature() bool {
    var t int = getCpuTemp()
    if (t >= fan.onThresh) {return true}
    if (t < fan.offThresh) {return false}
    if fan.state {
        return true
    } else {
        return false
    }
}

func (fan *Fan) getCpuTemperature() int {
    return getCpuTemp()
}

func Loop(onThresh int, offThresh int, fanPinName string, logger *Logger){
    // Load all the drivers:
    if _, err := host.Init(); err != nil {
        logger.Log("GPIO initialization error")
        os.Exit(1)
    }

    // Lookup a fan pin by its number:
    fanPin := gpioreg.ByName(fanPinName)
    if fanPin == nil {
        logger.Log("Failed to find Fan pin "+fanPinName)
        os.Exit(1)
    }

    // Fan off : pull down fan pin voltage 
    if err := fanPin.Out(gpio.Low); err != nil {
        logger.Log("Failed to initialize Fan pin "+fanPinName)
        os.Exit(1)
    }

    // init fan object state:
    var fan Fan
    fan.Init(onThresh, offThresh)

    // main loop
    for {
        var temp int = fan.getCpuTemperature()
	logger.Log("cpu temperature = " + strconv.Itoa(temp) + "\n")
        if fan.checkCpuTemperature() {
            // Fan on : pull up fan pin voltage 
            if err := fanPin.Out(gpio.High); err != nil {
                logger.Log("Failed to set Fan pin "+fanPinName)
                os.Exit(1)
            }
	} else {
            // Fan off : pull down fan pin voltage 
            if err := fanPin.Out(gpio.Low); err != nil {
                logger.Log("Failed to set Fan pin "+fanPinName)
                os.Exit(1)
            }
	}
	time.Sleep( time.Minute * 1 )
    }
}

func Usage() {
    fmt.Println("Usage: cpufan [ConfigFile]")
    os.Exit(1)
}

func main() {
    var configFileName string
    argv := os.Args
    if 2 < len(os.Args) { Usage() }
    if 2 == len(os.Args) {
        configFileName = argv[1]
	if _, err := os.Stat(configFileName); err != nil {
            fmt.Printf("Error: configfile \"%s\" does not exist.\n",configFileName)
            Usage()
	}
    } else {
        configFileName = defaultConfigFileName
    }
    // JSON形式configファイル読み込み
    texts, err := ioutil.ReadFile(configFileName)
    if err != nil {
        log.Fatal(err)
	os.Exit(1)
    }
    // configデータ(JSON)デコード
    var config Config
    if err := json.Unmarshal(texts, &config); err != nil {
        log.Fatal(err)
	os.Exit(1)
    }
    var sysLog *syslog.Writer
    if (true == config.UseSyslog) {
        sysLog, err = syslog.Dial("tcp", "localhost:514",
		syslog.LOG_WARNING|syslog.LOG_DAEMON, "ButtonAndShutdown")
        if err != nil {
            log.Fatal(err)
	    config.UseSyslog = false
        }
    }
    var filePointer *os.File
    if "" != config.LogFileName {
        filePointer, err = os.OpenFile(config.LogFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
        if err != nil {
            log.Fatal(err)
	    filePointer = nil
        }
    }
    var logger Logger
    logger.Init(config.UseSyslog, sysLog, config.UseStdout, filePointer)
    Loop(config.LowThreshold, config.HighThreshold , config.FanPin, &logger)
}
