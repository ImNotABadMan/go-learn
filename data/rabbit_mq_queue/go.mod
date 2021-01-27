module "data/rabbit_mq_queue"

go 1.15

require (
	data/rabbit_mq_queue/cp v1.0.0
)

replace (
	data/rabbit_mq_queue/cp => ./cp
)