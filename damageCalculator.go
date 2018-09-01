package main

import (
    "fmt"
)

type Weapon struct {
    Name string
    Kind string
    RawDamage float64
    ElemDamage float64
    Affinity float64
    Sharpness string
}

func main() {

    var testWeapon = Weapon{Name: "Katana KiKoup", Kind: "LongSword", RawDamage: 467, ElemDamage: 127, Affinity: 0.2, Sharpness: "green"}

    fmt.Println(testWeapon.getWeaponDamage())
}

func (this Weapon) getWeaponDamage() float64 {
    fmt.Println(this)

    var SharpRawDmgMult, SharpElemDmgMult float64

    switch(this.Sharpness) {
        case "red":
            SharpRawDmgMult = 0.3
            SharpElemDmgMult = 0.25
        case "orange":
            SharpRawDmgMult = 0.4
            SharpElemDmgMult = 0.50
        case "yellow":
            SharpRawDmgMult = 0.55
            SharpElemDmgMult = 0.75
        case "green":
            SharpRawDmgMult = 1
            SharpElemDmgMult = 1
        case "blue":
            SharpRawDmgMult = 1.13
            SharpElemDmgMult = 1.0625
        case "white":
            SharpRawDmgMult = 1.25
            SharpElemDmgMult = 1.125
    }

    var Diameter float64

    switch(this.Kind) {
        case "GreatSword":
            Diameter = 4.8
        case "SwordShield":
            Diameter = 1.4
        case "DualBlades":
            Diameter = 1.4
        case "LongSword":
            Diameter = 3.3
        case "Hammer":
            Diameter = 5.2
        case "HuntingHorn":
            Diameter = 4.2
        case "Lance":
            Diameter = 2.3
        case "GunLance":
            Diameter = 2.3
        case "SwitchAxe":
            Diameter = 3.5
        case "ChargeBlade":
            Diameter = 3.6
        case "InsectGlaive":
            Diameter = 3.1
        case "Bow":
            Diameter = 1.2
        case "LightBowGun":
            Diameter = 1.3
        case "HeavyBowGun":
            Diameter = 1.5
    }

    var rawdamage = (this.RawDamage / Diameter) * SharpRawDmgMult
    var elemDamage = (this.ElemDamage / 10) * SharpElemDmgMult
    var affinityDamage = (this.RawDamage / Diameter) * this.Affinity / 400

    var totalDamage = rawdamage + elemDamage + affinityDamage

    return totalDamage
}
