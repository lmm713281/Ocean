package Templates

var Viewer string = `
<!DOCTYPE html>
<!-- This site was created in Webflow. http://www.webflow.com-->
<!-- Last Published: Mon Feb 02 2015 20:11:43 GMT+0000 (UTC) -->
<html data-wf-site="547b44aa3e9ac2216ec5d048" data-wf-page="547b44aa3e9ac2216ec5d049">
<head>
  <meta charset="utf-8">
  <title>{{.Title}}</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="generator" content="Webflow">
  <link rel="stylesheet" type="text/css" href="/log/css/normalize.css">
  <link rel="stylesheet" type="text/css" href="/log/css/webflow.css">
  <link rel="stylesheet" type="text/css" href="/log/css/log.css">
  <script type="text/javascript" src="/log/js/modernizr.js"></script>
</head>
<body>
  <div class="headercontainer">
    <h1>Logbook</h1>
  </div>
  <div class="controlsection">
    <div class="w-row">
      <div class="w-col w-col-6">
        <h2 class="headercontrol">Options</h2>
        <div class="options"><a class="button optionbuttons" href="#">Enable live view</a>
        </div>
      </div>
      <div class="w-col w-col-6">
        <h2 class="headercontrol">Filtering</h2>
        <div class="filters">
          <div class="w-row">
            <div class="w-col w-col-6">
              <h3 class="subheader">By Time Range</h3><a class="button optionbuttons" href="#">Last 5 minutes</a><a class="button optionbuttons" href="#">Last 30 minutes</a><a class="button optionbuttons" href="#">Last 60 minutes</a><a class="button optionbuttons" href="#">Last 24 hours</a><a class="button optionbuttons" href="#">Last 7 days</a><a class="button optionbuttons" href="#">Last month</a><a class="button optionbuttons" href="#">All times</a>
            </div>
            <div class="w-col w-col-6">
              <h3 class="subheader">By Filter</h3>
              <div class="w-form">
                <form id="wf-form-Filters" name="wf-form-Filters" data-name="Filters">
                  <label for="Level">Level:</label>
                  <select class="w-select" id="Level" name="Level" data-name="Level">
                    <option value="*">Any</option>
                    <option value="INFO">Information</option>
                    <option value="WARN">Warning</option>
                    <option value="SECURITY">Security</option>
                    <option value="ERROR">Error</option>
                    <option value="DEBUG">Debug</option>
                    <option value="TALKATIVE">Talkative</option>
                  </select>
                  <label for="Category">Category:</label>
                  <select class="w-select" id="Category" name="Category" data-name="Category">
                    <option value="*">Any</option>
                    <option value="BUSINESS">Business</option>
                    <option value="APP">Application</option>
                    <option value="USER">User</option>
                    <option value="SYSTEM">System</option>
                  </select>
                  <label for="Impact">Impact:</label>
                  <select class="w-select" id="Impact" name="Impact" data-name="Impact">
                    <option value="*">Any</option>
                    <option value="LOW">Low</option>
                    <option value="MIDDLE">Middle</option>
                    <option value="HIGH">High</option>
                    <option value="CRITICAL">Critical</option>
                    <option value="UNKNOWN">Unknown</option>
                  </select>
                  <label for="Severity">Severity:</label>
                  <select class="w-select" id="Severity" name="Severity" data-name="Severity">
                    <option value="*">Any</option>
                    <option value="LOW">Low</option>
                    <option value="MIDDLE">Middle</option>
                    <option value="HIGH">High</option>
                    <option value="CRITICAL">Critical</option>
                    <option value="UNKNOWN">Unknown</option>
                  </select>
                  <label for="MSGName">Message Names:</label>
                  <select class="w-select" id="MSGName" name="MSGName" data-name="MSGName">
                    <option value="*">Any</option>
                    {{range .MessageNames}}
                      <option value="{{.}}">{{.}}</option>
                    {{end}}
                  </select>
                  <label for="Senders">Senders:</label>
                  <select class="w-select" id="Senders" name="Senders" data-name="Senders">
                    <option value="*">Any</option>
                    {{range .Sender}}
                      <option value="{{.}}">{{.}}</option>
                    {{end}}
                  </select>
                  <input class="w-button button optionbuttons applyfilters" type="submit" value="Apply filters" data-wait="Please wait...">
                </form>              
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div>
    <h2 class="headeroutput">Output</h2>
    <ul class="loglist">
      {{range .Events}}
      <li class="logline {{.AB}} {{.LogLevel}}">
        <div>{{.LogLine}}</div>
      </li>
      {{end}}
    </ul>
  </div>
  <div class="w-container pagecontainer"><a class="button changepagebutton" href="#">First</a><a class="button changepagebutton" href="#">-1</a>
    <div class="w-hidden-main w-hidden-medium newlineblock"></div>
    <div class="w-form formpageswrapper">
      <form class="formpages" id="wf-form-Pages" name="wf-form-Pages" data-name="Pages" method="post">
        <label class="formfield" for="CurrentPage">Page</label>
        <input class="w-input currentpage" id="CurrentPage" type="text" placeholder="1" name="CurrentPage" required="required" data-name="CurrentPage">
        <div class="textcountpages">of 1000 pages</div>
        <div class="w-hidden-main w-hidden-medium w-hidden-small newlineblock"></div>
        <input class="w-button button changepagebutton pagechangesubmit" type="submit" value="Change page" data-wait="Please wait...">
      </form>
    </div>
    <div class="w-hidden-main w-hidden-medium newlineblock"></div><a class="button changepagebutton" href="#">+1</a><a class="button changepagebutton" href="#">Last</a>
  </div>
  <script type="text/javascript" src="/log/js/jquery.min.js"></script>
  <script type="text/javascript" src="/log/js/webflow.js"></script>
</body>
</html>`
