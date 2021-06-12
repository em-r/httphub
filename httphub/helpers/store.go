package helpers

var JSONDoc = `{
	"whoami": "mehdi",
	"where": "33.6031505,-7.49227239",
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
		<whoami>xx</whoami>
		<where>xx</where>
		<links>
			<link xlink:type="simple" xlink:href="https://github.com/ElMehdi19">github</link>	
			<link xlink:type="simple" xlink:href="https://iammehdi.medium.com">medium</link>	
			<link xlink:type="simple" xlink:href="https://www.linkedin.com/in/el-mehdi-rami">linkedin</link>	
			<link xlink:type="simple" xlink:href="https://mehdi.codes">personal</link>	
		</links>
		<email xlink:type="simple" xlink:href="mailto:mehdi@httphub.io">contact</email>
	</me>	
`
