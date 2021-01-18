package main

import (
	"bytes"
	"fmt"
	"github.com/kardianos/service"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var appPath string

type program struct{}

func main() {
	var err error
	appPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("服务执行路径获取失败：%s", err.Error())
	}
	svcConfig := &service.Config{
		Name:        "nginx-service",
		DisplayName: "nginx-service",
		Description: "nginx-service",
	}
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatalln(err)
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err := s.Install()
			if err != nil {
				log.Fatalln(err.Error())
			}
			log.Println("服务安装成功")
			return
		case "remove":
			err := s.Uninstall()
			if err != nil {
				log.Fatalln(err.Error())
			}
			log.Print("服务卸载成功")
			return
		case "reload":
			prg.Reload()
			return
		}
	}
	err = s.Run()
}

func (p *program) run() {
	cmd := exec.Command(fmt.Sprintf(`%s\nginx.exe`, appPath), "-c", fmt.Sprintf(`%s\conf\nginx.conf`, appPath), "-p", appPath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		fmt.Println(stderr.String())
	}
	os.Exit(0)
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	cmd := exec.Command(fmt.Sprintf(`%s\nginx.exe`, appPath), "-s", "stop", "-p", appPath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		fmt.Println(stderr.String())
		return nil
	}
	log.Println("服务停止！")
	return nil
}

func (p *program) Reload() {
	cmd := exec.Command(fmt.Sprintf(`%s\nginx.exe`, appPath), "-s", "reload", "-p", appPath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		fmt.Println(stderr.String())
		return
	}
	log.Println("配置重载成功！")
}
