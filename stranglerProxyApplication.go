package main

import (
	log "log"
	env "pimblott.com/golang/stranglerProxy/environment"
)


/*
	Entry
*/
func main() {
	liveUrl := env.GetEnvOrStop("LIVE_URL" );
	testUrl := env.GetEnvOrStop( "TEST_URL")
	port := env.GetEnvWithFallback("PORT", "8085" );

	e := NewProxy(liveUrl, testUrl)

	log.Printf("Using live url: %s\n", e.LiveUrl())
	log.Printf("Comparing result to url: %s\n", e.TestUrl())

	e.StartProxy(port)
}