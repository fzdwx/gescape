# gescape

转义 \n 为换行输出，转义 \t 为 4 个空格。


```shell
go install github.com/fzdwx/gescape@main

xclip -o | gescape | xclip -selection clipboard
```


