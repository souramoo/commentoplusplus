package main

var version string

// omega service endpoints to update last mods for pages

const StagingUrl string = "https://omega-dot-getmega-app.appspot.com/twirp/consoleapi.pb.Website/UpdateUrlLastModTime"
const ProdUrl string = "https://omega-dot-mega-prod-ceef8.appspot.com/twirp/consoleapi.pb.Website/UpdateUrlLastModTime"

const CommentoProdUrl string = "https://commento.getmega.tech"
const CommentoStagingUrl string = "https://staging-commento.getmega.tech"
