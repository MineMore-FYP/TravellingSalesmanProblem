
		source /etc/profile
		source ~/.profile
		
process_worker_pool.py  --max_workers=5 -p 0 -c 1 -m None --poll 10 --task_url=tcp://10.22.200.65:54944 --result_url=tcp://10.22.200.65:54900 --logdir=/home/clusteruser/TravellingSalesmanProblem/PSO-GA//remote_htex --block_id=2 --hb_period=30 --hb_threshold=120 