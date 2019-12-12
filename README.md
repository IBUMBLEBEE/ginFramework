# Gin 框架基础库

## 框架结构

```text
Gin_Framework (主目录)
    main.go （文件，程序入口）
    cache （目录，集成缓存默认 Redis）
    conf （目录，配置文件）
    controllers (目录，业务逻辑控制器)
    logic (目录，待定)
    model (目录，数据库接口处理)
    .gitignore (隐藏文件，git 忽略目录或文件 log/_、_.log、_.log._、vendor/\*)
    Makefile (文件，待定)
    README.md (文件，文档说明)
```

## TODO

1. 前期先支持etcd和配置文件启动，后期添加命令行启动（参考docker项目 Cobra）
2. 独立出framework库，里面包含的功能是：
    1. logger初始化
    2. Gin框架初始化
    3. statsD初始化（集成接口监控）
    4. mysql初始化
    5. redis初始化
    6. kafka初始化（可选，默认不初始化，待定）
# gin_Framework
