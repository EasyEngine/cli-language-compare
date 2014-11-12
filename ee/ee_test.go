package main


import (
"testing"
"ee"
)

func TestEE_Run(t *testing.T) {

err := app.Run([]string{"site create 1.com --html"})
expect(t, err, nil)
}
