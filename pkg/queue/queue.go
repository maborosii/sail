package queue

import (
	"container/list"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var queueSize = 200

type Job struct {
	UUID      string
	DoneChan  chan struct{}
	HandleJob func() error //具体的处理逻辑
}

func NewJob(HandleJob func() error) *Job {
	return &Job{
		UUID:      uuid.NewString(),
		DoneChan:  make(chan struct{}, 1),
		HandleJob: HandleJob,
	}
}

// 消费者从队列中取出该job时 执行具体的处理逻辑
func (job *Job) Execute() error {
	// fmt.Println(time.Now())
	fmt.Println(job.UUID, "job start to execute ")
	return job.HandleJob()
}

// 执行完Execute后，调用该函数以通知主线程中等待的job
func (job *Job) Done() {
	job.DoneChan <- struct{}{}
	close(job.DoneChan)
}

// 工作单元等待自己被消费
func (job *Job) WaitDone() {
	<-job.DoneChan
}

type JobQueue struct {
	mu         sync.Mutex
	noticeChan chan struct{}
	queue      *list.List
	size       int
	capacity   int
}

// 初始化队列
func NewJobQueue(cap int) *JobQueue {
	return &JobQueue{
		capacity:   cap,
		queue:      list.New(),
		noticeChan: make(chan struct{}, 1),
	}
}

// 工作单元入队
func (q *JobQueue) PushJob(job *Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.size++
	if q.size > q.capacity {
		q.RemoveLeastJob()
	}

	q.queue.PushBack(job)
	q.noticeChan <- struct{}{}
}

// 工作单元出队
func (q *JobQueue) PopJob() *Job {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.size == 0 {
		return nil
	}

	// fmt.Println(q.size)
	q.size--
	return q.queue.Remove(q.queue.Front()).(*Job)
}

// 移除队列中的最后一个元素。
// 一般在容量满时，有新job加入时，会移除等待最久的一个job
func (q *JobQueue) RemoveLeastJob() {
	if q.queue.Len() != 0 {
		back := q.queue.Back()
		abandonJob := back.Value.(*Job)
		abandonJob.Done()
		q.queue.Remove(back)
	}
}

// 消费线程监听队列的该通道，查看是否有新的job需要消费
func (q *JobQueue) waitJob() <-chan struct{} {
	return q.noticeChan
}

type WorkerManager struct {
	jobQueue *JobQueue
}

func NewWorkerManager(jobQueue *JobQueue) *WorkerManager {
	return &WorkerManager{
		jobQueue: jobQueue,
	}
}
func (m *WorkerManager) createWorker() error {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("when execute job occured panic, panic: %s\n", err)
			}
		}()
		var currentJob *Job
		for range m.jobQueue.waitJob() {
			currentJob = m.jobQueue.PopJob()
			if err := currentJob.Execute(); err != nil {
				currentJob.Done()
				continue
			}
			currentJob.Done()
		}
	}()
	return nil
}

type FlowControl struct {
	jobQueue *JobQueue
	wm       *WorkerManager
}

func NewFlowControl() *FlowControl {
	jobQueue := NewJobQueue(queueSize)
	fmt.Println("init job queue success")

	m := NewWorkerManager(jobQueue)
	if err := m.createWorker(); err != nil {
		fmt.Printf("init job queue createworker occured err, err: %s", err)
		panic(err)
	}
	fmt.Println("init worker success")

	control := &FlowControl{
		jobQueue: jobQueue,
		wm:       m,
	}
	fmt.Println("init flowcontrol success")
	return control
}

func (c *FlowControl) CommitJob(job *Job) {
	c.jobQueue.PushJob(job)
}
