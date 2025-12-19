package oppbloom_test

import (
        "math/rand"
        "testing"
        "time"
        "fmt"

        "github.com/HoMuChen/opp-bloom"
)

func TestInit(t *testing.T) {
        set := oppbloom.New(1000)

        if set.Cap() != 1000 {
                t.Errorf("Expected capacity %v but got: %v", 1000, set.Cap())
        }
}

func TestAddAndContain(t *testing.T) {
        set := oppbloom.New(
                1000,
                oppbloom.WithHasher(&oppbloom.FNVHasher{}),
        )

        set.Add([]byte(`1`))
        exist := set.Contain([]byte(`1`))

        if exist != true {
                t.Errorf("Expected containg 1 but got false")
        }
}

func TestLimitedSize(t *testing.T) {
        set := oppbloom.New(
                1,
        )

        set.Add([]byte(`1`))
        set.Add([]byte(`2`))
        exist1 := set.Contain([]byte(`1`))
        exist2 := set.Contain([]byte(`2`))

        if exist1 != false {
                t.Errorf("Expected not containg 1 but got true")
        }
        if exist2 != true {
                t.Errorf("Expected containg 2 but got false")
        }
}

func testFalseNegative(size int, numKeys int) {
        rand.Seed(time.Now().UnixNano())
        set := oppbloom.New(
                size,
        )
        answer := make(map[string]bool)

        for i := 0; i < numKeys; i++ {
                key := make([]byte, 16)

                rand.Read(key)
                set.Add(key)
                answer[string(key)] = true
        }

        count := 0
        for key := range answer {
                if exist := set.Contain([]byte(key)); !exist {
                        count += 1
                }
        }

        fmt.Printf("Set size: %v\n", size)
        fmt.Printf("Numbe of keys: %v\n", numKeys)
        fmt.Printf("False negative rate %v %%\n\n", (float64(count) / float64(numKeys)) * 100)
}

func testRecentFalseNegative(size int, numKeys int, n int) {
        rand.Seed(time.Now().UnixNano())
        set := oppbloom.New(
                size,
        )

        for i := 0; i < n; i++ {
                for j := 0; j < numKeys; j++ {
                        key := make([]byte, 16)

                        rand.Read(key)
                        set.Add(key)
                }
        }

        answer := make(map[string]bool)
        for j := 0; j < numKeys; j++ {
                key := make([]byte, 16)

                rand.Read(key)
                set.Add(key)
                answer[string(key)] = true
        }

        count := 0
        for key := range answer {
                if exist := set.Contain([]byte(key)); !exist {
                        count += 1
                }
        }

        fmt.Printf("Set size: %v\n", size)
        fmt.Printf("Numbe of keys: %v\n", numKeys*(n+1))
        fmt.Printf("False negative rate of rencent %v keys: %v %%\n\n", numKeys, (float64(count) / float64(numKeys)) * 100)
}

func TestFalseNegative(t *testing.T) {
        testFalseNegative(10000, 100)
        testFalseNegative(10000, 500)
        testFalseNegative(10000, 1000)
        testFalseNegative(10000, 2000)
}

func TestRecentFalseNegative(t *testing.T) {
        testRecentFalseNegative(10000, 1000, 10)
        testRecentFalseNegative(10000, 2000, 5)
        testRecentFalseNegative(3000000, 100000, 30)
}

func BenchmarkAdd(b *testing.B) {
        set := oppbloom.New(100000)
        rand.Seed(time.Now().UnixNano())

        key := make([]byte, 16)
        for i := 0; i < b.N; i++ {
                rand.Read(key)
                set.Add(key)
        }
}
