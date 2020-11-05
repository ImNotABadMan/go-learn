package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mailUnit8 uint8
		sigalCh   chan uint8 = make(chan uint8)
		count     uint8
	)
	// 基于互斥锁
	// 创建互斥锁
	lock := sync.RWMutex{}
	// 发送的条件变量，基于写锁
	sendCond := sync.NewCond(&lock)
	// 接收的条件变量，基于读锁
	recvCond := sync.NewCond(lock.RLocker())

	// 接收程序
	go func() {
		// 先锁
		lock.RLock()
		// 不断地等待，防止接收到的不是自己想要的值，需要再次阻塞等待通知
		// 不为0才可以继续执行，不然就等待
		fmt.Println("wait receive 1 mail uint8", mailUnit8)
		for mailUnit8 != 1 {
			fmt.Println("    before receive 1 mail uint8", mailUnit8)
			// 阻塞当前协程，使当前协程的代码无法继续往下执行
			recvCond.Wait()
			fmt.Println("after-- receive 1 mail uint8", mailUnit8)
		}
		mailUnit8 = 2
		fmt.Println("receive 1 mail uint8", mailUnit8)
		lock.RUnlock()
		sendCond.Broadcast()
		sigalCh <- mailUnit8
	}()

	// 发送程序
	go func() {
		lock.Lock()
		fmt.Println("wait send 1 mail uint8", mailUnit8)
		for mailUnit8 != 0 {
			fmt.Println("    before send 1 mail uint8", mailUnit8)
			sendCond.Wait()
			fmt.Println("after-- wait send 1 mail uint8", mailUnit8)
		}
		mailUnit8 = 1
		fmt.Println("send 1 mail uint8", mailUnit8)
		lock.Unlock()
		// 需要被唤醒的队列的所有协程的第一个，也是最早的一个
		recvCond.Broadcast()
		sigalCh <- mailUnit8

	}()

	go func() {
		// 先锁
		lock.RLock()
		fmt.Println("wait receive 3 mail uint8", mailUnit8)
		// 不断地等待，防止接收到的不是自己想要的值，需要再次阻塞等待通知
		// 不为0才可以继续执行，不然就等待
		for mailUnit8 != 3 {
			fmt.Println("    before receive 3 mail uint8", mailUnit8)
			recvCond.Wait()
			fmt.Println("after-- receive 3 mail uint8", mailUnit8)
		}
		mailUnit8 = 4
		fmt.Println("receive 3 mail uint8", mailUnit8)
		lock.RUnlock()
		sendCond.Broadcast()
		sigalCh <- mailUnit8
	}()

	go func() {
		lock.Lock()
		fmt.Println("wait send 2 mail uint8", mailUnit8)
		for mailUnit8 != 2 {
			fmt.Println("    before wait send 2 mail uint8", mailUnit8)
			sendCond.Wait()
			fmt.Println("after-- wait send 2 mail uint8", mailUnit8)
		}
		mailUnit8 = 3
		fmt.Println("send 2 mail uint8", mailUnit8)
		lock.Unlock()
		// 需要被唤醒的队列的所有协程的第一个，也是最早的一个
		recvCond.Broadcast()
		sigalCh <- mailUnit8

	}()

	go func() {
		// 先锁
		lock.RLock()
		fmt.Println("wait receive 5 mail uint8", mailUnit8)
		// 不断地等待，防止接收到的不是自己想要的值，需要再次阻塞等待通知
		// 不为0才可以继续执行，不然就等待
		for mailUnit8 != 5 {
			fmt.Println("    before receive 5 mail uint8", mailUnit8)
			recvCond.Wait()
			fmt.Println("after-- receive 5 mail uint8", mailUnit8)
		}
		mailUnit8 = 6
		fmt.Println("receive 5 mail uint8", mailUnit8)
		lock.RUnlock()
		sendCond.Broadcast()
		sigalCh <- mailUnit8
	}()

	go func() {
		lock.Lock()
		fmt.Println("wait send 4 mail uint8", mailUnit8)
		for mailUnit8 != 4 {
			fmt.Println("    before wait send 4 mail uint8", mailUnit8)
			sendCond.Wait()
			fmt.Println("after-- wait send 4 mail uint8", mailUnit8)
		}
		mailUnit8 = 5
		fmt.Println("send 4 mail uint8", mailUnit8)
		lock.Unlock()
		// 需要被唤醒的队列的所有协程的第一个，也是最早的一个
		recvCond.Broadcast()
		sigalCh <- mailUnit8

	}()

	go func() {
		// 先锁
		lock.RLock()
		fmt.Println("wait receive 7 mail uint8", mailUnit8)
		// 不断地等待，防止接收到的不是自己想要的值，需要再次阻塞等待通知
		// 不为0才可以继续执行，不然就等待
		for mailUnit8 != 7 {
			fmt.Println("    before 7 mail uint8", mailUnit8)
			recvCond.Wait()
			fmt.Println("after-- receive 7 mail uint8", mailUnit8)
		}
		mailUnit8 = 8
		fmt.Println("receive 7 mail uint8", mailUnit8)
		lock.RUnlock()
		sendCond.Broadcast()
		sigalCh <- mailUnit8
	}()

	go func() {
		lock.Lock()
		fmt.Println("wait send 6 mail uint8", mailUnit8)
		for mailUnit8 != 6 {
			fmt.Println("    before wait send 6 mail uint8", mailUnit8)
			sendCond.Wait()
			fmt.Println("after-- wait send 6 mail uint8", mailUnit8)
		}
		mailUnit8 = 7
		fmt.Println("send 6 mail uint8", mailUnit8)
		lock.Unlock()
		//recvCond.Signal()  需要被唤醒的队列的所有协程的第一个，也是最早的一个
		recvCond.Broadcast()
		sigalCh <- mailUnit8

	}()

	for range sigalCh {
		if count++; count >= 8 {
			break
		}
	}
}
