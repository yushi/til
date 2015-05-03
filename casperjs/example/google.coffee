casper = require('casper').create()

casper.start 'http://google.com', ->
  @fill "form[action='/search']"
    q: 'coffee-script'
    true

casper.then ->
  @echo @getTitle()

casper.run()
