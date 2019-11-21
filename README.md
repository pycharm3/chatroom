* // 2019-11-20 15：40 
* // user:xy
* // 第一次提交
* 实现功能：
* 实现用户注册，用户信息入库（redis）

---------------------------------------------------
* // 2019-11-21 14：35 
* // user:xy
* // 第二次提交
* 实现功能：
* 用户登录成功显示当前在线用户
* 关键点：
* 在server维护一个map，field为用户Id，用户登录成功就把登录Id传到这个map作为field
* LoginResMsg新增一个UsersId，slice类型，用户登录成功遍历server维护的map将field追加到LoginResMsg.UsersId中
* client收到server回复将Msg反序列化得到LoginResMsg，再遍历其UsersId字段得到当前登录成功用户
