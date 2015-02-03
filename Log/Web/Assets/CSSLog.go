package Assets

var CSSLog string = `
body {
  font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;
  color: #333;
  font-size: 14px;
  line-height: 20px;
}
h1 {
  margin-top: 20px;
  margin-bottom: 10px;
  font-size: 38px;
  line-height: 44px;
  font-weight: 700;
}
h2 {
  margin-top: 20px;
  margin-bottom: 10px;
  font-size: 32px;
  line-height: 36px;
  font-weight: 700;
}
h3 {
  margin-top: 20px;
  margin-bottom: 10px;
  font-size: 24px;
  line-height: 30px;
  font-weight: 700;
}
h4 {
  margin-top: 10px;
  margin-bottom: 10px;
  font-size: 18px;
  line-height: 24px;
  font-weight: 700;
}
h5 {
  margin-top: 10px;
  margin-bottom: 10px;
  font-size: 14px;
  line-height: 20px;
  font-weight: 700;
}
h6 {
  margin-top: 10px;
  margin-bottom: 10px;
  font-size: 12px;
  line-height: 18px;
  font-weight: 700;
}
p {
  margin-bottom: 5px;
}
.button {
  display: inline-block;
  padding: 4px 15px;
  background-color: black;
  color: white;
  text-align: center;
  text-decoration: none;
}
.button.changepagebutton {
  margin-right: 3px;
  margin-left: 3px;
  background-color: #545454;
}
.button.changepagebutton.pagechangesubmit {
  margin-right: 30px;
}
.button.optionbuttons {
  margin: 6px;
  background-color: #545454;
}
.button.optionbuttons.applyfilters {
  margin-left: 0px;
}
.headercontainer {
  height: 50px;
  margin-top: -19px;
  background-color: black;
  color: white;
  text-align: center;
}
.logcontainer {
  margin-top: 0px;
}
.loglist {
  margin-right: 3px;
  margin-bottom: 3px;
  margin-left: 3px;
  padding-top: 0px;
  padding-left: 0px;
  list-style-type: none;
}
.logline {
  margin: 3px 6px;
  padding: 6px;
  border: 1px solid #9e9e9e;
  border-radius: 3px;
  font-family:'Source Code Pro', sans-serif;
  font-size: 10px;
  font-weight: 900;
}
.logline.loga {
  background-color: #ededed;
}
.logline.logb {
  background-color: #cfcfcf;
}
.logline.logwarn {
  color: #db7602;
}
.logline.logdebug {
  color: #04f;
}
.logline.logerror {
  color: #c00;
}
.logline.loginfo {
  color: black;
}
.logline.logtalkative {
  color: #a3a2a2;
}
.logline.logsecurity {
  color: #db4500;
}
.logheadercontainer {
  margin-top: 10px;
  text-align: center;
}
.showposition {
  margin-top: 10px;
  font-size: 20px;
}
.controlcontainer {
  margin-top: 10px;
}
.pagecontainer {
  margin-top: 10px;
  text-align: center;
  list-style-type: none;
}
.icons {
  position: absolute;
  display: block;
  width: 25px;
  height: 25px;
  margin-top: 15px;
  margin-left: 254px;
  padding-top: 0px;
  font-family: Glyphicons, sans-serif;
  color: black;
  font-size: 25px;
  font-weight: 400;
  text-decoration: none;
}
.icons.oneback {
  margin-top: 17px;
  margin-left: 219px;
}
.icons.tofirst {
  margin-top: 17px;
  margin-left: 255px;
}
.icons.next {
  margin-top: -38px;
  margin-left: 668px;
}
.icons.tolast {
  margin-top: -38px;
  margin-left: 702px;
}
.formfield {
  display: inline;
  margin-right: 6px;
  margin-left: 30px;
  color: black;
  font-weight: 400;
}
.formpages {
  display: inline;
}
.currentpage {
  display: inline-block;
  width: 100px;
  margin-top: 7px;
  margin-right: 6px;
  padding-top: 3px;
  padding-bottom: 3px;
  float: none;
  clear: none;
  color: black;
  text-align: center;
}
.formpageswrapper {
  display: inline;
}
.textcountpages {
  display: inline;
  margin-right: 6px;
  color: black;
}
.logchangepagebutton {
  height: 40px;
  margin-left: 10px;
  padding-top: 3px;
  padding-bottom: 3px;
  background-color: #a8a8a8;
  color: black;
  font-weight: 700;
}
.headercontrol {
  text-align: center;
}
.headeroutput {
  text-align: center;
}
.options {
  background-color: #ededed;
}
.filters {
  background-color: #ededed;
}
.labels {
  display: block;
}
@media (max-width: 991px) {
  .icons.oneback {
    margin-left: 113px;
  }
  .icons.tofirst {
    margin-left: 152px;
  }
  .icons.next {
    margin-left: 563px;
  }
  .icons.tolast {
    margin-left: 602px;
  }
}
@media (max-width: 767px) {
  .button.changepagebutton.pagechangesubmit {
    margin-right: 0px;
  }
  .icons.oneback {
    margin-left: 40px;
  }
  .icons.tofirst {
    margin-left: 73px;
  }
  .icons.next {
    margin-left: 483px;
  }
  .icons.tolast {
    margin-left: 518px;
  }
  .formfield {
    margin-left: 0px;
  }
  .newlineblock {
    height: 10px;
  }
}
@media (max-width: 479px) {
  .button.changepagebutton.pagechangesubmit {
    margin-right: 3px;
    margin-bottom: 20px;
    text-align: center;
  }
  .pagecontainer {
    margin-top: 20px;
  }
  .icons.oneback {
    margin-top: 113px;
    margin-left: 25px;
  }
  .icons.tofirst {
    margin-top: 113px;
    margin-left: 62px;
  }
  .icons.next {
    margin-top: 18px;
    margin-left: 242px;
  }
  .icons.tolast {
    margin-top: 18px;
    margin-left: 278px;
  }
  .formfield {
    margin-left: 0px;
  }
  .currentpage {
    margin-top: 20px;
  }
  .newlineblock {
    display: block;
    height: 10px;
  }
}

@font-face {
  font-family: 'Glyphicons';
  src: url('/binaryAssets/SourceCodePro-Black.eot') format('embedded-opentype'), url('/binaryAssets/SourceCodePro-Black.ttf') format('truetype'), url('/binaryAssets/SourceCodePro-Black.otf') format('opentype');
  font-weight: 400;
  font-style: normal;
}`
