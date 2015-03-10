RPUSH resque:queue:eat '{"class":"eat","args":["chicken"]}'
RPUSH resque:queue:drink '{"class":"drink","args":["water"]}'
go build && ./goworker-resque-example -queues=drink -interval=0.5
