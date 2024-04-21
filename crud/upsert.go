package crud

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Tea struct {
	Type     string
	Category string
	Toppings []string
	Price    float32
}
type Shops struct {
	ID            int    `bson:"id"`
	Count         int    `bson:"count"`
	ShopifyDomain string `bson:"shopify_domain"`
	UserID        int    `bson:"user_id"`
	ChargeID      struct {
		NumberLong string `bson:"$numberLong"`
	} `json:"charge_id"`
	Paid         bool    `bson:"paid"`
	PaymentDate  string  `bson:"payment_date"`
	PaymentPrice float64 `bson:"payment_price"`
	PaymentType  string  `bson:"payment_type"`
	CreatedAt    int64   `bson:"created_at"`
	UpdatedAt    string  `bson:"updated_at"`
	Fname        string  `bson:"fname"`
	Lname        string  `bson:"lname"`
	Country      string  `bson:"country"`
	Language     string  `bson:"language"`
	Email        string  `bson:"email"`
	Profile      struct {
		ID                                   float64     `bson:"id"`
		Name                                 string      `bson:"name"`
		Email                                string      `bson:"email"`
		Domain                               string      `bson:"domain"`
		Province                             string      `bson:"province"`
		Country                              string      `bson:"country"`
		Address1                             string      `bson:"address1"`
		Zip                                  string      `bson:"zip"`
		City                                 string      `bson:"city"`
		Source                               interface{} `bson:"source"`
		Phone                                string      `bson:"phone"`
		Latitude                             interface{} `bson:"latitude"`
		Longitude                            interface{} `bson:"longitude"`
		PrimaryLocale                        string      `bson:"primary_locale"`
		Address2                             string      `bson:"address2"`
		CreatedAt                            string      `bson:"created_at"`
		UpdatedAt                            string      `bson:"updated_at"`
		CountryCode                          string      `bson:"country_code"`
		CountryName                          string      `bson:"country_name"`
		Currency                             string      `bson:"currency"`
		CustomerEmail                        string      `bson:"customer_email"`
		Timezone                             string      `bson:"timezone"`
		IANATimezone                         string      `bson:"iana_timezone"`
		ShopOwner                            string      `bson:"shop_owner"`
		MoneyFormat                          string      `bson:"money_format"`
		MoneyWithCurrencyFormat              string      `bson:"money_with_currency_format"`
		WeightUnit                           string      `bson:"weight_unit"`
		ProvinceCode                         string      `bson:"province_code"`
		TaxesIncluded                        bool        `bson:"taxes_included"`
		AutoConfigureTaxInclusivity          interface{} `bson:"auto_configure_tax_inclusivity"`
		TaxShipping                          interface{} `bson:"tax_shipping"`
		CountyTaxes                          bool        `bson:"county_taxes"`
		PlanDisplayName                      string      `bson:"plan_display_name"`
		PlanName                             string      `bson:"plan_name"`
		HasDiscounts                         bool        `bson:"has_discounts"`
		HasGiftCards                         bool        `bson:"has_gift_cards"`
		MyshopifyDomain                      string      `bson:"myshopify_domain"`
		GoogleAppsDomain                     interface{} `bson:"google_apps_domain"`
		GoogleAppsLoginEnabled               interface{} `bson:"google_apps_login_enabled"`
		MoneyInEmailsFormat                  string      `bson:"money_in_emails_format"`
		MoneyWithCurrencyInEmailsFormat      string      `bson:"money_with_currency_in_emails_format"`
		EligibleForPayments                  bool        `bson:"eligible_for_payments"`
		RequiresExtraPaymentsAgreement       bool        `bson:"requires_extra_payments_agreement"`
		PasswordEnabled                      bool        `bson:"password_enabled"`
		HasStorefront                        bool        `bson:"has_storefront"`
		Finances                             bool        `bson:"finances"`
		PrimaryLocationID                    int64       `bson:"primary_location_id"`
		CookieConsentLevel                   string      `bson:"cookie_consent_level"`
		VisitorTrackingConsentPreference     string      `bson:"visitor_tracking_consent_preference"`
		CheckoutAPISupported                 bool        `bson:"checkout_api_supported"`
		MultiLocationEnabled                 bool        `bson:"multi_location_enabled"`
		SetupRequired                        bool        `bson:"setup_required"`
		PreLaunchEnabled                     bool        `bson:"pre_launch_enabled"`
		EnabledPresentmentCurrencies         []string    `bson:"enabled_presentment_currencies"`
		TransactionalSMSDisabled             bool        `bson:"transactional_sms_disabled"`
		MarketingSMSConsentEnabledAtCheckout bool        `bson:"marketing_sms_consent_enabled_at_checkout"`
	} `json:"profile"`
	WebhookVersion float64 `bson:"webhook_version"`
	FetchAuthor    bool    `bson:"fetch_author"`
	EndTrialDate   string  `bson:"end_trial_date"`
	Extras         struct {
		IsUseRender bool `bson:"isUseRender"`
	} `json:"extras"`
	StoreInfo struct {
		LazyLoad      bool `bson:"lazyLoad"`
		Shortcut      bool `bson:"shortcut"`
		ThemeScript   bool `bson:"themeScript"`
		PreloadScript bool `bson:"preloadScript"`
		ProductSchema bool `bson:"productSchema"`
	} `json:"store_info"`
	AccountInfo struct {
		EmailPreferences struct {
			Theme       bool `bson:"theme"`
			Updates     bool `bson:"updates"`
			Partners    bool `bson:"partners"`
			Research    bool `bson:"research"`
			Schedule    bool `bson:"schedule"`
			Analytics   bool `bson:"analytics"`
			Onboarding  bool `bson:"onboarding"`
			Newsletters bool `bson:"newsletters"`
		} `json:"emailPreferences"`
	} `json:"account_info"`
	Activity     string `bson:"activity"`
	IsGempagesV7 int    `bson:"is_gempages_v7"`
	ShopifyScope string `bson:"shopify_scope"`
}
type Output struct {
	ID            primitive.ObjectID `bson:"_id"`
	ShopifyDomain string             `bson:"shopify_domain"`
	UserID        int                `bson:"user_id"`
	PaymentType   string             `bson:"payment_type"`
	CreatedAt     string             `bson:"created_at"`
	Country       string             `bson:"profile.country"`
}

func MgShopAggreate() {

	uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("hello mongodb")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	colection := client.Database("db").Collection("shops")

	// match1 := bson.M{"country": "United Kingdom"}
	// match2 := bson.M{"shopify_domain": "art-tee-show.myshopify.com"}
	// // // Tạo một mảng chứa các phần $match
	// matchStages := []bson.M{{"$match": match1}, {"$match": match2}}
	// fmt.Println("len ", len(matchStages))
	matchStage := []bson.D{
		{{"$match", bson.D{{"country", "United Kingdom"}}}},
		{{"$match", bson.D{{"shopify_domain", "art-tee-show.myshopify.com"}}}},
	}

	//matchStage1 := bson.D{{"$match", bson.D{{"shopify_domain", "art-tee-show.myshopify.com"}}}}
	//find := bson.D{{{"country", "United Kingdom"}, {"shopify_domain", "art-tee-show.myshopify.com"}}
	cur, err := colection.Aggregate(context.TODO(),
		matchStage,
	)

	if err != nil {
		fmt.Println("err ", err.Error())
		return
	}
	defer cur.Close(context.TODO())

	var results []Shops
	if err := cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, item := range results {
		fmt.Println("shop ", item.Country, item.ShopifyDomain)
	}
}

type ShopifyDomain struct {
	Domain string `json:"domain"`
}

// CountryData struct
type CountryData struct {
	ID          string          `json:"_id"`
	ListDomains []ShopifyDomain `json:"list_domain"`
}

func PushExam() (interface{}, error) {
	uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("hello mongodb")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// filter := bson.M{}
	//build match
	// bson.D{
	// 	{"$group", bson.D{
	// 		{"_id", "$country"},
	// 		{"list_domain", bson.D{{"$push", "$shopify_domain"}}},
	// 	}},
	// }
	var results []bson.M

	// Định nghĩa pipeline aggregation
	countryPipeline := bson.D{
		{"$group", bson.D{
			{"_id", "$country"},
			{"list_domain", bson.D{
				{"$push", "$shopify_domain"},
			}},
		}},
	}

	// Chọn bảng và thực thi pipeline aggregation
	collection := client.Database("db").Collection("shops")
	cur, err := collection.Aggregate(context.Background(), mongo.Pipeline{countryPipeline})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var item bson.M
		if err := cur.Decode(&item); err != nil {
			continue
		}
		results = append(results, item)
	}
	// Decode kết quả vào một slice của CountryData
	// if err := cur.All(context.Background(), &results); err != nil {
	// 	return nil, err
	// }
	return results, nil
}
func UpdateTime1() {
	uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("hello mongodb")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	filter := bson.M{}
	// updated := bson.M{
	// 	"$set": bson.M{
	// 		"created_at":
	// 	}
	// }
	cur, _ := client.Database("db").Collection("shops").Find(context.TODO(), filter)
	for cur.Next(context.TODO()) {
		var doc bson.M
		if err := cur.Decode(&doc); err != nil {
			continue
		}
		if createdAt, ok := doc["created_at"].(int64); ok {
			// value, err := time.Parse("2006-01-02 15:04:05.000", createdAtStr)
			// if err != nil {
			// 	log.Printf("Error parsing created_at: %v\n", err)
			// 	continue
			// }
			createdAtInt64 := createdAt / 1000

			doc["created_at"] = createdAtInt64
		}

		newFilter := bson.M{"_id": doc["_id"]}
		updated := bson.M{"$set": doc}
		_, err = client.Database("db").Collection("shops").UpdateOne(context.TODO(),
			newFilter, updated)
		if err != nil {
			fmt.Println("err ", err)
		}
	}
}

func CreateShops() {
	uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("hello mongodb")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	colection := client.Database("db").Collection("shops")
	cur, err := colection.InsertOne(context.TODO(), &Shops{
		ShopifyDomain: "new_domain@shopify.com",
		UserID:        1991,
	}, nil)
	fmt.Println("err 1991 ", err)

	if err != nil {
		return
	}
	colection.UpdateByID(context.TODO(), cur.InsertedID, &Shops{
		UserID: 1992,
	})
	filter := bson.D{{"shopify_domain", "art-tee-show.myshopify.com"}}
	updated := bson.D{{"$set", &Shops{
		UserID: 2002,
	}}}
	_, err = colection.UpdateOne(context.TODO(),
		filter, updated)
	fmt.Println("err ", err)
}

type UpdateInput struct {
	NameShop string `json:"name_shop"`
	LastTime int64  `json:"last_time"`
	Country  string `json:"country"`
}

func UpdateShops(ctx context.Context, mapValue map[string]string, input *UpdateInput) error {
	uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("hello mongodb")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	colection := client.Database("db").Collection("shops")
	//find the shop first
	filter := bson.M{}
	if mapValue["id"] != "" {
		filter["id"] = mapValue["id"]
	}

	if mapValue["shopify_domain"] != "" {
		filter["shopify_domain"] = mapValue["shopify_domain"]
	}
	var shop bson.M

	cur := colection.FindOne(ctx, filter, nil)

	if err := cur.Decode(&shop); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("id ", shop["id"])
	if shop["id"] != "" {
		shop["country"] = input.Country
		shop["shopify_scope"] = input.NameShop
	}
	_, err = colection.UpdateOne(ctx, bson.M{"id": shop["id"]}, bson.M{"$set": shop}, nil)
	if err != nil {
		fmt.Println("update err ", err.Error())
		return nil
	}
	return nil
}

func UpdateShopD(ctx context.Context, mapValue map[string]string, input *UpdateInput) error {
	uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("hello mongodb")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	colection := client.Database("db").Collection("shops")
	//find the shop first
	filter := bson.D{}
	if mapValue["id"] != "" {
		filter = append(filter, primitive.E{"id", mapValue["id"]})
	}

	if mapValue["shopify_domain"] != "" {
		filter = append(filter, primitive.E{"shopify_domain", mapValue["shopify_domain"]})
	}
	var shop bson.M

	cur := colection.FindOne(ctx, filter, nil)

	if err := cur.Decode(&shop); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("id ", shop["id"])
	if shop["id"] != "" {
		shop["country"] = input.Country
		shop["shopify_scope"] = input.NameShop
	}
	_, err = colection.UpdateOne(ctx, bson.M{"id": shop["id"]}, bson.M{"$set": shop}, nil)
	if err != nil {
		fmt.Println("update err ", err.Error())
		return nil
	}
	return nil
}
