package Templates

var Overview = `
{{define "Overview"}}
<!DOCTYPE html>
<!-- This site was created in Webflow. http://www.webflow.com-->
<!-- Last Published: Mon Jun 22 2015 10:36:36 GMT+0000 (UTC) -->
<html data-wf-site="547b44aa3e9ac2216ec5d048" data-wf-page="5587d2abbf862f2179c4373b">
<head>
  <meta charset="utf-8">
  <title>Overview</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="generator" content="Webflow">
  <link rel="stylesheet" type="text/css" href="/admin/css/normalize.css">
  <link rel="stylesheet" type="text/css" href="/admin/css/webflow.css">
  <link rel="stylesheet" type="text/css" href="/admin/css/admin.css">
  <script type="text/javascript" src="/admin/js/modernizr.js"></script>
</head>
<body>
  <div class="w-section headercontainer">
    <h1>Administration</h1>
  </div>
  <div class="w-container adminsection">
    <div class="w-row">
      <div class="w-col w-col-4"><a class="button adminbutton" href="/log">Logging Viewer</a>
      </div>
      <div class="w-col w-col-4"><a class="button adminbutton" href="/upload">Upload a file</a>
      </div>
      <div class="w-col w-col-4"><a class="button adminbutton" href="/configuration">Configuration</a>
      </div>
    </div>
  <div class="admintextblock">The current Ocean's version is: {{.Version}}</div>
  </div>
  <script type="text/javascript" src="/admin/js/jquery.min.js"></script>
  <script type="text/javascript" src="/admin/js/webflow.js"></script>
</body>
</html>
{{end}}
`
