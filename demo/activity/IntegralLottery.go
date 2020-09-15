package activity

import (
	"context"
	"sync"
)

type IntegralLottery struct {
	ctx          context.Context
	userScore    map[int64]int64
	c            chan *UserScore
	currentScore int64
	round        int32
	awardUid     []int64 // 获奖名单 0 = 无人中奖
	sync.RWMutex
}

type IntegralLotteryResData struct {
	currentScore int64
	round        int32
	awardUid     []int64
}

type UserScore struct {
	Uid        int64
	PlusScore  int64
	TotalScore int64
	Candidate  bool
}

const ChanMax = 1 << 20
const CfgAwardScore = 50000

func NewIntegralLottery(ctx context.Context) *IntegralLottery {
	// 读取DB
	return &IntegralLottery{
		ctx:          ctx,
		userScore:    map[int64]int64{},
		c:            make(chan *UserScore, ChanMax),
		currentScore: 0,
		round:        1,
	}
}

func (il *IntegralLottery) ActStart() {
	go il.run()
}

func (il *IntegralLottery) AddUserScore(userScore *UserScore) {
	il.c <- userScore
}

func (il *IntegralLottery) run() {
	defer func() {
		// 如果关闭 说明活动结束 不处理后面的数据
	}()
	for {
		select {
		case <-il.ctx.Done():
			return
		case userScore, ok := <-il.c:
			if !ok {
				return
			}
			il.handle(userScore)
		}
	}
}

func (il *IntegralLottery) handle(userScore *UserScore) {
	il.currentScore += userScore.PlusScore
	if userScore.Candidate {
		il.userScore[userScore.Uid] += userScore.TotalScore
	}
	if il.currentScore >= CfgAwardScore {
		// 开始抽奖
		il.award()
	}
	left := il.currentScore - CfgAwardScore // 溢出
	il.currentScore = left
	il.round++
}

// 开始抽奖
func (il *IntegralLottery) award() {
	il.awardUid = append(il.awardUid, 1000)
}

func (il *IntegralLottery) GetCurrentInfo() *IntegralLotteryResData {
	data := &IntegralLotteryResData{
		currentScore: il.currentScore,
		round:        il.round,
	}
	data.awardUid = make([]int64, len(il.awardUid))
	for _, uid := range il.awardUid {
		data.awardUid = append(data.awardUid, uid)
	}
	return data
}
