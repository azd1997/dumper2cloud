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
