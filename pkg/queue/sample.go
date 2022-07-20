package queue

// request 包含的内容
type Request interface{}

// 请求队列, 本质是一个channel
type RequestQueue struct {
	Queue chan Request
}

var queue *RequestQueue

// 获取队列
func GetQueue() *RequestQueue {
	return queue
}

// 初始化队列
func InitRequestQueue(size int) {
	queue = &RequestQueue{
		Queue: make(chan Request, size),
	}
}

// 将请求放入队列
func (rq *RequestQueue) Enqueue(p Request) {
	rq.Queue <- p
}

// 请求队列服务, 一直等待接受和处理请求
func (rq *RequestQueue) Run() {
}
