package env

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//TODO 支持任何格式的文件
//TODO 读取固定格式的文件，支持文件 `#` 内容注释
//TODO 支持文件夹读取， 文件名.文件内关键词名

//默认加载配置到环境变量
func init() {
	Load()
}

//加载文件内容到环境变量中
func Load(paths ...string) {
	if len(paths) == 0 {
		paths = []string{".env"}
	}
	
	envVars := map[string]string{}
	for _, path := range paths {
		confMap := ReadFile(path)
		for k, v := range confMap {
			envVars[k] = v
		}
	}
	
	for envKey, envVal := range envVars {
		SetEnv(envKey, envVal)
	}
}

//获取环境变量
func GetEnv(name string, defaultVal string) string {
	envVar := os.Getenv(name)
	if envVar == "" {
		envVar = defaultVal
	}
	
	return envVar
}

//设置环境变量
func SetEnv(key string, value string) {
	if err := os.Setenv(key, value); err != nil {
		panic(err.Error())
	}
}

//读取文件配置
func ReadFile(filePath string) map[string]string {
	var sep string
	
	fileName, errRead := ioutil.ReadFile(filePath)
	if errRead != nil {
		panic(errRead.Error())
	}
	
	envs := make(map[string]string)
	configs := strings.Split(string(fileName), "\n")
	
	for _, config := range configs {
		if len(strings.TrimSpace(config)) == 0 || strings.Contains(config, "#") {
			continue
		}
		if strings.Contains(config, ":") {
			sep = ":"
		} else if strings.Contains(config, "=") {
			sep = "="
		}
		if 0 == len(sep) {
			continue
		}
		
		pairs := strings.Split(config, sep)
		envs[strings.ToUpper(strings.TrimSpace(pairs[0]))] = strings.TrimSpace(pairs[1])
	}
	
	return envs
}

//扫描文件夹
func ScanDir(dirPath string) []string {
	files := make([]string, 0, 0)
	
	if err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		
		return nil
	}); err != nil {
		panic(err.Error())
	}
	
	return files
}

//读取文件夹下文件配置
func ReadDir(dirPath string) map[string]string {
	confPairs := make(map[string]string)
	
	files := ScanDir(dirPath)
	for _, file := range files {
		filePathSlice := strings.Split(file, "\\")
		
		fullFileName := filePathSlice[len(filePathSlice)-1]
		pairs := ReadFile(file)
		for k, v := range pairs {
			newKey := fmt.Sprintf("%s.%s", k, strings.Split(fullFileName, "\\")[0])
			confPairs[newKey] = v
		}
	}
	
	return confPairs
}
