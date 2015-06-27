package Templates

var Configuration = `
{{define "Configuration"}}
<!DOCTYPE html>
<!-- This site was created in Webflow. http://www.webflow.com-->
<!-- Last Published: Fri Jun 26 2015 05:50:41 GMT+0000 (UTC) -->
<html data-wf-site="547b44aa3e9ac2216ec5d048" data-wf-page="558ce60e0939295c77a362db">
<head>
  <meta charset="utf-8">
  <title>Configuration</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="generator" content="Webflow">
  <link rel="stylesheet" type="text/css" href="/admin/css/normalize.css">
  <link rel="stylesheet" type="text/css" href="/admin/css/webflow.css">
  <link rel="stylesheet" type="text/css" href="/admin/css/admin.css">
  <script type="text/javascript" src="/admin/js/modernizr.js"></script>
</head>
<body>
  <div class="w-section headercontainer">
    <h1>Cluster Configuration</h1>
  </div>
  <div class="w-container adminsection">
    <p class="introtext"><strong>Attention:</strong> This configuration applies to the whole cluster of your Ocean servers. Therefore, <strong>this is not</strong> an individual configuration for any single server! Thus, please consider beforehand if the desired change matches all of your servers.
    </p>
    <div class="w-form">
      <form id="configuration" name="configuration" data-name="configuration" method="post" action="/configuration">
        {{range .Configuration}} 
          <label for="{{.Name}}">Configuration parameter:&nbsp;{{.Name}}</label>
          <input class="w-input" id="{{.Name}}" type="text" name="{{.Name}}" data-name="{{.Name}}" required="required" value="{{.Value}}" placeholder="{{.Value}}">
        {{end}}
        <input class="w-button button optionbuttons" type="submit" value="Apply all changes" data-wait="Please wait...">
      </form>
      <div class="w-form-done">
        <p>Thank you! Your submission has been received!</p>
      </div>
      <div class="w-form-fail">
        <p>Oops! Something went wrong while submitting the form</p>
      </div>
    </div>
  </div>
  <script type="text/javascript" src="/admin/js/jquery.min.js"></script>
  <script type="text/javascript" src="/admin/js/webflow.js"></script>
</body>
</html>
{{end}}
`
