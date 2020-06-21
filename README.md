# Commento++

DEMO: https://demo.souradip.com/chat.html

### Modified to Run on Heroku in Docker Container and autoupdate with GitLab upstream

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

This repo is set up to pull updates from the [GitLab master of Commento](https://gitlab.com/commento/commento/) and merge them with [fidiego's heroku deploy script](https://github.com/fidiego/commento-heroku). It's also set to autodeploy when there's a new commit onto my heroku, so my version should stay constantly updated ;)

Source: https://gitlab.com/commento/commento/

### Now with all my patches out of the box!

Changes from upstream:
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

(Have sent in merge requests, don't know when they'll be accepted, so here's a ready to use version!)


### Commento

##### [Homepage](https://commento.io) &nbsp;&ndash;&nbsp; [Demo](https://demo.commento.io) &nbsp;&ndash;&nbsp; [Documentation](https://docs.commento.io) &nbsp;&ndash;&nbsp; [Contributing](https://docs.commento.io/contributing/) &nbsp;&ndash;&nbsp; [#commento on Freenode](http://webchat.freenode.net/?channels=%23commento)

Commento is a platform that you can embed in your website to allow your readers to add comments. It's reasonably fast lightweight. Supports markdown, import from Disqus, voting, automated spam detection, moderation tools, sticky comments, thread locking, OAuth login, single sign-on, and email notifications.

###### How is this different from Disqus, Facebook Comments, and the rest?

Most other products in this space do not respect your privacy; showing ads is their primary business model and that nearly always comes at the users' cost. Commento has no ads; you're the customer, not the product. While Commento is [free software](https://www.gnu.org/philosophy/free-sw.en.html), in order to keep the service sustainable, the [hosted cloud version](https://commento.io) is not offered free of cost. Commento is also orders of magnitude lighter than alternatives.

###### Why should I care about my readers' privacy?

For starters, your readers value their privacy. Not caring about them is disrespectful and you will end up alienating your audience; they won't come back. Disqus still isn't GDPR-compliant (according to their <a href="https://help.disqus.com/terms-and-policies/privacy-faq" title="At the time of writing (28 December 2018)" rel="nofollow">privacy policy</a>). Disqus adds megabytes to your page size; what happens when a random third-party script that is injected into your website turns malicious?

#### Installation

Read the [documentation to get started](https://docs.commento.io/installation/).

#### Contributing

If this is your first contribution to Commento, please go through the [contribution guidelines](https://docs.commento.io/contributing/) before you begin. If you have any questions, join [#commento on Freenode](http://webchat.freenode.net/?channels=%23commento).
