package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	// 用户相关
	ERROR_AUTH_CHECK_TOKEN_FAIL:                      "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:                   "Token已超时",
	ERROR_AUTH_TOKEN:                                 "Token生成失败",
	ERROR_AUTH:                                       "Token错误",
	ERROR_AUTH_ADD_WRONG:                             "用户信息添加失败",
	ERROR_AUTH_WECHAT_LOGIN_FAILED:                   "用户微信小程序登陆失败",
	ERROR_AUTH_MOBILE_LOGIN_FAILED:                   "用户手机号登陆登陆失败",
	ERROR_AUTH_WECHAT_GET_INFO_FAILED:                "用户微信小程序获取信息失败",
	ERROR_AUTH_SAVE_COMPANY_INFO_ERR:                 "保存用户公司信息失败",
	ERROR_AUTH_COMPANY_INFO_NULL_ERR:                 "用户公司信息不存在",
	ERROR_AUTH_COMPANY_LAST_INFO_CANNOT_DELETE_ERR:   "公司信息至少一条,不能删除",
	ERROR_AUTH_EDUCATION_INFO_NULL_ERR:               "用户学历信息不存在",
	ERROR_AUTH_EDUCATION_LAST_INFO_CANNOT_DELETE_ERR: "学历信息至少一条,不能删除",
	ERROR_AUTH_SAVE_USR_INFO_ERR:                     "保存用户信息失败",
	ERROR_AUTH_SAVE_USR_ADDRESS_INFO_ERR:             "保存用户地址信息失败",
	ERROR_AUTH_ADD_WXINFO_WRONG:                      "添加用户微信信息失败",
	ERROR_AUTH_WXINFO_DIFF_WRONG:                     "校验用户微信信息失败",
	ERROR_LOGIN_TYPE:                                 "未授权的登陆类型",
	ERROR_AUTH_SAVE_USR_WALLET_INFO_ERR:              "保存用户钱包信息失败",
	ERROR_USER_TO_OTHER_IMPRESSION_NULL:              "该用户未对目标用户打印象标签",
	ERROR_AUTH_SAVE_IMPRESSION_INFO_ERR:              "用户保存印象失败",
	ERROR_AUTH_SAVE_EDUCATION_INFO_ERR:               "用户保存学历失败",
	ERROR_AUTH_SAVE_LGLT_INFO_ERR:                    "用户保存经纬度失败",
	ERROR_AUTH_LOCK_PHOTO_ERR:                        "锁相册失败",
	ERROR_AUTH_SAVE_USR_REPUTATION_INFO_ERR:          "保存声誉记录失败",
	ERROR_AUTH_SAVE_USR_Effect_INFO_ERR:              "保存影响记录失败",
	ERROR_AUTH_UNLOCK_PHOTO_ERR:                      "解相册失败",
	ERROR_AUTH_LOCKED_PHOTO_ERR:                      "相册已锁状态",
	ERROR_AUTH_UNLOCKED_PHOTO_ERR:                    "相册未锁状态",
	ERROR_AUTH_HASUNLOCKEDAUTH_PHOTO_ERR:             "已拥有查看相册权限",
	ERROR_AUTH_NOHASUNLOCKEDAUTH_PHOTO_ERR:           "无权限",
	ERROR_AUTH_INVITE_ERR:                            "该用户已经被邀请过",
	ERROR_TASK_SAVE_ERR:                              "发起认证信息失败",
	ERROR_TASK_MUNBER_ERR:                            "请先完成三项信息认证",
	ERROR_HASAUTH_ERR:                                "已经购买过查看权限",
	ERROR_ONUSERINFO_ERR:                             "无用户信息",
	ERROR_ONUSERWETCHATCODEINFOINFO_ERR:              "无用户微信号信息",
	ERROR_BUY_AUTH_ERR:                               "购买查看权限失败",
	ERROR_PHONE_EXSTS_ERR:                            "手机号已经存在",
	ERROR_CREATE_APPLAEUSER_FAILED:                   "注册苹果用户信息失败",
	ERROR_UPDATE_APPLAEUSER_FAILED:                   "更新苹果用户信息失败",

	ERROR_CREATE_RED_PACKET_FAILED:         "创建红包记录失败",
	ERROR_CREATE_RED_PACKET_LOG_FAILED:     "创建收红包记录日志失败",
	ERROR_NOEXSTS_RED_PACKET_FAILED:        "红包记录不存在",
	ERROR_RED_PACKET_GOLDCOIN_NOHAS_FAILED: "金币已经领完",
	ERROR_HAVE_RECEIVED_PACKET_FAILED:      "已经领过红包",
	ERROR_HAVE_RECEIVED_PACKET_END_FAILED:  "金币已经被人领完了",

	//广告
	ERROR_AD_UP_TO_FOUR:  "最多只能添加4四广告位",
	ERROR_AD_DELETE_FAIL: "群删除失败",
	ERROR_AD_NO_EXISTS:   "群不存在",

	// 验证码相关
	ERROR_CODE_SEND_FAILED:   "验证码发送失败",
	ERROR_CODE_VALID_FAILED:  "验证码不正确或已经失效",
	ERROR_CODE_LENGTH_FAILED: "验证码长度不正确",
	ERROR_CODE_MOBILE_FAILED: "手机号不正确",

	// 好友相关
	ERROR_USER_RELATION_IS_EXIST:  "你和用户已经是好友关系",
	ERROR_USER_RELATION_FAILD:     "添加好友关系失败",
	ERROR_USER_AREADEY_FRIEND:     "已经是好友了",
	ERROR_ADD_USER_RELATION_FAILD: "申请加好友失败",
	// 提现
	ERROR_USER_NOT_BIND_WECHAT: "您未绑定微信",
	ERROR_USER_SIGNED_DAY_NUM:  "您的签到天数小于系统要求",
	ERROR_USER_INVITE_PER_NUM:  "您的邀请好友数小于系统要求",
	ERROR_USER_NOT_ENOUGH_COIN: "您的金币余额不足",
	ERROR_USER_CASH_OUT_ERR:    "提现失败",

	// 上传文件
	ERROR_SAVE_SERVER_ERR: "文件上传服务器失败",
	ERROR_SAVE_OSS_ERR:    "文件上传云端失败",

	//认证相关
	ERROR_AUTH_TASK_ADD_FAILED: "发启认证任务失败",

	//充值购买相关
	ERROR_NOT_ENOUGH_COIN:    "金币不足，无法购买",
	ERROR_NOT_RIGHT_USER:     "当前用户与订单购买人不一致",
	ERROR_COMPLETE_ORDER:     "购买失败",
	ERROR_ORDER_STATUS:       "订单状态不正确",
	ERROR_ADDGOLDCOIN_STATUS: "添加金币失败",
	ERROR_CREATE_ORDER:       "订单创建失败",
	ERROR_PAY_ORDER:          "订单支付失败",
	ERROR_ORDER_NOT_EXIST:    "订单不存在",
	ERROR_TO_FASK:            "请慢一点",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
