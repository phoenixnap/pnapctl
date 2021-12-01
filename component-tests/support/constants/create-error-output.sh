invalidCommandPost="post"
invalidCommandCreate="createe" 
jsonFile="test.json"
yamlFile="test.yaml"
configFile="config.yaml"

expectedOutputCredentialError="Error: Post \"https://api.phoenixnap.com/bmc/v0/servers\": oauth2: cannot fetch token: 400 Bad Request
Response: {\"error\":\"invalid_client\",\"error_description\":\"Invalid client credentials\"}"

expectedOutputWrongFile="Error: Command 'create server' has been performed, but something went wrong. Error code: 0303"

expectedOutputWrongCommandError="Error: unknown command \"$invalidCommandCreate\" for \"pnapctl\"

Did you mean this?
	create

Run 'pnapctl --help' for usage."

expectedOutputFileNotSetError="Error: required flag(s) \"filename\" not set"

expectedOutputInvalidPostError="Error: unknown command \"$invalidCommandPost\" for \"pnapctl\"
Run 'pnapctl --help' for usage."

expectedOutputNonexistentJsonFile="Error: The file '$jsonFile' does not exist."

expectedOutputNonexistentYamlFile="Error: The file '$yamlFile' does not exist."

expectedOutputNotSpecifiedFilename="Error: required flag(s) \"filename\" not set"

expectedOutputInvalidFilenameCommand="Error: accepts 0 arg(s), received 2"
