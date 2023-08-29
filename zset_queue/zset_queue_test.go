package zsetqueue

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

// 使用zesed 实现延迟队列
const QueueName = "delayed_task"

func TestClient(t *testing.T) {
	AddQueue(QueueName)
}
func TestCosumer(t *testing.T) {
	r := NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	Cosumer(ctx, r.cli, QueueName)
	//time.Sleep(time.Second * 30)
}

type MyRedis struct {
	cli *redis.Client
}

func NewClient() MyRedis {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return MyRedis{
		cli: client,
	}
	// pong, err := client.Ping().Result()
	// fmt.Println(pong, err)
}

// ZAdd redis zadd 方法
func (m MyRedis) ZAdd(delayQ string, task string, score int) {
	err := m.cli.ZAdd(delayQ, redis.Z{
		Score:  float64(score),
		Member: task,
	}).Err()
	if err != nil {
		log.Println("add task failed", err)
	}
}

type Tasks struct {
	task map[string]int64
}

// AddTask  当前时间+延迟时间
func (t Tasks) AddTask(name string, n int) {
	delay := int(time.Second) * n
	t.task[name] = time.Now().Add(time.Duration(delay)).Unix()
}

func AddQueue(QueueName string) {
	task := Tasks{task: map[string]int64{}}
	task.AddTask("task1", 10)
	task.AddTask("task2", 10)
	task.AddTask("task3", 13)

	r := NewClient()
	for t, score := range task.task {
		r.ZAdd(QueueName, t, int(score))
	}
}

func Cosumer(ctx context.Context, cli *redis.Client, queueName string) {
	now := time.Now().Unix()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ok time out")
			return
		default:
			result, err := cli.ZRangeByScore(queueName, redis.ZRangeBy{
				Min:    "-1",
				Max:    fmt.Sprintf("%d", now),
				Offset: 0,
				Count:  1,
			}).Result()
			if err != nil {
				log.Println("get task failed", err)
				continue
			}
			if len(result) > 0 {
				task := result[0]
				fmt.Printf("run task:%s\n", task)
				_, err := cli.ZRem(queueName, task).Result()
				if err != nil {
					log.Println("delete task failed", err)
				}
			}
			time.Sleep(time.Second * 1)

		}

	}
}
