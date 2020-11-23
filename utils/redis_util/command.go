package redis_util

//部分常用的redis命令
const (
	//基本命令
	RcAuth   = "auth"
	RcPing   = "ping"
	RcSelect = "select"

	//key的操作命令
	RcDel       = "Del"       //DEL key   删除key
	RcDump      = "Dump"      //DUMP key  序列化key ，并返回被序列化的值
	RcExists    = "Exists"    //EXISTS key  检查key是否存在
	RcExpire    = "Expire"    //EXPIRE key seconds  设置过期时间，以秒计
	RcExpireAt  = "ExpireAt"  //EXPIREAT key timestamp  设置过期时间，指定时间戳(秒级)
	RcPexpire   = "Pexpire"   //PEXPIRE key milliseconds   设置过期时间以毫秒计
	RcPexpireAt = "PexpireAt" //PEXPIREAT key milliseconds-timestamp  设置过期时间，指定时间戳(毫秒级)
	RcKeys      = "Keys"      //KEYS pattern 查找所有符合给定模式( pattern)的key
	RcMove      = "Move"      //MOVE key db 将当前数据库的 key 移动到给定的数据库db
	RcPersist   = "Persist"   //PERSIST key 移除 key 的过期时间，key 将持久保持
	RcPttl      = "Pttl"      //PTTL key 以毫秒为单位返回 key 的剩余的过期时间
	RcTtl       = "ttl"       //TTL key 以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)
	RcRename    = "Rename"    //RENAME key newkey 修改 key 的名称
	RcRenameNx  = "RenameNx"  //仅当 newkey 不存在时，将 key 改名为 newkey
	RcType      = "type"      //TYPE key 返回 key 所储存的值的类型

	//字符串的操作命令
	RcSet          = "set"          //Set key value
	RcGet          = "get"          //Get key
	RcGetRange     = "GetRange"     //GETRANGE key start end  返回 key 中字符串值的子字符
	RcGetSet       = "GetSet"       //GETSET key value 将给定 key 的值设为 value ，并返回 key 的旧值(old value)
	RcGetBit       = "GetBit"       //GETBIT key offset 对 key 所储存的字符串值，获取指定偏移量上的位(bit)
	RcMget         = "Mget"         //MGET key1 [key2..] 获取所有(一个或多个)给定 key 的值
	RcSetBit       = "SetBit"       //SETBIT key offset value 对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)
	RcSetEx        = "SetEx"        //SETEX key seconds value
	RcSetNx        = "SetNx"        //SETNX key value 只有在 key 不存在时设置 key 的值
	RcIncr         = "incr"         //INCR key   将 key 中储存的数字值增一
	RcDecr         = "decr"         //DECR key  将 key 中储存的数字值减一
	RcHIncrByFloat = "hIncrByFloat" //HINCRBYFLOAT key field increment 为哈希表 key 中的指定字段的浮点数值加上增量 increment

	//hash 命令
	RcHdel    = "hdel"    //HDEL key field1 [field2]   删除一个或多个哈希表字段
	RcHexists = "hexists" //HEXISTS key field 查看哈希表 key 中，指定字段是否存在
	RcHget    = "hget"    //HGET key field
	RcHgetAll = "hgetall" //HGETALL key   获取在哈希表中指定 key 的所有字段和值
	RcHkeys   = "hkeys"   //HKEYS key  获取所有哈希表中的字段
	RcHlen    = "hlen"    //HLEN key 获取哈希表中字段的数量
	RcHmget   = "hmget"   //HMGET key field1 [field2] 获取所有给定字段的值
	RcHset    = "hset"    //HSET key field value
	RcHmset   = "hmset"   //HMSET key field1 value1 [field2 value2 ] 同时set多个 field-value
	RcHvals   = "hvals"   //HVALS key  获取哈希表中所有值
	RcHsetNx  = "hsetnx"  //HSETNX key field value 只有在字段 field 不存在时，设置哈希表字段的值
	RcHincrby = "Hincrby" //HINCRBY key field increment 为哈希表 key 中的指定字段的整数值加上增量 increment

	//list 命令
	RcBlpop   = "blpop"   //BLPOP key1 [key2 ] timeout 移出并获取列表的第一个元素 不存在则阻塞超时
	RcBrpop   = "brpop"   //BRPOP key1 [key2 ] timeout 移出并获取列表的最后一个元素 不存在则阻塞超时
	RcLindex  = "lindex"  //LINDEX key index  通过索引获取列表中的元素
	RcLinsert = "linsert" //LINSERT key BEFORE|AFTER pivot value 在列表的元素前或者后插入元素
	RcLlen    = "llen"    //LLEN key 获取列表长度
	RcLpop    = "lpop"    //LPOP key 移出并获取列表的第一个元素
	RcLpush   = "lpush"   // LPUSH key value1 [value2] 将一个或多个值插入到列表头部
	RcLrange  = "lrange"  //LRANGE key start stop 获取列表指定范围内的元素
	RcRpop    = "rpop"    //RPOP key 移除列表的最后一个元素，返回值为移除的元素
	RcRpush   = "rpush"   //RPUSH key value1 [value2] 在列表尾添加一个或多个值

	// count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
	//count < 0 : 从表尾开始向表头搜索， 移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
	//count = 0 : 移除表中所有与 VALUE 相等的值
	RcLrem = "lrem" //LREM key count value 移除列表元素

	RcLtrim = "ltrim" //LTRIM key start stop 对一个列表进行修剪(让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除)

	//set 命令
	RcSadd        = "sadd"        //SADD key member1 [member2] 向集合添加一个或多个成员
	RcScard       = "scard"       //SCARD key 获取集合的成员数(个数)
	RcSdiff       = "sdiff"       //SDIFF key1 [key2]  返回给定所有集合的差集
	RcSinter      = "sinter"      //SINTER key1 [key2] 返回给定所有集合的交集
	RcSisMember   = "sismember"   //SISMEMBER key member 判断 member 元素是否是集合 key 的成员
	RcSmembers    = "smembers"    //SMEMBERS key 返回集合中的所有成员
	RcSrem        = "srem"        //SREM key member1 [member2] 移除集合中一个或多个成员
	RcSunion      = "sunion"      //SUNION key1 [key2] 返回所有给定集合的并集
	RcSRandMember = "srandmember" // SRANDMEMBER key [count] 返回集合中一个或多个随机数

	//zset 部分常用命令
	RcZadd             = "zadd"             //ZADD key score1 member1 [score2 member2] 向有序集合添加一个或多个成员，或者更新已存在成员的分数
	RcZcard            = "zcard"            //ZCARD key 获取有序集合的成员数
	RcZcount           = "zcount"           //ZCOUNT key min max 计算在有序集合中指定区间分数的成员数
	RcZrange           = "zrange"           //ZRANGE key start stop [WITHSCORES] 通过索引区间返回有序集合成指定区间内的成员
	RcZrank            = "zrank"            //ZRANK key member 返回有序集合中指定成员的索引(排行)
	RcZrem             = "zrem"             //ZREM key member [member ...] 移除有序集合中的一个或多个成员
	RcZremRangeByLex   = "ZremRangeByLex"   //ZREMRANGEBYLEX key min max  移除有序集合中给定的字典区间的所有成员
	RcZremRangebyRank  = "ZremRangebyRank"  //ZREMRANGEBYRANK key start stop 移除有序集合中给定的排名区间的所有成员
	RcZremRangebyScore = "ZremRangebyScore" //移除有序集合中给定的分数区间的所有成员
	RcZrevRange        = "ZrevRange"        //ZREVRANGE key start stop [WITHSCORES] 返回有序集中指定区间内的成员，通过索引，分数从高到底
	RcZrevRangeByScore = "ZrevRangeByScore" //ZREVRANGEBYSCORE key max min [WITHSCORES]  返回有序集中指定分数区间内的成员，分数从高到低排序
	RcZrevRank         = "ZrevRank"         //ZREVRANK key member 返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序
	RcZscore           = "Zscore"           //ZSCORE key member  返回有序集中，成员的分数值
	RcZincrby          = "zincrby"          //ZINCRBY key increment member 有序集合中对指定成员的分数加上增量 increment

	//发布订阅命令
	RcPublish     = "publish"     //PUBLISH channel message  将信息发送到指定的频道
	RcSubscribe   = "subscribe"   //SUBSCRIBE channel [channel ...]  订阅给定的一个或多个频道的信息
	RcUnSubscribe = "unSubscribe" //UNSUBSCRIBE channel [channel ...]  退订指定的频道

	//事务命令
	RcMulti   = "multi"   //标记事务开始
	RcDisCard = "discard" //取消事务 放弃执行事务块内的所有命令
	RcExec    = "exec"    //执行所有事务块内的命令

	//script 脚本

	//管道
)
