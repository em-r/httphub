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

var HTMLDoc = `
<html>
	<head>
		<meta charset="utf-8" />
		<title>Sample HTML doc</title>
	</head>
	<body>
		<h1>About</h1>
		<ul>
			<li><h2>name: mehdi</h2></li>
			<li><h2>from: 33°25'17.0"N 7°50'32.6"W</h2></li>
		</ul>
		<h1>Links</h1>
		<ul>
			<li><h2>github <a href="https://github.com/ElMehdi19">https://github.com/ElMehdi19</a></h2></li>
			<li><h2>medium <a href="https://iammehdi.medium.com">https://iammehdi.medium.com</a></h2></li>
			<li><h2>linkedin <a href="https://www.linkedin.com/in/el-mehdi-rami">https://www.linkedin.com/in/el-mehdi-rami</a></h2></li>
			<li><h2>personal <a href="https://mehdi.codes">https://mehdi.codes</a></h2></li>
		</ul>
		<h1>contact <a href="mailto:mehdi@httphub.io">mehdi@httphub.io</a></h1>
	</body>
</html>
`

var TXTDoc = `
H      H TTTTTTTTTTTTT TTTTTTTTTTTTT PPPPPP	 	 H      H	U		   U  BBBBBB
H	   H	   T	   	     T	 	 P	   P	 H	    H	U		   U  B     B
H	   H	   T	   	     T	 	 P		P	 H	    H 	U		   U  B      B
H	   H	   T			 T		 P	   P	 H	    H	U		   U  B     B
HHHHHHHH	   T	   		 T		 PPPPPP		 HHHHHHHH	U		   U  BBBBBB
H      H	   T		     T		 P			 H	    H	U		   U  B     B
H	   H	   T			 T		 P			 H	    H	 U		  U   B      B
H	   H	   T			 T		 P			 H	    H	  U		 U    B     B
H	   H	   T			 T       P			 H	    H	   UUUUUU     BBBBBB
`
