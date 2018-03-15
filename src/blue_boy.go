package main

import (
    "flag"
    "github.com/tarm/goserial"
    "github.com/larspensjo/config"
    "os"
    "log"
)

var (
    conFile = flag.String("configfile","/config.ini","config file")
)

var TOPIC = make(map[string]string)

func serialOpt() {
        //获取当前路径
    file, _ := os.Getwd()
    cfg, err := config.ReadDefault(file + *conFile)

        //获取配置文件中的配置项
    id, err := cfg.String("COM","COMID")
    //设置串口编号
    c := &serial.Config{Name: id, Baud: 115200}
    //打开串口
    s, err := serial.OpenPort(c)

    if err != nil {
        log.Fatal(err)
    }

    command, err := cfg.String("COM","COMMAND")
    // 写入串口命令
    log.Printf("写入的指令 %s", command)
    n, err := s.Write([]byte(command))

    if err != nil {
        log.Fatal(err)
    }

    buf := make([]byte, 128)
    n, err = s.Read(buf)
    log.Printf("读取信息 %s", buf[:n])
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("%q", buf[:n])
}
