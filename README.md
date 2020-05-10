# Cruciforme

A form endpoint for static websites. Point any form to a cruciforme instance, and you will receive the answers by email.

## Usage

Simply point you form to https://crucifor.me and add a hidden field named `cruciform-mail`:

```html
<form action="https://crucifor.me" method="post" enctype="multipart/form-data">
	<!-- your inputs -->

	<input name="cruciforme-mail" value="your@awesome.mail" type="hidden">
</form>
```

You can redirect the user to a custom page using `cruciforme-success` and `cruciforme-error`:

```html
<input name="cruciforme-success" value="your.awesome.site/success" type="hidden">
<input name="cruciforme-error" value="your.awesome.site/error" type="hidden">
```

Finally, you can customize the subject of the email using `cruciforme-subject`:

```html
<input name="cruciforme-subject" value="Great, a form sent!" type="hidden">
```
