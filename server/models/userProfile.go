package models

type UserProfiles struct {
	UserId                 uint     `gorm:"unique;not null:true" json:"user_id"`
	Age                    *uint    `gorm:"not null:false" json:"age"`
	Height                 *uint    `gorm:"not null:false" json:"height"`
	Height_ft              *uint    `gorm:"not null:false" json:"height_ft"`
	Height_in              *uint    `gorm:"not null:false" json:"height_in"`
	Weight                 *float64 `gorm:"not null:false;type:float(3, 1)" json:"weight"`
	Weight_lbs             *float64 `gorm:"not null:false;type:float(3, 1)" json:"weight_lbs"`
	Sex                    *string  `gorm:"not null:false" json:"sex"`
	Activity_level         *string  `gorm:"not null:false" json:"activity_level"`
	MBR                    *uint    `gorm:"not null:false" json:"mbr"`
	Calorie_consumption    *uint    `gorm:"not null:false" json:"calorie_consumption"`
	HasWeightGoal          *bool    `gorm:"not null:false;defeault:false" json:"has_weigh_goal"`
	HasNutritionGoal       *bool    `gorm:"not null:false;defeault:false" json:"has_nutrition_goal"`
	Weight_goal            *string  `gorm:"not null:false;" json:"weight_goal"`
	Goal_weight            *uint    `gorm:"not null:false;" json:"goal_weight"`
	Calorie_goal           *uint    `gorm:"not null:false;" json:"calorie_goal"`
	Nutrition_protein_perc *uint    `gorm:"not null:false;" json:"nutrition_protein_perc"`
	Nutrition_carb_perc    *uint    `gorm:"not null:false;" json:"nutrition_carb_perc"`
	Nutrition_fat_perc     *uint    `gorm:"not null:false;" json:"nutrition_fat_perc"`
}
