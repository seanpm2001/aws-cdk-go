package awscdkappconfigalpha


// Example:
//   var application application
//   var alarm alarm
//
//
//   appconfig.NewEnvironment(this, jsii.String("MyEnvironment"), &EnvironmentProps{
//   	Application: Application,
//   	Monitors: []monitor{
//   		&monitor{
//   			Alarm: *Alarm,
//   		},
//   	},
//   })
//
// Experimental.
type EnvironmentProps struct {
	// The description of the environment.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The monitors for the environment.
	// Default: - No monitors.
	//
	// Experimental.
	Monitors *[]*Monitor `field:"optional" json:"monitors" yaml:"monitors"`
	// The name of the environment.
	// Default: - A name is generated.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The application to be associated with the environment.
	// Experimental.
	Application IApplication `field:"required" json:"application" yaml:"application"`
}

