/*
HTCAP - beta 1
Author: filippo.cavallarin@wearesegment.com

This program is free software; you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation; either version 2 of the License, or (at your option) any later
version.
*/

var system = require('system');
var fs = require('fs');



phantom.injectJs("functions.js");
phantom.injectJs("options.js");
phantom.injectJs("probe.js");


var startTime = Date.now();


var site = "";
var response = null;
//var showHelp = false;

var headers = {};
console.log("js args: " + system.args);
var args = getopt(system.args,"hVaftUJo:dICc:MSEp:Tsx:A:r:mHX:PD:R:");

var page = require('webpage').create();
var page_settings = {encoding: "utf8"};
var random = "IsHOulDb34RaNd0MsTR1ngbUt1mN0t";


if(typeof args == 'string'){
	console.log("Error: " + args);
	phantom.exit(-1);
}

for(var a = 0; a < args.opts.length; a++){
	switch(args.opts[a][0]){
		case "P":
			page_settings.operation = "POST";
			break;
		case "D":
			page_settings.data = args.opts[a][1];
			break;
		case "R":
			random = args.opts[a][1];
			break;
	}
}


parseArgsToOptions(args);

site = args.args[1];

if(!site){
	usage();
	phantom.exit(-1);
}

if(site.length < 4 || site.substring(0,4).toLowerCase() != "http"){
	site = "http://" + site;
}

console.log("[");

/* maximum execution time */
setTimeout(execTimedOut,options.maxExecTime);



phantom.onError = function(msg, trace) {
  var msgStack = ['PHANTOM ERROR: ' + msg];
  if (trace && trace.length) {
    msgStack.push('TRACE:');
    trace.forEach(function(t) {
      msgStack.push(' -> ' + (t.file || t.sourceURL) + ': ' + t.line + (t.function ? ' (in function ' + t.function +')' : ''));
    });
  }
  console.error(msgStack.join('\n'));
  phantom.exit(1);
};



page.onConsoleMessage = function(msg, lineNum, sourceId) {
	if(options.verbose)
		console.log("console: " + msg);
};
page.onError = function(msg, lineNum, sourceId) {
	if(options.verbose)
		console.log("console error: on   " + JSON.stringify(lineNum) + " " + msg);
};

page.onAlert = function(msg) {
	if(options.verbose)
  		console.log('ALERT: ' + msg);
};

page.settings.userAgent = options.userAgent;
page.settings.loadImages = options.loadImages;



page.onResourceReceived = function(resource) {
	if(window.response === null){
		window.response = resource;
		// @TODO sanytize response.contentType

	}
};


page.onResourceRequested = function(requestData, networkRequest) {

  var json = JSON.parse(JSON.stringify(requestData));
  var r = "";

  for (var i = 0; i < requestData.headers.length; ++i) {

    if (requestData.headers[i].name.toLowerCase() == "referer") {
      r = requestData.headers[i].value;
    }
  }
  var obj = {
    type: "link",
    method: json.method,
    url : json.url,
    refer: r,
    data: null
  };

  console.log('{"type":"request", "data":' + JSON.stringify(obj) + "}");
};

// to detect window.location= / document.location.href=
page.onNavigationRequested = onNavigationRequested;

page.onConfirm = function(msg) {return true;}; // recently changed


page.onInitialized = function(){
	// try to hide phantomjs
	page.evaluate(function(){
		window.__callPhantom = window.callPhantom;
		delete window.callPhantom;
	});

	startProbe(random);

};


page.onCallback = function(data) {
	switch(data.cmd){
		case "print":
			console.log(data.argument);
			break;
		case "die": // @TMP
			console.log(data.argument);
			phantom.exit(0);
		case "render":
			page.render("htcap_render.png");
			break;
		case "end":
			if(options.returnHtml){
				page.evaluate(function(options){
					window.__PROBE__.printPageHTML();
				}, options);
			}

			printStatus("ok", window.response.contentType);
			phantom.exit(0);
			break;

	}

};



if(options.httpAuth){
	headers['Authorization'] = 'Basic ' + btoa(options.httpAuth[0] + ":" + options.httpAuth[1]);
}

if(options.referer){
	headers['Referer'] = options.referer;
}

if (options.header) {
  for(var i = 0; i < options.setHeaders.length; ++i) {
    //console.log(options.setHeaders[i][0] + " : " + options.setHeaders[i][1]);
    headers[options.setHeaders[i][0]] = options.setHeaders[i][1];
  }
}

page.customHeaders = headers;


for(var a = 0; a < options.setCookies.length; a++){
	// maybe this is wrogn acconding to rfc .. but phantomjs cannot set cookie witout a domain...
	if(!options.setCookies[a].domain){
		var purl = document.createElement("a");
		purl.href=site;
		options.setCookies[a].domain = purl.hostname;
	}
	if(options.setCookies[a].expires)
		options.setCookies[a].expires *= 1000;

	phantom.addCookie(options.setCookies[a]);

}


page.onDOMContentLoaded = function() {
    // your code here
    console.log('DOMContentLoaded');
    phantom.exit();
};


page.open(site, page_settings, function(status) {
	var response = window.response; // just to be clear
	if (status !== 'success'){
		var mess = "";
		var out = {response: response};
		if(!response || response.headers.length === 0){
			printStatus("error", "load");
			phantom.exit(1);
		}

		// check for redirect first
		for(var a = 0; a < response.headers.length; a++){
			if(response.headers[a].name.toLowerCase() == 'location'){

				if(options.getCookies){
					printCookies(response.headers, site);
				}
				printStatus("ok", null, null, response.headers[a].value);
				phantom.exit(0);
			}
		}

		assertContentTypeHtml(response);

    // setTimeout(function () {
    //   phantomjs.exit(1);
    // }, 1000);
		phantom.exit(1);
	}


	if(options.getCookies){
		printCookies(response.headers, site);
	}

	assertContentTypeHtml(response);


	page.evaluate(function(){
		window.__PROBE__.waitAjax(function(xhrs){
			console.log("start");
			setTimeout(function() {
				window.__PROBE__.startAnalysis();
			}, 100);

		});
	});
  
});



