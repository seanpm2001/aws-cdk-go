// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Props for creating a Python Spark (ETL or Streaming) job executable.
//
// Example:
//   glue.NewJob(this, jsii.String("PythonSparkStreamingJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonStreaming(&pythonSparkJobExecutableProps{
//   		glueVersion: glue.glueVersion_V2_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
//   	}),
//   	description: jsii.String("an example Python Streaming job"),
//   })
//
// Experimental.
type PythonSparkJobExecutableProps struct {
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `field:"required" json:"glueVersion" yaml:"glueVersion"`
	// The Python version to use.
	// Experimental.
	PythonVersion PythonVersion `field:"required" json:"pythonVersion" yaml:"pythonVersion"`
	// The script that executes a job.
	// Experimental.
	Script Code `field:"required" json:"script" yaml:"script"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script. Only individual files are supported, directories are not supported.
	// See: `--extra-jars` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJars *[]Code `field:"optional" json:"extraJars" yaml:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See: `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJarsFirst *bool `field:"optional" json:"extraJarsFirst" yaml:"extraJarsFirst"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-py-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraPythonFiles *[]Code `field:"optional" json:"extraPythonFiles" yaml:"extraPythonFiles"`
}
