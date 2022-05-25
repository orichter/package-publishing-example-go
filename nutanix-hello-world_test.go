package nutanix_hello_world

import (
    "testing"
)

func TestHello(t *testing.T) {
    want := "hello nutanix world"
    if got := HelloWorld(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
