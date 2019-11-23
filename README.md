* // 2019-11-20 15：40 
* // user:xy
* // 第一次提交
* 实现功能：
* 实现用户注册，用户信息入库（redis）

---
* // 2019-11-21 14：35 
* // user:xy
* // 第二次提交
* 实现功能：
* 用户登录成功显示当前在线用户
* 关键点：
* 在server维护一个map，field为用户Id，用户登录成功就把登录Id传到这个map作为field
* LoginResMsg新增一个UsersId，slice类型，用户登录成功遍历server维护的map将field追加到LoginResMsg.UsersId中
* client收到server回复将Msg反序列化得到LoginResMsg，再遍历其UsersId字段得到当前登录成功用户

---
* // 2019-11-22 10：28
* // user:xy
* // 第三次提交
* 实现功能：
* 登录成功查看当前在线用户列表
* 关键点：
* 在client维护一个map，在message新建一个struct存放用户id和状态，在LoginResMsg里新增一个字段Users用来存放登录用户id，用户登录成功后反序列化并遍历该字* 段，将登录成功用户Id存入map中，遍历这个map得到登录成功的用户id

---
* // 2019-11-22 15:21
* // user:xy
* // 第四次提交
* 实现功能：
* 客户端发送群发消息给服务端（仅发送未作处理
* 关键点：
* 新增CurUser结构体存放Conn连接和message.User字段，message里新增SmsMsg结构体，作为发送消息结构体，存放content消息体和User实例，编写SendGroupMsg方法，参数为content，调用该方法传入要发送的消息，序列化处理，调用utils里的WritePkg发送序列化后的方法，完成消息发送到服务器

---
* // 2019-11-22 18：36
* // user:xy
* // 第五次提交
* 实现功能：
* 登录成功的用户能够群发消息给所有在线用户（自己除外）
* 关键点：
* client端发送消息给server已经完成，server接收到群发消息，在server/process2/smsProcess.go新增处理群发消息方法，client收到消息后反序列化并将消息打印出来，完成群发消息

---
* // 2019-11-23 15：02
* // user:xy
* // 第六次提交
* 实现功能：
* 实现点对点聊天（私聊）
* 关键点：
* 和群聊逻辑基本一样，私聊请求中新增一个对方Id字段，服务器接收到请求后进行字段遍历，把消息发送到对应Id用户连接
