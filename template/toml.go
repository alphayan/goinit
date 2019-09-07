package temp

const TOML = `#运行模式
[Runmode]
Runmode="dev"
#production
[pro]
#redis配置
RedisHost="192.168.0.100"
RedisPort="6379"
RedisPassword=""
RedisDB=0
#DB配置
DBHost="192.168.0.100"
DBPort="3306"
DBUsername="root"
DBPassword="password"
DBName="test"

#develop
[dev]
#redis配置
RedisHost="192.168.199.248"
RedisPort="6379"
RedisPassword=""
RedisDB=0
#DB配置
DBHost="192.168.199.248"
DBPort="3306"
DBUsername="root"
DBPassword="password"
DBName="test"
`
