package saml

const bindTemplate = `<html>
<body style="display: none" onload="document.saml.submit()">
	<form method="post" name="saml" action="{{.SsoUrl}}">
		<input type="hidden" name="SAMLRequest"
			value="{{.SAMLRequest}}"/>
		<input type="hidden" name="RelayState"
			value="{{.RelayState}}"/>
		<input type="submit" value="Submit"/>
	</form>
</body>
</html>`
