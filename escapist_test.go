package escapist

import (
	"fmt"
	"html"
	"testing"
)

var (
	exampleBasicStr      = "Hello there Mr. Joe <inject stuff>"
	exampleBasicNoRepStr = "Hello there Mr. Joe. This is a safe string"
	// This is the html from our github repo
	exampleHTMLPageStr = `
<html lang="en" class=" is-copy-enabled is-u2f-enabled"><head prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# object: http://ogp.me/ns/object# article: http://ogp.me/ns/article# profile: http://ogp.me/ns/profile#">
    <meta charset="utf-8">

    <link crossorigin="anonymous" href="https://assets-cdn.github.com/assets/frameworks-0dd18d999f1e3d1e515cf15c5628a70a3859d1af1c5daadeaa73ab4adfdeae9e.css" integrity="sha256-DdGNmZ8ePR5RXPFcViinCjhZ0a8cXareqnOrSt/erp4=" media="all" rel="stylesheet">
    <link crossorigin="anonymous" href="https://assets-cdn.github.com/assets/github-f89b14bb39d46baff97b6f66f5d6ca10c5dab4c17d37a2d884ecc3a5e474c6b5.css" integrity="sha256-+JsUuznUa6/5e29m9dbKEMXatMF9N6LYhOzDpeR0xrU=" media="all" rel="stylesheet">
    
    
    
    
    <link crossorigin="anonymous" href="https://assets-cdn.github.com/assets/system-fonts-d60dee4b68e3a38b01c0389551a0e38f04acd897ec85592f405d045ee853b3d1.css" integrity="sha256-1g3uS2jjo4sBwDiVUaDjjwSs2JfshVkvQF0EXuhTs9E=" media="all" rel="stylesheet">

    <link as="script" href="https://assets-cdn.github.com/assets/frameworks-149d9338c2665172870825c78fa48fdcca4d431d067cbf5fda7120d9e39cc738.js" rel="preload">
    
    <link as="script" href="https://assets-cdn.github.com/assets/github-c7401dbe3046797176f246d950cb4e6f80e1a4847c679d92c703af984d3698cd.js" rel="preload">

    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta http-equiv="Content-Language" content="en">
    <meta name="viewport" content="width=device-width">
    
    
    <title>itsmontoya/escapist: A byteslice-focused HTML escape library</title>
    <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="GitHub">
    <link rel="fluid-icon" href="https://github.com/fluidicon.png" title="GitHub">
    <link rel="apple-touch-icon" href="/apple-touch-icon.png">
    <link rel="apple-touch-icon" sizes="57x57" href="/apple-touch-icon-57x57.png">
    <link rel="apple-touch-icon" sizes="60x60" href="/apple-touch-icon-60x60.png">
    <link rel="apple-touch-icon" sizes="72x72" href="/apple-touch-icon-72x72.png">
    <link rel="apple-touch-icon" sizes="76x76" href="/apple-touch-icon-76x76.png">
    <link rel="apple-touch-icon" sizes="114x114" href="/apple-touch-icon-114x114.png">
    <link rel="apple-touch-icon" sizes="120x120" href="/apple-touch-icon-120x120.png">
    <link rel="apple-touch-icon" sizes="144x144" href="/apple-touch-icon-144x144.png">
    <link rel="apple-touch-icon" sizes="152x152" href="/apple-touch-icon-152x152.png">
    <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon-180x180.png">
    <meta property="fb:app_id" content="1401488693436528">

      <meta content="https://avatars1.githubusercontent.com/u/928954?v=3&amp;s=400" name="twitter:image:src"><meta content="@github" name="twitter:site"><meta content="summary" name="twitter:card"><meta content="itsmontoya/escapist" name="twitter:title"><meta content="escapist - A byteslice-focused HTML escape library" name="twitter:description">
      <meta content="https://avatars1.githubusercontent.com/u/928954?v=3&amp;s=400" property="og:image"><meta content="GitHub" property="og:site_name"><meta content="object" property="og:type"><meta content="itsmontoya/escapist" property="og:title"><meta content="https://github.com/itsmontoya/escapist" property="og:url"><meta content="escapist - A byteslice-focused HTML escape library" property="og:description">
      <meta name="browser-stats-url" content="https://api.github.com/_private/browser/stats">
    <meta name="browser-errors-url" content="https://api.github.com/_private/browser/errors">
    <link rel="assets" href="https://assets-cdn.github.com/">
    <link rel="web-socket" href="wss://live.github.com/_sockets/OTI4OTU0OmE0ZDJkMmVhYTY5ZmE4MjJkMDU1MGQ4ZWUyNjcwOTNkOjJiNTQwZmJjN2FlNTRkYjYxM2QzYmJhODA4Yzc4YzU2YTc5NTFmNGYxY2Q0YThjYTA3MDA0NWRlY2E0ZTQwNTk=--bf216a45406b1885dd422dde9d6fccf456b72124">
    <meta name="pjax-timeout" content="1000">
    <link rel="sudo-modal" href="/sessions/sudo_modal">

    <meta name="msapplication-TileImage" content="/windows-tile.png">
    <meta name="msapplication-TileColor" content="#ffffff">
    <meta name="selected-link" value="repo_source" data-pjax-transient="">

    <meta name="google-site-verification" content="KT5gs8h0wvaagLKAVWq8bbeNwnZZK1r1XQysX3xurLU">
<meta name="google-site-verification" content="ZzhVyEFwb7w3e0-uOTltm8Jsck2F5StVihD0exw2fsA">
    <meta name="google-analytics" content="UA-3769691-2">

<meta content="collector.githubapp.com" name="octolytics-host"><meta content="github" name="octolytics-app-id"><meta content="323575B1:22AB:13964D7:578927CF" name="octolytics-dimension-request_id"><meta content="928954" name="octolytics-actor-id"><meta content="itsmontoya" name="octolytics-actor-login"><meta content="c38240e87e8284c53eea885ace6e749af534ccd22e95960765cec87aed6b4fc8" name="octolytics-actor-hash">
<meta content="/<user-name>/<repo-name>" data-pjax-transient="true" name="analytics-location">



  <meta class="js-ga-set" name="dimension1" content="Logged In">



        <meta name="hostname" content="github.com">
    <meta name="user-login" content="itsmontoya">

        <meta name="expected-hostname" content="github.com">
      <meta name="js-proxy-site-detection-payload" content="N2ZiMzI5MzkwYTY0NDI2MWM2NDNhNzUyYmQ1YjZjMDJlZDFlZDk5ZmQ4Mjg1NDc1MGQ5YzM5MWQwNTc3M2M0N3x7InJlbW90ZV9hZGRyZXNzIjoiNTAuNTMuMTE3LjE3NyIsInJlcXVlc3RfaWQiOiIzMjM1NzVCMToyMkFCOjEzOTY0RDc6NTc4OTI3Q0YiLCJ0aW1lc3RhbXAiOjE0Njg2MDY0MTV9">


      <link rel="mask-icon" href="https://assets-cdn.github.com/pinned-octocat.svg" color="#4078c0">
      <link rel="icon" type="image/x-icon" href="https://assets-cdn.github.com/favicon.ico">

    <meta name="html-safe-nonce" content="65bfd46423f939d881f8c59e13cbdb8152d332ba">
    <meta content="98fbb59aa829e4abf83eee02b373688a17c53916" name="form-nonce">

    <meta http-equiv="x-pjax-version" content="de87b11f40bfe4609305bb57745744d3">
    

      
  <meta name="description" content="escapist - A byteslice-focused HTML escape library">
  <meta name="go-import" content="github.com/itsmontoya/escapist git https://github.com/itsmontoya/escapist.git">

  <meta content="928954" name="octolytics-dimension-user_id"><meta content="itsmontoya" name="octolytics-dimension-user_login"><meta content="63441746" name="octolytics-dimension-repository_id"><meta content="itsmontoya/escapist" name="octolytics-dimension-repository_nwo"><meta content="true" name="octolytics-dimension-repository_public"><meta content="false" name="octolytics-dimension-repository_is_fork"><meta content="63441746" name="octolytics-dimension-repository_network_root_id"><meta content="itsmontoya/escapist" name="octolytics-dimension-repository_network_root_nwo">
  <link href="https://github.com/itsmontoya/escapist/commits/master.atom" rel="alternate" title="Recent Commits to escapist:master" type="application/atom+xml">


      <link rel="canonical" href="https://github.com/itsmontoya/escapist" data-pjax-transient="">
  <style>[href^="https://www.googleadservices.com/pagead/aclk?"]
{display:none !important;}</style></head>


  <body class="logged-in   env-production linux vis-public">
    <div id="js-pjax-loader-bar" class="pjax-loader-bar"></div>
    <a href="#start-of-content" tabindex="1" class="accessibility-aid js-skip-to-content">Skip to content</a>

    
    
    



        <div class="header header-logged-in true" role="banner">
  <div class="container clearfix">

    <a class="header-logo-invertocat" href="https://github.com/" data-hotkey="g d" aria-label="Homepage" data-ga-click="Header, go to dashboard, icon:logo">
  <svg aria-hidden="true" class="octicon octicon-mark-github" height="28" version="1.1" viewBox="0 0 16 16" width="28"><path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0 0 16 8c0-4.42-3.58-8-8-8z"></path></svg>
</a>


        <div class="header-search scoped-search site-scoped-search js-site-search" role="search">
  <form accept-charset="UTF-8" action="/logout" class="logout-form" data-form-nonce="98fbb59aa829e4abf83eee02b373688a17c53916" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="✓"><input name="authenticity_token" type="hidden" value="vhreo1zjoIktgS1eUNxlXVDu7gA/9u34AWXs0odk7xYbzKd0nN6zoANlOos5X8Xo2tApKJp9OquIe7audoo+VA=="></div>
          <button class="dropdown-item dropdown-signout" data-ga-click="Header, sign out, icon:logout">
            Sign out
          </button>
</form>      </div>
    </div>
  </li>
</ul>


    
  </div>
</div>


      


    <div id="start-of-content" class="accessibility-aid"></div>

      <div id="js-flash-container">
</div>


    <div role="main">
        <div itemscope="" itemtype="http://schema.org/SoftwareSourceCode">
    <div id="js-repo-pjax-container" data-pjax-container="">
      
<div class="pagehead repohead instapaper_ignore readability-menu experiment-repo-nav">
  <div class="container repohead-details-container">

    

<ul class="pagehead-actions">

  <li>
        <form accept-charset="UTF-8" action="/itsmontoya/escapist/unstar" class="starred" data-form-nonce="98fbb59aa829e4abf83eee02b373688a17c53916" data-remote="true" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="✓"><input name="authenticity_token" type="hidden" value="dDEPAeYRHV/XiyBUKoj6CVvFSquudp1iho2qFFg9qAHKwGVCF2Msski/dlStyWC0+pcLzu8/PAEbb4ii1rMS8g=="></div>
      <button class="btn btn-sm btn-with-count js-toggler-target" aria-label="Unstar this repository" title="Unstar itsmontoya/escapist" data-ga-click="Repository, click unstar button, action:files#disambiguate; text:Unstar">
        <svg aria-hidden="true" class="octicon octicon-star" height="16" version="1.1" viewBox="0 0 14 16" width="14"><path d="M14 6l-4.9-.64L7 1 4.9 5.36 0 6l3.6 3.26L2.67 14 7 11.67 11.33 14l-.93-4.74z"></path></svg>
        Unstar
      </button>
        <a class="social-count js-social-count" href="/itsmontoya/escapist/stargazers">
          0
        </a>
</form>
   <form accept-charset="UTF-8" action="/users/set_protocol?protocol_selector=ssh&amp;protocol_type=push" data-form-nonce="98fbb59aa829e4abf83eee02b373688a17c53916" data-remote="true" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="✓"><input name="authenticity_token" type="hidden" value="BVH9bclXdKNCelTHINwxSz+DPhEpicj9KXMPxmBFORYhagedKuQoouOPHQZOsnDGAQK1h1JhQJv3klJzVc2vZw=="></div><button class="btn-link btn-change-protocol js-toggler-target right" type="submit">Use SSH</button></form>

        <h4 class="mb-1">
          Clone with HTTPS
          <a class="muted-link" href="https://help.github.com/articles/which-remote-url-should-i-use" target="_blank">
            <svg aria-hidden="true" class="octicon octicon-question" height="16" version="1.1" viewBox="0 0 14 16" width="14"><path d="M6 10h2v2H6v-2zm4-3.5C10 8.64 8 9 8 9H6c0-.55.45-1 1-1h.5c.28 0 .5-.22.5-.5v-1c0-.28-.22-.5-.5-.5h-1c-.28 0-.5.22-.5.5V7H4c0-1.5 1.5-3 3-3s3 1 3 2.5zM7 2.3c3.14 0 5.7 2.56 5.7 5.7s-2.56 5.7-5.7 5.7A5.71 5.71 0 0 1 1.3 8c0-3.14 2.56-5.7 5.7-5.7zM7 1C3.14 1 0 4.14 0 8s3.14 7 7 7 7-3.14 7-7-3.14-7-7-7z"></path></svg>
          </a>
        </h4>
        <p class="mb-2 get-repo-decription-text">
          Use Git or checkout with SVN using the web URL.
        </p>

        <div class="input-group js-zeroclipboard-container">
  <input type="text" class="form-control input-monospace input-sm js-zeroclipboard-target js-url-field" value="https://github.com/itsmontoya/escapist.git" aria-label="Clone this repository at https://github.com/itsmontoya/escapist.git" readonly="">
  <div class="input-group-button">
    <button aria-label="Copy to clipboard" class="js-zeroclipboard btn btn-sm zeroclipboard-button tooltipped tooltipped-s" data-copied-hint="Copied!" type="button"><svg aria-hidden="true" class="octicon octicon-clippy" height="16" version="1.1" viewBox="0 0 14 16" width="14"><path d="M2 13h4v1H2v-1zm5-6H2v1h5V7zm2 3V8l-3 3 3 3v-2h5v-2H9zM4.5 9H2v1h2.5V9zM2 12h2.5v-1H2v1zm9 1h1v2c-.02.28-.11.52-.3.7-.19.18-.42.28-.7.3H1c-.55 0-1-.45-1-1V4c0-.55.45-1 1-1h3c0-1.11.89-2 2-2 1.11 0 2 .89 2 2h3c.55 0 1 .45 1 1v5h-1V6H1v9h10v-2zM2 5h8c0-.55-.45-1-1-1H8c-.55 0-1-.45-1-1s-.45-1-1-1-1 .45-1 1-.45 1-1 1H3c-.55 0-1 .45-1 1z"></path></svg></button>
  </div>
</div>

      </div>

      <div class="clone-options ssh-clone-options">
       <form accept-charset="UTF-8" action="/itsmontoya/escapist/new/master" class="btn-group-form" data-form-nonce="98fbb59aa829e4abf83eee02b373688a17c53916" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="✓"><input name="authenticity_token" type="hidden" value="BbbVpORHHpT6tapyauI/nHJ+wsfAVH230z0l5cWHxm/wdHlG4RYbqvlw27JOqmqrFAEFeYAn8H/cCtAxyuxCZQ=="></div>
    <button class="btn btn-sm" type="submit" data-disable-with="Creating file…">
      Create new file
    </button>
</form>

      <a href="/itsmontoya/escapist/upload/master" class="btn btn-sm">
        Upload files
      </a>

    <a href="/itsmontoya/escapist/find/master" class="btn btn-sm empty-icon right" data-pjax="" data-hotkey="t" data-ga-click="Repository, find file, location:repo overview">
      Find file
    </a>
  </div>

  
<div class="select-menu branch-select-menu js-menu-container js-select-menu left">
  <button class="btn btn-sm select-menu-button js-menu-target css-truncate" data-hotkey="w" title="master" type="button" aria-label="Switch branches or tags" tabindex="0" aria-haspopup="true">
    <i>Branch:</i>
    <span class="js-select-button css-truncate-target">master</span>
  </button>

  <div class="select-menu-modal-holder js-menu-content js-navigation-container" data-pjax="" aria-hidden="true">

    <div class="select-menu-modal">
      <div class="select-menu-header">
        <svg aria-label="Close" class="octicon octicon-x js-menu-close" height="16" role="img" version="1.1" viewBox="0 0 12 16" width="12"><path d="M7.48 8l3.75 3.75-1.48 1.48L6 9.48l-3.75 3.75-1.48-1.48L4.52 8 .77 4.25l1.48-1.48L6 6.52l3.75-3.75 1.48 1.48z"></path></svg>
        <span class="select-menu-title">Switch branches/tags</span>
      </div>

      <div class="select-menu-filters">
        <div class="select-menu-text-filter">
          <input type="text" aria-label="Find or create a branch…" id="context-commitish-filter-field" class="form-control js-filterable-field js-navigation-enable" placeholder="Find or create a branch…">
        </div>
        <div class="select-menu-tabs">
          <ul>
            <li class="select-menu-tab">
              <a href="#" data-tab-filter="branches" data-filter-placeholder="Find or create a branch…" class="js-select-menu-tab" role="tab">Branches</a>
            </li>
            <li class="select-menu-tab">
              <a href="#" data-tab-filter="tags" data-filter-placeholder="Find a tag…" class="js-select-menu-tab" role="tab">Tags</a>
            </li>
          </ul>
        </div>
      </div>

      <div class="select-menu-list select-menu-tab-bucket js-select-menu-tab-bucket" data-tab-filter="branches" role="menu">

        <div data-filterable-for="context-commitish-filter-field" data-filterable-type="substring">


            <a class="select-menu-item js-navigation-item js-navigation-open selected" href="/itsmontoya/escapist/tree/master" data-name="master" data-skip-pjax="true" rel="nofollow">
              <svg aria-hidden="true" class="octicon octicon-check select-menu-item-icon" height="16" version="1.1" viewBox="0 0 12 16" width="12"><path d="M12 5l-8 8-4-4 1.5-1.5L4 10l6.5-6.5z"></path></svg>
              <span class="select-menu-item-text css-truncate-target js-select-menu-filter-text" title="master">
                master
              </span>
            </a>
        </div>

		<form accept-charset="UTF-8" action="/itsmontoya/escapist/branches" class="js-create-branch select-menu-item select-menu-new-item-form js-navigation-item js-new-item-form" data-form-nonce="98fbb59aa829e4abf83eee02b373688a17c53916" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="✓"><input name="authenticity_token" type="hidden" value="ibSCB+F6coIFFEXbmya4rTDIaJOnoDcvAsXXEf6vZF/1wAyS7eN/WOb0T6YslKHzZHjlpkTjxiQNNYSjnuavaw=="></div>
          <svg aria-hidden="true" class="octicon octicon-git-branch select-menu-item-icon" height="16" version="1.1" viewBox="0 0 10 16" width="10"><path d="M10 5c0-1.11-.89-2-2-2a1.993 1.993 0 0 0-1 3.72v.3c-.02.52-.23.98-.63 1.38-.4.4-.86.61-1.38.63-.83.02-1.48.16-2 .45V4.72a1.993 1.993 0 0 0-1-3.72C.88 1 0 1.89 0 3a2 2 0 0 0 1 1.72v6.56c-.59.35-1 .99-1 1.72 0 1.11.89 2 2 2 1.11 0 2-.89 2-2 0-.53-.2-1-.53-1.36.09-.06.48-.41.59-.47.25-.11.56-.17.94-.17 1.05-.05 1.95-.45 2.75-1.25S8.95 7.77 9 6.73h-.02C9.59 6.37 10 5.73 10 5zM2 1.8c.66 0 1.2.55 1.2 1.2 0 .65-.55 1.2-1.2 1.2C1.35 4.2.8 3.65.8 3c0-.65.55-1.2 1.2-1.2zm0 12.41c-.66 0-1.2-.55-1.2-1.2 0-.65.55-1.2 1.2-1.2.65 0 1.2.55 1.2 1.2 0 .65-.55 1.2-1.2 1.2zm6-8c-.66 0-1.2-.55-1.2-1.2 0-.65.55-1.2 1.2-1.2.65 0 1.2.55 1.2 1.2 0 .65-.55 1.2-1.2 1.2z"></path></svg>
            <div class="select-menu-item-text">
              <span class="select-menu-item-heading">Create branch: <span class="js-new-item-name"></span></span>
              <span class="description">from ‘master’</span>
            </div>
            <input type="hidden" name="name" id="name" class="js-new-item-value">
            <input type="hidden" name="branch" id="branch" value="master">
            <input type="hidden" name="path" id="path" value="">
</form>
      </div>

      <div class="select-menu-list select-menu-tab-bucket js-select-menu-tab-bucket" data-tab-filter="tags">
        <div data-filterable-for="context-commitish-filter-field" data-filterable-type="substring">


        </div>

        <div class="select-menu-no-results">Nothing to show</div>
      </div>

    </div>
  </div>
</div>


      <a href="/itsmontoya/escapist/pull/new/master" class="btn btn-sm new-pull-request-btn" data-pjax="" data-ga-click="Repository, new pull request, location:repo overview">
        New pull request
      </a>

  <div class="breadcrumb">
    
  </div>
</div>





  <div class="commit-tease js-details-container">
    <span class="right">
      Latest commit
      <a class="commit-tease-sha" href="/itsmontoya/escapist/commit/2f58730bc01a7d6ce030e0a419cac6e290c0626a" data-pjax="">
        2f58730
      </a>
      <span itemprop="dateModified"><relative-time datetime="2016-07-15T18:13:08Z" title="Jul 15, 2016, 11:13 AM PDT">a minute ago</relative-time></span>
    </span>


    <span class="commit-author-section">
      <img alt="@itsmontoya" class="avatar" height="20" src="https://avatars2.githubusercontent.com/u/928954?v=3&amp;s=40" width="20">
      <a href="/itsmontoya" class="user-mention" rel="author">itsmontoya</a>
    </span>

        <a href="/itsmontoya/escapist/commit/2f58730bc01a7d6ce030e0a419cac6e290c0626a" class="message" data-pjax="true" title="Small update to readme">Small update to readme</a>
    

  </div>


<div class="file-wrap">

  <a href="/itsmontoya/escapist/tree/2f58730bc01a7d6ce030e0a419cac6e290c0626a" class="hidden js-permalink-shortcut" data-hotkey="y">Permalink</a>

  <table class="files js-navigation-container js-active-navigation-container" data-pjax="">


    <tbody>
      <tr class="warning include-fragment-error">
        <td class="icon"><svg aria-hidden="true" class="octicon octicon-alert" height="16" version="1.1" viewBox="0 0 16 16" width="16"><path d="M8.865 1.52c-.18-.31-.51-.5-.87-.5s-.69.19-.87.5L.275 13.5c-.18.31-.18.69 0 1 .19.31.52.5.87.5h13.7c.36 0 .69-.19.86-.5.17-.31.18-.69.01-1L8.865 1.52zM8.995 13h-2v-2h2v2zm0-3h-2V6h2v4z"></path></svg></td>
        <td class="content" colspan="3">Failed to load latest commit information.</td>
      </tr>

        <tr class="js-navigation-item">
          <td class="icon">
            <svg aria-hidden="true" class="octicon octicon-file-text" height="16" version="1.1" viewBox="0 0 12 16" width="12"><path d="M6 5H2V4h4v1zM2 8h7V7H2v1zm0 2h7V9H2v1zm0 2h7v-1H2v1zm10-7.5V14c0 .55-.45 1-1 1H1c-.55 0-1-.45-1-1V2c0-.55.45-1 1-1h7.5L12 4.5zM11 5L8 2H1v12h10V5z"></path></svg>
            <img alt="" class="spinner" height="16" src="https://assets-cdn.github.com/images/spinners/octocat-spinner-32.gif" width="16">
          </td>
          <td class="content">
            <span class="css-truncate css-truncate-target"><a href="/itsmontoya/escapist/blob/master/LICENSE" class="js-navigation-open" id="9879d6db96fd29134fc802214163b95a-4be6530ebb91a4cf6cb7404616961f8e8cdedfda" itemprop="license" title="LICENSE">LICENSE</a></span>
          </td>
          <td class="message">
            <span class="css-truncate css-truncate-target">
                  <a href="/itsmontoya/escapist/commit/f6a01e7bef17a8d991876b907b2d3b344b08b2fc" class="message" data-pjax="true" title="Initial commit">Initial commit</a>
            </span>
          </td>
          <td class="age">
            <span class="css-truncate css-truncate-target"><time-ago datetime="2016-07-15T18:09:23Z" title="Jul 15, 2016, 11:09 AM PDT">5 minutes ago</time-ago></span>
          </td>
        </tr>
        <tr class="js-navigation-item">
          <td class="icon">
            <svg aria-hidden="true" class="octicon octicon-file-text" height="16" version="1.1" viewBox="0 0 12 16" width="12"><path d="M6 5H2V4h4v1zM2 8h7V7H2v1zm0 2h7V9H2v1zm0 2h7v-1H2v1zm10-7.5V14c0 .55-.45 1-1 1H1c-.55 0-1-.45-1-1V2c0-.55.45-1 1-1h7.5L12 4.5zM11 5L8 2H1v12h10V5z"></path></svg>
            <img alt="" class="spinner" height="16" src="https://assets-cdn.github.com/images/spinners/octocat-spinner-32.gif" width="16">
          </td>
          <td class="content">
            <span class="css-truncate css-truncate-target"><a href="/itsmontoya/escapist/blob/master/README.md" class="js-navigation-open" id="04c6e90faac2675aa89e2176d2eec7d8-3dadbca795e9fe33a1471b6a8d4b8e64d3d80c68" title="README.md">README.md</a></span>
          </td>
          <td class="message">
            <span class="css-truncate css-truncate-target">
                  <a href="/itsmontoya/escapist/commit/2f58730bc01a7d6ce030e0a419cac6e290c0626a" class="message" data-pjax="true" title="Small update to readme">Small update to readme</a>
            </span>
          </td>
          <td class="age">
            <span class="css-truncate css-truncate-target"><time-ago datetime="2016-07-15T18:13:08Z" title="Jul 15, 2016, 11:13 AM PDT">a minute ago</time-ago></span>
          </td>
        </tr>
        <tr class="js-navigation-item">
          <td class="icon">
            <svg aria-hidden="true" class="octicon octicon-file-text" height="16" version="1.1" viewBox="0 0 12 16" width="12"><path d="M6 5H2V4h4v1zM2 8h7V7H2v1zm0 2h7V9H2v1zm0 2h7v-1H2v1zm10-7.5V14c0 .55-.45 1-1 1H1c-.55 0-1-.45-1-1V2c0-.55.45-1 1-1h7.5L12 4.5zM11 5L8 2H1v12h10V5z"></path></svg>
            <img alt="" class="spinner" height="16" src="https://assets-cdn.github.com/images/spinners/octocat-spinner-32.gif" width="16">
          </td>
          <td class="content">
            <span class="css-truncate css-truncate-target"><a href="/itsmontoya/escapist/blob/master/escapist.go" class="js-navigation-open" id="7f54292bd8a7d9eb7f4de8cfe49d345e-c3407e5856f7770ff15135ec84b6550a84f2707c" title="escapist.go">escapist.go</a></span>
          </td>
          <td class="message">
            <span class="css-truncate css-truncate-target">
                  <a href="/itsmontoya/escapist/commit/f6a01e7bef17a8d991876b907b2d3b344b08b2fc" class="message" data-pjax="true" title="Initial commit">Initial commit</a>
            </span>
          </td>
          <td class="age">
            <span class="css-truncate css-truncate-target"><time-ago datetime="2016-07-15T18:09:23Z" title="Jul 15, 2016, 11:09 AM PDT">5 minutes ago</time-ago></span>
          </td>
        </tr>
        <tr class="js-navigation-item">
          <td class="icon">
            <svg aria-hidden="true" class="octicon octicon-file-text" height="16" version="1.1" viewBox="0 0 12 16" width="12"><path d="M6 5H2V4h4v1zM2 8h7V7H2v1zm0 2h7V9H2v1zm0 2h7v-1H2v1zm10-7.5V14c0 .55-.45 1-1 1H1c-.55 0-1-.45-1-1V2c0-.55.45-1 1-1h7.5L12 4.5zM11 5L8 2H1v12h10V5z"></path></svg>
            <img alt="" class="spinner" height="16" src="https://assets-cdn.github.com/images/spinners/octocat-spinner-32.gif" width="16">
          </td>
          <td class="content">
            <span class="css-truncate css-truncate-target"><a href="/itsmontoya/escapist/blob/master/escapist_test.go" class="js-navigation-open" id="db70169d88b8b40ec968d31a1c057c53-4034172adc3a1c9b58b2653e33ac20ec1a3e92c6" title="escapist_test.go">escapist_test.go</a></span>
          </td>
          <td class="message">
            <span class="css-truncate css-truncate-target">
                  <a href="/itsmontoya/escapist/commit/f6a01e7bef17a8d991876b907b2d3b344b08b2fc" class="message" data-pjax="true" title="Initial commit">Initial commit</a>
            </span>
          </td>
          <td class="age">
            <span class="css-truncate css-truncate-target"><time-ago datetime="2016-07-15T18:09:23Z" title="Jul 15, 2016, 11:09 AM PDT">5 minutes ago</time-ago></span>
          </td>
        </tr>
    </tbody>
  </table>

</div>

  <div class="repo-file-upload-tree-target js-document-dropzone js-upload-manifest-tree-view" data-drop-url="/itsmontoya/escapist/upload/master/" data-directory-upload="">
    <div class="repo-file-upload-outline">
      <div class="repo-file-upload-slate">
          <svg width="204" height="52" viewBox="0 0 204 52" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
            <g id="files-lg" class="files-lg" fill="#717171">
              <path class="file-graph" d="M40,42 L40,12 L42,12 L42,40 L64,40 L64,42 L40,42 L40,42 Z M45,37 L45,16 L49,16 L49,37 L45,37 L45,37 Z M51,37 L51,22 L55,22 L55,37 L51,37 L51,37 Z M68.5,48 L35.5,48 C34.7,48 34,47.3 34,46.5 L34,7.5 C34,6.7 34.7,6 35.5,6 L59,6 L70,17 L70,46.5 C70,47.3 69.3,48 68.5,48 L68.5,48 Z M68,18 L58,8 L36,8 L36,46 L68,46 L68,18 L68,18 Z M57,37 L57,27 L61,27 L61,37 L57,37 L57,37 Z"></path>
              <path class="file-zip" d="M17,14 L1,14 C0.4,14 0,14.5 0,15 L0,41 C0,41.5 0.4,42 1,42 L23,42 C23.6,42 24,41.5 24,41 L24,21 L17,14 Z M22,40 L2,40 L2,16 L8,16 L8,18 L10,18 L10,16 L16,16 L22,22 L22,40 Z M8,30.5 C6.8,31.2 6,32.5 6,34 L6,36 L14,36 L14,34 C14,31.8 12.2,30 10,30 L10,28 L8,28 L8,30.5 Z M12,32 L12,34 L8,34 L8,32 L12,32 Z M10,20 L10,18 L12,18 L12,20 L10,20 Z M8,20 L10,20 L10,22 L8,22 L8,20 Z M10,24 L10,22 L12,22 L12,24 L10,24 Z M8,24 L10,24 L10,26 L8,26 L8,24 Z M10,28 L10,26 L12,26 L12,28 L10,28 Z"></path>
              <path class="file-generic" d="M168.5,48 L135.5,48 C134.7,48 134,47.3 134,46.5 L134,7.5 C134,6.7 134.7,6 135.5,6 L159,6 L170,17 L170,46.5 C170,47.3 169.3,48 168.5,48 Z M168,18 L158,8 L136,8 L136,46 L168,46 L168,18 Z M140,35 L140,33 L161,33 L161,35 L140,35 Z M140,30 L140,28 L161,28 L161,30 L140,30 Z M140,25 L140,23 L161,23 L161,25 L140,25 Z M140,17 L140,15 L152,15 L152,17 L140,17 Z M140,40 L140,38 L161,38 L161,40 L140,40 Z"></path>
              <path class="file-acrobat" d="M181,14 C180.5,14 180,14.5 180,15 L180,41 C180,41.5 180.5,42 181,42 L203,42 C203.5,42 204,41.5 204,41 L204,23 L204,21 L197,14 L181,14 Z M200.8,29.9 C200.3,29.8 199.8,29.7 199.3,29.7 C198.5,29.7 197.7,29.8 196.8,29.9 C196.3,29.8 195.7,29.3 194.8,28.6 C193.9,27.9 193.1,26.3 192.2,23.9 C192.5,22.2 192.6,20.9 192.7,19.9 C192.8,18.9 192.8,18.4 192.7,18.4 C192.8,17.6 192.6,17 192.3,16.6 C192,16.2 191.4,16 191,16 L196,16 L202,22 L202,30.4 C201.6,30.2 201.2,30 200.8,29.9 Z M182,16 L190,16 C189.8,16.1 189.6,16.2 189.4,16.4 C189.2,16.6 189,16.9 188.9,17.4 C188.7,18.2 188.6,19.2 188.7,20.3 C188.8,21.5 189.1,22.7 189.4,23.9 C188.9,25.4 188.2,27.2 187.2,29.3 C186.2,31.4 185.6,32.6 185.4,33 C185.1,33.1 184.7,33.3 184,33.6 C183.3,33.9 182.7,34.4 182,35 L182,16 L182,16 Z M195.1,31 C193.8,31.2 192.6,31.4 191.5,31.7 C190.3,32 189.1,32.4 187.8,32.9 L189,30.4 C189.8,28.7 190.4,27.2 190.8,25.8 L190.8,25.9 C191.7,28.2 192.5,29.6 193.1,30.1 C193.8,30.5 194.5,30.8 195.1,31 L195.1,31 Z M184.1,39.2 C185,38.4 186.2,36.9 187.7,34.4 C188.3,34.1 188.9,33.9 189.3,33.7 L190.1,33.4 C191.1,33.1 192,32.9 193,32.7 C194,32.5 195,32.4 196,32.3 C196.9,32.7 197.9,33.1 198.8,33.3 C199.8,33.6 200.6,33.7 201.3,33.8 C201.5,33.8 201.8,34 202,34 L202,40 L182,40 C182.4,39.9 183.6,39.6 184.1,39.2 Z"></path>
              <path class="file-code" d="M111,0 L82,0 C80.9,0 80,0.9 80,2 L80,50 C80,51.1 80.9,52 82,52 L122,52 C123.1,52 124,51.1 124,50 L124,13 L111,0 Z M122,50 L82,50 L82,2 L110,2 L122,14 L122,50 Z M107,18 L116,28 L107,38 L104,35 L111,28 L104,21 L107,18 Z M100,21 L93,28 L100,35 L97,38 L88,28 L97,18 L100,21 Z"></path>
            </g>
          </svg>

          <h2>Drop to upload your files</h2>
      </div>
    </div>
  </div>


  <div id="readme" class="readme boxed-group clearfix announce instapaper_body md">
    <h3>
      <svg aria-hidden="true" class="octicon octicon-book" height="16" version="1.1" viewBox="0 0 16 16" width="16"><path d="M3 5h4v1H3V5zm0 3h4V7H3v1zm0 2h4V9H3v1zm11-5h-4v1h4V5zm0 2h-4v1h4V7zm0 2h-4v1h4V9zm2-6v9c0 .55-.45 1-1 1H9.5l-1 1-1-1H2c-.55 0-1-.45-1-1V3c0-.55.45-1 1-1h5.5l1 1 1-1H15c.55 0 1 .45 1 1zm-8 .5L7.5 3H2v9h6V3.5zm7-.5H9.5l-.5.5V12h6V3z"></path></svg>
      README.md
    </h3>

      <article class="markdown-body entry-content" itemprop="text"><h1><a id="user-content-escapist--" class="anchor" href="#escapist--" aria-hidden="true"><svg aria-hidden="true" class="octicon octicon-link" height="16" version="1.1" viewBox="0 0 16 16" width="16"><path d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Escapist <a href="https://godoc.org/github.com/itsmontoya/escapist"><img src="https://camo.githubusercontent.com/b0b01ed3504223a4acb9601ae4fc48a2656a2bcd/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f6974736d6f6e746f79612f65736361706973743f7374617475732e737667" alt="GoDoc" data-canonical-src="https://godoc.org/github.com/itsmontoya/escapist?status.svg" style="max-width:100%;"></a> <a href="https://camo.githubusercontent.com/a83179c68eea82c5c5a3301c42a8caf31f8c88d5/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f7374617475732d616c7068612d7265642e737667" target="_blank"><img src="https://camo.githubusercontent.com/a83179c68eea82c5c5a3301c42a8caf31f8c88d5/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f7374617475732d616c7068612d7265642e737667" alt="Status" data-canonical-src="https://img.shields.io/badge/status-alpha-red.svg" style="max-width:100%;"></a></h1>

<p>Escapist is a byteslice-focused HTML escaping library</p>

<h2><a id="user-content-benchmarks" class="anchor" href="#benchmarks" aria-hidden="true"><svg aria-hidden="true" class="octicon octicon-link" height="16" version="1.1" viewBox="0 0 16 16" width="16"><path d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Benchmarks</h2>

<div class="highlight highlight-source-shell"><pre><span class="pl-c"># Basic test, very simple replacement</span>
BenchmarkBasic-4        10000000               142 ns/op              48 B/op          1 allocs/op
BenchmarkHTMLBasic-4     5000000               271 ns/op              96 B/op          2 allocs/op
BenchmarkAdvBasic-4     10000000               173 ns/op              48 B/op          1 allocs/op

<span class="pl-c"># No replacement</span>
BenchmarkNoRep-4        30000000                58.3 ns/op             0 B/op          0 allocs/op
BenchmarkHTMLNoRep-4    20000000                64.1 ns/op             0 B/op          0 allocs/op
BenchmarkAdvNoRep-4     20000000               100 ns/op               0 B/op          0 allocs/op

<span class="pl-c"># This is considered to be a very basic escape.</span>

<span class="pl-c"># In contrast, escapist.EscapeAdv looks at:</span>

<h2><a id="user-content-usage" class="anchor" href="#usage" aria-hidden="true"><svg aria-hidden="true" class="octicon octicon-link" height="16" version="1.1" viewBox="0 0 16 16" width="16"><path d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Usage</h2>

<div class="highlight highlight-source-go"><pre><span class="pl-k">package</span> main

<span class="pl-k">import</span> <span class="pl-s"><span class="pl-pds">"</span>fmt<span class="pl-pds">"</span></span>

<span class="pl-k">func</span> <span class="pl-en">main</span>() {
    <span class="pl-smi">b</span> <span class="pl-k">:=</span> <span class="pl-c1">Escape</span>([]<span class="pl-k">byte</span>(<span class="pl-s"><span class="pl-pds">"</span>Hello there Mr. Joe &lt;inject stuff&gt;<span class="pl-pds">"</span></span>))
    fmt.<span class="pl-c1">Println</span>(<span class="pl-k">string</span>(b))
}</pre></div>
</article>
  </div>


  </div>
  <div class="modal-backdrop js-touch-events"></div>
</div>


    </div>
  </div>

    </div>

        <div class="container site-footer-container">
  <div class="site-footer" role="contentinfo">
    <ul class="site-footer-links right">
        <li><a href="https://github.com/contact" data-ga-click="Footer, go to contact, text:contact">Contact GitHub</a></li>
      <li><a href="https://developer.github.com" data-ga-click="Footer, go to api, text:api">API</a></li>
      <li><a href="https://training.github.com" data-ga-click="Footer, go to training, text:training">Training</a></li>
      <li><a href="https://shop.github.com" data-ga-click="Footer, go to shop, text:shop">Shop</a></li>
        <li><a href="https://github.com/blog" data-ga-click="Footer, go to blog, text:blog">Blog</a></li>
        <li><a href="https://github.com/about" data-ga-click="Footer, go to about, text:about">About</a></li>

    </ul>

    <a href="https://github.com" aria-label="Homepage" class="site-footer-mark" title="GitHub">
      <svg aria-hidden="true" class="octicon octicon-mark-github" height="24" version="1.1" viewBox="0 0 16 16" width="24"><path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0 0 16 8c0-4.42-3.58-8-8-8z"></path></svg>
</a>
    <ul class="site-footer-links">
      <li>© 2016 <span title="0.10313s from github-fe137-cp1-prd.iad.github.net">GitHub</span>, Inc.</li>
        <li><a href="https://github.com/site/terms" data-ga-click="Footer, go to terms, text:terms">Terms</a></li>
        <li><a href="https://github.com/site/privacy" data-ga-click="Footer, go to privacy, text:privacy">Privacy</a></li>
        <li><a href="https://github.com/security" data-ga-click="Footer, go to security, text:security">Security</a></li>
        <li><a href="https://status.github.com/" data-ga-click="Footer, go to status, text:status">Status</a></li>
        <li><a href="https://help.github.com" data-ga-click="Footer, go to help, text:help">Help</a></li>
    </ul>
  </div>
</div>



    

    <div id="ajax-error-message" class="ajax-error-message flash flash-error">
      <svg aria-hidden="true" class="octicon octicon-alert" height="16" version="1.1" viewBox="0 0 16 16" width="16"><path d="M8.865 1.52c-.18-.31-.51-.5-.87-.5s-.69.19-.87.5L.275 13.5c-.18.31-.18.69 0 1 .19.31.52.5.87.5h13.7c.36 0 .69-.19.86-.5.17-.31.18-.69.01-1L8.865 1.52zM8.995 13h-2v-2h2v2zm0-3h-2V6h2v4z"></path></svg>
      <button type="button" class="flash-close js-flash-close js-ajax-error-dismiss" aria-label="Dismiss error">
        <svg aria-hidden="true" class="octicon octicon-x" height="16" version="1.1" viewBox="0 0 12 16" width="12"><path d="M7.48 8l3.75 3.75-1.48 1.48L6 9.48l-3.75 3.75-1.48-1.48L4.52 8 .77 4.25l1.48-1.48L6 6.52l3.75-3.75 1.48 1.48z"></path></svg>
      </button>
      Something went wrong with that request. Please try again.
    </div>


      
      <script crossorigin="anonymous" integrity="sha256-FJ2TOMJmUXKHCCXHj6SP3MpNQx0GfL9f2nEg2eOcxzg=" src="https://assets-cdn.github.com/assets/frameworks-149d9338c2665172870825c78fa48fdcca4d431d067cbf5fda7120d9e39cc738.js"></script>
      <script async="async" crossorigin="anonymous" integrity="sha256-x0AdvjBGeXF28kbZUMtOb4DhpIR8Z52SxwOvmE02mM0=" src="https://assets-cdn.github.com/assets/github-c7401dbe3046797176f246d950cb4e6f80e1a4847c679d92c703af984d3698cd.js"></script>
      
      
      
      
      
      
    <div class="js-stale-session-flash stale-session-flash flash flash-warn flash-banner hidden">
      <svg aria-hidden="true" class="octicon octicon-alert" height="16" version="1.1" viewBox="0 0 16 16" width="16"><path d="M8.865 1.52c-.18-.31-.51-.5-.87-.5s-.69.19-.87.5L.275 13.5c-.18.31-.18.69 0 1 .19.31.52.5.87.5h13.7c.36 0 .69-.19.86-.5.17-.31.18-.69.01-1L8.865 1.52zM8.995 13h-2v-2h2v2zm0-3h-2V6h2v4z"></path></svg>
      <span class="signed-in-tab-flash">You signed in with another tab or window. <a href="">Reload</a> to refresh your session.</span>
      <span class="signed-out-tab-flash">You signed out in another tab or window. <a href="">Reload</a> to refresh your session.</span>
    </div>
    <div class="facebox" id="facebox" style="display:none;">
  <div class="facebox-popup">
    <div class="facebox-content" role="dialog" aria-labelledby="facebox-header" aria-describedby="facebox-description">
    </div>
    <button type="button" class="facebox-close js-facebox-close" aria-label="Close modal">
      <svg aria-hidden="true" class="octicon octicon-x" height="16" version="1.1" viewBox="0 0 12 16" width="12"><path d="M7.48 8l3.75 3.75-1.48 1.48L6 9.48l-3.75 3.75-1.48-1.48L4.52 8 .77 4.25l1.48-1.48L6 6.52l3.75-3.75 1.48 1.48z"></path></svg>
    </button>
  </div>
</div>

  


</body></html>
`

	exampleBasic      = []byte(exampleBasicStr)
	exampleBasicNoRep = []byte(exampleBasicNoRepStr)
	exampleHTMLPage   = []byte(exampleHTMLPageStr)

	outB   []byte
	outStr string
)

func TestBasic(t *testing.T) {
	b := Escape(exampleBasic)
	fmt.Println(string(b))
}

func TestBasicComplex(t *testing.T) {
	b := Escape(exampleHTMLPage)
	fmt.Println(string(b))
}

func TestBasicNoRep(t *testing.T) {
	b := Escape(exampleBasicNoRep)
	fmt.Println(string(b))
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outB = Escape(exampleBasic)
	}

	b.ReportAllocs()
}

func BenchmarkComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outB = Escape(exampleHTMLPage)
	}

	b.ReportAllocs()
}

func BenchmarkNoRep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outB = Escape(exampleBasicNoRep)
	}

	b.ReportAllocs()
}

func BenchmarkHTMLBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outStr = html.EscapeString(exampleBasicStr)
	}

	b.ReportAllocs()
}

func BenchmarkHTMLComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outStr = html.EscapeString(exampleHTMLPageStr)
	}

	b.ReportAllocs()
}

func BenchmarkHTMLNoRep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outStr = html.EscapeString(exampleBasicNoRepStr)
	}

	b.ReportAllocs()
}

func BenchmarkAdvBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outB = EscapeAdv(exampleBasic)
	}

	b.ReportAllocs()
}

func BenchmarkAdvComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outB = EscapeAdv(exampleHTMLPage)
	}

	b.ReportAllocs()
}

func BenchmarkAdvNoRep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outB = EscapeAdv(exampleBasicNoRep)
	}

	b.ReportAllocs()
}
