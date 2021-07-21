RmRedisKeys
===========

[English manual](README.MD)

RmRedisKeys用一款帮助批量删除redis里边符合某些特定模式的缓存项的工具。

如果不是使用的默认的redis(localhost:6379)，可以在可执行文件所在目录增加config.yml，对redis进行配置。

config.yml
----------

```yaml
host: 192.168.2.3
port: 6379
password: xxxxx
database: 0
```

示例
-------

```bash
rmrediskeys "pattern.*"
```
这个命令在你进行手工确认后，将删除所有匹配`pattern.*`这个模式的redis里边的缓存项，如：pattern.1333。