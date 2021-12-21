package main

// The wrapper of your app
func yourApp(s server) {

	// This is just some sample code to do something
	s.winlog.Info(1, "In Xpert PointSix Parse")
	s.winlog.Info(1, "Still parsing...")
	s.winlog.Info(1, "And parsing...")
	s.winlog.Info(1, "And the service will keep parsing...")

	// Notice that if this exits, the service continues to run - you can launch web servers, etc.
}
