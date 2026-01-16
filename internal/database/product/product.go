package product

type Product struct {
	ID_                      string        `json:"_id" bson:"_id"`
	Complete                 int           `json:"complete" bson:"complete,truncate"`
	CategoriesPropertiesTags []string      `json:"categories_properties_tags" bson:"categories_properties_tags"`
	BrandsHierarchy          []string      `json:"brands_hierarchy" bson:"brands_hierarchy"`
	Lc                       string        `json:"lc" bson:"lc"`
	EnvironmentalScoreTags   []string      `json:"environmental_score_tags" bson:"environmental_score_tags"`
	NutriscoreVersion        string        `json:"nutriscore_version" bson:"nutriscore_version"`
	LanguagesTags            []string      `json:"languages_tags" bson:"languages_tags"`
	Code                     string        `json:"code" bson:"code"`
	ServingQuantity          float64       `json:"serving_quantity" bson:"serving_quantity,truncate"`
	ProductNameDe            string        `json:"product_name_de" bson:"product_name_de"`
	ID                       string        `json:"id" bson:"id"`
	AllergensHierarchy       []interface{} `json:"allergens_hierarchy" bson:"allergens_hierarchy,truncate"`
	NutritionGradesTags      []string      `json:"nutrition_grades_tags" bson:"nutrition_grades_tags"`
	TracesFromIngredients    string        `json:"traces_from_ingredients" bson:"traces_from_ingredients"`
	Traces                   string        `json:"traces" bson:"traces"`
	BrandsOld                string        `json:"brands_old" bson:"brands_old"`
	NutritionDataPer         string        `json:"nutrition_data_per" bson:"nutrition_data_per"`
	NutritionDataPreparedPer string        `json:"nutrition_data_prepared_per" bson:"nutrition_data_prepared_per"`
	Nutriscore2023Tags       []string      `json:"nutriscore_2023_tags" bson:"nutriscore_2023_tags"`
	ServingSize              string        `json:"serving_size" bson:"serving_size"`
	NutritionScoreBeverage   int           `json:"nutrition_score_beverage" bson:"nutrition_score_beverage,truncate"`
	Lang                     string        `json:"lang" bson:"lang"`
	CountriesTags            []string      `json:"countries_tags" bson:"countries_tags"`
	AllergensFromIngredients string        `json:"allergens_from_ingredients" bson:"allergens_from_ingredients"`
	CreatedT                 int           `json:"created_t" bson:"created_t,truncate"`
	NutritionGradeFr         string        `json:"nutrition_grade_fr" bson:"nutrition_grade_fr"`
	NutritionGrades          string        `json:"nutrition_grades" bson:"nutrition_grades"`
	NutritionData            string        `json:"nutrition_data" bson:"nutrition_data"`
	DataQualityTags          []string      `json:"data_quality_tags" bson:"data_quality_tags"`
	Countries                string        `json:"countries" bson:"countries"`
	MiscTags                 []string      `json:"misc_tags" bson:"misc_tags"`
	CodesTags                []string      `json:"codes_tags" bson:"codes_tags"`
	UniqueScansN             int           `json:"unique_scans_n" bson:"unique_scans_n,truncate"`
	FoodGroupsTags           []interface{} `json:"food_groups_tags" bson:"food_groups_tags,truncate"`
	States                   string        `json:"states" bson:"states"`
	ProductQuantity          float64       `json:"product_quantity" bson:"product_quantity,truncate"`
	Brands                   string        `json:"brands" bson:"brands"`
	Quantity                 string        `json:"quantity" bson:"quantity"`
	EnvironmentalScoreData   struct {
		Status  string `json:"status" bson:"status"`
		Missing struct {
			Categories  int `json:"categories" bson:"categories,truncate"`
			Labels      int `json:"labels" bson:"labels,truncate"`
			Ingredients int `json:"ingredients" bson:"ingredients,truncate"`
			Packagings  int `json:"packagings" bson:"packagings,truncate"`
			Origins     int `json:"origins" bson:"origins,truncate"`
		} `json:"missing" bson:"missing"`
		MissingAgribalyseMatchWarning int `json:"missing_agribalyse_match_warning" bson:"missing_agribalyse_match_warning,truncate"`
		Agribalyse                    struct {
			Warning string `json:"warning" bson:"warning"`
		} `json:"agribalyse" bson:"agribalyse"`
		MissingKeyData int `json:"missing_key_data" bson:"missing_key_data,truncate"`
		Adjustments    struct {
			OriginsOfIngredients struct {
				EpiValue int `json:"epi_value" bson:"epi_value,truncate"`
				Values   struct {
					Lb    int `json:"lb" bson:"lb,truncate"`
					Es    int `json:"es" bson:"es,truncate"`
					Fr    int `json:"fr" bson:"fr,truncate"`
					Lu    int `json:"lu" bson:"lu,truncate"`
					De    int `json:"de" bson:"de,truncate"`
					Nl    int `json:"nl" bson:"nl,truncate"`
					Is    int `json:"is" bson:"is,truncate"`
					Al    int `json:"al" bson:"al,truncate"`
					Xk    int `json:"xk" bson:"xk,truncate"`
					Ro    int `json:"ro" bson:"ro,truncate"`
					Ch    int `json:"ch" bson:"ch,truncate"`
					Dz    int `json:"dz" bson:"dz,truncate"`
					Ps    int `json:"ps" bson:"ps,truncate"`
					Mc    int `json:"mc" bson:"mc,truncate"`
					Ax    int `json:"ax" bson:"ax,truncate"`
					Md    int `json:"md" bson:"md,truncate"`
					Hr    int `json:"hr" bson:"hr,truncate"`
					Fi    int `json:"fi" bson:"fi,truncate"`
					Dk    int `json:"dk" bson:"dk,truncate"`
					Tr    int `json:"tr" bson:"tr,truncate"`
					Cz    int `json:"cz" bson:"cz,truncate"`
					Cy    int `json:"cy" bson:"cy,truncate"`
					No    int `json:"no" bson:"no,truncate"`
					At    int `json:"at" bson:"at,truncate"`
					Fo    int `json:"fo" bson:"fo,truncate"`
					Us    int `json:"us" bson:"us,truncate"`
					Tn    int `json:"tn" bson:"tn,truncate"`
					Lv    int `json:"lv" bson:"lv,truncate"`
					Je    int `json:"je" bson:"je,truncate"`
					Ly    int `json:"ly" bson:"ly,truncate"`
					Bg    int `json:"bg" bson:"bg,truncate"`
					Sk    int `json:"sk" bson:"sk,truncate"`
					World int `json:"world" bson:"world,truncate"`
					Sm    int `json:"sm" bson:"sm,truncate"`
					Gg    int `json:"gg" bson:"gg,truncate"`
					Me    int `json:"me" bson:"me,truncate"`
					It    int `json:"it" bson:"it,truncate"`
					Be    int `json:"be" bson:"be,truncate"`
					Rs    int `json:"rs" bson:"rs,truncate"`
					Si    int `json:"si" bson:"si,truncate"`
					Im    int `json:"im" bson:"im,truncate"`
					Pl    int `json:"pl" bson:"pl,truncate"`
					Gr    int `json:"gr" bson:"gr,truncate"`
					Pt    int `json:"pt" bson:"pt,truncate"`
					Lt    int `json:"lt" bson:"lt,truncate"`
					Mk    int `json:"mk" bson:"mk,truncate"`
					Sj    int `json:"sj" bson:"sj,truncate"`
					Hu    int `json:"hu" bson:"hu,truncate"`
					Li    int `json:"li" bson:"li,truncate"`
					Se    int `json:"se" bson:"se,truncate"`
					Uk    int `json:"uk" bson:"uk,truncate"`
					Il    int `json:"il" bson:"il,truncate"`
					Sy    int `json:"sy" bson:"sy,truncate"`
					Ma    int `json:"ma" bson:"ma,truncate"`
					Ie    int `json:"ie" bson:"ie,truncate"`
					Mt    int `json:"mt" bson:"mt,truncate"`
					Ba    int `json:"ba" bson:"ba,truncate"`
					Eg    int `json:"eg" bson:"eg,truncate"`
					Va    int `json:"va" bson:"va,truncate"`
					Gi    int `json:"gi" bson:"gi,truncate"`
					Ua    int `json:"ua" bson:"ua,truncate"`
					Ee    int `json:"ee" bson:"ee,truncate"`
					Ad    int `json:"ad" bson:"ad,truncate"`
				} `json:"values" bson:"values"`
				TransportationValues struct {
					Cy    int `json:"cy" bson:"cy,truncate"`
					Tr    int `json:"tr" bson:"tr,truncate"`
					Cz    int `json:"cz" bson:"cz,truncate"`
					Dk    int `json:"dk" bson:"dk,truncate"`
					At    int `json:"at" bson:"at,truncate"`
					No    int `json:"no" bson:"no,truncate"`
					Mc    int `json:"mc" bson:"mc,truncate"`
					Ps    int `json:"ps" bson:"ps,truncate"`
					Fi    int `json:"fi" bson:"fi,truncate"`
					Hr    int `json:"hr" bson:"hr,truncate"`
					Md    int `json:"md" bson:"md,truncate"`
					Ax    int `json:"ax" bson:"ax,truncate"`
					Us    int `json:"us" bson:"us,truncate"`
					Fo    int `json:"fo" bson:"fo,truncate"`
					Tn    int `json:"tn" bson:"tn,truncate"`
					Nl    int `json:"nl" bson:"nl,truncate"`
					Is    int `json:"is" bson:"is,truncate"`
					Es    int `json:"es" bson:"es,truncate"`
					Lb    int `json:"lb" bson:"lb,truncate"`
					Fr    int `json:"fr" bson:"fr,truncate"`
					De    int `json:"de" bson:"de,truncate"`
					Lu    int `json:"lu" bson:"lu,truncate"`
					Dz    int `json:"dz" bson:"dz,truncate"`
					Ch    int `json:"ch" bson:"ch,truncate"`
					Xk    int `json:"xk" bson:"xk,truncate"`
					Al    int `json:"al" bson:"al,truncate"`
					Ro    int `json:"ro" bson:"ro,truncate"`
					Se    int `json:"se" bson:"se,truncate"`
					Uk    int `json:"uk" bson:"uk,truncate"`
					Mk    int `json:"mk" bson:"mk,truncate"`
					Sj    int `json:"sj" bson:"sj,truncate"`
					Hu    int `json:"hu" bson:"hu,truncate"`
					Pt    int `json:"pt" bson:"pt,truncate"`
					Lt    int `json:"lt" bson:"lt,truncate"`
					Li    int `json:"li" bson:"li,truncate"`
					Ee    int `json:"ee" bson:"ee,truncate"`
					Ua    int `json:"ua" bson:"ua,truncate"`
					Gi    int `json:"gi" bson:"gi,truncate"`
					Va    int `json:"va" bson:"va,truncate"`
					Ad    int `json:"ad" bson:"ad,truncate"`
					Sy    int `json:"sy" bson:"sy,truncate"`
					Il    int `json:"il" bson:"il,truncate"`
					Ba    int `json:"ba" bson:"ba,truncate"`
					Eg    int `json:"eg" bson:"eg,truncate"`
					Mt    int `json:"mt" bson:"mt,truncate"`
					Ie    int `json:"ie" bson:"ie,truncate"`
					Ma    int `json:"ma" bson:"ma,truncate"`
					Gg    int `json:"gg" bson:"gg,truncate"`
					World int `json:"world" bson:"world,truncate"`
					Sm    int `json:"sm" bson:"sm,truncate"`
					Be    int `json:"be" bson:"be,truncate"`
					Rs    int `json:"rs" bson:"rs,truncate"`
					It    int `json:"it" bson:"it,truncate"`
					Me    int `json:"me" bson:"me,truncate"`
					Bg    int `json:"bg" bson:"bg,truncate"`
					Ly    int `json:"ly" bson:"ly,truncate"`
					Je    int `json:"je" bson:"je,truncate"`
					Lv    int `json:"lv" bson:"lv,truncate"`
					Sk    int `json:"sk" bson:"sk,truncate"`
					Gr    int `json:"gr" bson:"gr,truncate"`
					Pl    int `json:"pl" bson:"pl,truncate"`
					Im    int `json:"im" bson:"im,truncate"`
					Si    int `json:"si" bson:"si,truncate"`
				} `json:"transportation_values" bson:"transportation_values"`
				AggregatedOrigins []struct {
					Percent int    `json:"percent" bson:"percent,truncate"`
					Origin  string `json:"origin" bson:"origin"`
				} `json:"aggregated_origins" bson:"aggregated_origins"`
				OriginsFromOriginsField []string `json:"origins_from_origins_field" bson:"origins_from_origins_field"`
				TransportationScores    struct {
					Ch    float64 `json:"ch" bson:"ch,truncate"`
					Dz    float64 `json:"dz" bson:"dz,truncate"`
					Ro    float64 `json:"ro" bson:"ro,truncate"`
					Xk    float64 `json:"xk" bson:"xk,truncate"`
					Al    float64 `json:"al" bson:"al,truncate"`
					Is    float64 `json:"is" bson:"is,truncate"`
					Nl    float64 `json:"nl" bson:"nl,truncate"`
					De    float64 `json:"de" bson:"de,truncate"`
					Lu    float64 `json:"lu" bson:"lu,truncate"`
					Fr    float64 `json:"fr" bson:"fr,truncate"`
					Lb    float64 `json:"lb" bson:"lb,truncate"`
					Es    float64 `json:"es" bson:"es,truncate"`
					Tn    float64 `json:"tn" bson:"tn,truncate"`
					Us    float64 `json:"us" bson:"us,truncate"`
					Fo    float64 `json:"fo" bson:"fo,truncate"`
					At    float64 `json:"at" bson:"at,truncate"`
					No    float64 `json:"no" bson:"no,truncate"`
					Cy    float64 `json:"cy" bson:"cy,truncate"`
					Dk    float64 `json:"dk" bson:"dk,truncate"`
					Tr    float64 `json:"tr" bson:"tr,truncate"`
					Cz    float64 `json:"cz" bson:"cz,truncate"`
					Hr    float64 `json:"hr" bson:"hr,truncate"`
					Fi    float64 `json:"fi" bson:"fi,truncate"`
					Ax    float64 `json:"ax" bson:"ax,truncate"`
					Md    float64 `json:"md" bson:"md,truncate"`
					Ps    float64 `json:"ps" bson:"ps,truncate"`
					Mc    float64 `json:"mc" bson:"mc,truncate"`
					Pl    float64 `json:"pl" bson:"pl,truncate"`
					Gr    float64 `json:"gr" bson:"gr,truncate"`
					Si    float64 `json:"si" bson:"si,truncate"`
					Im    float64 `json:"im" bson:"im,truncate"`
					Rs    float64 `json:"rs" bson:"rs,truncate"`
					Be    float64 `json:"be" bson:"be,truncate"`
					Me    float64 `json:"me" bson:"me,truncate"`
					It    float64 `json:"it" bson:"it,truncate"`
					Sm    float64 `json:"sm" bson:"sm,truncate"`
					World float64 `json:"world" bson:"world,truncate"`
					Gg    float64 `json:"gg" bson:"gg,truncate"`
					Sk    float64 `json:"sk" bson:"sk,truncate"`
					Lv    float64 `json:"lv" bson:"lv,truncate"`
					Je    float64 `json:"je" bson:"je,truncate"`
					Ly    float64 `json:"ly" bson:"ly,truncate"`
					Bg    float64 `json:"bg" bson:"bg,truncate"`
					Ad    float64 `json:"ad" bson:"ad,truncate"`
					Ee    float64 `json:"ee" bson:"ee,truncate"`
					Va    float64 `json:"va" bson:"va,truncate"`
					Gi    float64 `json:"gi" bson:"gi,truncate"`
					Ua    float64 `json:"ua" bson:"ua,truncate"`
					Eg    float64 `json:"eg" bson:"eg,truncate"`
					Ba    float64 `json:"ba" bson:"ba,truncate"`
					Ie    float64 `json:"ie" bson:"ie,truncate"`
					Ma    float64 `json:"ma" bson:"ma,truncate"`
					Mt    float64 `json:"mt" bson:"mt,truncate"`
					Il    float64 `json:"il" bson:"il,truncate"`
					Sy    float64 `json:"sy" bson:"sy,truncate"`
					Uk    float64 `json:"uk" bson:"uk,truncate"`
					Se    float64 `json:"se" bson:"se,truncate"`
					Li    float64 `json:"li" bson:"li,truncate"`
					Sj    float64 `json:"sj" bson:"sj,truncate"`
					Mk    float64 `json:"mk" bson:"mk,truncate"`
					Hu    float64 `json:"hu" bson:"hu,truncate"`
					Pt    float64 `json:"pt" bson:"pt,truncate"`
					Lt    float64 `json:"lt" bson:"lt,truncate"`
				} `json:"transportation_scores" bson:"transportation_scores"`
				Warning               string   `json:"warning" bson:"warning"`
				EpiScore              int      `json:"epi_score" bson:"epi_score,truncate"`
				OriginsFromCategories []string `json:"origins_from_categories" bson:"origins_from_categories"`
			} `json:"origins_of_ingredients" bson:"origins_of_ingredients"`
			ThreatenedSpecies struct {
				Warning string `json:"warning" bson:"warning"`
			} `json:"threatened_species" bson:"threatened_species"`
			ProductionSystem struct {
				Labels  []interface{} `json:"labels" bson:"labels,truncate"`
				Value   int           `json:"value" bson:"value,truncate"`
				Warning string        `json:"warning" bson:"warning"`
			} `json:"production_system" bson:"production_system"`
			Packaging struct {
				Warning string `json:"warning" bson:"warning"`
				Value   int    `json:"value" bson:"value,truncate"`
			} `json:"packaging" bson:"packaging"`
		} `json:"adjustments" bson:"adjustments"`
		Grade string `json:"grade" bson:"grade"`
	} `json:"environmental_score_data" bson:"environmental_score_data"`
	UnknownNutrientsTags []interface{} `json:"unknown_nutrients_tags" bson:"unknown_nutrients_tags,truncate"`
	ScansN               int           `json:"scans_n" bson:"scans_n,truncate"`
	WeighersTags         []interface{} `json:"weighers_tags" bson:"weighers_tags,truncate"`
	CategoriesProperties struct {
	} `json:"categories_properties" bson:"categories_properties"`
	BrandsLc            string        `json:"brands_lc" bson:"brands_lc"`
	ProductType         string        `json:"product_type" bson:"product_type"`
	NoNutritionData     string        `json:"no_nutrition_data" bson:"no_nutrition_data"`
	Packagings          []interface{} `json:"packagings" bson:"packagings,truncate"`
	PackagingsMaterials struct {
	} `json:"packagings_materials" bson:"packagings_materials"`
	DataSourcesTags    []string      `json:"data_sources_tags" bson:"data_sources_tags"`
	NutriscoreGrade    string        `json:"nutriscore_grade" bson:"nutriscore_grade"`
	NovaGroupsTags     []string      `json:"nova_groups_tags" bson:"nova_groups_tags"`
	AllergensFromUser  string        `json:"allergens_from_user" bson:"allergens_from_user"`
	LastModifiedT      int           `json:"last_modified_t" bson:"last_modified_t,truncate"`
	CountriesHierarchy []string      `json:"countries_hierarchy" bson:"countries_hierarchy"`
	TracesTags         []interface{} `json:"traces_tags" bson:"traces_tags,truncate"`
	Rev                int           `json:"rev" bson:"rev,truncate"`
	Nutriments         struct {
		Energy100G              int     `json:"energy_100g" bson:"energy_100g,truncate"`
		Fat                     int     `json:"fat" bson:"fat,truncate"`
		EnergyKcal              int     `json:"energy-kcal" bson:"energy-kcal,truncate"`
		SugarsValue             float64 `json:"sugars_value" bson:"sugars_value,truncate"`
		SugarsUnit              string  `json:"sugars_unit" bson:"sugars_unit"`
		Carbohydrates           float64 `json:"carbohydrates" bson:"carbohydrates,truncate"`
		Carbohydrates100G       float64 `json:"carbohydrates_100g" bson:"carbohydrates_100g,truncate"`
		Energy                  int     `json:"energy" bson:"energy,truncate"`
		SugarsServing           float64 `json:"sugars_serving" bson:"sugars_serving,truncate"`
		Fat100G                 int     `json:"fat_100g" bson:"fat_100g,truncate"`
		ProteinsValue           int     `json:"proteins_value" bson:"proteins_value,truncate"`
		EnergyKcalUnit          string  `json:"energy-kcal_unit" bson:"energy-kcal_unit"`
		CarbohydratesValue      float64 `json:"carbohydrates_value" bson:"carbohydrates_value,truncate"`
		Sugars100G              float64 `json:"sugars_100g" bson:"sugars_100g,truncate"`
		FatUnit                 string  `json:"fat_unit" bson:"fat_unit"`
		CarbohydratesServing    float64 `json:"carbohydrates_serving" bson:"carbohydrates_serving,truncate"`
		Proteins                int     `json:"proteins" bson:"proteins,truncate"`
		EnergyValue             int     `json:"energy_value" bson:"energy_value,truncate"`
		EnergyUnit              string  `json:"energy_unit" bson:"energy_unit"`
		ProteinsUnit            string  `json:"proteins_unit" bson:"proteins_unit"`
		EnergyKcalValue         int     `json:"energy-kcal_value" bson:"energy-kcal_value,truncate"`
		EnergyKcal100G          int     `json:"energy-kcal_100g" bson:"energy-kcal_100g,truncate"`
		EnergyKcalServing       float64 `json:"energy-kcal_serving" bson:"energy-kcal_serving,truncate"`
		FatServing              int     `json:"fat_serving" bson:"fat_serving,truncate"`
		CarbohydratesUnit       string  `json:"carbohydrates_unit" bson:"carbohydrates_unit"`
		FatValue                int     `json:"fat_value" bson:"fat_value,truncate"`
		Proteins100G            int     `json:"proteins_100g" bson:"proteins_100g,truncate"`
		ProteinsServing         int     `json:"proteins_serving" bson:"proteins_serving,truncate"`
		EnergyKcalValueComputed int     `json:"energy-kcal_value_computed" bson:"energy-kcal_value_computed,truncate"`
		EnergyServing           int     `json:"energy_serving" bson:"energy_serving,truncate"`
		Sugars                  float64 `json:"sugars" bson:"sugars,truncate"`
	} `json:"nutriments" bson:"nutriments"`
	ServingQuantityUnit string   `json:"serving_quantity_unit" bson:"serving_quantity_unit"`
	CountriesLc         string   `json:"countries_lc" bson:"countries_lc"`
	NovaGroupDebug      string   `json:"nova_group_debug" bson:"nova_group_debug"`
	NutriscoreTags      []string `json:"nutriscore_tags" bson:"nutriscore_tags"`
	Nutriscore          struct {
		Num2021 struct {
			Data struct {
				Sugars                                   float64     `json:"sugars" bson:"sugars,truncate"`
				Energy                                   int         `json:"energy" bson:"energy,truncate"`
				IsBeverage                               int         `json:"is_beverage" bson:"is_beverage,truncate"`
				IsWater                                  int         `json:"is_water" bson:"is_water,truncate"`
				Sodium                                   interface{} `json:"sodium" bson:"sodium,truncate"`
				IsCheese                                 int         `json:"is_cheese" bson:"is_cheese,truncate"`
				IsFat                                    int         `json:"is_fat" bson:"is_fat,truncate"`
				Proteins                                 int         `json:"proteins" bson:"proteins,truncate"`
				SaturatedFat                             int         `json:"saturated_fat" bson:"saturated_fat,truncate"`
				Fiber                                    int         `json:"fiber" bson:"fiber,truncate"`
				SaturatedFatRatio                        int         `json:"saturated_fat_ratio" bson:"saturated_fat_ratio,truncate"`
				FruitsVegetablesNutsColzaWalnutOliveOils int         `json:"fruits_vegetables_nuts_colza_walnut_olive_oils" bson:"fruits_vegetables_nuts_colza_walnut_olive_oils,truncate"`
			} `json:"data" bson:"data"`
			Grade                string `json:"grade" bson:"grade"`
			NutriscoreComputed   int    `json:"nutriscore_computed" bson:"nutriscore_computed,truncate"`
			NutriscoreApplicable int    `json:"nutriscore_applicable" bson:"nutriscore_applicable,truncate"`
			NutrientsAvailable   int    `json:"nutrients_available" bson:"nutrients_available,truncate"`
			CategoryAvailable    int    `json:"category_available" bson:"category_available,truncate"`
		} `json:"2021" bson:"2021"`
		Num2023 struct {
			NutriscoreComputed int    `json:"nutriscore_computed" bson:"nutriscore_computed,truncate"`
			Grade              string `json:"grade" bson:"grade"`
			Data               struct {
				Proteins                int         `json:"proteins" bson:"proteins,truncate"`
				IsFatOilNutsSeeds       int         `json:"is_fat_oil_nuts_seeds" bson:"is_fat_oil_nuts_seeds,truncate"`
				SaturatedFat            int         `json:"saturated_fat" bson:"saturated_fat,truncate"`
				Fiber                   interface{} `json:"fiber" bson:"fiber,truncate"`
				FruitsVegetablesLegumes interface{} `json:"fruits_vegetables_legumes" bson:"fruits_vegetables_legumes,truncate"`
				SaturatedFatRatio       int         `json:"saturated_fat_ratio" bson:"saturated_fat_ratio,truncate"`
				Salt                    interface{} `json:"salt" bson:"salt,truncate"`
				Sugars                  float64     `json:"sugars" bson:"sugars,truncate"`
				IsRedMeatProduct        int         `json:"is_red_meat_product" bson:"is_red_meat_product,truncate"`
				IsBeverage              int         `json:"is_beverage" bson:"is_beverage,truncate"`
				Energy                  int         `json:"energy" bson:"energy,truncate"`
				IsWater                 int         `json:"is_water" bson:"is_water,truncate"`
				IsCheese                int         `json:"is_cheese" bson:"is_cheese,truncate"`
			} `json:"data" bson:"data"`
			CategoryAvailable    int `json:"category_available" bson:"category_available,truncate"`
			NutrientsAvailable   int `json:"nutrients_available" bson:"nutrients_available,truncate"`
			NutriscoreApplicable int `json:"nutriscore_applicable" bson:"nutriscore_applicable,truncate"`
		} `json:"2023" bson:"2023"`
	} `json:"nutriscore" bson:"nutriscore"`
	Completeness                                float64       `json:"completeness" bson:"completeness,truncate"`
	ProductName                                 string        `json:"product_name" bson:"product_name"`
	PackagingMaterialsTags                      []interface{} `json:"packaging_materials_tags" bson:"packaging_materials_tags,truncate"`
	Keywords                                    []string      `json:"_keywords" bson:"_keywords"`
	LastUpdatedT                                int           `json:"last_updated_t" bson:"last_updated_t,truncate"`
	EntryDatesTags                              []string      `json:"entry_dates_tags" bson:"entry_dates_tags"`
	LanguagesHierarchy                          []string      `json:"languages_hierarchy" bson:"languages_hierarchy"`
	MainCountriesTags                           []interface{} `json:"main_countries_tags" bson:"main_countries_tags,truncate"`
	DataQualityInfoTags                         []string      `json:"data_quality_info_tags" bson:"data_quality_info_tags"`
	Nutriscore2021Tags                          []string      `json:"nutriscore_2021_tags" bson:"nutriscore_2021_tags"`
	MaxImgid                                    string        `json:"max_imgid" bson:"max_imgid"`
	StatesTags                                  []string      `json:"states_tags" bson:"states_tags"`
	NutritionScoreWarningNoFruitsVegetablesNuts int           `json:"nutrition_score_warning_no_fruits_vegetables_nuts" bson:"nutrition_score_warning_no_fruits_vegetables_nuts,truncate"`
	EnvironmentalScoreGrade                     string        `json:"environmental_score_grade" bson:"environmental_score_grade"`
	PopularityTags                              []string      `json:"popularity_tags" bson:"popularity_tags"`
	PackagingRecyclingTags                      []interface{} `json:"packaging_recycling_tags" bson:"packaging_recycling_tags,truncate"`
	CheckersTags                                []interface{} `json:"checkers_tags" bson:"checkers_tags,truncate"`
	NutrientLevels                              struct {
	} `json:"nutrient_levels" bson:"nutrient_levels"`
	PnnsGroups1Tags              []string      `json:"pnns_groups_1_tags" bson:"pnns_groups_1_tags"`
	LastEditDatesTags            []string      `json:"last_edit_dates_tags" bson:"last_edit_dates_tags"`
	NutritionScoreWarningNoFiber int           `json:"nutrition_score_warning_no_fiber" bson:"nutrition_score_warning_no_fiber,truncate"`
	InformersTags                []string      `json:"informers_tags" bson:"informers_tags"`
	PnnsGroups2                  string        `json:"pnns_groups_2" bson:"pnns_groups_2"`
	BrandsTags                   []string      `json:"brands_tags" bson:"brands_tags"`
	Allergens                    string        `json:"allergens" bson:"allergens"`
	PnnsGroups1                  string        `json:"pnns_groups_1" bson:"pnns_groups_1"`
	EcoscoreTags                 []string      `json:"ecoscore_tags" bson:"ecoscore_tags"`
	PnnsGroups2Tags              []string      `json:"pnns_groups_2_tags" bson:"pnns_groups_2_tags"`
	PopularityKey                int           `json:"popularity_key" bson:"popularity_key,truncate"`
	AllergensTags                []interface{} `json:"allergens_tags" bson:"allergens_tags,truncate"`
}
