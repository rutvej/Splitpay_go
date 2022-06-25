build:
	sudo docker build -t splitpay .;
	
run:
	sudo docker run -d --ulimit nofile=20000:20000 --name splitpay --restart always splitpay;

rebuild:
	sudo docker build -t splitpay .;
	sudo docker rm -f splitpay ;
	sudo docker run -d --ulimit nofile=20000:20000 --net=host --name splitpay --restart always splitpay; 