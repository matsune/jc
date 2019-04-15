# jc
`jc` is a colorized JSON processing library and command-line tool.


## Usage
### CLI
![2019-02-16 15 33 32](https://user-images.githubusercontent.com/12775019/52895845-63132300-3203-11e9-8161-bcc89f068f7a.png)

`jc` command gets plain json string through pipe and processes pretty formatted json to Stdout.  
`jc` has default color pattern but you can change them by creating config file under your $HOME directory.

You can apply color and styles by describing the ANSI/VT100 format code in the config file. 

```toml
# ~/.jc.conf
Key     = "1,31"  # Bold, Red
Number  = "34"    # Blue
String  = "33"    # Yellow
Bool    = "36"    # Cyan
Null    = "37,42" # LightGray, Green Background
```
|name|desc|
|---|---|
|Key|Object key|
|Number|Value of number type|
|String|Value of string type|
|Bool|Value of bool type|
|Null|Value of null|

### Library
```go
jsonStr := `
{
  "foo": "bar",
  "num": 100
}
`
j := jc.New()
j.Colorize(jsonStr)
```
