# Commento++

### 💬 Try it out and deploy your own
[LIVE DEMO](https://demo.souradip.com/chat.html)

[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/souramoo/commentoplusplus) 

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template/hame0C)

### ❓ About
Commento++ is a free, open source, fast & lightweight comments box that you can embed in your static website instead of Disqus.

### ⚡ Features
- Markdown support
- Import from Disqus
- Voting
- Automated spam detection (Askimet and Perspective integration)
- Moderation tools
- Sticky comments
- Thread locking
- OAuth login (Google, Github, Twitter) and single sign-on
- Hot-reloading of comments
- Email notifications.

### 🤝 Support
Please [(donate)](https://paypal.me/souramoo) if you find my work helpful (this will always remain free and open source)!

### 📷 Screenshots
![Commento++ in action](https://i.imgur.com/x4IA22n.gif)

### 🤔 How is this different from Disqus, Facebook Comments, and the rest?

- 🐱‍👤  Respects your privacy and no adverts
- 💄 Prettier comments box compared to other FOSS alternatives
- ⚡ Orders of magnitude lighter and faster than alternatives
- 🕐 One click to deploy your own instance to a free Heroku account in seconds
- 🔌 You can self-host too for maximum control!

### Get started

To start you just need to launch an instance.

The button below will work for a Heroku account:

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/souramoo/commentoplusplus)

The button below will work for a free Railway account:

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template/hame0C)

Otherwise, most of the below is the same as documented at https://docs.commento.io

If you want to self-host, you will need a PostgreSQL server handy and then:
1) Use this repo's Dockerfile if you're into that kind of thing
2) Download the plug and play pre-compiled version from the [releases](https://github.com/souramoo/commentoplusplus/releases)
3) To build yourself, you can clone this repo (you will require `nodejs`, `yarn`, `golang` installed) and run `make prod` and you will generate `./build/prod/commento`


To launch, you should configure the following environment variables below:
```
$ export COMMENTO_ORIGIN=http://commento.example.com:8080
$ export COMMENTO_PORT=8080
$ export COMMENTO_POSTGRES=postgres://username:password@postgres.example.com:5432/commento?sslmode=disable
$ export COMMENTO_CDN_PREFIX=$COMMENTO_ORIGIN

```

And then you can run the `commento` binary.

#### Logging and graphing page views

The logging defaults to off to preseve disk space, but you can specify the `COMMENTO_ENABLE_LOGGING` environment variable to true to enable each page view being logged and a nice graph generated on your dashboard. This will however use up a lot of space quickly if you have a high traffic website; you may want to consider a more fully-featured analytics solution for your website.

e.g.

```
$ export COMMENTO_ENABLE_LOGGING=true
```

to turn this feature on.

#### Wildcard domain support
A new feature added recently, with better edge-case handling of domain names, etc.

This feature however will open up your commento instance to abuse if it is shared between a lot of people (e.g. people registering `e%` to register *every domain beginning with e*...)

As most of commento++ instances are serving one user only, I have assumed you will be sensible about this and enabled wildcard domain support by default. 

If you want the old behaviour, you can disable this with an environment variable:

```
$ export COMMENTO_ENABLE_WILDCARDS=false
```

#### Perspective spam detection 
To enable Perspective (https://www.perspectiveapi.com/) spam detection add following environment variables:

```
$ export COMMENTO_PERSPECTIVE_KEY=YOUR_API_KEY_FROM_PERSPECTIVE
$ export COMMENTO_PERSPECTIVE_LIMIT=0.5
$ export COMMENTO_PERSPECTIVE_LANGUAGE=en
```

COMMENTO_PERSPECTIVE_KEY:
To create a new Perspective API key follow the instructions at https://developers.perspectiveapi.com/s/docs-get-started

COMMENTO_PERSPECTIVE_LIMIT:
The limit defines the minimum value for the Perspective probability summary score. Everything above the value will be flagged in Commentoplus. (Default 0.5)

COMMENTO_PERSPECTIVE_LANGUAGE:
Configure the language to your requirements. (Default: en)

Make sure that you have enabled the automatic spam detection in the dashboard.


#### Disabling SMTP Host verification check

Commento++ allows configuration of the tlsConfig for both SMTPS as well as StartTLS for email sending.
For context, this is required for the [https://cloudron.io/](Cloudron) app package.

To skip SMTP hostname verification, use:

```
$ export SMTP_SKIP_HOST_VERIFY=true
```

#### STARTTLS
If you have any issues with email sending relating to not using STARTTLS, consider:

```
$ export USE_STARTTLS=true
```


#### Docker setup
Alternatively you can use the pre-build images from:
- https://gitlab.com/caroga/commentoplusplus-docker
- https://hub.docker.com/r/caroga/commentoplusplus

Instructions for configuring the docker image can be found [here](https://docs.commento.io/installation/self-hosting/on-your-server/docker.html). Are you missing a version? Please contact @caroga [here](https://gitlab.com/caroga/commentoplusplus-docker).


### Finally

Once you have created an account in your commento instance, it should give you instructions on how to embed this into your site! It should be as simple as:

```
<script defer src="https://(server url)/js/commento.js"></script>
<div id="commento"></div>
```

### If you're running this behind nginx/another reverse proxy
Remember to either forward the websockets through to commento in your nginx config, e.g.:

```
location / {
    proxy_pass http://commento;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "Upgrade";
    proxy_set_header Host $host;
}
```

Or if you'd rather not do that, disable websockets in favour of HTTP polling by adding `data-no-websockets="true"` to the commento <script> tag (or `data-no-livereload="true"`` to only load comments on page load, see below!)

### SSL Support
Commento++ supports native SSL without use of an nginx proxy. Three properties are required for Native SSL:

- COMMENTO_SSL
- COMMENTO_SSL_CERT
- COMMENTO_SSL_KEY

`COMMENTO_SSL=true` enables native SSL. Default is false.

If `COMMENTO_SSL=true` then `COMMENTO_SSL_CERT` and `COMMENTO_SSL_KEY` must be set to the path to a valid SSL Certificate and Key pair.

### More options to configure commento's frontend

You can add the following to commento's script tag:

- `data-css-override="http://server/styles.css"` - A URL to a CSS file with overriding styles. Defaults to no override and uses Commento's default theme.
- `data-auto-init="false"` - Commento automatically initialises itself when the page is loaded. If you prefer to load Commento dynamically (for example, after the user clicks a button), you can disable this. You will be required to call `window.commento.main()` when you want to load Commento. By default, this is true.
- `data-id-root="notcommento"` - By default, Commento looks for a `<div>` with `id="commento"`. If you want to load Commento in a different element, you can set this attribute to the ID of that element.
- `data-no-fonts="true"` - By default, Commento uses the Source Sans Pro font to present a good design out-of-the-box. If you'd like to disable this so that Commento never loads the font files, you can set this to true. By default, this is true.
- `data-hide-deleted` - By default, deleted comments with undeleted replies are shown with a "[deleted]" tag. If you'd like to disable this, setting this to true will hide deleted comments even if there are legitimate replies underneath. Deleted comments without any undeleted comments underneath are hidden irrespective of the value of this function. By default, this is false.
- `data-no-websockets="true"` - Disables websocket functionality in favour of HTTP polling to have the same live reload functionality in a situation where websockets aren't allowed (e.g. a reverse proxy)
- `data-no-livereload="true"` - Disabled all hot reload functionality (this supercedes the above flag) - all comments are loaded once and only once on page load.

e.g. Usage example:
```
<script defer src="https://chat.mookerj.ee/js/commento.js" data-no-websockets="true"></script>
```

### How is this different to the original Commento?
Original source is from @adtac at https://gitlab.com/commento/commento/ - this fork is largely a result of me getting carried away fixing a lot of bugs but the original maintainer seemingly disappearing!

(Inconclusive) list of changes from upstream:
- [NEW FEATURE: Auto refreshing comments with WebSockets for push updates](https://gitlab.com/commento/commento/-/merge_requests/168)
- NEW FEATURE: Window title updates when there's new activity
- NEW FEATURE: Permalinks, and a subtle yellow highlight animation for new comments when they come in live
- NEW FEATURE: Smooth scrolling
- NEW FEATURE: Hide +/- if no children
- NEW FEATURE: Errors now slide down from the top rather than the ugly error system before
- [NEW FEATURE: Guests can leave their name](https://gitlab.com/commento/commento/-/merge_requests/169)
- [FIXED: Twitter profile photo bug](https://gitlab.com/commento/commento/-/merge_requests/159)
- [FIXED: Duplicate comment bug on login](https://gitlab.com/commento/commento/-/merge_requests/160)
- [FIXED: Add target="_blank" to all external links, while also adding "noopener" to prevent XSS](https://gitlab.com/commento/commento/-/merge_requests/161)
- [FIXED: Allow anchor links onto same page](https://gitlab.com/commento/commento/-/merge_requests/162)
- [NEW FEATURE: Comment moderation dashboard, to approve/delete comments across your entire domain from one place](https://gitlab.com/commento/commento/-/merge_requests/163)
- [NEW FEATURE: MathJax support hook, will plug in to any MathJax library included on the same page commento is on](https://gitlab.com/commento/commento/-/merge_requests/164)
- [NEW FEATURE: Press enter to log in after entering your password](https://gitlab.com/commento/commento/-/merge_requests/167)
- [FIXED: Deleted comments not returned in array](https://gitlab.com/commento/commento/-/merge_requests/170)
- [NEW FEATURE: Reinit widget functionality for Single Page Applications](https://gitlab.com/commento/commento/-/merge_requests/182)
- NEW FEATURE: Wildcards possible in domain name (so can serve %.example.com)
- NEW FEATURE: Support of the Perspective API for spam detection (https://www.perspectiveapi.com/)

I've sent in merge requests for a lot of the above but I don't know when they'll be accepted, so here's a ready to use version with all batteries included to help out fellow bloggers!

### How to use this in a SPA (Single Page Application)

Commento++ runs a bit of code on page load to initialize the widget. This widget can be customized by using data attributes on the script tag. When using commento++ in a SPA you might want to change the pageId for the widget when navigating to a new blog post without a browser page load. Below you'll find an example for an Commento++ component in React:

```js
import React, { useEffect } from 'react'

const Commento = ({ pageId }) => {
  useEffect(() => {
    if (typeof window !== 'undefined' && !window.commento) {
      // init empty object so commento.js script extends this with global functions
      window.commento = {}
      const script = document.createElement('script')
      // Replace this with the url to your commento instance's commento.js script
      script.src = `http://localhost:8080/js/commento.js`
      script.defer = true
      // Set default attributes for first load
      script.setAttribute('data-auto-init', false)
      script.setAttribute('data-page-id', pageId)
      script.setAttribute('data-id-root', 'commento-box')
      script.onload = () => {
        // Tell commento.js to load the widget
        window.commento.main()
      }
      document.getElementsByTagName('head')[0].appendChild(script)
    } else if (typeof window !== 'undefined' && window.commento) {
      // In-case the commento.js script has already been loaded reInit the widget with a new pageId
      window.commento.reInit({
        pageId: pageId,
      })
    }
  }, [])

  return <div id="commento-box" />
}

export default Commento
```

Commento initializes itself and extends the `window.commento` object. When you have an HTML element with the id `commento` this will live on the `window.commento` namespace. Replacing the HTML element (as SPAs do) the `window.commento` is reset to the new element, losing all extended functionality provided by the commento++ script. Make sure to provide a `data-id-root` other than `commento` for this to work, see `commento-box` in the example above. 

The `window.commento.reInit` function can be called with the following updated options (all optional):

```js
{
    pageId: "string", // eg: "path/to/page"
    idRoot: "string", // eg: "new-element-id"
    noFonts: "string", // Boolean string, "true" or "false"
    hideDeleted: "string", // Boolean string, "true" or "false"
    cssOverride: "string" // or null to reset to undefined
}
```
