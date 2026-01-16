package product

type Product struct {
	ID_                      string        `json:"_id" bson:"_id"`
	Complete                 int           `json:"complete" bson:"complete"`
	CategoriesPropertiesTags []string      `json:"categories_properties_tags" bson:"categories_properties_tags"`
	BrandsHierarchy          []string      `json:"brands_hierarchy" bson:"brands_hierarchy"`
	Lc                       string        `json:"lc" bson:"lc"`
	EnvironmentalScoreTags   []string      `json:"environmental_score_tags" bson:"environmental_score_tags"`
	NutriscoreVersion        string        `json:"nutriscore_version" bson:"nutriscore_version"`
	LanguagesTags            []string      `json:"languages_tags" bson:"languages_tags"`
	Code                     string        `json:"code" bson:"code"`
	ServingQuantity          float64       `json:"serving_quantity" bson:"serving_quantity"`
	ProductNameDe            string        `json:"product_name_de" bson:"product_name_de"`
	ID                       string        `json:"id" bson:"id"`
	AllergensHierarchy       []interface{} `json:"allergens_hierarchy" bson:"allergens_hierarchy"`
	NutritionGradesTags      []string      `json:"nutrition_grades_tags" bson:"nutrition_grades_tags"`
	TracesFromIngredients    string        `json:"traces_from_ingredients" bson:"traces_from_ingredients"`
	Traces                   string        `json:"traces" bson:"traces"`
	BrandsOld                string        `json:"brands_old" bson:"brands_old"`
	NutritionDataPer         string        `json:"nutrition_data_per" bson:"nutrition_data_per"`
	NutritionDataPreparedPer string        `json:"nutrition_data_prepared_per" bson:"nutrition_data_prepared_per"`
	Nutriscore2023Tags       []string      `json:"nutriscore_2023_tags" bson:"nutriscore_2023_tags"`
	ServingSize              string        `json:"serving_size" bson:"serving_size"`
	NutritionScoreBeverage   int           `json:"nutrition_score_beverage" bson:"nutrition_score_beverage"`
	Lang                     string        `json:"lang" bson:"lang"`
	CountriesTags            []string      `json:"countries_tags" bson:"countries_tags"`
	AllergensFromIngredients string        `json:"allergens_from_ingredients" bson:"allergens_from_ingredients"`
	CreatedT                 int           `json:"created_t" bson:"created_t"`
	NutritionGradeFr         string        `json:"nutrition_grade_fr" bson:"nutrition_grade_fr"`
	NutritionGrades          string        `json:"nutrition_grades" bson:"nutrition_grades"`
	NutritionData            string        `json:"nutrition_data" bson:"nutrition_data"`
	DataQualityTags          []string      `json:"data_quality_tags" bson:"data_quality_tags"`
	Countries                string        `json:"countries" bson:"countries"`
	MiscTags                 []string      `json:"misc_tags" bson:"misc_tags"`
	CodesTags                []string      `json:"codes_tags" bson:"codes_tags"`
	UniqueScansN             int           `json:"unique_scans_n" bson:"unique_scans_n"`
	FoodGroupsTags           []interface{} `json:"food_groups_tags" bson:"food_groups_tags"`
	States                   string        `json:"states" bson:"states"`
	ProductQuantity          float64       `json:"product_quantity" bson:"product_quantity"`
	Brands                   string        `json:"brands" bson:"brands"`
	Quantity                 string        `json:"quantity" bson:"quantity"`
	EnvironmentalScoreData   struct {
		Status  string `json:"status" bson:"status"`
		Missing struct {
			Categories  int `json:"categories" bson:"categories"`
			Labels      int `json:"labels" bson:"labels"`
			Ingredients int `json:"ingredients" bson:"ingredients"`
			Packagings  int `json:"packagings" bson:"packagings"`
			Origins     int `json:"origins" bson:"origins"`
		} `json:"missing" bson:"missing"`
		MissingAgribalyseMatchWarning int `json:"missing_agribalyse_match_warning" bson:"missing_agribalyse_match_warning"`
		Agribalyse                    struct {
			Warning string `json:"warning" bson:"warning"`
		} `json:"agribalyse" bson:"agribalyse"`
		MissingKeyData int `json:"missing_key_data" bson:"missing_key_data"`
		Adjustments    struct {
			OriginsOfIngredients struct {
				EpiValue int `json:"epi_value" bson:"epi_value"`
				Values   struct {
					Lb    int `json:"lb" bson:"lb"`
					Es    int `json:"es" bson:"es"`
					Fr    int `json:"fr" bson:"fr"`
					Lu    int `json:"lu" bson:"lu"`
					De    int `json:"de" bson:"de"`
					Nl    int `json:"nl" bson:"nl"`
					Is    int `json:"is" bson:"is"`
					Al    int `json:"al" bson:"al"`
					Xk    int `json:"xk" bson:"xk"`
					Ro    int `json:"ro" bson:"ro"`
					Ch    int `json:"ch" bson:"ch"`
					Dz    int `json:"dz" bson:"dz"`
					Ps    int `json:"ps" bson:"ps"`
					Mc    int `json:"mc" bson:"mc"`
					Ax    int `json:"ax" bson:"ax"`
					Md    int `json:"md" bson:"md"`
					Hr    int `json:"hr" bson:"hr"`
					Fi    int `json:"fi" bson:"fi"`
					Dk    int `json:"dk" bson:"dk"`
					Tr    int `json:"tr" bson:"tr"`
					Cz    int `json:"cz" bson:"cz"`
					Cy    int `json:"cy" bson:"cy"`
					No    int `json:"no" bson:"no"`
					At    int `json:"at" bson:"at"`
					Fo    int `json:"fo" bson:"fo"`
					Us    int `json:"us" bson:"us"`
					Tn    int `json:"tn" bson:"tn"`
					Lv    int `json:"lv" bson:"lv"`
					Je    int `json:"je" bson:"je"`
					Ly    int `json:"ly" bson:"ly"`
					Bg    int `json:"bg" bson:"bg"`
					Sk    int `json:"sk" bson:"sk"`
					World int `json:"world" bson:"world"`
					Sm    int `json:"sm" bson:"sm"`
					Gg    int `json:"gg" bson:"gg"`
					Me    int `json:"me" bson:"me"`
					It    int `json:"it" bson:"it"`
					Be    int `json:"be" bson:"be"`
					Rs    int `json:"rs" bson:"rs"`
					Si    int `json:"si" bson:"si"`
					Im    int `json:"im" bson:"im"`
					Pl    int `json:"pl" bson:"pl"`
					Gr    int `json:"gr" bson:"gr"`
					Pt    int `json:"pt" bson:"pt"`
					Lt    int `json:"lt" bson:"lt"`
					Mk    int `json:"mk" bson:"mk"`
					Sj    int `json:"sj" bson:"sj"`
					Hu    int `json:"hu" bson:"hu"`
					Li    int `json:"li" bson:"li"`
					Se    int `json:"se" bson:"se"`
					Uk    int `json:"uk" bson:"uk"`
					Il    int `json:"il" bson:"il"`
					Sy    int `json:"sy" bson:"sy"`
					Ma    int `json:"ma" bson:"ma"`
					Ie    int `json:"ie" bson:"ie"`
					Mt    int `json:"mt" bson:"mt"`
					Ba    int `json:"ba" bson:"ba"`
					Eg    int `json:"eg" bson:"eg"`
					Va    int `json:"va" bson:"va"`
					Gi    int `json:"gi" bson:"gi"`
					Ua    int `json:"ua" bson:"ua"`
					Ee    int `json:"ee" bson:"ee"`
					Ad    int `json:"ad" bson:"ad"`
				} `json:"values" bson:"values"`
				TransportationValues struct {
					Cy    int `json:"cy" bson:"cy"`
					Tr    int `json:"tr" bson:"tr"`
					Cz    int `json:"cz" bson:"cz"`
					Dk    int `json:"dk" bson:"dk"`
					At    int `json:"at" bson:"at"`
					No    int `json:"no" bson:"no"`
					Mc    int `json:"mc" bson:"mc"`
					Ps    int `json:"ps" bson:"ps"`
					Fi    int `json:"fi" bson:"fi"`
					Hr    int `json:"hr" bson:"hr"`
					Md    int `json:"md" bson:"md"`
					Ax    int `json:"ax" bson:"ax"`
					Us    int `json:"us" bson:"us"`
					Fo    int `json:"fo" bson:"fo"`
					Tn    int `json:"tn" bson:"tn"`
					Nl    int `json:"nl" bson:"nl"`
					Is    int `json:"is" bson:"is"`
					Es    int `json:"es" bson:"es"`
					Lb    int `json:"lb" bson:"lb"`
					Fr    int `json:"fr" bson:"fr"`
					De    int `json:"de" bson:"de"`
					Lu    int `json:"lu" bson:"lu"`
					Dz    int `json:"dz" bson:"dz"`
					Ch    int `json:"ch" bson:"ch"`
					Xk    int `json:"xk" bson:"xk"`
					Al    int `json:"al" bson:"al"`
					Ro    int `json:"ro" bson:"ro"`
					Se    int `json:"se" bson:"se"`
					Uk    int `json:"uk" bson:"uk"`
					Mk    int `json:"mk" bson:"mk"`
					Sj    int `json:"sj" bson:"sj"`
					Hu    int `json:"hu" bson:"hu"`
					Pt    int `json:"pt" bson:"pt"`
					Lt    int `json:"lt" bson:"lt"`
					Li    int `json:"li" bson:"li"`
					Ee    int `json:"ee" bson:"ee"`
					Ua    int `json:"ua" bson:"ua"`
					Gi    int `json:"gi" bson:"gi"`
					Va    int `json:"va" bson:"va"`
					Ad    int `json:"ad" bson:"ad"`
					Sy    int `json:"sy" bson:"sy"`
					Il    int `json:"il" bson:"il"`
					Ba    int `json:"ba" bson:"ba"`
					Eg    int `json:"eg" bson:"eg"`
					Mt    int `json:"mt" bson:"mt"`
					Ie    int `json:"ie" bson:"ie"`
					Ma    int `json:"ma" bson:"ma"`
					Gg    int `json:"gg" bson:"gg"`
					World int `json:"world" bson:"world"`
					Sm    int `json:"sm" bson:"sm"`
					Be    int `json:"be" bson:"be"`
					Rs    int `json:"rs" bson:"rs"`
					It    int `json:"it" bson:"it"`
					Me    int `json:"me" bson:"me"`
					Bg    int `json:"bg" bson:"bg"`
					Ly    int `json:"ly" bson:"ly"`
					Je    int `json:"je" bson:"je"`
					Lv    int `json:"lv" bson:"lv"`
					Sk    int `json:"sk" bson:"sk"`
					Gr    int `json:"gr" bson:"gr"`
					Pl    int `json:"pl" bson:"pl"`
					Im    int `json:"im" bson:"im"`
					Si    int `json:"si" bson:"si"`
				} `json:"transportation_values" bson:"transportation_values"`
				AggregatedOrigins []struct {
					Percent int    `json:"percent" bson:"percent"`
					Origin  string `json:"origin" bson:"origin"`
				} `json:"aggregated_origins" bson:"aggregated_origins"`
				OriginsFromOriginsField []string `json:"origins_from_origins_field" bson:"origins_from_origins_field"`
				TransportationScores    struct {
					Ch    float64 `json:"ch" bson:"ch"`
					Dz    float64 `json:"dz" bson:"dz"`
					Ro    float64 `json:"ro" bson:"ro"`
					Xk    float64 `json:"xk" bson:"xk"`
					Al    float64 `json:"al" bson:"al"`
					Is    float64 `json:"is" bson:"is"`
					Nl    float64 `json:"nl" bson:"nl"`
					De    float64 `json:"de" bson:"de"`
					Lu    float64 `json:"lu" bson:"lu"`
					Fr    float64 `json:"fr" bson:"fr"`
					Lb    float64 `json:"lb" bson:"lb"`
					Es    float64 `json:"es" bson:"es"`
					Tn    float64 `json:"tn" bson:"tn"`
					Us    float64 `json:"us" bson:"us"`
					Fo    float64 `json:"fo" bson:"fo"`
					At    float64 `json:"at" bson:"at"`
					No    float64 `json:"no" bson:"no"`
					Cy    float64 `json:"cy" bson:"cy"`
					Dk    float64 `json:"dk" bson:"dk"`
					Tr    float64 `json:"tr" bson:"tr"`
					Cz    float64 `json:"cz" bson:"cz"`
					Hr    float64 `json:"hr" bson:"hr"`
					Fi    float64 `json:"fi" bson:"fi"`
					Ax    float64 `json:"ax" bson:"ax"`
					Md    float64 `json:"md" bson:"md"`
					Ps    float64 `json:"ps" bson:"ps"`
					Mc    float64 `json:"mc" bson:"mc"`
					Pl    float64 `json:"pl" bson:"pl"`
					Gr    float64 `json:"gr" bson:"gr"`
					Si    float64 `json:"si" bson:"si"`
					Im    float64 `json:"im" bson:"im"`
					Rs    float64 `json:"rs" bson:"rs"`
					Be    float64 `json:"be" bson:"be"`
					Me    float64 `json:"me" bson:"me"`
					It    float64 `json:"it" bson:"it"`
					Sm    float64 `json:"sm" bson:"sm"`
					World float64 `json:"world" bson:"world"`
					Gg    float64 `json:"gg" bson:"gg"`
					Sk    float64 `json:"sk" bson:"sk"`
					Lv    float64 `json:"lv" bson:"lv"`
					Je    float64 `json:"je" bson:"je"`
					Ly    float64 `json:"ly" bson:"ly"`
					Bg    float64 `json:"bg" bson:"bg"`
					Ad    float64 `json:"ad" bson:"ad"`
					Ee    float64 `json:"ee" bson:"ee"`
					Va    float64 `json:"va" bson:"va"`
					Gi    float64 `json:"gi" bson:"gi"`
					Ua    float64 `json:"ua" bson:"ua"`
					Eg    float64 `json:"eg" bson:"eg"`
					Ba    float64 `json:"ba" bson:"ba"`
					Ie    float64 `json:"ie" bson:"ie"`
					Ma    float64 `json:"ma" bson:"ma"`
					Mt    float64 `json:"mt" bson:"mt"`
					Il    float64 `json:"il" bson:"il"`
					Sy    float64 `json:"sy" bson:"sy"`
					Uk    float64 `json:"uk" bson:"uk"`
					Se    float64 `json:"se" bson:"se"`
					Li    float64 `json:"li" bson:"li"`
					Sj    float64 `json:"sj" bson:"sj"`
					Mk    float64 `json:"mk" bson:"mk"`
					Hu    float64 `json:"hu" bson:"hu"`
					Pt    float64 `json:"pt" bson:"pt"`
					Lt    float64 `json:"lt" bson:"lt"`
				} `json:"transportation_scores" bson:"transportation_scores"`
				Warning               string   `json:"warning" bson:"warning"`
				EpiScore              int      `json:"epi_score" bson:"epi_score"`
				OriginsFromCategories []string `json:"origins_from_categories" bson:"origins_from_categories"`
			} `json:"origins_of_ingredients" bson:"origins_of_ingredients"`
			ThreatenedSpecies struct {
				Warning string `json:"warning" bson:"warning"`
			} `json:"threatened_species" bson:"threatened_species"`
			ProductionSystem struct {
				Labels  []interface{} `json:"labels" bson:"labels"`
				Value   int           `json:"value" bson:"value"`
				Warning string        `json:"warning" bson:"warning"`
			} `json:"production_system" bson:"production_system"`
			Packaging struct {
				Warning string `json:"warning" bson:"warning"`
				Value   int    `json:"value" bson:"value"`
			} `json:"packaging" bson:"packaging"`
		} `json:"adjustments" bson:"adjustments"`
		Grade string `json:"grade" bson:"grade"`
	} `json:"environmental_score_data" bson:"environmental_score_data"`
	UnknownNutrientsTags []interface{} `json:"unknown_nutrients_tags" bson:"unknown_nutrients_tags"`
	ScansN               int           `json:"scans_n" bson:"scans_n"`
	WeighersTags         []interface{} `json:"weighers_tags" bson:"weighers_tags"`
	CategoriesProperties struct {
	} `json:"categories_properties" bson:"categories_properties"`
	BrandsLc            string        `json:"brands_lc" bson:"brands_lc"`
	ProductType         string        `json:"product_type" bson:"product_type"`
	NoNutritionData     string        `json:"no_nutrition_data" bson:"no_nutrition_data"`
	Packagings          []interface{} `json:"packagings" bson:"packagings"`
	PackagingsMaterials struct {
	} `json:"packagings_materials" bson:"packagings_materials"`
	DataSourcesTags    []string      `json:"data_sources_tags" bson:"data_sources_tags"`
	NutriscoreGrade    string        `json:"nutriscore_grade" bson:"nutriscore_grade"`
	NovaGroupsTags     []string      `json:"nova_groups_tags" bson:"nova_groups_tags"`
	AllergensFromUser  string        `json:"allergens_from_user" bson:"allergens_from_user"`
	LastModifiedT      int           `json:"last_modified_t" bson:"last_modified_t"`
	CountriesHierarchy []string      `json:"countries_hierarchy" bson:"countries_hierarchy"`
	TracesTags         []interface{} `json:"traces_tags" bson:"traces_tags"`
	Rev                int           `json:"rev" bson:"rev"`
	Nutriments         struct {
		Energy100G              int     `json:"energy_100g" bson:"energy_100g"`
		Fat                     int     `json:"fat" bson:"fat"`
		EnergyKcal              int     `json:"energy-kcal" bson:"energy-kcal"`
		SugarsValue             float64 `json:"sugars_value" bson:"sugars_value"`
		SugarsUnit              string  `json:"sugars_unit" bson:"sugars_unit"`
		Carbohydrates           float64 `json:"carbohydrates" bson:"carbohydrates"`
		Carbohydrates100G       float64 `json:"carbohydrates_100g" bson:"carbohydrates_100g"`
		Energy                  int     `json:"energy" bson:"energy"`
		SugarsServing           float64 `json:"sugars_serving" bson:"sugars_serving"`
		Fat100G                 int     `json:"fat_100g" bson:"fat_100g"`
		ProteinsValue           int     `json:"proteins_value" bson:"proteins_value"`
		EnergyKcalUnit          string  `json:"energy-kcal_unit" bson:"energy-kcal_unit"`
		CarbohydratesValue      float64 `json:"carbohydrates_value" bson:"carbohydrates_value"`
		Sugars100G              float64 `json:"sugars_100g" bson:"sugars_100g"`
		FatUnit                 string  `json:"fat_unit" bson:"fat_unit"`
		CarbohydratesServing    float64 `json:"carbohydrates_serving" bson:"carbohydrates_serving"`
		Proteins                int     `json:"proteins" bson:"proteins"`
		EnergyValue             int     `json:"energy_value" bson:"energy_value"`
		EnergyUnit              string  `json:"energy_unit" bson:"energy_unit"`
		ProteinsUnit            string  `json:"proteins_unit" bson:"proteins_unit"`
		EnergyKcalValue         int     `json:"energy-kcal_value" bson:"energy-kcal_value"`
		EnergyKcal100G          int     `json:"energy-kcal_100g" bson:"energy-kcal_100g"`
		EnergyKcalServing       float64 `json:"energy-kcal_serving" bson:"energy-kcal_serving"`
		FatServing              int     `json:"fat_serving" bson:"fat_serving"`
		CarbohydratesUnit       string  `json:"carbohydrates_unit" bson:"carbohydrates_unit"`
		FatValue                int     `json:"fat_value" bson:"fat_value"`
		Proteins100G            int     `json:"proteins_100g" bson:"proteins_100g"`
		ProteinsServing         int     `json:"proteins_serving" bson:"proteins_serving"`
		EnergyKcalValueComputed int     `json:"energy-kcal_value_computed" bson:"energy-kcal_value_computed"`
		EnergyServing           int     `json:"energy_serving" bson:"energy_serving"`
		Sugars                  float64 `json:"sugars" bson:"sugars"`
	} `json:"nutriments" bson:"nutriments"`
	ServingQuantityUnit string   `json:"serving_quantity_unit" bson:"serving_quantity_unit"`
	CountriesLc         string   `json:"countries_lc" bson:"countries_lc"`
	NovaGroupDebug      string   `json:"nova_group_debug" bson:"nova_group_debug"`
	NutriscoreTags      []string `json:"nutriscore_tags" bson:"nutriscore_tags"`
	Nutriscore          struct {
		Num2021 struct {
			Data struct {
				Sugars                                   float64     `json:"sugars" bson:"sugars"`
				Energy                                   int         `json:"energy" bson:"energy"`
				IsBeverage                               int         `json:"is_beverage" bson:"is_beverage"`
				IsWater                                  int         `json:"is_water" bson:"is_water"`
				Sodium                                   interface{} `json:"sodium" bson:"sodium"`
				IsCheese                                 int         `json:"is_cheese" bson:"is_cheese"`
				IsFat                                    int         `json:"is_fat" bson:"is_fat"`
				Proteins                                 int         `json:"proteins" bson:"proteins"`
				SaturatedFat                             int         `json:"saturated_fat" bson:"saturated_fat"`
				Fiber                                    int         `json:"fiber" bson:"fiber"`
				SaturatedFatRatio                        int         `json:"saturated_fat_ratio" bson:"saturated_fat_ratio"`
				FruitsVegetablesNutsColzaWalnutOliveOils int         `json:"fruits_vegetables_nuts_colza_walnut_olive_oils" bson:"fruits_vegetables_nuts_colza_walnut_olive_oils"`
			} `json:"data" bson:"data"`
			Grade                string `json:"grade" bson:"grade"`
			NutriscoreComputed   int    `json:"nutriscore_computed" bson:"nutriscore_computed"`
			NutriscoreApplicable int    `json:"nutriscore_applicable" bson:"nutriscore_applicable"`
			NutrientsAvailable   int    `json:"nutrients_available" bson:"nutrients_available"`
			CategoryAvailable    int    `json:"category_available" bson:"category_available"`
		} `json:"2021" bson:"2021"`
		Num2023 struct {
			NutriscoreComputed int    `json:"nutriscore_computed" bson:"nutriscore_computed"`
			Grade              string `json:"grade" bson:"grade"`
			Data               struct {
				Proteins                int         `json:"proteins" bson:"proteins"`
				IsFatOilNutsSeeds       int         `json:"is_fat_oil_nuts_seeds" bson:"is_fat_oil_nuts_seeds"`
				SaturatedFat            int         `json:"saturated_fat" bson:"saturated_fat"`
				Fiber                   interface{} `json:"fiber" bson:"fiber"`
				FruitsVegetablesLegumes interface{} `json:"fruits_vegetables_legumes" bson:"fruits_vegetables_legumes"`
				SaturatedFatRatio       int         `json:"saturated_fat_ratio" bson:"saturated_fat_ratio"`
				Salt                    interface{} `json:"salt" bson:"salt"`
				Sugars                  float64     `json:"sugars" bson:"sugars"`
				IsRedMeatProduct        int         `json:"is_red_meat_product" bson:"is_red_meat_product"`
				IsBeverage              int         `json:"is_beverage" bson:"is_beverage"`
				Energy                  int         `json:"energy" bson:"energy"`
				IsWater                 int         `json:"is_water" bson:"is_water"`
				IsCheese                int         `json:"is_cheese" bson:"is_cheese"`
			} `json:"data" bson:"data"`
			CategoryAvailable    int `json:"category_available" bson:"category_available"`
			NutrientsAvailable   int `json:"nutrients_available" bson:"nutrients_available"`
			NutriscoreApplicable int `json:"nutriscore_applicable" bson:"nutriscore_applicable"`
		} `json:"2023" bson:"2023"`
	} `json:"nutriscore" bson:"nutriscore"`
	Completeness                                float64       `json:"completeness" bson:"completeness"`
	ProductName                                 string        `json:"product_name" bson:"product_name"`
	PackagingMaterialsTags                      []interface{} `json:"packaging_materials_tags" bson:"packaging_materials_tags"`
	Keywords                                    []string      `json:"_keywords" bson:"_keywords"`
	LastUpdatedT                                int           `json:"last_updated_t" bson:"last_updated_t"`
	EntryDatesTags                              []string      `json:"entry_dates_tags" bson:"entry_dates_tags"`
	LanguagesHierarchy                          []string      `json:"languages_hierarchy" bson:"languages_hierarchy"`
	MainCountriesTags                           []interface{} `json:"main_countries_tags" bson:"main_countries_tags"`
	DataQualityInfoTags                         []string      `json:"data_quality_info_tags" bson:"data_quality_info_tags"`
	Nutriscore2021Tags                          []string      `json:"nutriscore_2021_tags" bson:"nutriscore_2021_tags"`
	MaxImgid                                    string        `json:"max_imgid" bson:"max_imgid"`
	StatesTags                                  []string      `json:"states_tags" bson:"states_tags"`
	NutritionScoreWarningNoFruitsVegetablesNuts int           `json:"nutrition_score_warning_no_fruits_vegetables_nuts" bson:"nutrition_score_warning_no_fruits_vegetables_nuts"`
	EnvironmentalScoreGrade                     string        `json:"environmental_score_grade" bson:"environmental_score_grade"`
	PopularityTags                              []string      `json:"popularity_tags" bson:"popularity_tags"`
	PackagingRecyclingTags                      []interface{} `json:"packaging_recycling_tags" bson:"packaging_recycling_tags"`
	CheckersTags                                []interface{} `json:"checkers_tags" bson:"checkers_tags"`
	NutrientLevels                              struct {
	} `json:"nutrient_levels" bson:"nutrient_levels"`
	PnnsGroups1Tags              []string      `json:"pnns_groups_1_tags" bson:"pnns_groups_1_tags"`
	LastEditDatesTags            []string      `json:"last_edit_dates_tags" bson:"last_edit_dates_tags"`
	NutritionScoreWarningNoFiber int           `json:"nutrition_score_warning_no_fiber" bson:"nutrition_score_warning_no_fiber"`
	InformersTags                []string      `json:"informers_tags" bson:"informers_tags"`
	PnnsGroups2                  string        `json:"pnns_groups_2" bson:"pnns_groups_2"`
	BrandsTags                   []string      `json:"brands_tags" bson:"brands_tags"`
	Allergens                    string        `json:"allergens" bson:"allergens"`
	PnnsGroups1                  string        `json:"pnns_groups_1" bson:"pnns_groups_1"`
	EcoscoreTags                 []string      `json:"ecoscore_tags" bson:"ecoscore_tags"`
	PnnsGroups2Tags              []string      `json:"pnns_groups_2_tags" bson:"pnns_groups_2_tags"`
	PopularityKey                int           `json:"popularity_key" bson:"popularity_key"`
	AllergensTags                []interface{} `json:"allergens_tags" bson:"allergens_tags"`
}
