package main

import (
  "net/http"
  "./example"
  "io/ioutil"
  "io"
  "os"
  "bufio"
  "strings"
  "fmt"
)

// http handler
func helloHandler(writer http.ResponseWriter, r *http.Request)  {
  key := r.FormValue("key")
  value := r.FormValue("value")
  if key == "" || value == "" {
    fmt.Fprint(writer, "no param")
    return
  } else {
    fmt.Printf("key=%s,value=%s\n", key, value)
    io.WriteString(writer, "hello, world!!")
  }
}

// ioutil
func readFile(fileName string) (err error) {
  data, err := ioutil.ReadFile(fileName)

  if err == nil {
    var content = string(data)
    fmt.Printf("content -> %s", content)
  }
  return err

}

var dictMap = make(map[string]string, 2)

func readLineByLine(fileName string) (error error) {

  f, err := os.Open(fileName)
  if err != nil {
    return err
  }

  defer  f.Close()

  reader := bufio.NewReader(f)

  for  {
    line, _ , err := reader.ReadLine()

    if err != nil {
      if err == io.EOF {
        return nil
      }
      return err
    }

    var lineContent = string(line)
    fmt.Printf("line content -> %s\n", lineContent)

    var sp []string = strings.Split(lineContent, "\t")

    fmt.Println("sp[0]", sp[0])

    dictMap[sp[0]] = sp[1]

  }

  return err

}

func main() {

  fmt.Println("go hello, world!!")

  // 正常情况
  if result, errorMsg := example.Divide(100, 10); errorMsg == "" {
    fmt.Println("100/10 = ", result)
  }
  // 当被除数为零的时候会返回错误信息
  if _, errorMsg := example.Divide(100, 0); errorMsg != "" {
    fmt.Println("errorMsg is: ", errorMsg)
  }

  fmt.Printf("listen port:%d...\n", 12345)

  readFile("/Users/devops/workspace/go-pro/go-example/text")

  err := readLineByLine("/Users/devops/workspace/go-pro/go-example/text")

  // 读取文件成功返回
  if err == nil {
    fmt.Println("length map -> ", len(dictMap))
    for key := range dictMap {
      fmt.Println("map item -> " + dictMap[key])
    }
    fmt.Printf("read success")
  }
  var resumeInffo = example.NewResume("CJYOU", "188", "SH")

  fmt.Print("info: name -> " + resumeInffo.GetName())

  // http resquest
  // http.HandleFunc("/go", helloHandler)
  //  http.ListenAndServe(":12345", nil)

  // server filesystem
  // http.ListenAndServe(":12345", http.FileServer(http.Dir(".")))


}
