package models

import (
    "bufio"
    "crypto/md5"
    "regexp"
    "log"
    "fmt"
    "strings"
    "os"
    "net/http"
)
func CheckLogin(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("_admin_")
    realUsers := ParseAuthUser()
    if err != nil || cookie.Value == "" || authUser(realUsers, cookie.Value) == false {
        http.Redirect(w, r, "/login/", http.StatusFound)
    }

    //解析r.Referer()，非登录url则转入，否则进入home
    refer := r.Referer()
    if strings.Count(refer, "/login/") > 0 {
        http.Redirect(w, r, "/home/", http.StatusFound)
    }
    http.Redirect(w, r, r.Referer(), http.StatusFound)
}

/**
 * 验证用户
 */
func authUser(realUsers map[string] string, md5CV string) (flag bool) {
    flag = false
    userStr := realUsers[md5CV]
    if len(userStr) > 5 {
        flag = true
    }

    return flag
}

/**
 * 加载并解析配置文件设置
 */
func ParseCfg(fileName string) (cfg map[string]string){
    gopath := os.Getenv("GOPATH");
    var file = gopath + "/conf/" + fileName + ".ini"
    //filePath := flag.String("config", file, "load config contents....")

    cfg = make(map[string]string)
    f, err := os.Open(file)
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

func Md5Crypt(args ...interface{}) (md5Str string){
    var beforeMd5 string

    for _, arg := range args {
        //interface to string
        newOne := arg.(string)
        beforeMd5 += string(newOne)
    }
    md5Byte := md5.Sum([]byte(beforeMd5))
    //I need to cast byte to string
    md5Str = fmt.Sprintf("%x", md5Byte)

    return md5Str
}
/**
 * 解析用户授权配置文件
 */
func ParseAuthUser() (realUsers map[string] string) {
    realUsers = make(map[string] string)
    userString := ParseCfg("authorized")
    var md5Str string

    userPwds := strings.Split(userString["user"], ",")
    if len(userPwds) > 0 {
        for _, up := range userPwds {
            realOne := strings.Split(up, ":")
            u := strings.Trim(realOne[0], "\"")
            p := strings.Trim(realOne[1], "\"")
            md5Str = Md5Crypt(u, p, u, "u")

      //      beforeMd5 = []byte(u + p + u + "u")
      //      md5Byte := md5.Sum(beforeMd5)
            //I need to cast byte to string
      //      md5Str = fmt.Sprintf("%x", md5Byte)

            realUsers[md5Str] = u
        }
    }

    return realUsers
}

func PrintOut(args ...interface{}) {
    log.Println(args)
}
