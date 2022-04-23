package qreki

import "testing"

func TestAdd(t *testing.T) {
  type args struct {
    a int
    b int
  }

  tests := []struct {
    name string
    args args
    want int
  }{
    {
      name: "normal",
      args: args{a: 1, b: 2},
      want: 3,
    },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := Add(tt.args.a, tt.args.b); got != tt.want {
        t.Errorf("add() = %v, want %v", got, tt.want)
      }
    })
  }
}
