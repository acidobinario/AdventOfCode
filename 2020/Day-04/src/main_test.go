package main

import (
	"testing"
)

var byr = "2002"
var iyr = "2012"
var eyr = "2022"
var hgt = "60in"
var hcl = "#123abc"
var ecl = "brn"
var pid = "000000001"

func Test_passport_checkHgt(t *testing.T) {

	type fields struct {
		byr *string
		iyr *string
		eyr *string
		hgt *string
		hcl *string
		ecl *string
		pid *string
		cid *string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "test ok",
			fields: fields{
				byr: &byr,
				iyr: &iyr,
				eyr: &eyr,
				hgt: &hgt,
				hcl: &hcl,
				ecl: &ecl,
				pid: &pid,
				cid: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p Passport
			p = Passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.checkHgt(); got != tt.want {
				t.Errorf("checkHgt() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_passport_checkEcl(t *testing.T) {
	type fields struct {
		byr *string
		iyr *string
		eyr *string
		hgt *string
		hcl *string
		ecl *string
		pid *string
		cid *string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "UWU",
			fields: fields{
				byr: &byr,
				iyr: &iyr,
				eyr: &eyr,
				hgt: &hgt,
				hcl: &hcl,
				ecl: &ecl,
				pid: &pid,
				cid: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.checkEcl(); got != tt.want {
				t.Errorf("checkEcl() = %v, want %v", got, tt.want)
			}
		})
	}
}