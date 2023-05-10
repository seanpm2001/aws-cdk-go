// An experiment to bundle the entire CDK into a single module
package awscdk


// The AWS OrganizationalUnitIds or Accounts for which to create stack instances in the specified Regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentTargetsProperty := &DeploymentTargetsProperty{
//   	AccountFilterType: jsii.String("accountFilterType"),
//   	Accounts: []*string{
//   		jsii.String("accounts"),
//   	},
//   	OrganizationalUnitIds: []*string{
//   		jsii.String("organizationalUnitIds"),
//   	},
//   }
//
type CfnStackSet_DeploymentTargetsProperty struct {
	// Limit deployment targets to individual accounts or include additional accounts with provided OUs.
	//
	// The following is a list of possible values for the `AccountFilterType` operation.
	//
	// - `INTERSECTION` : StackSets deploys to the accounts specified in `Accounts` parameter.
	// - `DIFFERENCE` : StackSets excludes the accounts specified in `Accounts` parameter. This enables user to avoid certain accounts within an OU such as suspended accounts.
	// - `UNION` : StackSets includes additional accounts deployment targets.
	//
	// This is the default value if `AccountFilterType` is not provided. This enables user to update an entire OU and individual accounts from a different OU in one request, which used to be two separate requests.
	// - `NONE` : Deploys to all the accounts in specified organizational units (OU).
	AccountFilterType *string `field:"optional" json:"accountFilterType" yaml:"accountFilterType"`
	// The names of one or more AWS accounts for which you want to deploy stack set updates.
	//
	// *Pattern* : `^[0-9]{12}$`.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// The organization root ID or organizational unit (OU) IDs to which StackSets deploys.
	//
	// *Pattern* : `^(ou-[a-z0-9]{4,32}-[a-z0-9]{8,32}|r-[a-z0-9]{4,32})$`.
	OrganizationalUnitIds *[]*string `field:"optional" json:"organizationalUnitIds" yaml:"organizationalUnitIds"`
}

