package Templates

var FileUpload = `
{{define "FileUpload"}}
<!DOCTYPE html>
<!-- This site was created in Webflow. http://www.webflow.com-->
<!-- Last Published: Fri Jun 26 2015 05:50:41 GMT+0000 (UTC) -->
<html data-wf-site="547b44aa3e9ac2216ec5d048" data-wf-page="558ce167e20f1f4d31c64577">
<head>
  <meta charset="utf-8">
  <title>Upload a file</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="generator" content="Webflow">
  <link rel="stylesheet" type="text/css" href="/admin/css/normalize.css">
  <link rel="stylesheet" type="text/css" href="/admin/css/webflow.css">
  <link rel="stylesheet" type="text/css" href="/admin/css/admin.css">
  <script type="text/javascript" src="/admin/js/modernizr.js"></script>
</head>
<body>
  <div class="w-section headercontainer">
    <h1>Upload a file</h1>
  </div>
  <div class="w-container adminsection">
    <p class="introtext">This function enables you to upload a file to the distributed file system of the MongoDB database. If the desired file is already present, a new revision of this file is created. Therefore, an already existing file gets never overwritten! Please consider, that the configured maximum size of the header of the admin web server (see configuration) forces the maximum file size. Thus, please check the current maximum. The default maximum is&nbsp;approx. 10 MB!</p>
    <div class="w-form">
      <form id="upload" name="upload" data-name="upload" method="post" action="/upload" enctype="multipart/form-data">
        <label for="file">Please select a file to upload:</label>
        <input class="w-input" id="file" type="file" name="file" data-name="file">
        <input class="w-button button optionbuttons" type="submit" value="Upload this file" data-wait="Please wait...">
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
