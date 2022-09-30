# utils for remote development based http with no client need

##  design
tbd


## feature
### set host copy

```bash
curl -H "Content-Type:text/plain" --data-binary @./main.go http://192.168.56.1:8377/setclip
```
### get host image

```bash
curl  192.168.56.1:8090/getimg  -o test.png
```

### open url in host browser

```bash
curl http://192.168.56.1:8377/openurl -d '{"url":"https://www.baidu.com"}' -X POST -H "Content-Type:application/json"
```
### sync clip command
```
ropreate -h 192.168.56.1
```

## thoughs from repos
[clipper](https://github.com/wincent/clipper)
[lemonade](https://github.com/lemonade-command/lemonade)

