# Typing game by Golang
## 目的
* goroutineやchannelについての理解を深める
* パッケージ化してみる

## 概要
```bash
go run main.go [-limit n]
```
limitはオプションでタイムリミットのこと.
デフォルトでは20secになっている.

こんな感じ
```bash
❯ go run main.go -limit 10
start typing game.
type the word!
time limit is 10sec
>> speed    speed 
>> lie    liee
missed!
>> front    front
>> be    be
>> measure    measu
game ends! the number of correct answer: 3
```

## TODO
* contextなどにも触れてみたい
* 構造体やinterfaceをどんどん使えるようになりたい.