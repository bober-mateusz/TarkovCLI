// models package contains various types used in the project.
package models

// Ammo represents the structure of each ammo entry
type Ammo struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	ShortName       string  `json:"shortName"`
	Weight          float64 `json:"weight"`
	Caliber         string  `json:"caliber"`
	StackMaxSize    int     `json:"stackMaxSize"`
	Tracer          bool    `json:"tracer"`
	TracerColor     string  `json:"tracerColor"`
	AmmoType        string  `json:"ammoType"`
	ProjectileCount int     `json:"projectileCount"`
	Ballistics      struct {
		Damage              int     `json:"damage"`
		ArmorDamage         int     `json:"armorDamage"`
		FragmentationChance float64 `json:"fragmentationChance"`
		RicochetChance      float64 `json:"ricochetChance"`
		PenetrationChance   float64 `json:"penetrationChance"`
		PenetrationPower    int     `json:"penetrationPower"`
		Accuracy            int     `json:"accuracy"`
		Recoil              int     `json:"recoil"`
		InitialSpeed        int     `json:"initialSpeed"`
	} `json:"ballistics"`
}
