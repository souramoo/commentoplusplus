package main

var version string

// omega service endpoints to update last mods for pages

const StagingEndPoint string = "https://omega-dot-getmega-app.appspot.com/twirp/consoleapi.pb.Website/UpdateUrlLastModTime"
const ProdEndPoint string = "https://omega-dot-mega-prod-ceef8.appspot.com/twirp/consoleapi.pb.Website/UpdateUrlLastModTime"

const CommentoProdDomain string = "www.getmega.com"
const CommentoStagingDomain string = "getmega-app.web.app"
