package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math/rand"
	"time"
)

var (
	Plate  [4][4]int
	count  int
	maxNum int
)

func init() {
	filling()
	printPlate()
}

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		//fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		//fmt.Println(key,"|",char)
		count++
		switch key {
		case 65517:
			//fmt.Println("up")
			reUp()
			printPlate()
		case 65515:
			//fmt.Println("left")
			reLeft()
			printPlate()
		case 65516:
			//fmt.Println("down")
			reDown()
			printPlate()
		case 65514:
			//fmt.Println("right")
			reRight()
			printPlate()
		}

		if key == keyboard.KeyEsc {
			break
		}
		ok := fillingNew()
		if !ok {
			fmt.Printf("game over,max num is %d\noperand is %d ", maxNum,count)
			break
		}
	}
}

// 打印数阵
func printPlate() {

	for k := range Plate {
		fmt.Println(Plate[k])
	}
}

// 随机获取一个随机数，为1到9
func random() (int64, int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(3), r.Int63n(3)
}

// 新填充数字,返回true则表示仍存在空白格，如果返回false则说明已经满了，游戏结束
func filling() bool {

	defer fmt.Printf("count is %d\n", count)
	num1, num2 := random()

	ok := true
	for ok {
		if Plate[num1][num2] == 0 {
			Plate[num1][num2] = 2
			ok = false
		}
		zero := checkZero()
		if !zero {
			return false
		}
		num1, num2 = random()
	}
	return true
}

// 新填充数字的方法
// 因为原来填充数字的方法存在缺陷，原逻辑为随机生成一个坐标填入数字，如果已存在则再次生成
// 导致了操作数多了之后，投入概率过低
func fillingNew() bool {
	defer fmt.Printf("You did it %d times\n", count)
	zero := checkZero()
	if !zero {
		return false
	}

	// 获取当前空白格的队列
	queue := getZeroQueue()

	singleRandom := getSingleRandom(int64(len(queue)))
	fmt.Println(queue[singleRandom])
	Plate[queue[singleRandom][0]][queue[singleRandom][1]] = 2
	return true
}

// 获取空白格队列
func getZeroQueue() [][]int {
	var queue [][]int
	for k, v := range Plate {
		for k2, v2 := range v {
			if v2 == 0 {
				queue = append(queue, []int{k, k2})
			}
		}
	}
	return queue
}

// 为队列中的空白格随机填充数字
func getSingleRandom(n int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(n)
}

// 验证是否还有空白
func checkZero() bool {
	for _, v := range Plate {
		for k := range v {
			if v[k] == 0 {
				return true
			}
			if v[k] > maxNum {
				maxNum = v[k]
			}
		}
	}
	return false
}

// 向上
func reUp() {
	reverse(1)
	moveLeft()
	reverse(3)
}

// 向下
func reDown() {
	reverse(3)
	moveLeft()
	reverse(1)
}

// 向左
func reLeft() {
	moveLeft()
}

// 向右
func reRight() {
	reverse(2)
	moveLeft()
	reverse(2)
}

// 向左反转90度
func reverse(n int) {
	for x := 0; x < n; x++ {
		var newPlate [4][4]int

		num := len(Plate) - 1
		for k := range newPlate {

			for i := range newPlate[k] {
				newPlate[k][i] = Plate[i][num]
			}
			num--
		}

		Plate = newPlate
	}

}

// 向左移动数字
func moveLeft() {

	num := 1
	for j := range Plate {
		num = 1
		for num != 0 {
			num = 0
			for k := 0; k < len(Plate)-1; k++ {
				if Plate[j][k] == 0 && Plate[j][k+1] != 0 {
					Plate[j][k], Plate[j][k+1] = Plate[j][k+1], Plate[j][k]
					num++
					continue
				}
				if Plate[j][k] == Plate[j][k+1] && Plate[j][k] != 0 {
					Plate[j][k] += Plate[j][k+1]
					Plate[j][k+1] = 0
					num++
					continue
				}
			}

		}
	}

}
