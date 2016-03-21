package main

import (
    "fmt"
    "os"
   // "strconv"
    "reflect"
    "xmlib/xmjson"
    "encoding/json"
    "io"
    "strings"
)

func main(){
    const jsonStream = `
        {"Name": "Ed", "Text": "Knock knock."}
        {"Name": "Sam", "Text": "Who's there?"}
        {"Name": "Ed", "Text": "Go fmt."}
        {"Name": "Sam", "Text": "Go fmt who?"}
        {"Name": "Ed", "Text": "Go fmt yourself!"}
    `
    type Message struct {
        Name, Text string
    }
    dec := json.NewDecoder(strings.NewReader(jsonStream))
    for {
        var m Message
        if err := dec.Decode(&m); err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("err")
        }
        fmt.Printf("%s: %s\n", m.Name, m.Text)
    }



    var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
    type Animal struct {
        Name  string
        Order string
    }
    var animals []Animal
    err := json.Unmarshal(jsonBlob, &animals)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v", animals)
    fmt.Println(animals[0].Name)




    var jst = []byte(`[{"a":"ddddd","n":9223372036854775807},{"a":"ssssss","n":223372036854775807}]`)
    type va struct {
        A string
        N int64
    }

    var vt []va
    err = json.Unmarshal(jst,&vt)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v", vt)
    fmt.Println(vt[1].N, reflect.TypeOf(vt[0].N))

    var r []map[string]interface{}
    err = json.Unmarshal(jst, &r)
    gobook := r[0]
    for k, v := range gobook {
        switch v2 := v.(type) {
            case string:
                fmt.Println(k, "is string", v2)
            case int:
                fmt.Println(k, "is int", v2)
            case float64:
                fmt.Println(k, "is float64", v2)
            case bool:
                fmt.Println(k, "is bool", v2)
            case []interface{}:
                fmt.Println(k, "is an array: ")
                for i, iv := range v2{
                    fmt.Println(i,iv)
                }
            default:
                fmt.Println(k, "is another")
        }
    }
    fmt.Println("\n","*************","\n")




    var jstr string = `{
        "test": { 
            "string_slice": ["asdf", "ghjk", "zxcv"],
            "silce": [1, "2", 3],
            "silcewithsubs": [{"subkeyone": 1},
            {"subkeytwo": 2, "subkeythree": 3}],
            "int": 10,
            "float": 5.150,
            "bignum": 9223372036854775807,
            "string": "simplejson",
            "bool": true 
        },
        "test1": { 
            "string_slice": "sb"
        }
    }`

    //xmjson
    js, err := xmjson.NewJSON([]byte(jstr))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(js)
    
    val, err := js.Get("test1").Get("bignum").String()
    fmt.Println(val, err, "\n")

    silce, err := js.Get("test").Get("silce").Slice()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(silce)

    aws := js.Get("test").Get("silcewithsubs")
    fmt.Println(aws)

    awsval, err := aws.GetIndex(0).Get("subkeyone").Int()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(awsval)

    awsval, err = aws.GetIndex(1).Get("subkeytwo").Int()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(awsval)

    awsval, err = aws.GetIndex(1).Get("subkeythree").Int()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(awsval)

    i, err := js.Get("test").Get("int").Int()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(i)

    f, err := js.Get("test").Get("float").Float64()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(f)

    s, err := js.Get("test").Get("string").String()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(s)

    b, err := js.Get("test").Get("bool").Bool()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(b)

    strs, err := js.Get("test").Get("string_slice").StringSlice()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(strs[0])

    gp, err := js.GetLot("test", "string").String()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(gp)

    gp1, err := js.GetLot("test", "int").Int()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(gp1)

    js.Set("test", "setTest")
    fmt.Println(js.Get("test")) 


    os.Exit(1)
}
