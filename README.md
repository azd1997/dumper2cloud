# dumper2cloud
- 将数据库备份至云端

## 使用说明

1. 编译可执行文件`go build -o ./bin/d2c .`
2. 根据要备份的数据库自行填写`./conf/config.ini`配置文件
3. 执行`./bin/d2c`

## 测试说明

在没有安装mysql/mydumper等的情况下，进行测试，
可以在配置文件中启用`fakedumper`和`fakecloud`进行测试。
`fakedumper`需要提前编译`go build -o ./bin/fakedumper ./test/fakedumper/.`

测试示例：
```shell script
eiger@eiger-pad:~/gopath-default/src/github.com/azd1997/dumper2cloud$ ./bin/d2c 
2020/09/16 20:25:31 config init finish. configpath=./conf/config.ini
2020/09/16 20:25:31 config:
 dumper2cloud:  [fakedumper /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/bin/fakedumper minio dumper-sql 1024 0.6 0.3 10 10] 
 minio:  [play.min.io Q3AM3UQ867SPQQA43P2F zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG true us-east-1] 
 mysql:  [127.0.0.1 3306 root pwd xx /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data 128 ]
2020/09/16 20:25:36 Successfully created dumper-sql-2020916-2025
2020/09/16 20:25:36 scheduler init finish ...
2020/09/16 20:25:36 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-1 succ, delete it now
2020/09/16 20:25:41 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-2 succ, delete it now
2020/09/16 20:25:46 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-3 succ, delete it now
2020/09/16 20:25:51 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-4 succ, delete it now
2020/09/16 20:25:53 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-5 succ, delete it now
2020/09/16 20:25:56 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-6 succ, delete it now
2020/09/16 20:26:00 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-7 succ, delete it now
2020/09/16 20:26:04 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-8 succ, delete it now
2020/09/16 20:26:07 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-9 succ, delete it now
2020/09/16 20:26:11 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-10 succ, delete it now

d.cmd [/home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/bin/fakedumper -o /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data] Stdout/Stderr: 
everything seems ok.

2020/09/16 20:26:14 dumpLoop quit ...
2020/09/16 20:26:14 detectLoop quit ...
2020/09/16 20:26:14 uploadLoop quit ...
2020/09/16 20:26:14 dumper2cloud finish ...
```

## 补充说明

由于文件上传是个比较耗时的操作，后续需要考虑将文件上传做成可配置的多goroutine执行，以匹配dumper的速度