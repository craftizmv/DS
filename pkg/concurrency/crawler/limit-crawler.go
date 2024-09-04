package crawler

/*
	Given is a crawler (modified from the Go tour) that requests pages excessively.
	However, we don't want to burden the webserver too much.
	Your task is to change the code to limit the crawler to at most one page per second,
	while maintaining concurrency (in other words, Crawl() must be called concurrently)
*/
