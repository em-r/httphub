package helpers

var JSONDoc = `{
	"whoami": "mehdi",
	"where": "33°25'17.0''N 7°50'32.6''W",
	"links": {
		"github": "https://github.com/ElMehdi19",
		"medium": "https://iammehdi.medium.com",
		"linkedin": "https://www.linkedin.com/in/el-mehdi-rami",
		"personal": "https://mehdi.codes"
	},
	"email": "mehdi@httphub.io"
}`

var XMLDoc = `<?xml version='1.0' encoding='us-ascii'?>
	<me xmlns:xlink="http://www.w3.org/1999/xlink">	
		<whoami>mehdi</whoami>
		<from>33°25'17.0"N 7°50'32.6"W</from>
		<links>
			<link xlink:type="simple" xlink:href="https://github.com/ElMehdi19">github</link>	
			<link xlink:type="simple" xlink:href="https://iammehdi.medium.com">medium</link>	
			<link xlink:type="simple" xlink:href="https://www.linkedin.com/in/el-mehdi-rami">linkedin</link>	
			<link xlink:type="simple" xlink:href="https://mehdi.codes">personal</link>	
		</links>
		<email xlink:type="simple" xlink:href="mailto:mehdi@httphub.io">contact</email>
	</me>	
`
