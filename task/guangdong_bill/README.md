广东账单模板消息推送 6838
> 每月20-30号，早上07:00-20:00执行，每20分钟执行一次 
0 * 10-23 * * *

```bash
# 脚本
./crontab -taskName=guangdong_bill -limit=100000
```

测试
```bash
# 脚本
./crontab -taskName=guangdong_bill -limit=100000 -test=1
```