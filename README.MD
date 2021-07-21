RmRedisKeys
===========

[Chinese manual](README_CN.MD)

A tool help to delete redis keys with some pattern.

If the redis configuration is not the default(localhost:6379), a config.yml can be added to the same directory of the executable file.

config.yml
----------

```yaml
host: 192.168.2.3
port: 6379
password: xxxxx
database: 0
```

Example
-------

```bash
rmrediskeys "pattern.*"
```
The keys with pattern `pattern.*` will be deleted with your confirm, such as `pattern.1333`.