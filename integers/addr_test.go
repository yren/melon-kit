package integers
import "testing"
func TestAddr(t *testing.T) {
    got := Addr(2, 2)
    expected := 4
    if got != expected {
        t.Errorf("got: %d, expected: %d", got, expected)
    }
}