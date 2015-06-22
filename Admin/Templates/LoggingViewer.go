package Templates

// The template for the web log viewer:
var LoggingViewer string = `
{{define "WebLog"}}
<!DOCTYPE html>
<!-- This site was created in Webflow. http://www.webflow.com-->
<!-- Last Published: Mon Feb 02 2015 20:11:43 GMT+0000 (UTC) -->
<html data-wf-site="547b44aa3e9ac2216ec5d048" data-wf-page="547b44aa3e9ac2216ec5d049">
<head>
  <meta charset="utf-8">
  <title>{{.Title}}</title>
  {{if .SetLiveView}}
  <meta http-equiv="refresh" content="30; URL=/log?Level={{.CurrentLevel}}&TimeRange={{.CurrentTimeRange}}&Category={{.CurrentCategory}}&Impact={{.CurrentImpact}}&Severity={{.CurrentSeverity}}&MSGName={{.CurrentMessageName}}&Sender={{.CurrentSender}}&CurrentPage={{.CurrentPage}}&LiveView={{.SetLiveView}}">
  {{end}}
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="generator" content="Webflow">
  <link rel="stylesheet" type="text/css" href="/admin/css/normalize.css">
  <link rel="stylesheet" type="text/css" href="/admin/css/webflow.css">
  <link rel="stylesheet" type="text/css" href="/admin/css/admin.css">
  <script type="text/javascript" src="/admin/js/modernizr.js"></script>
</head>
<body>
  <div class="headercontainer">
    <h1>Logging Viewer</h1>
  </div>
  <div class="w-section controlsection">
    <div class="w-row">
      <div class="w-col w-col-6">
        <h2 class="headercontrol">Options</h2>
        <div class="options">
          {{if .SetLiveView}}
            <a class="button optionbuttons currentoption" href="/log?Level={{.CurrentLevel}}&TimeRange={{.CurrentTimeRange}}&Category={{.CurrentCategory}}&Impact={{.CurrentImpact}}&Severity={{.CurrentSeverity}}&MSGName={{.CurrentMessageName}}&Sender={{.CurrentSender}}&CurrentPage={{.CurrentPage}}&LiveView=false">
          {{else}}
            <a class="button optionbuttons" href="/log?Level={{.CurrentLevel}}&TimeRange={{.CurrentTimeRange}}&Category={{.CurrentCategory}}&Impact={{.CurrentImpact}}&Severity={{.CurrentSeverity}}&MSGName={{.CurrentMessageName}}&Sender={{.CurrentSender}}&CurrentPage={{.CurrentPage}}&LiveView=true">
          {{end}}
              {{if .SetLiveView}}
              Disable live view
              {{else}}
              Enable live view
              {{end}}
            </a>
        </div>
      </div>
      <div class="w-col w-col-6">
        <h2 class="headercontrol">Filtering</h2>
        <div class="filters">
          <div class="w-form">
            <form class="filterformcontainer" id="wf-form-Filters" name="wf-form-Filters" data-name="Filters">
              <div class="w-row">
                <div class="w-col w-col-6">
                  <div class="columns">
                    <h3 class="subheader">By Time Range</h3>
                    <label for="TimeRange">Time range:</label>
                    <select class="w-select" id="TimeRange" name="TimeRange" data-name="TimeRange">
                      <option value="*"{{.IsSelected "*" .CurrentTimeRange | .Safe}}>Whole time range</option>
                      <option value="last5min"{{.IsSelected "last5min" .CurrentTimeRange | .Safe}}>Last 5 minutes</option>
                      <option value="last30min"{{.IsSelected "last30min" .CurrentTimeRange | .Safe}}>Last 30 minutes</option>
                      <option value="last60min"{{.IsSelected "last60min" .CurrentTimeRange | .Safe}}>Last 60 minutes</option>
                      <option value="last24h"{{.IsSelected "last24h" .CurrentTimeRange | .Safe}}>Last 24 hours</option>
                      <option value="last7d"{{.IsSelected "last7d" .CurrentTimeRange | .Safe}}>Last 7 days</option>
                      <option value="lastMonth"{{.IsSelected "lastMonth" .CurrentTimeRange | .Safe}}>Last month</option>
                    </select>
                  </div>
                </div>
                <div class="w-col w-col-6">
                  <div class="columns">
                    <h3 class="subheader">By Filter</h3>
                    <label for="Level">Level:</label>
                    <select class="w-select" id="Level" name="Level" data-name="Level">
                      <option value="*"{{.IsSelected "*" .CurrentLevel | .Safe}}>Any</option>
                      <option value="INFO"{{.IsSelected "INFO" .CurrentLevel | .Safe}}>Information</option>
                      <option value="WARN"{{.IsSelected "WARN" .CurrentLevel | .Safe}}>Warning</option>
                      <option value="SECURITY"{{.IsSelected "SECURITY" .CurrentLevel | .Safe}}>Security</option>
                      <option value="ERROR"{{.IsSelected "ERROR" .CurrentLevel | .Safe}}>Error</option>
                      <option value="DEBUG"{{.IsSelected "DEBUG" .CurrentLevel | .Safe}}>Debug</option>
                      <option value="TALKATIVE"{{.IsSelected "TALKATIVE" .CurrentLevel | .Safe}}>Talkative</option>
                    </select>
                    <label for="Category">Category:</label>
                    <select class="w-select" id="Category" name="Category" data-name="Category">
                      <option value="*"{{.IsSelected "*" .CurrentCategory | .Safe}}>Any</option>
                      <option value="BUSINESS"{{.IsSelected "BUSINESS" .CurrentCategory | .Safe}}>Business</option>
                      <option value="APP"{{.IsSelected "APP" .CurrentCategory | .Safe}}>Application</option>
                      <option value="USER"{{.IsSelected "USER" .CurrentCategory | .Safe}}>User</option>
                      <option value="SYSTEM"{{.IsSelected "SYSTEM" .CurrentCategory | .Safe}}>System</option>
                    </select>
                    <label for="Impact">Impact:</label>
                    <select class="w-select" id="Impact" name="Impact" data-name="Impact">
                      <option value="*"{{.IsSelected "*" .CurrentImpact | .Safe}}>Any</option>
                      <option value="LOW"{{.IsSelected "LOW" .CurrentImpact | .Safe}}>Low</option>
                      <option value="MIDDLE"{{.IsSelected "MIDDLE" .CurrentImpact | .Safe}}>Middle</option>
                      <option value="HIGH"{{.IsSelected "HIGH" .CurrentImpact | .Safe}}>High</option>
                      <option value="CRITICAL"{{.IsSelected "CRITICAL" .CurrentImpact | .Safe}}>Critical</option>
                      <option value="UNKNOWN"{{.IsSelected "UNKNOWN" .CurrentImpact | .Safe}}>Unknown</option>
                    </select>
                    <label for="Severity">Severity:</label>
                    <select class="w-select" id="Severity" name="Severity" data-name="Severity">
                      <option value="*"{{.IsSelected "*" .CurrentSeverity | .Safe}}>Any</option>
                      <option value="LOW"{{.IsSelected "LOW" .CurrentSeverity | .Safe}}>Low</option>
                      <option value="MIDDLE"{{.IsSelected "MIDDLE" .CurrentSeverity | .Safe}}>Middle</option>
                      <option value="HIGH"{{.IsSelected "HIGH" .CurrentSeverity | .Safe}}>High</option>
                      <option value="CRITICAL"{{.IsSelected "CRITICAL" .CurrentSeverity | .Safe}}>Critical</option>
                      <option value="UNKNOWN"{{.IsSelected "UNKNOWN" .CurrentSeverity | .Safe}}>Unknown</option>
                    </select>
                    <label for="MSGName">Message Names:</label>
                    <select class="w-select" id="MSGName" name="MSGName" data-name="MSGName">
                      <option value="*"{{.IsSelected "*" .CurrentMessageName | .Safe}}>Any</option>
                      {{$currentMessageName := .CurrentMessageName}}
                      {{range .MessageNames}}
                        <option value="{{.}}"{{.IsSelected . $currentMessageName | .Safe}}>{{.}}</option>
                      {{end}}
                    </select>
                    <label for="Sender">Sender:</label>
                    <select class="w-select" id="Sender" name="Sender" data-name="Sender">
                      <option value="*"{{.IsSelected "*" .CurrentSender | .Safe}}>Any</option>
                      {{$currentSender := .CurrentSender}}
                      {{range .Sender}}
                        <option value="{{.}}"{{.IsSelected . $currentSender | .Safe}}>{{.}}</option>
                      {{end}}
                    </select>
                  </div>
                </div>
              </div>
              <input class="w-button button optionbuttons applyfilters" type="submit" value="Apply filters" data-wait="Please wait...">
            </form>
          </div>
          <div class="w-form-done">
            <p>Thank you! Your submission has been received!</p>
          </div>
          <div class="w-form-fail">
            <p>Oops! Something went wrong while submitting the form :(</p>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="w-section">
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
      <div class="w-form-done">
        <p>Thank you! Your submission has been received!</p>
      </div>
      <div class="w-form-fail">
        <p>Oops! Something went wrong while submitting the form :(</p>
      </div>
    </div>
    <div class="w-hidden-main w-hidden-medium newlineblock"></div><a class="button changepagebutton" href="#">+1</a><a class="button changepagebutton" href="#">Last</a>
  </div>
  <script type="text/javascript" src="/admin/js/jquery.min.js"></script>
  <script type="text/javascript" src="/admin/js/webflow.js"></script>
</body>
</html>
{{end}}`
