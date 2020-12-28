syno_build:
	# rm ./syno_raw_matcher
	GOOS=linux GOARCH=arm GOARM=5 cd cmd && go build -o ../data/syno_raw_matcher command.go

syno_deploy:
	bash data/syno_deploy.sh

syno_release: | syno_build syno_deploy
	echo './syno_raw_matcher -path "./2020/" -output test.sh'

pc_build:
	cd cmd && go build -o ../data/raw_matcher command.go

test:
	cd raw_matcher && go test -v
