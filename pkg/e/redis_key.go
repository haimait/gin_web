package e

/*定义redis中的常量*/
const (
	REDIS_WECHAT_FORMID_BY_USER string = "WeChat_FormId_List_%d" //存储用户的所有formId
	REDIS_WECHAT_TASKID_LIST    string = "WeChat_Task_List"      //存储需要发送消息的taskid列表
	REDIS_LINKBOOK              string = "LinkBook"              //用户的通讯录数据
	REDIS_LINKBOOK_WITH_USER    string = "LinkBook_%d"           //用户的通讯录数据
	REDIS_LOCK_TASK             string = "Lock_User_Task_%d"     //缓存锁定用户每个任务只能提交一次

	REDIS_RED_PACKET_WITH_ID     string = "Red_Packet_%d"     //测试
	REDIS_RED_PACKET_LOG_WITH_ID string = "Red_Packet_Log_%d" //测试

	RED_PACKET_FROM_USER_INFO_KEY_WITH_ORDER_CODE string = "red_packet_form_user_info_order_code_%s"  //发红包人的信息
	RED_PACKET_POOL_LIST_KEY_WITH_ORDER_CODE      string = "red_packet_pool_list_key_order_code_%s"   //红包池key
	RED_PACKET_DETAIL_LIST_KEY_WITH_ORDER_CODE    string = "red_packet_detail_list_key_order_code_%s" //红包详情key
	RED_PACKET_HOLD_LIST_KEY_WITH_ORDER_CODE      string = "red_packet_hold_list_key_order_code_%s"   //已经抢红包用户的key
	RED_PACKET_HOLD_LIST_KEY_WITH_USER_ID         string = "red_packet_hold_list_key_user_id_%d"      //已经抢红包用户的ID key

	//用户相关
	REDIS_USER_DETIAL string = "User_Detail_%d" //用户的各种详细信息存储到缓存中

	//用户职业,兴趣
	REDIS_HOBBY_LIST      string = "HOBBY_LIST"      //用户兴趣列表
	REDIS_PROFESSION_LIST string = "PROFESSION_LIST" //用户职业信息列表
)
