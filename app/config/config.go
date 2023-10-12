// 生成全局配置和admin SSl证书
package config

import (
	"akBlog/app/util"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// 配置文件路径
const configPath = "config/main.conf"

// 检测是否存在配置文件，不存在则新建
func init() {
	if _, err := os.ReadFile(configPath); err != nil {
		fmt.Println("未检测到配置文件，是否创建(y/n)")
		var is string
		fmt.Scanln(&is)
		if is != "y" {
			fmt.Println("退出成功")
			os.Exit(1)
		}
		createConfig()
	}
}

// 配置信息
type ConfigInfo struct {
	// 服务器IP
	ServerIP string
	// 主域名
	Domain string
	// 公开web服务器端口
	Port string
	// 管理员web服务器端口
	AdminPort string
	// 公开web服务器是否使用HTTPS
	IsHTTPS string
}

// 快捷获取配置文件信息
func Get(name string) string {

	data, err := os.ReadFile(configPath)
	util.ErrprDisplay(err)

	c := ConfigInfo{}
	err = json.Unmarshal(data, &c)
	util.ErrprDisplay(err)

	switch name {
	case "port":
		return c.Port

	case "serverIP":
		return c.ServerIP

	case "domain":
		return c.Domain

	case "adminPort":
		return c.AdminPort

	case "isHTTPS":
		return c.IsHTTPS
	}

	return ""

}

// akBlog配置创建
func createConfig() {
	//配置文件夹创建
	os.Mkdir("config/cert/", 0700)
	os.Mkdir("config/cert/Ca", 0700)
	os.Mkdir("config/cert/adminCa", 0700)

	// 配置信息
	configInfo := ConfigInfo{}

	//写入配置文件
	fmt.Println("输入你的服务器IP(默认:127.0.0.1)")
	fmt.Scanln(&configInfo.ServerIP)
	if configInfo.ServerIP == "" {
		configInfo.ServerIP = "127.0.0.1"
	}

	fmt.Println("输入你的域名(默认:localhost)")
	fmt.Scanln(&configInfo.Domain)
	if configInfo.Domain == "" {
		configInfo.Domain = "localhost"
	}

	fmt.Println("设置你的web端口(默认:8080)")
	fmt.Scanln(&configInfo.Port)
	if configInfo.Port == "" {
		configInfo.Port = ":8080"
	} else {
		configInfo.Port = ":" + configInfo.Port
	}

	fmt.Println("设置你的网站管理员端口(默认:59812)")
	fmt.Scanln(&configInfo.AdminPort)
	if configInfo.AdminPort == "" {
		configInfo.AdminPort = ":59812"
	} else {
		configInfo.AdminPort = ":" + configInfo.AdminPort
	}

	// 公开端口是否使用HTTPS协议
	fmt.Println(configInfo.Port + "端口是否使用HTTPS协议(y/n)")
	var isHTTPS string
	fmt.Scanln(&isHTTPS)
	if isHTTPS == "y" {
		configInfo.IsHTTPS = "y"
		fmt.Println(configInfo.Port + "端口将使用https协议")
		fmt.Println("!!!")
		fmt.Println("请将证书放在config/cert/Ca/路径")
		fmt.Println("证书文件名格式：\n 域名.crt\n域名.key\n不规范命名将无法正确导入证书")
		fmt.Println("!!!")
	}

	// 将数据写入配置文件
	data, err := json.MarshalIndent(configInfo, "", "	")
	util.ErrprDisplay(err)
	err = os.WriteFile(configPath, data, 0600)
	util.ErrprDisplay(err)

	// 设置配置文件夹为只读
	util.Command("chmod 400 config/cert/adminCa/*")
	util.Command("chmod 400 " + configPath)

	// 控制台输出配置信息
	fmt.Println("网站管理员端口设置为" + Get("adminPort"))
	fmt.Println("web端口为" + Get("port"))
	fmt.Println("服务器域名是" + Get("domain"))
	fmt.Println("服务器IP是" + Get("serverIP"))

	// 创建管理员证书
	createCA(configInfo.Domain)

}

// 创建管理员https证书
func createCA(doMain string) {
	domain := Get("domain")
	shell("openssl genrsa -out ./config/cert/adminCa/ca.key 4096")
	fmt.Println("生成证书进度：35%")

	shell(`openssl req -x509 -new -nodes -sha512 -days 3650 \
	-subj "/C=CN/ST=Beijing/L=Beijing/O=example/OU=Personal/CN=` + domain + `" \
	-key ./config/cert/adminCa/ca.key \
	-out ./config/cert/adminCa/ca.crt`)
	fmt.Println("生成证书进度：50%")

	shell(`openssl genrsa -out ./config/cert/adminCa/` + domain + `.key 4096`)
	fmt.Println("生成证书进度：65%")

	shell(`openssl req -sha512 -new \
	-subj "/C=CN/ST=Beijing/L=Beijing/O=example/OU=Personal/CN=` + domain + `" \
	-key ./config/cert/adminCa/` + domain + `.key \
	-out ./config/cert/adminCa/` + domain + `.csr`)
	fmt.Println("生成证书进度：80%")

	shell(`cat > ./config/cert/adminCa/v3.ext <<-EOF
	authorityKeyIdentifier=keyid,issuer
	basicConstraints=CA:FALSE
	keyUsage=digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
	extendedKeyUsage=serverAuth
	subjectAltName=@alt_names
	
	[alt_names]
	DNS.1=` + domain + `
	EOF`)
	fmt.Println("生成证书进度：95%")

	shell(`openssl x509 -req -sha512 -days 3650 \
	-extfile ./config/cert/adminCa/v3.ext \
	-CA ./config/cert/adminCa/ca.crt -CAkey ./config/cert/adminCa/ca.key -CAcreateserial \
	-in ./config/cert/adminCa/` + domain + `.csr \
	-out ./config/cert/adminCa/` + domain + `.crt`)
	fmt.Println("生成证书进度：100%")
	// 删除不需要的证书
	// shell("rm -rf ./config/cert/adminCa/*.csr")
	// shell("rm -rf ./config/cert/adminCa/v3.ext")
	// // shell("rm -rf ./config/cert/adminCa/ca.crt")
	// shell("rm -rf ./config/cert/adminCa/ca.key")
	// shell("rm -rf ./config/cert/adminCa/ca.srl")
}

// 执行命令
func shell(cmd string) {

	tmp := exec.Command("bash", "-c", cmd)
	tmp.Stdout = os.Stdout
	tmp.Stderr = os.Stderr
	tmp.Run()
}
