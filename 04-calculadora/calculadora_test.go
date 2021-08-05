package main

import "testing"

func Test_adicionarpercentual(t *testing.T) {
	type args struct {
		valor       float64
		porcentagem float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adicionarpercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
				t.Errorf("adicionarpercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_descobrirpercentual(t *testing.T) {
	type args struct {
		valor  float64
		valor2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := descobrirpercentual(tt.args.valor, tt.args.valor2); got != tt.want {
				t.Errorf("descobrirpercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divisao(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("divisao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplicacao(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplicacao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("multiplicacao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_soma(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soma(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("soma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtracao(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "test 2 - 1 ",
			args: args{2, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtracao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("subtracao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow(t *testing.T) {
	type args struct {
		valor    float64
		multiplo float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.valor, tt.args.multiplo); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}