viashared:
	$(MAKE) -C viashared all

libRandom:
	$(MAKE) -C libRandom all

clean:
	$(MAKE) -C viashared clean
	$(MAKE) -C libRandom clean
	rm -rf bin

.PHONY: viashared libRandom
