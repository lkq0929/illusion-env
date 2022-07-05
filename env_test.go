package env

import (
	"reflect"
	"testing"
)

//测试读文件
func TestReadFile(t *testing.T) {
	type ReadFileTest struct {
		input  string
		output map[string]string
	}
	
	tests := map[string]ReadFileTest{
		"readFile_simple": {
			input:  "./examples/.env",
			output: map[string]string{"LANGUAGE": "Golang", "AUTHOR": "li.kq", "FUNCTION": "ReadConf", "REDIS_HOST": "127.0.0.1", "REDIS_PORT": "6379"},
		},
	}
	
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ReadFile(test.input)
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("except:%v, return: %v", test.output, result)
			}
		})
	}
}

func TestScanDir(t *testing.T) {
	type scanDirTests struct {
		input  string
		output []string
	}
	tests := map[string]scanDirTests{
		"scanDir_relative_path": {"./examples",
			[]string{"./examples", "examples\\.env"}},
		"scanDir_absolute_path": {"D:\\workspace\\goprojects\\src\\github.com\\li.kaiqiang3\\env\\examples",
			[]string{"D:\\workspace\\goprojects\\src\\github.com\\li.kaiqiang3\\env\\examples",
				"D:\\workspace\\goprojects\\src\\github.com\\li.kaiqiang3\\env\\examples\\.env"}},
	}
	
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ScanDir(test.input)
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("except:%v, return: %v", test.output, result)
			}
		})
	}
}

/*func TestReadDir(t *testing.T) {

}*/

