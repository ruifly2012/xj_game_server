package msg

//登陆并且坐下
type Game_C_LoginDown struct {
	Token     string `json:"token"`
	MachineID string `json:"machine_id"`
	TableID   int32  `json:"table_id"`
	ChairID   int32  `json:"chair_id"`
}

//登陆消息
type Game_C_TokenLogin struct {
	Token     string `json:"token"`
	MachineID string `json:"machine_id"`
}

//机器人登陆
type Game_C_RobotLogin struct {
	UserID  int32 `json:"user_id"`
	BatchID int32 `json:"batch_id"`
}

//用户坐下
type Game_C_UserSitDown struct {
	TableID int32 `json:"table_id"`
	ChairID int32 `json:"chair_id"`
}

//用户起立
type Game_C_UserStandUp struct {
}

//用户下注
type Game_C_UserJetton struct {
	JettonArea  int32   `json:"jetton_area"`
	JettonScore float32 `json:"jetton_score"`
}

//获取用户列表 分页
type Game_C_UserList struct {
	Page int32 `json:"page"`
	Size int32 `json:"size"`
}

//请求失败
type Game_S_ReqlyFail struct {
	ErrorCode int32  `json:"errno"`
	ErrorMsg  string `json:"errmsg"`
}

//登陆成功
type Game_S_LoginSuccess struct {
	Data            []*Game_S_MsgHallHistory `json:"data"`
	Status          int32                    `json:"status"`
	TableID         int32                    `json:"table_id"`
	GameJettonTime  int32                    `json:"game_jetton_time"`
	GameLotteryTime int32                    `json:"game_lottery_time"`
}

//大厅场景
type Game_S_MsgHallHistory struct {
	TableID         int32       `json:"table_id"`
	GameJettonTime  int32       `json:"game_getton_time"`
	GameLotteryTime int32       `json:"game_lottery_time"`
	JettonList      []float32   `json:"jetton_list"`
	LotteryRecord   [][]float32 `json:"lottery_record"`
	GameStatus      int32       `json:"game_status"`
	SceneStartTime  int32       `son:"scene_start_time"`
	UserCount       int32       `json:"user_count"`
}

//开始游戏消息
type Game_S_GameStart struct {
	Data []*Game_S_User `json:"data"`
}

//user
type Game_S_User struct {
	UserID       int32   `json:"user_id"`
	NikeName     string  `json:"nike_name"`
	UserGold     float32 `json:"user_gold"`
	UserDiamond  float32 `json:"user_diamond"`
	MemberOrder  int32   `json:"member_order"`
	HeadImageUrl string  `json:"head_image_url"`
	FaceID       int32   `json:"face_id"`
	RoleID       int32   `json:"role_id"`
	SuitID       int32   `json:"suit_id"`
	PhotoFrameID int32   `json:"photo_frame_id"`
	TableID      int32   `json:"table_id"`
	ChairID      int32   `json:"chair_id"`
	Status       int32   `json:"status"`
	Gender       int32   `json:"gender"`
}

//结束游戏消息
type Game_S_GameConclude struct {
	LotteryPoker  []*Game_S_LotteryPoker `json:"lottery_poker"`
	WinArea       []float32              `json:"win_area"`
	UserListLoss  map[int32]float32      `json:"user_list_loss"`
	UserListMoney map[int32]float32      `json:"user_list_money"`
	ThisMoney     float32                `json:"this_money"`
	UserLoss      float32                `json:"user_loss"`
}

type Game_S_LotteryPoker struct {
	PokerCode    int32   `json:"poker_code"`
	PokerType    int32   `json:"poker_type"`
	LotteryPoker []int32 `json:"lottery_poker"`
}

//下注场景消息
type Game_S_JettonScene struct {
	SceneStartTime int32                `json:"scene_start_time"`
	UserChairID    int32                `json:"user_chair_id"`
	UserList       []*Game_S_User       `json:"user_list"`
	LotteryRecord  [][]float32          `json:"lottery_record"`
	UserArraJetton []*Game_C_AreaJetton `json:"user_arra_jetton"`
	UserJetton     []*Game_C_AreaJetton `json:"user_jetton"`
	RecordID       string               `json:"record_id"` // 局号
}

// 下注结构体
type Game_C_AreaJetton struct {
	Area   int32   `json:"area"`
	Jetton float32 `json:"jetton"`
}

//游戏结束大厅场景消息
type Game_S_Hall struct {
	TableID       int32     `json:"table_id"`
	LotteryRecord []float32 `json:"lottery_record"`
	UserCount     int32     `json:"user_count"`
}

//当前下注状况,每个区域,每个玩家的下注情况
type Game_S_AreaJetton struct {
	TableID        int32                `json:"table_id"`
	UserArraJetton []*Game_C_AreaJetton `json:"user_arra_jetton"`
	UserCount      int32                `json:"user_count"`
}

//开奖场景消息
type Game_S_LotteryScene struct {
	SceneStartTime int32                  `json:"scene_start_time"`
	UserChairID    int32                  `json:"user_chair_id"`
	UserList       []*Game_S_User         `json:"user_list"`
	LotteryRecord  [][]float32            `json:"lottery_record"`
	UserArraJetton []*Game_C_AreaJetton   `json:"user_arra_jetton"`
	LotteryPoker   []*Game_S_LotteryPoker `json:"lottery_poker"`
	WinArea        []float32              `json:"win_area"`
	UserJetton     []*Game_C_AreaJetton   `json:"user_jetton"`
}

//坐下通知消息
type Game_S_SitDownNotify struct {
	UserChairID int32        `json:"user_chair_id"`
	User        *Game_S_User `json:"user"`
}

//起立通知消息
type Game_S_StandUpNotify struct {
	ChairID int32 `json:"chair_id"`
}

//掉线通知消息
type Game_S_OffLineNotify struct {
	ChairID int32 `json:"chair_id"`
}

//上线通知消息
type Game_S_OnLineNotify struct {
	ChairID int32 `json:"chair_id"`
}

// 下注通知
type Game_S_UserJetton struct {
	ChairID     int32   `json:"chair_id"`
	JettonArea  int32   `json:"jetton_area"`
	JettonScore float32 `json:"jetton_score"`
	UserScore   float32 `json:"user_score"`
}

// 返回用户列表数据
type Game_S_UserList struct {
	Data []*Game_S_TempUser `json:"data"`
}

type Game_S_TempUser struct {
	User        *Game_S_User `json:"user"`
	TotalJetton float32      `json:"total_jetton"`
	TotalWin    int32        `json:"total_win"`
}
