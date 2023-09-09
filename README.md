# 爬取巨潮资讯网公告
## 关于参数如何填写的问题
### stock
股票代码
可以为空,或填写如"000001,gssz0000001",逗号前后分别为./info/szse_stock.txt中的code和orgId
### searchkey
标题关键词
可以为空,或填写如"处罚"或"处罚;公告",分号表示或
### category
报告类型
可以为空,或填写./info/category.txt中的key
### trade
行业分类
可以为空,或填写./info/trade.txt中的key
### seDate
日期范围
不可以为空, 填写如:"2021-01-01~2021-12-31"

## 输入参数前最好看看会有多少