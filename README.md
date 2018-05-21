# Quick Setup

- Clone this repo
- `dep ensure`
- Start a MySQL instance (v5.6+) and create an empty DB
- Import db_scripts/db-build.sql to create the tables.
- Copy env.json.example to env.json and fill in your DB details
- In the Session section, you should generate new passwords for the following keys:
-- AuthKey should be a 64 byte password and then base64 encoded
-- EncryptKey should be a 32 byte password and then base64 encoded
-- CSRFKey should be a 32 byte password and then base64 encoded
- Run the application using the command: go run contractors.go

# Assets

The asset folder contains a dynamic folder and a static folder.

The dynamic (private) folder contains the SASS and JavaScript source files, and a large PNG image which is used to generate favicons for different platforms. This is the folder in which you want to make your changes.

The static (public) folder contains the minified CSS and JavaScript as well as the generated favicons. The static folder is designed to be served up to the browser.

To regenerate static files after making changes to dynamic files:

```bash
# Install Gulp globally
npm install -g gulp-cli

# Install Gulp locally and dependencies from package.json
npm install
```

Once the environment is set up, you should have your terminal open to the root of the project folder. There are a couple commands you can use with Gulp that are in the gulpfile.js.

```bash
# Compile the SASS from asset/dynamic/sass and store CSS in asset/static/css/all.css
gulp sass

# Concat the JavaScript from asset/dynamic/js and store JS in asset/static/js/all.js
gulp javascript

# Copy the jQuery files from node_modules/jquery to asset/static/js
gulp jquery

# Copy the Bootstrap files from node_modules/bootstrap to asset/static
gulp bootstrap

# Copy the Underscore files from node_modules/underscore to asset/static/js
gulp underscore

# Run tasks favicon-generate and favicon-inject
gulp favicon

# Generate favicons from asset/dynamic/logo.png and copy to /asset/static/favicon
gulp favicon-generate

# Generate view/partial/favicon.tmpl with favicon tags
gulp favicon-inject

# Update the asset/dynamic/favicon/data.json file with the latest version from the RealFaviconGenerator website
gulp favicon-update

# Run the sass and javascript tasks when any of the files change
gulp watch

# Run all the tasks once
gulp init

# Run just the sass and javascript tasks once
gulp default
```

It is best to run gulp watch so when you are working with the SASS and JavaScript files so they will automatically generate in the static folder for you.
