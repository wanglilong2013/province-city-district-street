## 获取全国省、市、区、街道数据
数据来源于高德开放平台中的`行政区域查询`

--------------------------------

### 1. 准备数据表

```
CREATE TABLE `areas` (
  `id` int(10) unsigned NOT NULL,
  `name` varchar(128) NOT NULL,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `level` varchar(32) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`,`name`),
  KEY `parent_id_index` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

###2. 修改配置文件

根据自身情况修改conf.go文件中的数据库配置信息

```
    MysqlDbUser = "root"  //数据库连接用户名
    MysqlDbPwd  = "root"  //数据库连接密码
    MysqlDbHost = "127.0.0.1"   //数据库连接地址
    MysqlDbPort = 3306   //数据库连接端口
    MysqlDbName = "test"  //数据库名称
    Key = "XXXX" //高德API KEY,高德开放平台申请
```

### 3.编译并执行main.go文件



#### 4.数据验证

-----------

#### 4.1 获取所有省份

```
   root@localhost:3306 test[2] 16:05:43>select * from areas where level ='province';
   +--------+--------------------------+-----------+----------+---------------------+---------------------+
   | id     | name                     | parent_id | level    | created_at          | updated_at          |
   +--------+--------------------------+-----------+----------+---------------------+---------------------+
   | 110000 | 北京市                   |         0 | province | 2020-09-24 15:32:24 | 2020-09-24 15:32:24 |
   | 120000 | 天津市                   |         0 | province | 2020-09-24 15:32:11 | 2020-09-24 15:32:11 |
   | 130000 | 河北省                   |         0 | province | 2020-09-24 15:31:34 | 2020-09-24 15:31:34 |
   | 140000 | 山西省                   |         0 | province | 2020-09-24 15:32:12 | 2020-09-24 15:32:12 |
   | 150000 | 内蒙古自治区             |         0 | province | 2020-09-24 15:29:56 | 2020-09-24 15:29:56 |
   | 210000 | 辽宁省                   |         0 | province | 2020-09-24 15:30:07 | 2020-09-24 15:30:07 |
   | 220000 | 吉林省                   |         0 | province | 2020-09-24 15:32:06 | 2020-09-24 15:32:06 |
   | 230000 | 黑龙江省                 |         0 | province | 2020-09-24 15:29:31 | 2020-09-24 15:29:31 |
   | 310000 | 上海市                   |         0 | province | 2020-09-24 15:30:30 | 2020-09-24 15:30:30 |
   | 320000 | 江苏省                   |         0 | province | 2020-09-24 15:31:15 | 2020-09-24 15:31:15 |
   | 330000 | 浙江省                   |         0 | province | 2020-09-24 15:31:21 | 2020-09-24 15:31:21 |
   | 340000 | 安徽省                   |         0 | province | 2020-09-24 15:30:39 | 2020-09-24 15:30:39 |
   | 350000 | 福建省                   |         0 | province | 2020-09-24 15:31:03 | 2020-09-24 15:31:03 |
   | 360000 | 江西省                   |         0 | province | 2020-09-24 15:31:26 | 2020-09-24 15:31:26 |
   | 370000 | 山东省                   |         0 | province | 2020-09-24 15:30:32 | 2020-09-24 15:30:32 |
   | 410000 | 河南省                   |         0 | province | 2020-09-24 15:29:46 | 2020-09-24 15:29:46 |
   | 420000 | 湖北省                   |         0 | province | 2020-09-24 15:30:13 | 2020-09-24 15:30:13 |
   | 430000 | 湖南省                   |         0 | province | 2020-09-24 15:30:52 | 2020-09-24 15:30:52 |
   | 440000 | 广东省                   |         0 | province | 2020-09-24 15:29:39 | 2020-09-24 15:29:39 |
   | 450000 | 广西壮族自治区           |         0 | province | 2020-09-24 15:31:09 | 2020-09-24 15:31:09 |
   | 460000 | 海南省                   |         0 | province | 2020-09-24 15:31:00 | 2020-09-24 15:31:00 |
   | 500000 | 重庆市                   |         0 | province | 2020-09-24 15:30:45 | 2020-09-24 15:30:45 |
   | 510000 | 四川省                   |         0 | province | 2020-09-24 15:31:50 | 2020-09-24 15:31:50 |
   | 520000 | 贵州省                   |         0 | province | 2020-09-24 15:30:24 | 2020-09-24 15:30:24 |
   | 530000 | 云南省                   |         0 | province | 2020-09-24 15:32:18 | 2020-09-24 15:32:18 |
   | 540000 | 西藏自治区               |         0 | province | 2020-09-24 15:30:49 | 2020-09-24 15:30:49 |
   | 610000 | 陕西省                   |         0 | province | 2020-09-24 15:30:19 | 2020-09-24 15:30:19 |
   | 620000 | 甘肃省                   |         0 | province | 2020-09-24 15:31:44 | 2020-09-24 15:31:44 |
   | 630000 | 青海省                   |         0 | province | 2020-09-24 15:31:02 | 2020-09-24 15:31:02 |
   | 640000 | 宁夏回族自治区           |         0 | province | 2020-09-24 15:31:08 | 2020-09-24 15:31:08 |
   | 650000 | 新疆维吾尔自治区         |         0 | province | 2020-09-24 15:30:01 | 2020-09-24 15:30:01 |
   | 710000 | 台湾省                   |         0 | province | 2020-09-24 15:31:34 | 2020-09-24 15:31:34 |
   | 810000 | 香港特别行政区           |         0 | province | 2020-09-24 15:31:44 | 2020-09-24 15:31:44 |
   | 820000 | 澳门特别行政区           |         0 | province | 2020-09-24 15:31:44 | 2020-09-24 15:31:44 |
   +--------+--------------------------+-----------+----------+---------------------+---------------------+
   34 rows in set (0.03 sec)
   
   ```
#### 4.2 根据省份ID获取城市数据（四川）
   
```
    root@localhost:3306 test[4] 16:06:01>select * from areas where parent_id =510000 and level = 'city';
    +--------+-----------------------------+-----------+-------+---------------------+---------------------+
    | id     | name                        | parent_id | level | created_at          | updated_at          |
    +--------+-----------------------------+-----------+-------+---------------------+---------------------+
    | 510100 | 成都市                      |    510000 | city  | 2020-09-24 15:31:51 | 2020-09-24 15:31:51 |
    | 510300 | 自贡市                      |    510000 | city  | 2020-09-24 15:32:01 | 2020-09-24 15:32:01 |
    | 510400 | 攀枝花市                    |    510000 | city  | 2020-09-24 15:32:03 | 2020-09-24 15:32:03 |
    | 510500 | 泸州市                      |    510000 | city  | 2020-09-24 15:32:00 | 2020-09-24 15:32:00 |
    | 510600 | 德阳市                      |    510000 | city  | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510700 | 绵阳市                      |    510000 | city  | 2020-09-24 15:31:55 | 2020-09-24 15:31:55 |
    | 510800 | 广元市                      |    510000 | city  | 2020-09-24 15:31:54 | 2020-09-24 15:31:54 |
    | 510900 | 遂宁市                      |    510000 | city  | 2020-09-24 15:31:54 | 2020-09-24 15:31:54 |
    | 511000 | 内江市                      |    510000 | city  | 2020-09-24 15:31:59 | 2020-09-24 15:31:59 |
    | 511100 | 乐山市                      |    510000 | city  | 2020-09-24 15:32:02 | 2020-09-24 15:32:02 |
    | 511300 | 南充市                      |    510000 | city  | 2020-09-24 15:31:58 | 2020-09-24 15:31:58 |
    | 511400 | 眉山市                      |    510000 | city  | 2020-09-24 15:32:02 | 2020-09-24 15:32:02 |
    | 511500 | 宜宾市                      |    510000 | city  | 2020-09-24 15:31:59 | 2020-09-24 15:31:59 |
    | 511600 | 广安市                      |    510000 | city  | 2020-09-24 15:31:56 | 2020-09-24 15:31:56 |
    | 511700 | 达州市                      |    510000 | city  | 2020-09-24 15:31:57 | 2020-09-24 15:31:57 |
    | 511800 | 雅安市                      |    510000 | city  | 2020-09-24 15:32:01 | 2020-09-24 15:32:01 |
    | 511900 | 巴中市                      |    510000 | city  | 2020-09-24 15:31:51 | 2020-09-24 15:31:51 |
    | 512000 | 资阳市                      |    510000 | city  | 2020-09-24 15:31:55 | 2020-09-24 15:31:55 |
    | 513200 | 阿坝藏族羌族自治州          |    510000 | city  | 2020-09-24 15:32:00 | 2020-09-24 15:32:00 |
    | 513300 | 甘孜藏族自治州              |    510000 | city  | 2020-09-24 15:32:05 | 2020-09-24 15:32:05 |
    | 513400 | 凉山彝族自治州              |    510000 | city  | 2020-09-24 15:32:03 | 2020-09-24 15:32:03 |
    +--------+-----------------------------+-----------+-------+---------------------+---------------------+
    21 rows in set (0.00 sec)
    
```

#### 4.3 根据城市ID 获取区县数据 （成都）
    
```
    root@localhost:3306 test[5] 16:06:54>select * from areas where parent_id = 510100 and level = 'district';
    +--------+--------------+-----------+----------+---------------------+---------------------+
    | id     | name         | parent_id | level    | created_at          | updated_at          |
    +--------+--------------+-----------+----------+---------------------+---------------------+
    | 510104 | 锦江区       |    510100 | district | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510105 | 青羊区       |    510100 | district | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510106 | 金牛区       |    510100 | district | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510107 | 武侯区       |    510100 | district | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510108 | 成华区       |    510100 | district | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510112 | 龙泉驿区     |    510100 | district | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510113 | 青白江区     |    510100 | district | 2020-09-24 15:31:52 | 2020-09-24 15:31:52 |
    | 510114 | 新都区       |    510100 | district | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510115 | 温江区       |    510100 | district | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510116 | 双流区       |    510100 | district | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510117 | 郫都区       |    510100 | district | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510121 | 金堂县       |    510100 | district | 2020-09-24 15:31:52 | 2020-09-24 15:31:52 |
    | 510129 | 大邑县       |    510100 | district | 2020-09-24 15:31:52 | 2020-09-24 15:31:52 |
    | 510131 | 蒲江县       |    510100 | district | 2020-09-24 15:31:52 | 2020-09-24 15:31:52 |
    | 510132 | 新津区       |    510100 | district | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
    | 510181 | 都江堰市     |    510100 | district | 2020-09-24 15:31:52 | 2020-09-24 15:31:52 |
    | 510182 | 彭州市       |    510100 | district | 2020-09-24 15:31:51 | 2020-09-24 15:31:51 |
    | 510183 | 邛崃市       |    510100 | district | 2020-09-24 15:31:52 | 2020-09-24 15:31:52 |
    | 510184 | 崇州市       |    510100 | district | 2020-09-24 15:31:52 | 2020-09-24 15:31:52 |
    | 510185 | 简阳市       |    510100 | district | 2020-09-24 15:31:52 | 2020-09-24 15:31:52 |
    +--------+--------------+-----------+----------+---------------------+---------------------+
    20 rows in set (0.00 sec)
    
 ```
   
#### 4.4 根据区县ID 获取街道数据（双流）
       
```  
   root@localhost:3306 test[6] 16:07:39>select * from areas where parent_id = 510116 and level = 'street';
   +--------+-----------------+-----------+--------+---------------------+---------------------+
   | id     | name            | parent_id | level  | created_at          | updated_at          |
   +--------+-----------------+-----------+--------+---------------------+---------------------+
   | 510116 | 万安镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 三星镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 东升街道        |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 中和街道        |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 九江街道        |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 公兴街道        |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 兴隆镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 华阳镇街道      |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 协和街道        |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 合江镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 大林镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 太平镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 彭镇            |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 新兴镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 正兴镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 永兴镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 永安镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 煎茶镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 白沙镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 籍田镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 胜利镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 西航港街道      |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 金桥镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 黄水镇          |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 黄甲街道        |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   | 510116 | 黄龙溪镇        |    510116 | street | 2020-09-24 15:31:53 | 2020-09-24 15:31:53 |
   +--------+-----------------+-----------+--------+---------------------+---------------------+
   26 rows in set (0.00 sec)
```
