package global

// table 相关
const (
	LotteryCount   = 20 //开机结果记录100局 保存
	AreaCount      = 8   //  下注区域
	WinRecordCount = 20  // 用户赢的记录20局
	PokerCount     = 1   // 奔驰宝马
	ListSize       = 2   //富豪榜/神算子
)

// 游戏状态
const (
	GameStatusJetton  = iota //下注状态
	GameStatusLottery        //开奖状态
)

// 错误码
const (
	ServerError = 500

	// 登录
	LoginError      = 1001
	LoginTokenError = 1002

	//坐下
	SitDownError1 = 2001
	SitDownError2 = 2002
	SitDownError3 = 2003

	// 下注
	JettonError1 = 3001
	JettonError2 = 3002

	//起立
	StandUpError1 = 4001
	StandUpError2 = 4002

	//用户列表
	UserListError1 = 5001
	UserListError2 = 5001
)

//区域赔率
var AreaMultiple = [AreaCount]float32{
	40, 5, 30, 5, 20, 5, 10, 5,
}
