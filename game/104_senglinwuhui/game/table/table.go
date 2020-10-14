package table

import (
	"fmt"
	"math"
	"net"
	"sort"
	"sync"
	"xj_game_server/game/104_senglinwuhui/conf"
	"xj_game_server/game/104_senglinwuhui/game/logic"
	"xj_game_server/game/104_senglinwuhui/global"
	"xj_game_server/game/104_senglinwuhui/model"
	"xj_game_server/game/104_senglinwuhui/msg"
	"xj_game_server/game/public/common"
	"xj_game_server/game/public/mysql"
	"xj_game_server/game/public/redis"
	"xj_game_server/game/public/store"
	"xj_game_server/game/public/user"
	"xj_game_server/util/leaf/log"

	rand "xj_game_server/util/leaf/util"

	"time"
)

/*
森林舞会
*/
var List = make([]*Item, 0)

//初始化桌子
func OnInit() {
	//初始化桌子
	for i := int32(0); i < store.GameControl.GetGameInfo().TableCount; i++ {
		temp := &Item{
			tableID:              i,
			chair:                &model.Chair{},
			gameStatus:           0,
			sceneStartTime:       0,
			userList:             sync.Map{},
			userCount:            0,
			androidCount:         0,
			randomColor:          make([]int32, 0),
			lotteryRecord:        make([]int32, 0),
			lotterySpecialRecord: make([]int32, 0),
			userListJetton:       sync.Map{},
			userListJettons:      []sync.Map{},
			userListAreaJetton:   sync.Map{},
			userListWinRecord:    sync.Map{},
			lotteryPoker:         make([]int32, 0),
			winArea:              make([]bool, 0),
			userListLoss:         make(map[int32]float32, 0),
			userTax:              make(map[int32]float32, 0),
			systemScore:          0,
			specialUserList:      make([]*msg.Game_S_User, 0),
			areaJettions:         make([]*msg.Game_C_AreaJetton, 0),
			cronTimer:            &time.Timer{},
		}
		temp.chair.OnInit(store.GameControl.GetGameInfo().ChairCount)
		List = append(List, temp)
		go List[i].sendUserAreaLottery()
		time.Sleep(1 * time.Second)
		go List[i].onEventGameStart()
	}
}

type Item struct {
	tableID              int32                    //桌子号
	chair                *model.Chair             //椅子
	gameStatus           int32                    //游戏状态
	sceneStartTime       int64                    //场景开始时间
	userList             sync.Map                 //玩家列表 座位号-->uid map[int32]int32
	userCount            int32                    //玩家数量
	androidCount         int32                    //机器人数量
	randomColor          []int32                  //开奖记录 随机颜色
	colorIndex           int32                    //开奖记录 颜色下标
	lotteryRecord        []int32                  //开奖记录 global.LotteryCount
	lotterySpecialRecord []int32                  //特殊开奖记录 global.LotteryCount
	userListJetton       sync.Map                 //玩家总下注 1局 座位号 map[int32]float32
	userListJettons      []sync.Map               //玩家总下注 20局 [20]userListJetton
	userListAreaJetton   sync.Map                 //玩家每个区域下注信息 座位号 map[int32][global.AreaCount]float32
	userListWinRecord    sync.Map                 //玩家输赢记录 --->座位号map[int32][global.WinRecordCount]bool
	lotteryPoker         []int32                  //开奖颜色和动物global.PokerCount
	lotterySpecial       int32                    //特殊开奖标识
	winArea              []bool                   //输赢区域global.AreaCount
	userListLoss         map[int32]float32        //用户盈亏 座位号
	userTax              map[int32]float32        //用户税收
	systemScore          float32                  //系统盈亏
	specialUserList      []*msg.Game_S_User       //富豪榜/神算子
	areaJettions         []*msg.Game_C_AreaJetton //每个区域的下注数
	cronTimer            *time.Timer              //定时器下注
	drawID               string                   // 游戏记录id
	roundOrder           string
}

// 定时发送发送每个区域每个人的总下注
func (it *Item) sendUserAreaLottery() {
	for {

		it.cronTimer = time.NewTimer(time.Second * 1)
		select {
		case <-it.cronTimer.C:
			if it.gameStatus == global.GameStatusJetton {

				var areaJettions = make([]*msg.Game_C_AreaJetton, 0)

				for i := 0; i < global.AreaCount; i++ {
					var tempArearJettion msg.Game_C_AreaJetton
					var tempJetton float32

					it.userListAreaJetton.Range(func(chairID, value interface{}) bool {

						tempJetton += value.([global.AreaCount]float32)[i]
						return true
					})
					tempArearJettion = msg.Game_C_AreaJetton{
						Area:   int32(i),
						Jetton: tempJetton,
					}

					areaJettions = append(areaJettions, &tempArearJettion)
				}
				it.areaJettions = areaJettions
				// 发送每个区域每个人的总下注
				//data, _ := json.Marshal(model.AreaJettonMsg{
				//	TableID:     it.tableID,
				//	AreaJettons: areaJettions,
				//})
				it.sendAllUser(&msg.Game_S_AreaJetton{
					//TableID:        it.tableID,
					UserArraJetton: areaJettions,
					//UserCount:      it.userCount,
				})
			}
		}
	}
}

//是否下注
func (it *Item) IsUserBet(user *user.Item) bool {
	load, ok := it.userListJetton.Load(user.ChairID)
	if ok && it.gameStatus == global.GameStatusJetton {
		if load.(float32) > 0 {
			return true
		}
	}
	return false
}

func (it *Item) GetTableID() int32 {
	return it.tableID
}

func (it *Item) GetLotteryRecord() []int32 {
	return it.lotteryRecord
}

func (it *Item) GetGameStatus() int32 {
	return it.gameStatus
}
func (it *Item) GetSceneStartTime() int64 {
	return it.sceneStartTime
}
func (it *Item) GetUserCount() int32 {
	return it.userCount
}

// 游戏大厅场景发送 登录或者游戏结束发送
//func (it *Item) SendLotteryRecord(args ...interface{}) {
//	userItem := args[0].(*user.Item)
//	userItem.WriteMsg(&msg.Game_S_Hall{
//		TableID:       it.tableID,
//		LotteryRecord: it.lotteryRecord[len(it.lotteryRecord)-1:][0],
//		UserCount:     it.userCount,
//		RandomColor:   it.randomColor,
//	})
//}

//开始游戏
func (it *Item) onEventGameStart() {
	//清除上局的数据
	it.userListAreaJetton = sync.Map{}
	it.userListJetton = sync.Map{}
	//设置游戏状态
	it.gameStatus = global.GameStatusJetton
	//随机颜色
	it.randomColor = logic.RandomColor()
	it.cronTimer.Reset(300 * time.Microsecond)
	it.sceneStartTime = time.Now().Unix()
	// 计算神算子或富豪榜
	it.specialUser()
	//发送开始游戏给所有用户
	var colorIndex = make([]int32, 0)
	for _, v := range it.randomColor {
		x1 := v % 0x0f
		colorIndex = append(colorIndex, x1)
	}
	it.sendAllUser(&msg.Game_S_GameStart{
		//Data:    it.specialUserList,
		GemList: colorIndex,
		//RandomColor:    it.randomColor,
		//SceneStartTime: it.sceneStartTime,
	})

	// 记录总开局数 1天
	key := fmt.Sprintf("kind-%d:game-%d:table-%d:", conf.GetServer().KindID, conf.GetServer().GameID, it.tableID)
	err := redis.GameClient.Client.Incr(key).Err()
	//设置 超时时间 每天凌晨到期
	day := rand.EndOfDay(time.Now()).Sub(time.Now())
	err = redis.GameClient.Client.Expire(key, day).Err()
	if err != nil {
		_ = log.Logger.Errorf("write total game err %v", err)
	}
	fmt.Printf("[森林舞会]桌子%d开始游戏,游戏状态%d,游戏人数:%d,神算子富豪榜%d,开奖颜色:%v\n", it.tableID, it.gameStatus, it.userCount, len(it.specialUserList), it.randomColor)

	//下注定时器
	t := time.NewTimer(time.Second * time.Duration(conf.GetServer().GameJettonTime))
	select {
	case <-t.C:
		it.onEventGameConclude()
	}
}

//神算子和富豪榜
func (it *Item) specialUser() {
	var richUser model.Special         //富豪
	var operator model.SpecialOperator //神算子
	it.specialUserList = make([]*msg.Game_S_User, 0)
	it.userListWinRecord.Range(func(key, value interface{}) bool {
		var temp *model.Operator
		uid, ok := it.userList.Load(key.(int32))
		if !ok {
			//_ = log.Logger.Errorf("specialUser it.userList err %d",key.(int32))
			return true
		}
		userItem, ok := user.List.Load(uid.(int32))
		if !ok {
			_ = log.Logger.Errorf("specialUser user.List err %d", uid.(int32))
			return true
		}
		for _, v := range value.([global.WinRecordCount]bool) {
			var i = 0
			if v {
				i++
				temp = &model.Operator{
					User: &msg.Game_S_User{
						UserID:       userItem.(*user.Item).Info.UserID,
						NikeName:     userItem.(*user.Item).Info.NikeName,
						UserDiamond:  userItem.(*user.Item).Info.UserDiamond,
						HeadImageUrl: userItem.(*user.Item).Info.HeadImageUrl,
						ChairID:      userItem.(*user.Item).Info.ChairID,
					},
					Count: i,
				}
			}
		}
		if temp != nil {
			operator = append(operator, temp)
		}
		//富豪榜
		richUser = append(richUser, &msg.Game_S_User{
			UserID:       userItem.(*user.Item).Info.UserID,
			NikeName:     userItem.(*user.Item).Info.NikeName,
			UserDiamond:  userItem.(*user.Item).Info.UserDiamond,
			HeadImageUrl: userItem.(*user.Item).Info.HeadImageUrl,
			ChairID:      userItem.(*user.Item).Info.ChairID,
		})
		return true
	})
	// 前2名
	sort.Sort(richUser)
	// 前2名
	sort.Sort(operator)
	if len(richUser) >= global.ListSize {
		it.specialUserList = append(it.specialUserList, richUser[0:global.ListSize]...)
	}
	if len(operator) >= global.ListSize {
		for i := 0; i < global.ListSize; i++ {
			it.specialUserList = append(it.specialUserList, operator[i].User)
		}
	}

}

//结束游戏
func (it *Item) onEventGameConclude() {
	it.cronTimer.Reset(time.Second*time.Duration(conf.GetServer().GameLotteryTime) + 5)
	it.gameStatus = global.GameStatusLottery
	it.sceneStartTime = time.Now().Unix()
	//randNumber := rand.RandInterval(0, 101)
	//var min float32
	//var count int32
	////开奖
	//for {
	//	// 发牌
	//	it.lotteryPoker, it.lotterySpecial, it.colorIndex = logic.Client.DispatchTableCard(it.randomColor)
	//	// 获取赢的区域
	//	it.winArea = logic.Client.GetWinArea(it.lotteryPoker, it.lotterySpecial)
	//	// 系统盈亏和用户盈亏
	//	var userWinSum float32
	//	it.systemScore, it.userListLoss, it.userTax, userWinSum = logic.Client.GetSystemLoss(it.winArea, it.userListAreaJetton, it.userList, it.lotterySpecial)
	//
	//	if float32(randNumber) < store.GameControl.GetUserWinRate() { //用户赢
	//		// 系统库存不够的时候用户输
	//		if store.GameControl.GetStore() < 0 {
	//			break
	//		}
	//		if min > userWinSum {
	//			min = userWinSum
	//			count++
	//		}
	//		if count == 5 {
	//			break
	//		}
	//		if userWinSum >= 0 {
	//			break
	//		}
	//	} else { //用户输
	//		if userWinSum <= 0 {
	//			break
	//		}
	//	}
	//}

	var winAreas = make([]store.WinAreas, 0)

	for i := 0; i < 5; i++ {
		var w = store.WinAreas{}

		// 发牌
		w.LotteryPoker, w.LotterySpecial, w.ColorIndex = logic.Client.DispatchTableCard(it.randomColor)
		// 获取赢的区域
		w.WinArea = logic.Client.GetWinArea(w.LotteryPoker, w.LotterySpecial)
		// 系统盈亏和用户盈亏
		w.SystemScore, w.UserListLoss, w.UserTax, _ = logic.Client.GetSystemLoss(w.WinArea, it.userListAreaJetton, it.userList, w.LotterySpecial)

		w.Stores = store.GameControl.GetStore1() + w.SystemScore

		winAreas = append(winAreas, w)
	}

	sort.Sort(store.IntSlice(winAreas))

	it.lotteryPoker = winAreas[0].LotteryPoker
	it.winArea = winAreas[0].WinArea
	it.systemScore = winAreas[0].SystemScore
	it.userListLoss = winAreas[0].UserListLoss
	it.userTax = winAreas[0].UserTax
	it.lotterySpecial = winAreas[0].LotterySpecial
	it.colorIndex = winAreas[0].ColorIndex

	// 更新系统库存
	store.GameControl.ChangeStore(it.systemScore)

	//更新开奖记录
	it.onUpdateWins()
	//更新用户输赢记录
	it.onUpdateUsersWins()
	//记录游戏记录
	it.onWriteGameRecord()
	//用户写分
	it.onWriteGameScore()

	// 热更数据
	it.onUpdateAgentData()

	var gameConclude msg.Game_S_GameConclude

	var special int32
	switch it.lotterySpecial {
	case 0x30:
		special = 0
	case 0x40:
		special = 1
	case 0x50:
		special = 2
	}
	gameConclude.LotterySpecial = special
	//gameConclude.ColorIndex = it.colorIndex
	//logic.GetAnimalName(it.lotteryPoker[1])  //动物
	//gameConclude.Animal = it.lotteryPoker[1] //动物

	//gameConclude.LotteryPoker = logic.GetWinAreaId(it.winArea)
	//gameConclude.LotterySpecial = it.lotterySpecial
	//gameConclude.WinArea = it.winArea
	//gameConclude.ColorIndex = it.colorIndex
	// 发自己和特殊的四人的记录
	var newListLoss = make(map[int32]float32)
	var new2ListLoss = make(map[int32]float32)
	//神算子/富豪
	for _, v := range it.specialUserList {

		uid, ok := it.userList.Load(v.ChairID)
		if ok {
			value, ok := user.List.Load(uid)
			if ok {
				if store.GameControl.GetGameInfo().DeductionsType == 0 {
					newListLoss[v.ChairID] = value.(*user.Item).UserGold
				} else {
					newListLoss[v.ChairID] = value.(*user.Item).UserDiamond
				}
			}
			//新盈亏 加上下注数
			load, o := it.userListJetton.Load(v.ChairID)
			if o {
				new2ListLoss[v.ChairID] = it.userListLoss[v.ChairID] + load.(float32)
			}
		}

	}

	it.userList.Range(func(chairID, uid interface{}) bool {
		value, ok := user.List.Load(uid)
		if ok {
			mySelf := value.(*user.Item)
			// 如果游戏结束的时候用户在离线状态 解锁用户
			if mySelf.Status == user.StatusOffline {
				// 起立 强制退出
				it.OnActionUserStandUp(value.(*user.Item), true)
				// map 中移除
				user.List.Delete(uid)
				return true
			}
			_, ok := newListLoss[mySelf.ChairID]
			if !ok {
				if store.GameControl.GetGameInfo().DeductionsType == 0 {
					newListLoss[mySelf.ChairID] = value.(*user.Item).UserGold
				} else {
					newListLoss[mySelf.ChairID] = value.(*user.Item).UserDiamond
				}
			}
			//新盈亏 加上下注数
			load, o := it.userListJetton.Load(mySelf.ChairID)
			if o {
				new2ListLoss[mySelf.ChairID] = it.userListLoss[mySelf.ChairID] + load.(float32)
			}

			gameConclude.AnimalIndex = it.lotteryPoker[1]
			gameConclude.GemIndex = it.colorIndex
			gameConclude.Money = new2ListLoss[mySelf.ChairID]
			// 重新赋值 减去税收
			//gameConclude.UserListMoney = newListLoss
			//gameConclude.UserListLoss = new2ListLoss
			value.(*user.Item).WriteMsg(&gameConclude)
			if !ok {
				delete(newListLoss, mySelf.ChairID)
			}
		}
		return true
	})

	//user.List.Range(func(key, value interface{}) bool {
	//	// 空闲状态
	//	if value.(*user.Item).Status == user.StatusFree {
	//		//发送空闲状态数据
	//		//it.SendLotteryRecord(value)
	//		return true
	//	}
	//	return true
	//})
	//看是否记录了20局
	if len(it.userListJettons) == global.WinRecordCount {
		it.userListJettons = it.userListJettons[1:]
		it.userListJettons = append(it.userListJettons, it.userListJetton)
	} else {
		it.userListJettons = append(it.userListJettons, it.userListJetton)
	}

	fmt.Printf("[森林舞会]桌子%d结束游戏,游戏状态%d\n,系统损耗%f\n,用户损耗%v\n开奖号码%s\n,特殊开奖列表%v\n,开奖列表% X\n,中奖区域%v\n,开奖颜色：%v,开奖颜色下标：%d,开奖颜色下标：%s\n", it.tableID, it.gameStatus, it.systemScore, it.userListLoss, logic.GetColorName(it.lotteryPoker[0])+"-"+logic.GetAnimalName(it.lotteryPoker[1]), it.lotterySpecialRecord, it.lotteryRecord, it.winArea, it.randomColor, it.colorIndex, logic.GetColorName(it.randomColor[it.colorIndex]))
	//清空上局数据
	it.userListLoss = make(map[int32]float32)
	it.systemScore = 0
	it.lotterySpecial = 0
	it.lotteryPoker = make([]int32, 0)
	it.drawID = ""
	//开奖定时器
	t := time.NewTimer(time.Second * time.Duration(conf.GetServer().GameLotteryTime))
	select {
	case <-t.C:
		it.onEventGameStart()
	}
}

//更新开奖记录
func (it *Item) onUpdateWins() {
	//self.winArea
	//
	if len(it.lotteryRecord) >= global.LotteryCount {
		it.lotteryRecord = it.lotteryRecord[1:]
	}
	for key, v := range it.winArea {
		if v {
			it.lotteryRecord = append(it.lotteryRecord, int32(key))
			break
		}
	}

	if len(it.lotterySpecialRecord) >= global.LotteryCount {
		it.lotterySpecialRecord = it.lotterySpecialRecord[1:]
	}
	it.lotterySpecialRecord = append(it.lotterySpecialRecord, it.lotterySpecial)

}

//更新用户输赢记录
func (it *Item) onUpdateUsersWins() {
	it.userList.Range(func(key, value interface{}) bool {
		// 取出20句的所有用户输赢记录
		v, _ := it.userListWinRecord.LoadOrStore(key.(int32), [global.WinRecordCount]bool{})
		// 迁移以为
		temp := v.([global.WinRecordCount]bool)
		for i := 1; i < len(temp); i++ {
			temp[i], temp[i-1] = temp[i-1], temp[i]
		}
		// 收过税的人肯定是赢了
		temp[len(temp)-1] = it.userListLoss[key.(int32)] > 0
		it.userListWinRecord.Store(key.(int32), temp)

		return true
	})
}

// 热更数据
func (it *Item) onUpdateAgentData() {
	var decPercentValue float32
	var intDataType = 2 //数据类型：1 注册，2 游戏输赢，3 返佣，4 充值，5 兑换，6 领取佣金
	for key, v := range it.userListLoss {
		decAmount := v // 输赢金额
		uid, ok := it.userList.Load(key)
		if !ok {
			_ = log.Logger.Errorf("onUpdateAgentData it.userList err: %d", key)
			return
		}
		userItem, ok := user.List.Load(uid.(int32))
		// 过滤机器人
		if userItem.(*user.Item).BatchID != -1 {
			continue
		}
		mysql.GameClient.WriteAgentData(uid.(int32),
			intDataType,
			decAmount,
			decPercentValue,
		)
	}
}

//用户写分
func (it *Item) onWriteGameScore() {
	//GSP_WriteGameScore
	var endTime = time.Now().Format("2006-01-02 15:04:05")
	for key, v := range it.userListLoss {
		var intWinCount, //胜利盘数
			intLostCount, //失败盘数
			intDrawCount, //和局盘数
			intFleeCount, //逃跑数目
			tintTaskForward int32 // 任务跟进
		uid, ok := it.userList.Load(key)
		// 过滤机器人
		if !ok {
			_ = log.Logger.Errorf("onWriteGameScore it.userList err: %d", key)
			return
		}
		userItem, ok := user.List.Load(uid.(int32))
		if !ok {
			_ = log.Logger.Errorf("onWriteGameScore user.List err:%d", uid.(int32))
			return
		}
		if v > 0 {
			intWinCount = 1
		}
		if v == 0 {
			intDrawCount = 1
		}
		if v < 0 {
			intDrawCount = 1
		}
		if userItem.(*user.Item).Status == user.StatusOffline {
			intFleeCount = 1
		}
		if store.GameControl.GetGameInfo().DeductionsType == 0 {
			userItem.(*user.Item).UserGold += v
		} else {
			if !userItem.(*user.Item).IsRobot() {
				if !redis.GameClient.IsExistsDiamond(userItem.(*user.Item).UserID) {
					scoreInfo, _ := mysql.GetGameScoreInfoByUserId(mysql.GameClient.GetXJGameDB, userItem.(*user.Item).UserID)
					redis.GameClient.SetDiamond(userItem.(*user.Item).UserID, scoreInfo.Diamond)
				}

				userDiamond, err := redis.GameClient.GetDiamond(userItem.(*user.Item).UserID)
				if err != nil {
					log.Logger.Error("GetDiamond err:", err)
				} else {
					var tempScore float32
					v, ok := it.userListJetton.Load(userItem.(*user.Item).ChairID)
					if ok {
						tempScore = v.(float32)
					}
					userItem.(*user.Item).UserDiamond = float32(userDiamond) + tempScore
				}
				var dia = v + it.userTax[key]
				userItem.(*user.Item).Jackpot += dia + float32(math.Abs(float64(store.GameControl.GetGameInfo().UmRevenueRatio*dia)))
			}
			userItem.(*user.Item).UserDiamond += v
			//userItem.(*user.Item).UserDiamond += v
		}
		// 过滤机器人
		if userItem.(*user.Item).BatchID != -1 {
			continue
		}
		randNum := rand.Krand(6, 3)
		it.roundOrder = fmt.Sprintf("%v%v%s", conf.GetServer().GameID, time.Now().Unix(), randNum)

		host, _, _ := net.SplitHostPort(userItem.(*user.Item).Agent.RemoteAddr().String())
		errorCode, errorMsg := mysql.GameClient.WriteUserScore(uid.(int32),
			v,
			store.GameControl.GetGameInfo().DeductionsType,
			it.userTax[key], //  税收写分
			intWinCount,
			intLostCount,
			intDrawCount,
			intFleeCount,
			conf.GetServer().GameJettonTime,
			tintTaskForward,
			store.GameControl.GetGameInfo().KindID,
			store.GameControl.GetGameInfo().GameID,
			v,
			host,
			time.Unix(it.sceneStartTime, 0).Format("2006-01-02 15:04:05"),
			endTime,
			it.drawID,
			userItem.(*user.Item).Jackpot,
			userItem.(*user.Item).UserDiamond,
			it.roundOrder,
		)
		if !userItem.(*user.Item).IsRobot() && store.GameControl.GetGameInfo().DeductionsType == 1 {
			redis.GameClient.SetDiamond(userItem.(*user.Item).UserID, userItem.(*user.Item).UserDiamond)
			redis.GameClient.RegisterRecharge(userItem.(*user.Item).UserID)
		}
		if errorCode != common.StatusOK {
			_ = log.Logger.Errorf(" mysql GSP_WriteGameScore存储过程 %n %s ", errorCode, errorMsg)
			return
		}
	}
}

//游戏记录
func (it *Item) onWriteGameRecord() {
	//GSP_RecordDrawInfo
	//-- intTableID：桌子ID
	//-- intUserCount：用户数量
	//-- intAndroidCount：机器人数量
	//-- decWasteCount：损耗数目
	//-- decResveueCount：税收数目
	//-- timeEnterTime：游戏开始时间
	//-- timeLeaveTime：游戏结束时间
	//-- tintScoreType：金币类型

	var endTime = time.Now().Format("2006-01-02 15:04:05")

	var startTime = time.Unix(it.sceneStartTime, 0).Format("2006-01-02 15:04:05")
	var taxSum float32 = 0
	for _, v := range it.userTax {
		taxSum += v
	}
	errorCode, errorMsg, drawID := mysql.GameClient.WriteGameRecord(
		it.tableID,
		it.userCount,
		it.androidCount,
		it.systemScore,
		taxSum,
		startTime,
		endTime,
		store.GameControl.GetGameInfo().DeductionsType, "",
	)
	if errorCode != common.StatusOK {
		_ = log.Logger.Errorf(" mysql GSP_RecordDrawInfo存储过程报错 %n %s ", errorCode, errorMsg)
		return
	}
	it.drawID = drawID
}

//发送场景
func (it *Item) onEventSendGameScene(args ...interface{}) {
	userItem := args[0].(*user.Item)
	switch it.gameStatus {
	case global.GameStatusJetton: //下注场景消息
		var jettonScene msg.Game_S_JettonScene
		sceneTime := int64(conf.GetServer().GameJettonTime) - (time.Now().Unix() - it.sceneStartTime)
		if sceneTime < 0 {
			sceneTime = 0
		}
		jettonScene.SceneStartTime = float64(sceneTime)
		//所有人
		// 发送每个区域每个人的总下注
		//data, _ := json.Marshal(model.AreaJettonMsg{
		//	TableID:     it.tableID,
		//	AreaJettons: areaJettions,
		//})

		var userJettonList = make([]*msg.JettonInfo, 0)
		load, ok := it.userListAreaJetton.Load(userItem.ChairID)
		if ok {
			for k, v := range load.([global.AreaCount]float32) {
				userJetton := &msg.JettonInfo{
					Area:  int32(k),
					Money: v,
				}
				userJettonList = append(userJettonList, userJetton)
			}
		} else {
			for i := 0; i < global.AreaCount; i++ {
				userJetton := &msg.JettonInfo{
					Area:  int32(i),
					Money: 0,
				}
				userJettonList = append(userJettonList, userJetton)
			}
		}
		jettonScene.UserJetton = userJettonList

		var areaJettions = make([]*msg.JettonInfo, 0)

		for i := 0; i < global.AreaCount; i++ {
			var tempArearJettion msg.JettonInfo
			var tempJetton float32

			it.userListAreaJetton.Range(func(chairID, value interface{}) bool {
				tempJetton += value.([global.AreaCount]float32)[i]
				return true
			})
			tempArearJettion = msg.JettonInfo{
				Area:  int32(i),
				Money: tempJetton,
			}
			areaJettions = append(areaJettions, &tempArearJettion)
		}
		// 发送每个区域每个人的总下注
		jettonScene.UserArraJetton = areaJettions
		var colorIndex = make([]int32, 0)
		for _, v := range it.randomColor {
			x1 := v % 0x0f
			colorIndex = append(colorIndex, x1)
		}
		jettonScene.GemList = colorIndex
		//jettonScene.GemList = it.randomColor
		//jettonScene.UserCount = it.userCount
		//随机颜色
		//jettonScene.RandomColor = it.randomColor
		//data, err := json.Marshal(jettonScene)
		//if err != nil {
		//	_ = log.Logger.Errorf("onEventSendGameScene err %v", err)
		//	userItem.Close()
		//	return
		//}
		userItem.WriteMsg(&jettonScene)
	case global.GameStatusLottery: //开奖场景消息
		var lotteryScene msg.Game_S_ConcludeScene
		sceneTime := int64(conf.GetServer().GameLotteryTime) - (time.Now().Unix() - it.sceneStartTime)
		if sceneTime < 0 {
			sceneTime = 0
		}
		lotteryScene.SurplusTime = float64(sceneTime)

		//lotteryScene.UserArraJetton = it.areaJettions
		var colorIndex = make([]int32, 0)
		for _, v := range it.randomColor {
			x1 := v % 0x0f
			colorIndex = append(colorIndex, x1)
		}
		//var jettonInfo = make([]msg.JettonInfo, 0)
		//for i := 0; i < global.AreaCount; i++ {
		//	var tempArearJettion msg.JettonInfo
		//	var tempJetton float32
		//	it.userListAreaJetton.Range(func(chairID, value interface{}) bool {
		//		tempJetton += value.([global.AreaCount]float32)[i]
		//		return true
		//	})
		//	tempArearJettion = msg.JettonInfo{
		//		Area:  int32(i),
		//		Money: tempJetton,
		//	}
		//	jettonInfo = append(jettonInfo, tempArearJettion)
		//}
		//lotteryScene.UserJetton = jettonInfo

		var userJettonList = make([]*msg.JettonInfo, 0)
		load, ok := it.userListAreaJetton.Load(userItem.ChairID)
		if ok {
			for k, v := range load.([global.AreaCount]float32) {
				userJetton := &msg.JettonInfo{
					Area:  int32(k),
					Money: v,
				}
				userJettonList = append(userJettonList, userJetton)
			}
		} else {
			for i := 0; i < global.AreaCount; i++ {
				userJetton := &msg.JettonInfo{
					Area:  int32(i),
					Money: 0,
				}
				userJettonList = append(userJettonList, userJetton)
			}
		}
		lotteryScene.UserJetton = userJettonList

		var areaJettions = make([]*msg.JettonInfo, 0)

		for i := 0; i < global.AreaCount; i++ {
			var tempArearJettion msg.JettonInfo
			var tempJetton float32

			it.userListAreaJetton.Range(func(chairID, value interface{}) bool {
				tempJetton += value.([global.AreaCount]float32)[i]
				return true
			})
			tempArearJettion = msg.JettonInfo{
				Area:  int32(i),
				Money: tempJetton,
			}
			areaJettions = append(areaJettions, &tempArearJettion)
		}
		// 发送每个区域每个人的总下注
		lotteryScene.UserArraJetton = areaJettions
		//玩家每个区域的下注数
		lotteryScene.GemList = colorIndex
		var gameConclude msg.Game_S_GameConclude
		//fmt.Println("转换之前11", it.lotteryRecord[len(it.lotteryRecord)-1])
		//x1 := it.randomColor[len(it.lotteryRecord)-1] % 0xf0
		//fmt.Println("转换之后", x1)
		//
		//fmt.Println("转换之前", it.lotteryRecord[len(it.lotteryRecord)-1])
		//fmt.Println("转换之后", it.lotteryRecord[len(it.lotteryRecord)-1]%0x0f/10)
		gameConclude.AnimalIndex = it.lotteryRecord[len(it.lotteryRecord)-1] % 4
		//gameConclude.AnimalIndex = it.lotteryPoker[1]
		gameConclude.GemIndex = it.colorIndex
		//fmt.Println("金额2", it.userListLoss)
		gameConclude.Money = 0

		var special int32
		switch it.lotterySpecial {
		case 0x30:
			special = 0
		case 0x40:
			special = 1
		case 0x50:
			special = 2
		}
		gameConclude.LotterySpecial = special
		lotteryScene.Result = gameConclude
		userItem.WriteMsg(&lotteryScene)
	}
}

//用户坐下
func (it *Item) OnActionUserSitDown(args ...interface{}) {
	userItem := args[0].(*user.Item)
	// 检查是否锁定
	lock := mysql.GameClient.IsLock(userItem.UserID)
	if lock {
		_ = log.Logger.Errorf("坐下失败, 上局游戏未结束")

		userItem.WriteMsg(&msg.Game_S_ReqlyFail{
			ErrorCode: global.SitDownError1,
			ErrorMsg:  "坐下失败, 上局游戏未结束",
		})
		userItem.Close()
		return
	}

	//校验是否满人
	if it.userCount >= store.GameControl.GetGameInfo().ChairCount || it.chair.IsFull() {
		_ = log.Logger.Errorf("坐下失败, 房间人数已满!")

		userItem.WriteMsg(&msg.Game_S_ReqlyFail{
			ErrorCode: global.SitDownError2,
			ErrorMsg:  "坐下失败, 房间人数已满!",
		})
		userItem.Close()
		return
	}
	host, _, _ := net.SplitHostPort(userItem.Agent.RemoteAddr().String())
	//strings.Split(userItem.Agent.RemoteAddr().String(), ":")[0]
	// 锁定
	if err := mysql.GameClient.Lock(userItem.UserID, host); err != nil {
		_ = log.Logger.Errorf("锁定用户失败 err %v", err)
		userItem.WriteMsg(&msg.Game_S_ReqlyFail{
			ErrorCode: global.SitDownError3,
			ErrorMsg:  "坐下失败, 服务器繁忙!",
		})
		userItem.Close()
		return
	}
	// 取出椅子
	chair := it.chair.GetChair()
	if chair < 0 {
		_ = log.Logger.Errorf("坐下失败, 房间人数已满!")
		userItem.WriteMsg(&msg.Game_S_ReqlyFail{
			ErrorCode: global.SitDownError2,
			ErrorMsg:  "坐下失败, 房间人数已满!",
		})
		userItem.Close()
		return
	}
	//加入游戏用户列表
	it.userList.Store(chair, userItem.UserID)
	it.userCount++
	if userItem.IsRobot() {
		it.androidCount++
	}
	userItem.SitDown(it.tableID, chair)

	//发送坐下通知给其他玩家
	//it.sendAllUser(&msg.Game_S_SitDownNotify{
	//	UserChairID: chair,
	//	User: &msg.Game_S_User{
	//		UserID:       userItem.GetUserInfo().UserID,
	//		NikeName:     userItem.GetUserInfo().NikeName,
	//		UserDiamond:  userItem.GetUserInfo().UserDiamond,
	//		HeadImageUrl: userItem.GetUserInfo().HeadImageUrl,
	//		ChairID:      userItem.GetUserInfo().ChairID,
	//	},
	//})

	var his = make([]msg.SlwhLottery, 0)
	for k, v := range it.lotteryRecord {
		var h msg.SlwhLottery
		h.Index = v % 0x0f
		var special int32
		switch it.lotterySpecialRecord[k] {
		case 0x30:
			special = 0
		case 0x40:
			special = 1
		case 0x50:
			special = 2
		}
		h.Special = special
		his = append(his, h)
	}
	//h.Special = it.lotterySpecialRecord[k] % 0x0f

	userItem.WriteMsg(&msg.Game_S_LoginSuccess{
		GameJettonTime:  conf.GetServer().GameJettonTime,
		GameLotteryTime: conf.GetServer().GameLotteryTime,
		JettonList:      conf.GetServer().JettonList,
		History:         his,
		RewardRate:      global.AreaMultiple,
		Status:          1,
	})

	//发送场景消息
	it.onEventSendGameScene(userItem)

}

//用户起立
func (it *Item) OnActionUserStandUp(args ...interface{}) {
	userItem := args[0].(*user.Item)
	flag := args[1].(bool)
	if !flag {
		//检测是否已押注
		v, ok := it.userListAreaJetton.Load(userItem.ChairID)
		if ok && it.gameStatus == global.GameStatusJetton {
			for _, v1 := range v.([global.AreaCount]float32) {
				if v1 != 0 {
					userItem.WriteMsg(&msg.Game_S_ReqlyFail{
						ErrorCode: global.StandUpError1,
						ErrorMsg:  "押注状态下不允许退出!",
					})
					return
				}
			}
		}

	}

	//移出游戏用户列表
	oldChairID := userItem.ChairID
	it.userList.Delete(oldChairID)
	it.chair.AddChair(oldChairID)
	it.userCount--
	if userItem.IsRobot() {
		it.androidCount--
	}
	userItem.StandUp()
	//解锁
	if err := mysql.GameClient.UnLock(userItem.UserID); err != nil {
		_ = log.Logger.Errorf("解锁用户失败 err %v", err)
		return
	}

	//删除这个座位的输赢记录
	//delete(it.userListWinRecord, userItem.ChairID)
	it.userListWinRecord.Delete(oldChairID)
	//起立通知桌面上的人通知
	if it.isOnTable(userItem) {
		it.sendAllUser(&msg.Game_S_StandUpNotify{
			ChairID: oldChairID,
		})
	}

}

//用户断线
func (it *Item) OnActionUserOffLine(args ...interface{}) {
	userItem := args[0].(*user.Item)
	//设置用户状态
	userItem.Status = user.StatusOffline
	//断线通知桌面上的人
	if it.isOnTable(userItem) {
		it.sendOtherUser(userItem.UserID, &msg.Game_S_OffLineNotify{
			ChairID: userItem.ChairID,
		})
	}
}

//用户重入
func (it *Item) OnActionUserReconnect(args ...interface{}) {
	userItem := args[0].(*user.Item)
	if userItem.Status == user.StatusOffline {
		//设置用户状态
		userItem.Status = user.StatusPlaying
		//发送上线通知给其他玩家
		if it.isOnTable(userItem) {
			it.sendOtherUser(userItem.UserID, &msg.Game_S_OnLineNotify{
				ChairID: userItem.ChairID,
			})
		}
	}

	lengh := len(it.lotteryRecord)
	if lengh > 15 {
		lengh = 15
	}

	var his = make([]msg.SlwhLottery, 0)
	for k, v := range it.lotteryRecord {
		var h msg.SlwhLottery
		h.Index = v % 0x0f

		var special int32
		switch it.lotterySpecialRecord[k] {
		case 0x30:
			special = 0
		case 0x40:
			special = 1
		case 0x50:
			special = 2
		}
		h.Special = special
		//h.Special = it.lotterySpecialRecord[k] % 0x0f
		his = append(his, h)
	}

	userItem.WriteMsg(&msg.Game_S_LoginSuccess{
		GameJettonTime:  conf.GetServer().GameJettonTime,
		GameLotteryTime: conf.GetServer().GameLotteryTime,
		JettonList:      conf.GetServer().JettonList,
		History:         his,
		RewardRate:      global.AreaMultiple,
		Status:          2,
	})

	//发送场景消息
	it.onEventSendGameScene(userItem)
}

//下注事件
func (it *Item) OnUserPlaceJetton(args ...interface{}) {
	userItem := args[0].(*user.Item)
	m := args[1].(*msg.Game_C_UserJetton)
	//检测数据是否异常
	if m.JettonArea < 0 || m.JettonArea >= global.AreaCount || m.JettonScore <= 0 {
		_ = log.Logger.Errorf("OnUserPlaceJetton err %s", "下注失败, 无效的数据")
		userItem.WriteMsg(&msg.Game_S_ReqlyFail{
			ErrorCode: global.JettonError1,
			ErrorMsg:  "下注失败, 无效的数据",
		})
		userItem.Close()
		return
	}

	//检验是否是下注状态
	if it.gameStatus != global.GameStatusJetton {
		_ = log.Logger.Errorf("OnUserPlaceJetton err %s", "下注失败, 非下注状态")
		//userItem.WriteMsg(&msg.Game_S_ReqlyFail{
		//	ErrorCode: global.JettonError1,
		//	ErrorMsg:  "下注失败, 非下注状态",
		//})
		//userItem.Close()
		return
	}
	value, ok := it.userList.Load(userItem.ChairID)
	//检验用户是否在用户列表里
	if ok && userItem.UserID != value.(int32) {
		_ = log.Logger.Errorf("OnUserPlaceJetton err %s", "下注失败, 用户不在用户列表里")
		userItem.WriteMsg(&msg.Game_S_ReqlyFail{
			ErrorCode: global.JettonError1,
			ErrorMsg:  "下注失败, 用户不在用户列表里",
		})
		userItem.Close()
		return
	}

	var tempScore float32
	v, ok := it.userListJetton.Load(userItem.ChairID)
	if ok {
		tempScore = v.(float32)
	}
	if !userItem.IsRobot() {

		if !redis.GameClient.IsExistsDiamond(userItem.UserID) {
			scoreInfo, _ := mysql.GetGameScoreInfoByUserId(mysql.GameClient.GetXJGameDB, userItem.UserID)
			redis.GameClient.SetDiamond(userItem.UserID, scoreInfo.Diamond)
		}

		userDiamond, err := redis.GameClient.GetDiamond(userItem.UserID)
		if err != nil {
			log.Logger.Error("GetDiamond err:", err)
		} else {

			userItem.UserDiamond = float32(userDiamond) + tempScore
		}
	}

	//判断下注积分是否足够
	if store.GameControl.GetGameInfo().DeductionsType == 0 {

		if userItem.UserGold < tempScore+m.JettonScore {
			_ = log.Logger.Errorf("OnUserPlaceJetton err %s", "下注失败, 金币不足!")
			userItem.WriteMsg(&msg.Game_S_ReqlyFail{
				ErrorCode: global.JettonError1,
				ErrorMsg:  "下注失败, 金币不足!",
			})
			return
		}
	} else {
		if userItem.UserDiamond < tempScore+m.JettonScore {
			_ = log.Logger.Errorf("OnUserPlaceJetton err %s", "下注失败, 余额不足!")
			userItem.WriteMsg(&msg.Game_S_ReqlyFail{
				ErrorCode: global.JettonError2,
				ErrorMsg:  "下注失败, 余额不足",
			})
			return
		}
	}

	if !userItem.IsRobot() {
		err := redis.GameClient.SetDiamond(userItem.UserID, userItem.UserDiamond-(tempScore+m.JettonScore))
		if err != nil {
			userItem.WriteMsg(&msg.Game_S_ReqlyFail{
				ErrorCode: global.JettonError2,
				ErrorMsg:  "下注失败",
			})
			return
		}

		redis.GameClient.RegisterRecharge(userItem.UserID)
	}

	//记录下注
	userJetton, ok1 := it.userListAreaJetton.Load(userItem.ChairID)
	if ok1 {
		temp := userJetton.([global.AreaCount]float32)
		temp[m.JettonArea] += m.JettonScore
		it.userListAreaJetton.Store(userItem.ChairID, temp)
	} else {
		var temp [global.AreaCount]float32
		temp[m.JettonArea] += m.JettonScore
		it.userListAreaJetton.Store(userItem.ChairID, temp)
	}

	if score, ok2 := it.userListJetton.Load(userItem.ChairID); ok2 {
		temp := score.(float32)
		temp += m.JettonScore
		it.userListJetton.Store(userItem.ChairID, temp)
	} else {
		it.userListJetton.Store(userItem.ChairID, m.JettonScore)
	}
	//富豪榜和神算子下注 才通知
	//if it.isOnTable(userItem) {
	//	it.sendAllUser(&msg.Game_S_UserJetton{
	//		Info: &msg.JettonInfo{
	//			Area:  m.JettonArea,
	//			Money: m.JettonScore,
	//		},
	//	})
	//
	//} else {
	//	userItem.WriteMsg(&msg.Game_S_UserJetton{
	//		Info: &msg.JettonInfo{
	//			Area:  m.JettonArea,
	//			Money: m.JettonScore,
	//		},
	//		//JettonArea:  m.JettonArea,
	//		//JettonScore: m.JettonScore,
	//	})
	//}

	userItem.WriteMsg(&msg.Game_S_UserJetton{
		Info: &msg.JettonInfo{
			Area:  m.JettonArea,
			Money: m.JettonScore,
		},
		//JettonArea:  m.JettonArea,
		//JettonScore: m.JettonScore,
	})

}

// 获取用户列表
func (it *Item) GetUserList(args ...interface{}) {
	userItem := args[0].(*user.Item)
	m := args[1].(*msg.Game_C_UserList)
	if m.Page <= 0 {
		m.Page = 1
	}
	// 控制数量
	if m.Size >= 10 {
		m.Size = 10
	}
	var list = make([]*user.Info, 0)
	it.userList.Range(func(key, uid interface{}) bool {
		item, ok := user.List.Load(uid)
		if !ok {
			_ = log.Logger.Errorf("GetUserList err %d", uid)
			return false
		}
		list = append(list, item.(*user.Item).GetUserInfo())
		return true
	})
	start := (m.Page - 1) * m.Size
	end := start + m.Size
	lenSize := int32(len(list))
	if lenSize <= m.Size {
		start = 0
		end = lenSize
	} else {
		if start >= lenSize {
			start = lenSize - 1
			if start <= 0 {
				start = 0
			}
		}
		if end >= lenSize {
			end = lenSize - 1
			if end <= 0 {
				end = start
			}
		}
	}

	var dataTempUser = make([]*msg.Game_S_TempUser, 0)
	for _, v := range list[start:end] {
		var totalJetton float32
		var totalWin int32
		for _, jetton := range it.userListJettons {
			value, ok := jetton.Load(v.ChairID)
			if ok {
				totalJetton += value.(float32)
			}
		}
		it.userListWinRecord.Range(func(key, value interface{}) bool {
			if key.(int32) == v.ChairID {
				for _, win := range value.([global.WinRecordCount]bool) {
					if win {
						totalWin++
					}
				}
			}
			return true
		})
		dataTempUser = append(dataTempUser, &msg.Game_S_TempUser{
			User: &msg.Game_S_User{
				UserID:       v.UserID,
				NikeName:     v.NikeName,
				UserDiamond:  v.UserDiamond,
				HeadImageUrl: v.HeadImageUrl,
				ChairID:      v.ChairID,
			},
			TotalJetton: totalJetton,
			TotalWin:    totalWin,
		})
	}

	userItem.WriteMsg(&msg.Game_S_UserList{
		Data: dataTempUser,
	})
}

// 发送所有人
func (it *Item) sendAllUser(data interface{}) {
	it.userList.Range(func(chairID, uid interface{}) bool {
		value, ok := user.List.Load(uid)
		if !ok {
			_ = log.Logger.Errorf("sendAll user.List err %d", uid)
			return true
		}
		value.(*user.Item).WriteMsg(data)
		return true
	})
}

// 发送其他人
func (it *Item) sendOtherUser(userID int32, data interface{}) {
	it.userList.Range(func(chairID, uid interface{}) bool {
		//过滤userID
		if userID == uid {
			return true
		}

		value, ok := user.List.Load(uid)
		if !ok {
			_ = log.Logger.Errorf("sendAll user.List err %d", uid)
			return true
		}
		value.(*user.Item).WriteMsg(data)
		return true
	})
}

// 判断当前人是否在桌子显示
func (it *Item) isOnTable(userItem *user.Item) bool {
	isSend := false
	for _, v := range it.specialUserList {
		if v.UserID == userItem.UserID {
			isSend = true
			break
		}
	}
	return isSend
}
