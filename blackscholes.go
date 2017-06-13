package main

import (
	"math"
	"github.com/chobie/go-gaussian"
)

type Option struct {
	price		float64
	exercise	float64
	time		float64
	vol		float64
	rf		float64
	call		bool
}


func newOption (p float64, x float64, t float64, sigma float64, r float64, isCall bool) Option {
	newOption :=	Option{
			price:		p,
			exercise:	x,
			time:		t/365,
			vol:		sigma,
			rf:		r,
			call:		isCall,
			}
	return newOption
}

func OptionPrice (p float64, x float64, t float64, vol float64, r float64, iscall bool) float64 {
	var price float64
	if iscall == true {
		price = vC(p, x, t, vol, r)
	} else {
		price = vP(p, x, t, vol, r)
	}
	return price
}

func vC (p float64, x float64, t float64, vol float64, r float64) float64 {
	d1 := calcD1(p, x, t, vol, r)
	d2 := calcD2(d1, t, vol)
	stdnorm := gaussian.NewGaussian(0, 1)
	return (p * stdnorm.Cdf(d1)) - ((x / (math.Pow(2.71828, r * t))) * stdnorm.Cdf(d2))
}

func vP (p float64, x float64, t float64, vol float64, r float64) float64 {
	return vC(p, x, t, vol, r) + (x/math.Pow(2.71828, r*t))-p
}


func calcD1 (p float64, x float64, t float64, vol float64, r float64) float64 {
	return (math.Log(p/x) + (r + math.Pow(vol, 2)*0.5)*t)/(vol * math.Sqrt(t))
}

func calcD2 (d1 float64, t float64, vol float64) float64 {
	return (d1 - (vol * math.Sqrt(t)))
}

func CalcDelta (p float64, x float64, t float64, vol float64, r float64, call bool) float64 {
	var delta float64
	callDelta := vC(p+1, x, t, vol, r) - vC(p, x, t, vol, r)
	if call == true {
		delta = callDelta
		} else {
		delta = 1.0 - callDelta
		} 
	return delta
}

