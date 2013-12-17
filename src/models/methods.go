package models

import (
    "bufio"
    "flag"
    "regexp"
//   "log"
    "strings"
    "os"
    "fmt"
)
/**
 * 加载并解析配置文件设置
 */
func ParseCfg(fileName string) (cfg map[string]string){
    gopath := os.Getenv("GOPATH");
    var file = gopath + "/conf/" + fileName + ".ini"
    filePath := flag.String("config", file, "load config contents....")
    //fmt.Println(*filePath)
    cfg = make(map[string]string)
    f, err := os.Open(*filePath)
    //文件加载出错
    if err != nil {
        PrintOut("error on opening >> " + file)
        os.Exit(-1)
    }
    re := regexp.MustCompile(`[\t ]*([0-9A-Za-z_]+)[\t ]*=[\t ]*([^\t\n\f\r# ]+)[\t #]*`)
    //using scanner to read file
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        //查找匹配的
        slice := re.FindStringSubmatch(line)

        if slice != nil {
            cfg[slice[1]] = slice[2]
        }
    }

    return cfg
}

func PrintOut(contents string) {
    fmt.Println(contents)
}
