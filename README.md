# Log
简单的一个对以下功能的封装
1. 个人常用的logrus logFormat
2. lumberjack的log rotate功能

P.S. 提供了module级别的日志记录函数；查询日志输出文件&行号的逻辑被我改写了，功相比于原生的性能和实用性上可能有出入；介意的话可以直接用Logger句柄记录日志而不直接使用API
