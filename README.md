# command

Specify redis server

    ./goworker-resque-example -queues=hnap -interval=0.5 -uri="redis://cache.test.0001.usw2.cache.amazonaws.com:6379/"

> 注意, 即使是預設的 port 也要寫上去

# Test

    RPUSH resque:queue:eat '{"class":"eat","args":["chicken"]}'
    RPUSH resque:queue:drink '{"class":"drink","args":["water"]}'
    go build && ./goworker-resque-example -queues=drink -interval=0.5




