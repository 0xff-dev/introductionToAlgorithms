package dp

import (
	"fmt"
	"log"
)

/*
据说在很久很久以前，可怜的兔子经历了人生中最大的打击——赛跑输给乌龟后，心中郁闷，发誓要报仇雪恨，于是躲进了杭州下沙某农业园卧薪尝胆潜心修炼，终于练成了绝技，能够毫不休息得以恒定的速度(VR m/s)一直跑。兔子一直想找机会好好得教训一下乌龟，以雪前耻。
最近正值HDU举办50周年校庆，社会各大名流齐聚下沙，兔子也趁此机会向乌龟发起挑战。虽然乌龟深知获胜希望不大，不过迫于舆论压力，只能接受挑战。
比赛是设在一条笔直的道路上，长度为L米，规则很简单，谁先到达终点谁就算获胜。
无奈乌龟自从上次获胜以后，成了名龟，被一些八卦杂志称为“动物界的刘翔”，广告不断，手头也有了不少积蓄。为了能够再赢兔子，乌龟不惜花下血本买了最先进的武器——“"小飞鸽"牌电动车。这辆车在有电的情况下能够以VT1 m/s的速度“飞驰”，可惜电池容量有限，每次充满电最多只能行驶C米的距离，以后就只能用脚来蹬了，乌龟用脚蹬时的速度为VT2 m/s。更过分的是，乌龟竟然在跑道上修建了很多很多（N个)的供电站，供自己给电动车充电。其中，每次充电需要花费T秒钟的时间。当然，乌龟经过一个充电站的时候可以选择去或不去充电。
比赛马上开始了，兔子和带着充满电的电动车的乌龟并列站在起跑线上。你的任务就是写个程序，判断乌龟用最佳的方案进军时，能不能赢了一直以恒定速度奔跑的兔子。
*/

const INF = 0xffff

func tortoiseCanWin() {
	var (
		L int // 总的路径长度
		N int // 充电站个数
		C int // 充满电后，乌龟的车子可以形式的距离
		T int // 充电需要消耗的时间
		VR int // 兔子的速度
		VT1 int // 乌龟的车速
		VT2 int // 乌龟的蹬车速度
	)
	_, err := fmt.Scanf("%d", &L)
	if err != nil {
		log.Fatal("read load length error")
	}
	_, err = fmt.Scanf("%d %d %d", &N, &C, &T)
	if err != nil {
		log.Fatal("read N, C, T error")
	}
	_, err = fmt.Scanf("%d %d %d", &VR, &VT1, &VT2)
	if err != nil {
		log.Fatal("read vr, vt1, vt2 error")
	}
	station := make([]int, N+2)
	station[0], station[N+1] = 0, L
	for i := 1; i <= N; i++ {
		fmt.Scanf("%d", &station[i])
	}
	dp := make([]float64, N+2)
	dp[0] = 0
	for pos := 1; pos < N+2; pos ++ {
		dp[pos] = INF
		for j := 0; j < pos; j++ {
			length := station[pos] - station[j]
			var cost float64
			if length > C {
				cost = float64(C*1.0/VT1) + float64(length*1.0/VT2)
			} else {
				cost = float64(length*1.0/VT1)
			}
			if j != 0 {
				cost += float64(T)
			}
			if cost < dp[pos] {
				dp[pos] =cost
			}
		}
	}
	if dp[N+1] > float64(L*1.0/VR){
		fmt.Println("Good job,rabbit!")
	} else {
		fmt.Println("What a pity rabbit!")
	}
}