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
2020/09/16 20:36:08 config init finish. configpath=./conf/config.ini
2020/09/16 20:36:08 config:
 dumper2cloud:  [fakedumper /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/bin/fakedumper minio dumper-sql 1024 0.6 0.3 10 10] 
 minio:  [play.min.io Q3AM3UQ867SPQQA43P2F zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG true us-east-1] 
 mysql:  [127.0.0.1 3306 root pwd xx /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data 128 ]
2020/09/16 20:36:10 Successfully created dumper-sql-2020916-2036
2020/09/16 20:36:10 scheduler init finish ...
2020/09/16 20:36:10 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-1 succ, delete it now
2020/09/16 20:36:14 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-2 succ, delete it now
2020/09/16 20:36:18 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-3 succ, delete it now
2020/09/16 20:36:21 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-4 succ, delete it now
2020/09/16 20:36:25 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-5 succ, delete it now
2020/09/16 20:36:29 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-6 succ, delete it now
2020/09/16 20:36:33 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-7 succ, delete it now
2020/09/16 20:36:36 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-8 succ, delete it now
2020/09/16 20:36:40 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-9 succ, delete it now
2020/09/16 20:36:45 upload /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data/fake-dumper-10 succ, delete it now

d.cmd [/home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/bin/fakedumper -o /home/eiger/gopath-default/src/github.com/azd1997/dumper2cloud/test/data] Stdout/Stderr: 
2020/09/16 20:36:13 dump file [fake-dumper-1] write finish. size 10485770 B
2020/09/16 20:36:17 dump file [fake-dumper-2] write finish. size 10485770 B
2020/09/16 20:36:21 dump file [fake-dumper-3] write finish. size 10485770 B
2020/09/16 20:36:25 dump file [fake-dumper-4] write finish. size 10485770 B
2020/09/16 20:36:28 dump file [fake-dumper-5] write finish. size 10485770 B
2020/09/16 20:36:32 dump file [fake-dumper-6] write finish. size 10485770 B
2020/09/16 20:36:36 dump file [fake-dumper-7] write finish. size 10485770 B
2020/09/16 20:36:39 dump file [fake-dumper-8] write finish. size 10485770 B
2020/09/16 20:36:43 dump file [fake-dumper-9] write finish. size 10485770 B
2020/09/16 20:36:47 dump file [fake-dumper-10] write finish. size 10485670 B
2020/09/16 20:36:47 dump finish. size 104857600 B


2020/09/16 20:36:47 dumpLoop quit ...
2020/09/16 20:36:47 detectLoop quit ...
2020/09/16 20:36:47 uploadLoop quit ...
2020/09/16 20:36:47 dumper2cloud finish ...

```

## 补充说明

由于文件上传是个比较耗时的操作，后续需要考虑将文件上传做成可配置的多goroutine执行，以匹配dumper的速度