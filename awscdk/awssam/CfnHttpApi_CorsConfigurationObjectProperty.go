package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsConfigurationObjectProperty := &CorsConfigurationObjectProperty{
//   	AllowCredentials: jsii.Boolean(false),
//   	AllowHeaders: []*string{
//   		jsii.String("allowHeaders"),
//   	},
//   	AllowMethods: []*string{
//   		jsii.String("allowMethods"),
//   	},
//   	AllowOrigins: []*string{
//   		jsii.String("allowOrigins"),
//   	},
//   	ExposeHeaders: []*string{
//   		jsii.String("exposeHeaders"),
//   	},
//   	MaxAge: jsii.Number(123),
//   }
//
type CfnHttpApi_CorsConfigurationObjectProperty struct {
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowCredentials`.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowHeaders`.
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowMethods`.
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowOrigins`.
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.ExposeHeaders`.
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.MaxAge`.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}
