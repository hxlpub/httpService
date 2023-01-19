# httpService
a simple http service

# license
httpService is licensed under the MIT License. Please see the LICENSE  files for more information.
# compile 
sh build.sh
# simple to use 
#启动服务
./httpService
#发送hello请求
curl -X POST -d '{"dialog": "hello"}' http://localhost:12345/hello
