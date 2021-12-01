invalidCommandPost="post"
invalidCommandCreate="createe" 

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