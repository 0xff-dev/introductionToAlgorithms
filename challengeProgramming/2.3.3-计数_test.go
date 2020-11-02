package challengeProgramming

import "testing"


func TestDivideN2MGroups(t *testing.T) {
    n, m, M := 4, 3, 10000
    if res := DivideN2MGroups(n, m, M); res != 4 {
        t.Fatalf("expect %d get %d", 4, res)
    }
}
