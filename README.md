# ECOSystem

## Modular e-commerce server written in Golang

### Note: this project is in very early stage development - the API is likely to change considerably and the feature set is currently too incomplete to implement a fully working e-commerce site.  Having said that, it's under active development and I hope to be using it in production soon.  If you're interested in the project (using, helping etc), please contact me directly to discuss.

### Quick Summary

ECOSystem arises from a frustration with existing and traditional approaches to building e-commerce websites.  I experienced these frustrations over 10 years as the owner of an e-commerce business and the main developer responsible for our site.

I believe that the software a merchant chooses to run their e-commerce site should be as independent as possible from the other software, tools and processes they use to run their business.  This approach of loosely-coupled modules allows a merchant to pick and choose the very best software available for each discrete function within their business.  ECOsystem is designed to implement just one layer in a loosely-coupled 'ecosystem' of e-commerce components - it would typically sit on top of a merchant's enterprise software and other backends, as well as interface with other cloud-based APIs, like content providers, review software, analytics, payment processing etc.

*This is a very rough and ready summary of the main things you need to know about ECOSystem, including architecture, design philosophy and motivation:*

- Not just a backend: ECOSystem includes functionality to actually generate and display an e-commerce website.  It does this in a data-driven, CMS-less way - using a JSON feed of product information which is supposed to be taken directly from the merchant's business backend software.  There is no replication of data required - no secondary 'website' database to maintain - integration is primary.  Merchant's manage their inventory in their usual software, not a blogging platform, bloated CMS or some other proprietory database.

- Traditional client-server architecture: ECOSystem is not a RESTful JSON API, nor is it a client-side Javascript app.  ECOSystem generates pages and partials.  The 'slow' information on a page (think product descriptions, images, category layout etc) is rendered on the server and delivered in one HTML request.  Faster moving 'components' (e.g the cart) are also rendered server-side and delivered as partials to be loaded into the relevent portions of the page.  The result is an SEO-friendly site, with fast initial load and snappy, app-like interface.
