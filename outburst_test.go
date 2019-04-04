package outburst

import "testing"

func buildTest(t *testing.T) {
	ob := NewOutBurst()
	if ob.Conf.Vulgar == nil {
		t.Errorf("Vulgar not initialized")
	}
	if ob.Conf.Debug == nil {
		t.Errorf("Debug not initialized")
	}
	if ob.Conf.Emojis == nil {
		t.Errorf("Emojis not initialized")
	}
	if ob.Conf.TimeFormat == "" {
		t.Errorf("TimeFormat not initialized")
	}
	if ob.Conf.Padding == nil {
		t.Errorf("Padding not initialized")
	}
}

func outTest(t *testing.T) {
	ob := NewOutBurst()
	ob.Out(Knot{"Scooby Doo": "Dog", "Age": 4})
}
