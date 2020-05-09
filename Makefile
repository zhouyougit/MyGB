build:
	go build -gcflags "all=-N -l" -o mygb

run:
	./mygb run SuperMarioLand.gb
